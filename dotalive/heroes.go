package dotalive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const heroesURL = "https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v1"

type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // Internal name: e.g., "npc_dota_hero_antimage"
}

type HeroResponse struct {
	Result struct {
		Heroes []Hero `json:"heroes"`
	} `json:"result"`
}

type HeroMapper struct {
	mu    sync.RWMutex
	cache map[int]string
}

func NewHeroMapper() *HeroMapper {
	return &HeroMapper{
		cache: make(map[int]string),
	}
}

// Update fetches the latest hero list from Steam
func (m *HeroMapper) Update(apiKey string, httpClient *http.Client) error {
	url := fmt.Sprintf("%s?key=%s&language=en_us", heroesURL, apiKey)

	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data HeroResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	for _, h := range data.Result.Heroes {
		m.cache[h.ID] = h.Name
	}
	return nil
}

func (m *HeroMapper) GetName(id int) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if name, ok := m.cache[id]; ok {
		// Cleans "npc_dota_hero_axe" -> "axe"
		clean := strings.TrimPrefix(name, "npc_dota_hero_")
		replaceUnderscore := strings.ReplaceAll(clean, "_", " ")
		capitalize := cases.Title(language.English).String(replaceUnderscore)
		return capitalize
	}

	return "Unknown"
}
