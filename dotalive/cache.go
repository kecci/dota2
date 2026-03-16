package dotalive

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

const RedisKeyLiveMatches = "dota:matches:live"

type CacheService struct {
	rdb    *redis.Client
	client *Client
}

func NewCacheService(rdb *redis.Client, client *Client) *CacheService {
	return &CacheService{rdb: rdb, client: client}
}

// StartWorker runs a background ticker to refresh the cache.
// This prevents your frontend from ever being blocked by Steam's latency.
func (s *CacheService) StartWorker(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			matches, err := s.client.GetLiveMatches()
			if err != nil {
				continue // In production, log this to Sentry/Grafana
			}

			data, _ := json.Marshal(matches)
			// Cache for slightly longer than the ticker to avoid gaps
			s.rdb.Set(ctx, RedisKeyLiveMatches, data, interval+time.Second*5)
		}
	}
}

// GetCachedMatches returns the latest data from Redis instantly
func (s *CacheService) GetCachedMatches(ctx context.Context) ([]Game, error) {
	val, err := s.rdb.Get(ctx, RedisKeyLiveMatches).Result()
	if err != nil {
		return nil, err
	}

	var games []Game
	err = json.Unmarshal([]byte(val), &games)
	return games, err
}
