package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kecci/dota2/opendota"
)

func main() {
	// 1. Setup your mapper for hero names
	// OpenDota uses the same Hero IDs as Valve
	apiKey := os.Getenv("OPENDOTA_API_KEY") // Optional for OpenDota free tier

	// 2. Initialize the OpenDota client
	// OpenDota doesn't require a key for low-volume testing (2k requests/day)
	client := opendota.NewClient(apiKey)

	fmt.Println("🚀 Starting Dota 2 Live Match Monitor (OpenDota)...")

	// 3. Simple Polling Loop
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	matchDetail, err := client.GetMatchDetail(8732939854)
	if err != nil {
		log.Printf("Error fetching matches: %v", err)
	}
	b, _ := json.Marshal(matchDetail)
	println(string(b))
}
