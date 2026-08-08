package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/parse"
)

type ChallengeState struct {
	Target  int64 `json:"target"`
	Current int64 `json:"current"`
}

func RunISKChallenge() {
	switch len(os.Args) {
	case 2:
		state, err := loadChallenge()
		if err != nil {
			fmt.Println("No saved challenge.")
			return
		}

		printChallenge(state)

	case 4:
		if os.Args[2] == "add" {
			state, err := loadChallenge()
			if err != nil {
				fmt.Println("No saved challenge.")
				return
			}

			amount, err := parse.ISK(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}

			state.Current += int64(amount)

			if err := saveChallenge(state); err != nil {
				fmt.Println(err)
				return
			}

			printChallenge(state)
			return
		}

		target, err := parse.ISK(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		current, err := parse.ISK(os.Args[3])
		if err != nil {
			fmt.Println(err)
			return
		}

		state := ChallengeState{
			Target:  int64(target),
			Current: int64(current),
		}

		if err := saveChallenge(state); err != nil {
			fmt.Println(err)
			return
		}

		printChallenge(state)

	default:
		fmt.Printf("Usage:\n")
		fmt.Printf("  %s isk-challenge\n", os.Args[0])
		fmt.Printf("  %s isk-challenge <target> <current>\n", os.Args[0])
		fmt.Printf("  %s isk-challenge add <amount>\n", os.Args[0])
	}
}

func challengePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".eve-trader/isk-challenge.json"
	}

	return filepath.Join(home, ".eve-trader", "isk-challenge.json")
}

func loadChallenge() (ChallengeState, error) {
	data, err := os.ReadFile(challengePath())
	if err != nil {
		return ChallengeState{}, err
	}

	var state ChallengeState

	if err := json.Unmarshal(data, &state); err != nil {
		return ChallengeState{}, err
	}

	return state, nil
}

func saveChallenge(state ChallengeState) error {
	path := challengePath()

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func printChallenge(state ChallengeState) {
	remaining := max(state.Target-state.Current, 0)
	progress := float64(state.Current) / float64(state.Target) * 100

	fmt.Println("🚀 EVE ISK CHALLENGE")
	fmt.Println("===================================")
	fmt.Printf("Target:      %s\n", format.ISK(state.Target))
	fmt.Printf("Exploration: %s\n", format.ISK(state.Current))
	fmt.Printf("Remaining:   %s\n", format.ISK(remaining))
	fmt.Printf("Progress:    %.1f%%\n", progress)
}
