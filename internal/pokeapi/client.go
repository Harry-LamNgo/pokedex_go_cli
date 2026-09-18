package pokeapi

import (
	"net/http"
	"time"

	"github.com/Harry-LamNgo/pokdex_go_cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(intervalCache, timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(intervalCache),
	}
}
