package dotalive

import "encoding/json"

type OpenDotaLiveMatch struct {
	MatchID       string `json:"match_id"`
	ServerSteamID string `json:"server_steam_id"`
	LobbyID       string `json:"lobby_id"`
	GameTime      int    `json:"game_time"`
	AverageMmr    int    `json:"average_mmr"`
	RadiantScore  int    `json:"radiant_score"`
	DireScore     int    `json:"dire_score"`
	Players       []struct {
		AccountID int    `json:"account_id"`
		HeroID    int    `json:"hero_id"`
		Name      string `json:"name"`
		Level     int    `json:"level"`
		Gold      int    `json:"gold"`
		NetWorth  int    `json:"net_worth"`
		Xp        int    `json:"xp"`
	} `json:"players"`
}

func (c *Client) GetOpenDotaLive() ([]OpenDotaLiveMatch, error) {
	resp, err := c.httpClient.Get("https://api.opendota.com/api/live")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var matches []OpenDotaLiveMatch
	err = json.NewDecoder(resp.Body).Decode(&matches)
	return matches, err
}
