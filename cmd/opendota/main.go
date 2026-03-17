package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/kecci/dota2/dotalive"
	"github.com/kecci/dota2/opendota"
)

// STEAM_API_KEY=<your-steam-key> go run cmd/opendota/main.go -match_id=8733022580
func main() {
	// 1. Setup your mapper for hero names
	// OpenDota uses the same Hero IDs as Valve
	apiKey := os.Getenv("OPENDOTA_API_KEY") // Optional for OpenDota free tier
	steamKey := os.Getenv("STEAM_API_KEY")  // Required for SteamApiKey

	matchID := flag.Int("match_id", 0, "match_id")
	flag.Parse()

	// Hydrate the hero names from a reliable source
	// Initialize and hydrate the hero cache
	mapper := dotalive.NewHeroMapper()
	if err := mapper.Update(steamKey, http.DefaultClient); err != nil {
		log.Printf("Warning: could not fetch hero list: %v", err)
	}

	// 2. Initialize the OpenDota client
	// OpenDota doesn't require a key for low-volume testing (2k requests/day)
	client := opendota.NewClient(apiKey)

	if matchID != nil && *matchID > 0 {
		detail, err := client.GetMatchDetail(int64(*matchID))
		if err != nil {
			panic(err)
		}
		b, _ := json.MarshalIndent(detail, "", "   ")
		println(string(b))
		return
	}

	fmt.Println("🚀 Starting Dota 2 Live Match Monitor (OpenDota)...")

	matches, err := client.GetOpenDotaLive()
	if err != nil {
		log.Printf("Error fetching matches: %v", err)
	}

	if len(matches) == 0 {
		fmt.Println("No live professional matches found at the moment.")
	}

	slices.SortFunc(matches, func(a, b opendota.OpenDotaLiveMatch) int {
		if a.LobbyID < b.LobbyID {
			return 1
		}
		if a.LobbyID == b.LobbyID {
			return 0
		}
		return -1
	})

	fmt.Printf("\n--- Found %d Live Matches ---\n", len(matches))
	for _, match := range matches {
		fmt.Printf("Match [%s] | %s vs %s  | Avg MMR: %d | Score: %d - %d\n",
			match.MatchID, match.TeamNameRadiant, match.TeamNameDire, match.AverageMmr, match.RadiantScore, match.DireScore)

		for _, p := range match.Players {
			heroName := mapper.GetName(p.HeroID)

			// OpenDota /live returns these details per player
			fmt.Printf("  - %-15s | Lvl: %-2d | Hero: %-12s | NetWorth: %d\n",
				p.Name, p.Level, heroName, p.NetWorth)
		}
	}
}
