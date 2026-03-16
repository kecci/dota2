package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/kecci/dota2/dotalive"
)

func main() {
	apiKey := os.Getenv("STEAM_API_KEY")
	client := dotalive.NewClient(apiKey)

	// Initialize and hydrate the hero cache
	mapper := dotalive.NewHeroMapper()
	if err := mapper.Update(apiKey, http.DefaultClient); err != nil {
		log.Printf("Warning: could not fetch hero list: %v", err)
	}

	// Fetch matches
	games, err := client.GetLiveMatches()
	if err != nil {
		panic(err)
	}

	// Sort by LobbyID
	slices.SortFunc(games, func(a, b dotalive.Game) int {
		if a.LobbyID > b.LobbyID {
			return 1
		}
		return -1
	})

	for _, game := range games {
		fmt.Println("==========================================")
		fmt.Printf("Match: %s vs %s\n", game.RadiantTeam.TeamName, game.DireTeam.TeamName)

		fmt.Printf("LobbyID: %v\n", game.LobbyID)
		fmt.Printf("MatchID: %v\n", game.MatchID)
		fmt.Printf("Spectators: %d\n", game.Spectators)
		fmt.Printf("LeagueID: %d\n", game.LeagueID)
		stramDelay := time.Duration(game.StreamDelay) * time.Second
		fmt.Printf("StreamDelay: %s\n", stramDelay.String())
		gameDuration := time.Duration(game.Scoreboard.Duration) * time.Second
		fmt.Printf("Game Duration: %v\n", gameDuration.String())
		roshanRespawnTimer := time.Duration(game.Scoreboard.RoshanRespawnTimer) * time.Second
		fmt.Printf("RoshanRespawnTimer: %v\n", roshanRespawnTimer.String())
		fmt.Printf("Score: %d VS %d\n", game.Scoreboard.Radiant.Score, game.Scoreboard.Dire.Score)
		fmt.Printf("Radiant Barrack %v & Tower %v | Dire Barrack %v & Tower %v\n", game.Scoreboard.Radiant.BarracksState, game.Scoreboard.Radiant.TowerState, game.Scoreboard.Dire.BarracksState, game.Scoreboard.Dire.TowerState)

		teamPlayers := map[int][]dotalive.Player{}
		for _, p := range game.Players {
			if p.HeroID != 0 {
				teamPlayers[p.Team] = append(teamPlayers[p.Team], p)
			}
		}

		for team := 0; team < 2; team++ {
			fmt.Printf("[%s][%s]\n", game.GetSide(team), game.GetTeam(team).TeamName)

			scoreboard := game.Scoreboard.Radiant
			if team == 1 {
				scoreboard = game.Scoreboard.Dire
			}

			netSorted := []dotalive.PlayerStatus{}
			for _, p := range teamPlayers[team] {
				// Get Score Player
				scorePlayer := scoreboard.GetScorePlayer(p.AccountID, p.HeroID)
				netSorted = append(netSorted, dotalive.PlayerStatus{
					Player:           p,
					ScoreboardPlayer: scorePlayer,
				})
			}

			slices.SortFunc(netSorted, func(a, b dotalive.PlayerStatus) int {
				if a.ScoreboardPlayer.NetWorth < b.ScoreboardPlayer.NetWorth {
					return 1
				}
				return -1
			})

			for _, status := range netSorted {
				// 3. Map the ID to a readable name
				heroName := mapper.GetName(status.Player.HeroID)
				fmt.Printf("  %s (%s) \t|\t Lvl%d Net%d KDA=%d/%d/%d\n", status.Player.Name, heroName, status.ScoreboardPlayer.Level, status.ScoreboardPlayer.NetWorth, status.ScoreboardPlayer.Kills, status.ScoreboardPlayer.Deaths, status.ScoreboardPlayer.Assists)
			}

		}

		fmt.Println()
	}
}
