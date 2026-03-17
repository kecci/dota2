package opendota

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

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

type MatchDetail struct {
	MatchID               int64       `json:"match_id"`
	Players               []Players   `json:"players"`
	SeriesID              int         `json:"series_id"`
	SeriesType            int         `json:"series_type"`
	Cluster               int         `json:"cluster"`
	ReplaySalt            int         `json:"replay_salt"`
	RadiantWin            bool        `json:"radiant_win"`
	Duration              int         `json:"duration"`
	PreGameDuration       int         `json:"pre_game_duration"`
	StartTime             int         `json:"start_time"`
	MatchSeqNum           int64       `json:"match_seq_num"`
	TowerStatusRadiant    int         `json:"tower_status_radiant"`
	TowerStatusDire       int         `json:"tower_status_dire"`
	BarracksStatusRadiant int         `json:"barracks_status_radiant"`
	BarracksStatusDire    int         `json:"barracks_status_dire"`
	FirstBloodTime        int         `json:"first_blood_time"`
	LobbyType             int         `json:"lobby_type"`
	HumanPlayers          int         `json:"human_players"`
	Leagueid              int         `json:"leagueid"`
	GameMode              int         `json:"game_mode"`
	Flags                 int         `json:"flags"`
	Engine                int         `json:"engine"`
	RadiantScore          int         `json:"radiant_score"`
	DireScore             int         `json:"dire_score"`
	RadiantTeamID         int         `json:"radiant_team_id"`
	RadiantName           string      `json:"radiant_name"`
	RadiantLogo           int64       `json:"radiant_logo"`
	RadiantTeamComplete   int         `json:"radiant_team_complete"`
	DireTeamID            int         `json:"dire_team_id"`
	DireName              string      `json:"dire_name"`
	DireLogo              int64       `json:"dire_logo"`
	DireTeamComplete      int         `json:"dire_team_complete"`
	RadiantCaptain        int         `json:"radiant_captain"`
	DireCaptain           int         `json:"dire_captain"`
	PicksBans             []PicksBans `json:"picks_bans"`
	OdData                OdData      `json:"od_data"`
	League                League      `json:"league"`
	RadiantTeam           RadiantTeam `json:"radiant_team"`
	DireTeam              DireTeam    `json:"dire_team"`
	Metadata              any         `json:"metadata"`
	ReplayURL             string      `json:"replay_url"`
	Patch                 int         `json:"patch"`
	Region                int         `json:"region"`
}
type PermanentBuffs struct {
	PermanentBuff int `json:"permanent_buff"`
	StackCount    int `json:"stack_count"`
	GrantTime     int `json:"grant_time"`
}
type GoldPerMin struct {
	Raw int     `json:"raw"`
	Pct float64 `json:"pct"`
}
type XpPerMin struct {
	Raw int     `json:"raw"`
	Pct float64 `json:"pct"`
}
type KillsPerMin struct {
	Raw float64 `json:"raw"`
	Pct float64 `json:"pct"`
}
type LastHitsPerMin struct {
	Raw float64 `json:"raw"`
	Pct float64 `json:"pct"`
}
type HeroDamagePerMin struct {
	Raw float64 `json:"raw"`
	Pct float64 `json:"pct"`
}
type HeroHealingPerMin struct {
	Raw int     `json:"raw"`
	Pct float64 `json:"pct"`
}
type TowerDamage struct {
	Raw int     `json:"raw"`
	Pct float64 `json:"pct"`
}
type Benchmarks struct {
	GoldPerMin        GoldPerMin        `json:"gold_per_min"`
	XpPerMin          XpPerMin          `json:"xp_per_min"`
	KillsPerMin       KillsPerMin       `json:"kills_per_min"`
	LastHitsPerMin    LastHitsPerMin    `json:"last_hits_per_min"`
	HeroDamagePerMin  HeroDamagePerMin  `json:"hero_damage_per_min"`
	HeroHealingPerMin HeroHealingPerMin `json:"hero_healing_per_min"`
	TowerDamage       TowerDamage       `json:"tower_damage"`
}
type Players struct {
	AccountID          int              `json:"account_id"`
	PlayerSlot         int              `json:"player_slot"`
	PartyID            int              `json:"party_id"`
	PermanentBuffs     []PermanentBuffs `json:"permanent_buffs"`
	PartySize          int              `json:"party_size"`
	TeamNumber         int              `json:"team_number"`
	TeamSlot           int              `json:"team_slot"`
	HeroID             int              `json:"hero_id"`
	HeroVariant        int              `json:"hero_variant"`
	Item0              int              `json:"item_0"`
	Item1              int              `json:"item_1"`
	Item2              int              `json:"item_2"`
	Item3              int              `json:"item_3"`
	Item4              int              `json:"item_4"`
	Item5              int              `json:"item_5"`
	Backpack0          int              `json:"backpack_0"`
	Backpack1          int              `json:"backpack_1"`
	Backpack2          int              `json:"backpack_2"`
	ItemNeutral        int              `json:"item_neutral"`
	ItemNeutral2       int              `json:"item_neutral2"`
	Kills              int              `json:"kills"`
	Deaths             int              `json:"deaths"`
	Assists            int              `json:"assists"`
	LeaverStatus       int              `json:"leaver_status"`
	LastHits           int              `json:"last_hits"`
	Denies             int              `json:"denies"`
	GoldPerMin         int              `json:"gold_per_min"`
	XpPerMin           int              `json:"xp_per_min"`
	Level              int              `json:"level"`
	NetWorth           int              `json:"net_worth"`
	AghanimsScepter    int              `json:"aghanims_scepter"`
	AghanimsShard      int              `json:"aghanims_shard"`
	Moonshard          int              `json:"moonshard"`
	HeroDamage         int              `json:"hero_damage"`
	TowerDamage        int              `json:"tower_damage"`
	HeroHealing        int              `json:"hero_healing"`
	Gold               int              `json:"gold"`
	GoldSpent          int              `json:"gold_spent"`
	AbilityUpgradesArr []int            `json:"ability_upgrades_arr"`
	Personaname        string           `json:"personaname"`
	Name               any              `json:"name"`
	LastLogin          time.Time        `json:"last_login"`
	RankTier           int              `json:"rank_tier"`
	ComputedMmr        any              `json:"computed_mmr"`
	IsSubscriber       bool             `json:"is_subscriber"`
	RadiantWin         bool             `json:"radiant_win"`
	StartTime          int              `json:"start_time"`
	Duration           int              `json:"duration"`
	Cluster            int              `json:"cluster"`
	LobbyType          int              `json:"lobby_type"`
	GameMode           int              `json:"game_mode"`
	IsContributor      bool             `json:"is_contributor"`
	Patch              int              `json:"patch"`
	Region             int              `json:"region"`
	IsRadiant          bool             `json:"isRadiant"`
	Win                int              `json:"win"`
	Lose               int              `json:"lose"`
	TotalGold          int              `json:"total_gold"`
	TotalXp            int              `json:"total_xp"`
	KillsPerMin        float64          `json:"kills_per_min"`
	Kda                int              `json:"kda"`
	Abandons           int              `json:"abandons"`
	Benchmarks         Benchmarks       `json:"benchmarks"`
}
type PicksBans struct {
	IsPick bool `json:"is_pick"`
	HeroID int  `json:"hero_id"`
	Team   int  `json:"team"`
	Order  int  `json:"order"`
}
type OdData struct {
	HasAPI     bool `json:"has_api"`
	HasGcdata  bool `json:"has_gcdata"`
	HasParsed  bool `json:"has_parsed"`
	HasArchive bool `json:"has_archive"`
}
type League struct {
	Leagueid int    `json:"leagueid"`
	Ticket   any    `json:"ticket"`
	Banner   any    `json:"banner"`
	Tier     string `json:"tier"`
	Name     string `json:"name"`
}
type RadiantTeam struct {
	TeamID  int    `json:"team_id"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	LogoURL string `json:"logo_url"`
}
type DireTeam struct {
	TeamID  int    `json:"team_id"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	LogoURL any    `json:"logo_url"`
}

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

// 8732939854
func (c *Client) GetMatchDetail(matchID int64) (MatchDetail, error) {
	md := MatchDetail{}
	req, err := http.NewRequest("GET", "https://api.opendota.com/api/matches/"+strconv.Itoa(int(matchID)), nil)
	if err != nil {
		return md, err
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Cookie", "session=e30=; session.sig=9A3MANd7pSadSYrfALkOJK-1IPc")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return md, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return md, err
	}
	_ = json.Unmarshal(bodyText, &md)
	return md, nil
}
