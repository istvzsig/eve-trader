#!/usr/bin/env bash
# dump-project.sh
#
# Creates a single “project dump” text file by recursively concatenating the
# contents of files in the current directory.
#
# Core behavior
# - Removes any existing output file (project-dump.txt).
# - Writes a header containing the current date/time.
# - Recursively finds all regular files under the current directory.
# - Excludes:
#   - .git contents
#   - the output file itself
#   - common build/vendor directories (configurable below)
# - Optionally skips:
#   - binary files (best-effort, using the `file` command when available)
#   - oversized files (configurable byte limit)
# - Protects against accidentally dumping sensitive content by skipping files
#   whose paths match common secret/credential/private-key patterns and common
#   sensitive file extensions.
#
# Output format
# - Each included file is appended in full, preceded by separators and a line:
#     FILE: <path>
#
# Usage
# - ./dump-project.sh
#
# Configuration knobs (edit in the script)
# - OUTPUT: output filename
# - SKIP_BINARY: true/false
# - MAX_BYTES: max file size in bytes (0 disables size limiting)
# - SENSITIVE_PATTERNS: substrings; if present in the path, the file is skipped
# - SENSITIVE_EXTENSIONS: file suffixes; if matched, the file is skipped
# - SKIP_DIRS: directories skipped entirely (and their contents)

set -euo pipefail

OUTPUT="project-dump.txt"
rm -f "$OUTPUT"

echo "Project dump: $(date)" >> "$OUTPUT"
echo "==========================" >> "$OUTPUT"

SKIP_BINARY=true
MAX_BYTES=0

SENSITIVE_PATTERNS=(
  ".env" "env." "dotenv"
  ".env." "secrets" "secret"
  "token" "tokens" "apikey" "api-key" "api_key"
  "accesskey" "access-key" "access_key"
  "private" "privkey" "private_key"
  "credential" "credentials"
  "bearer" "oauth" "oidc"
  "cookie" "cookies"
  "jwt"
  "id_rsa" "id_ed25519"
  "server.key" "private.pem" "public.pem"
  "keystore" "jks" "p12" "pfx"
  "backup" "backup.tar" "dump" "backup.zip"
)

SENSITIVE_EXTENSIONS=(
  ".pem" ".key" ".p12" ".pfx" ".jks" ".jceks" ".crt"
  ".cer" ".der" ".csr" ".ovpn"
)

SKIP_DIRS=(
  "./.git"
  "./.github"
  "./.gitlab"
  "./.vscode"
  "./node_modules"
  "./dist"
  "./build"
  "./target"
  "./vendor"
)

is_sensitive_path() {
  local file="$1"

  for p in "${SENSITIVE_PATTERNS[@]}"; do
    [[ "$file" == *"$p"* ]] && return 0
  done
  for ext in "${SENSITIVE_EXTENSIONS[@]}"; do
    [[ "$file" == *"$ext" ]] && return 0
  done
  return 1
}

is_binary_file() {
  local file="$1"
  command -v file >/dev/null 2>&1 || return 1

  file -b --mime "$file" 2>/dev/null | grep -qiE \
    'application/(octet-stream|x-binary|x-msdownload)|image/|audio/|video/' && return 0

  file -b "$file" 2>/dev/null | grep -qiE 'charset=binary|binary data' && return 0

  return 1
}

# Build prune expression safely (no eval).
# We'll pass it to find as positional args.
prune_args=()
for d in "${SKIP_DIRS[@]}"; do
  # prune the directory itself and everything under it
  prune_args+=( -path "$d" -o -path "$d/*" )
done

# Parenthesize the prune logic for correct precedence.
# We do: find . \( ( -path d1 -o -path d1/* ) -o ( -path d2 -o -path d2/* ) ... \) -prune -o -type f ...
# To achieve that without eval, we group using explicit parentheses and careful assembly.
find_args=(
  . 
  \( 
)

# Assemble grouped -path clauses: ((d1) -o (d1/*)) -o ((d2) -o (d2/*)) ...
# We'll iterate again to insert parentheses per directory.
for i in "${!SKIP_DIRS[@]}"; do
  d="${SKIP_DIRS[$i]}"
  find_args+=( \( -path "$d" -o -path "$d/*" \) )
  if [[ "$i" -lt $((${#SKIP_DIRS[@]}-1)) ]]; then
    find_args+=( -o )
  fi
done

find_args+=(
  \)
  -prune
  -o
  -type f
  ! -name "$OUTPUT"
  -print0
)

# Note: We append -print0 at the end; reading handles NUL-separated paths.
find "${find_args[@]}" | while IFS= read -r -d '' file; do
  if is_sensitive_path "$file"; then
    continue
  fi

  if [[ "$SKIP_BINARY" == "true" ]] && is_binary_file "$file"; then
    continue
  fi

  if [[ "$MAX_BYTES" -gt 0 ]]; then
    bytes="$(wc -c < "$file" 2>/dev/null || echo 0)"
    [[ "$bytes" -gt "$MAX_BYTES" ]] && continue
  fi

  {
    echo ""
    echo "========================================"
    echo "FILE: $file"
    echo "========================================"
    cat "$file"
    echo ""
  } >> "$OUTPUT"
done

echo "Created $OUTPUT"
