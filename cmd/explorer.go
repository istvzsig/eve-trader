package main

import "fmt"

func RunExplorer() {
	fmt.Println(`
╔══════════════════════════════════════════════════════════════╗
║                    EVE EXPLORER GUIDE                        ║
╚══════════════════════════════════════════════════════════════╝

NULLSEC ROUNDTRIP
──────────────────────────────────────────────────────────────

1. NOISE-5 'NEEDLEJACK'
   K-space → random Nullsec

   Use:
     • Skip many jumps into Nullsec.
     • Good for exploration.
     • Always have an exit strategy.

   Remember:
     • Scan immediately after landing.
     • Create a SAFE bookmark.
     • Create a PING bookmark near important locations.
     • Never assume the destination is safe.

2. SIGNAL-5 'NEEDLEJACK'
   K-space → active Nullsec

   Use:
     • Nullsec with higher player activity.
     • More likely to find players.
     • More PvP risk.

   Good for:
     • PvP hunting.
     • Active regions.
     • Exploration with more activity.

3. HOME-5 'POCHVEN'
   K-space → Pochven

   Use:
     • Pochven Express.
     • Travel shortcut.
     • Access Triglavian space.

   BEFORE ENTERING:
     • Carry an EXIT plan.
     • Carry Core Scanner Probes.
     • Know which filament you will use to leave.

POCHVEN EXPRESS
──────────────────────────────────────────────────────────────

Carry:

   HOME-5
      K-space → Pochven

   PROXIMITY-5 'EXTRACTION'
      Pochven → nearby K-space

   GLORIFICATION-1 'DEVANA'
      Pochven → Triglavian Minor Victory system

Normal exit:

   Pochven
      ↓
   WAIT FOR FILAMENT TIMER
      ↓
   Proximity-5
      ↓
   Nearby K-space
      ↓
   Navigate to Jita / home

If you land in a BAD Pochven location:

   Pochven
      ↓
   WAIT
      ↓
   Proximity-5
      ↓
   K-space
      ↓
   Home-5
      ↓
   Pochven
      ↓
   REROLL
      ↓
   Try again

If a Minor Victory system is conveniently located:

   Pochven
      ↓
   WAIT
      ↓
   Glorification
      ↓
   Minor Victory system
      ↓
   K-space
      ↓
   Jita / home

EXPLORATION SURVIVAL
──────────────────────────────────────────────────────────────

ALWAYS:

   ✓ Carry Core Scanner Probes
   ✓ Have an MWD
   ✓ Bookmark SAFE
   ✓ Bookmark PING
   ✓ Watch Local in Nullsec
   ✓ D-scan regularly
   ✓ Check gates before warping
   ✓ Keep an escape plan
   ✓ Don't carry everything you own

NEVER:

   ✗ Assume a Nullsec system is empty
   ✗ Warp blindly to a gate
   ✗ Enter a bubble without checking
   ✗ Filament without knowing how you will leave
   ✗ Carry your entire wallet in one Heron

BOOKMARK NAMING
──────────────────────────────────────────────────────────────

Wormhole:

   NS - ENTER - C4/C5
   WH - EXIT - HS

Nullsec:

   SYSTEM - SAFE
   SYSTEM - PING - GATE
   SYSTEM - ENTER - C4/C5

Example:

   MMUF-8 - SAFE
   MMUF-8 - PING - JX0-S1
   MMUF-8 - ENTER - C4/C5

BASIC EXPLORATION LOOP
──────────────────────────────────────────────────────────────

   K-SPACE
      ↓
   Noise-5 / Signal-5
      ↓
   NULLSEC
      ↓
   SAFE
      ↓
   Scan signatures
      ↓
   Relic / Data sites
      ↓
   Loot
      ↓
   Find EXIT
      ↓
   K-SPACE
      ↓
   SELL LOOT
      ↓
   REPEAT

MAIN RULE
──────────────────────────────────────────────────────────────

   "Never enter dangerous space
    without knowing how you are getting out."`)
}
