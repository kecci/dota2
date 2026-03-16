package dotalive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.steampowered.com/IDOTA2Match_570/GetLiveLeagueGames/v1"

// LiveMatch represents the top-level structure of the Valve API response
type LiveMatchResponse struct {
	Result struct {
		Games []Game `json:"games"`
	} `json:"result"`
}

type Game struct {
	LobbyID     uint64   `json:"lobby_id"`
	MatchID     uint64   `json:"match_id"`
	Spectators  int      `json:"spectators"`
	LeagueID    int      `json:"league_id"`
	StreamDelay int      `json:"stream_delay_s"`
	RadiantTeam TeamInfo `json:"radiant_team"`
	DireTeam    TeamInfo `json:"dire_team"`
	Players     []Player `json:"players"`
	Scoreboard  Score    `json:"scoreboard"`
}

func (g Game) GetTeam(t int) TeamInfo {
	switch t {
	case 1:
		return g.DireTeam
	default:
		return g.RadiantTeam
	}
}

func (g Game) GetSide(t int) string {
	switch t {
	case 1:
		return "Dire"
	default:
		return "Radiant"
	}
}

type TeamInfo struct {
	TeamName string `json:"team_name"`
	TeamLogo uint64 `json:"team_logo"`
}

type Player struct {
	AccountID int    `json:"account_id"`
	Name      string `json:"name"`
	HeroID    int    `json:"hero_id"`
	Team      int    `json:"team"` // 0 for Radiant, 1 for Dire
}

type Score struct {
	Duration           float64   `json:"duration"`
	RoshanRespawnTimer int       `json:"roshan_respawn_timer"`
	Radiant            TeamScore `json:"radiant"`
	Dire               TeamScore `json:"dire"`
}

type ScoreboardPlayer struct {
	AccountID  int    `json:"account_id"`
	HeroID     int    `json:"hero_id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Gold       int    `json:"gold"`         // Current reliable + unreliable gold
	NetWorth   int    `json:"net_worth"`    // Total value (Gold + Items)
	XPPerMin   int    `json:"xp_per_min"`   // XPM
	GoldPerMin int    `json:"gold_per_min"` // GPM
	LastHits   int    `json:"last_hits"`
	Denies     int    `json:"denies"`
	Kills      int    `json:"kills"`
	Deaths     int    `json:"deaths"`
	Assists    int    `json:"assists"`
}

type PlayerStatus struct {
	Player
	ScoreboardPlayer
}

type TeamScore struct {
	Score         int                `json:"score"`
	TowerState    uint32             `json:"tower_state"`
	BarracksState uint32             `json:"barracks_state"`
	Players       []ScoreboardPlayer `json:"players"` // Nested player stats
}

func (ts TeamScore) GetScorePlayer(accountID int, heroID int) ScoreboardPlayer {
	for _, player := range ts.Players {
		if player.AccountID == accountID && player.HeroID == heroID {
			return player
		}
	}
	return ScoreboardPlayer{}
}

type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Dota 2 Live Match client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetLiveMatches fetches all currently active league matches
func (c *Client) GetLiveMatches() ([]Game, error) {
	url := fmt.Sprintf("%s?key=%s", baseURL, c.apiKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to reach Steam API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("steam API returned status: %d", resp.StatusCode)
	}

	var data LiveMatchResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return data.Result.Games, nil
}
