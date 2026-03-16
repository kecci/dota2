package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kecci/dota2/dotalive"
)

func main() {
	// 1. Setup your mapper for hero names
	// OpenDota uses the same Hero IDs as Valve
	apiKey := os.Getenv("OPENDOTA_API_KEY") // Optional for OpenDota free tier
	mapper := dotalive.NewHeroMapper()

	// Hydrate the hero names from a reliable source
	if err := mapper.Update(apiKey, http.DefaultClient); err != nil {
		log.Printf("Warning: Hero mapper failed to hydrate: %v", err)
	}

	// 2. Initialize the OpenDota client
	// OpenDota doesn't require a key for low-volume testing (2k requests/day)
	client := dotalive.NewClient(apiKey)

	fmt.Println("🚀 Starting Dota 2 Live Match Monitor (OpenDota)...")

	// 3. Simple Polling Loop
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		matches, err := client.GetOpenDotaLive()
		if err != nil {
			log.Printf("Error fetching matches: %v", err)
			continue
		}

		if len(matches) == 0 {
			fmt.Println("No live professional matches found at the moment.")
			continue
		}

		fmt.Printf("\n--- Found %d Live Matches ---\n", len(matches))
		for _, match := range matches {
			fmt.Printf("Match [%s] | Avg MMR: %d | Score: %d - %d\n",
				match.MatchID, match.AverageMmr, match.RadiantScore, match.DireScore)

			for _, p := range match.Players {
				heroName := mapper.GetName(p.HeroID)

				// OpenDota /live returns these details per player
				fmt.Printf("  - %-15s | Lvl: %-2d | Hero: %-12s | NetWorth: %d\n",
					p.Name, p.Level, heroName, p.NetWorth)
			}
		}
	}
}
