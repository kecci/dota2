# Dota 2 Live Match Tracker (Go)

A high-performance Go package for tracking live professional Dota 2 matches. This package provides a thread-safe hero mapper and a Redis-backed caching service to ensure minimal latency and protection against Steam API rate limits.

**Rate Limiting:** You only call Steam  **1,440 times/day** , well within the standard 100,000 limit, regardless of how many users visit your site.

## 💡 Features

* **Steam Web API Integration** : Direct connection to `GetLiveLeagueGames`.
* **Intelligent Hero Mapping** : Automatic hydration of hero IDs to human-readable names with `npc_dota_hero_` prefix cleaning.
* **Performance First** :
* **In-Memory Cache** : Thread-safe (RWMutex) hero lookup.
* **Redis Provider** : Background worker pattern to separate API polling from request handling.
* **Production Ready** : Built-in support for context-based cancellation and graceful worker shutdowns.

---

## 🚀 Getting Started

## Installation

**Bash**

```
go get github.com/your-repo/dotalive
```

## Basic Usage (API Client)

Initialize the client with your Steam Web API Key to fetch raw match data.

**Go**

```
client := dotalive.NewClient("YOUR_STEAM_API_KEY")

// Fetch all live professional matches
games, err := client.GetLiveMatches()
if err != nil {
    log.Fatal(err)
}
```

## Advanced Usage (Worker + Redis Cache)

For high-load environments (like financial or trading dashboards), use the `CacheService`. This prevents blocking your API on Steam's latency.

**Go**

```
// 1. Setup Hero Mapper
mapper := dotalive.NewHeroMapper()
mapper.Update(apiKey, http.DefaultClient)

// 2. Setup Redis and Cache Service
rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
cacheService := dotalive.NewCacheService(rdb, dotaClient)

// 3. Start the Background Producer (Polling)
go cacheService.StartWorker(ctx, 1*time.Minute)

// 4. Consume data instantly in your Handlers
matches, err := cacheService.GetCachedMatches(ctx)
```

---

## 🛠 Project Structure

| **File** | **Responsibility**                                     |
| -------------- | ------------------------------------------------------------ |
| `client.go`  | Core Steam API interactions and Struct definitions.          |
| `heroes.go`  | Hero ID to Name mapping and string sanitization.             |
| `cache.go`   | Redis worker logic and SOTW (State of the World) management. |

---

## 🔧 Technical Specifications

## Hero Name Sanitization

The mapper automatically cleans internal Valve strings:

* `npc_dota_hero_antimage` → `antimage`
* `npc_dota_hero_nevermore` → `nevermore`

## Performance Metrics

* [ ] **Steam API Latency** : 200ms – 800ms (Varies by Valve's server load).
* [ ] **Redis Cache Latency** : <5ms (O(1) lookup).
* [ ] **Memory Footprint** : Minimal (Hero map is <200 entries).

## Example Response

```
  [L1GA TEAM] nobody (unknown)
  [L1GA TEAM] Robbnroll (unknown)
  [Pipsqueak+4] Alle (kez)
  [L1GA TEAM] Mormagni (unknown)
  [L1GA TEAM] O'Block (treant)
  [L1GA TEAM] aadaam.musiaal (unknown)
  [Pipsqueak+4] sagyndym (largo)
  [Pipsqueak+4] 3.7 (tiny)
  [L1GA TEAM] player (shadow_demon)
  [L1GA TEAM] hairy_freak (unknown)
  [Pipsqueak+4] delusion (shadow_shaman)
  [Pipsqueak+4] ряф^^ (queenofpain)
  [L1GA TEAM] Layton (storm_spirit)
  [L1GA TEAM] kUBAJlDA (centaur)
  [L1GA TEAM] kill them all (life_stealer)
```

---

## 📝 License

MIT License.
