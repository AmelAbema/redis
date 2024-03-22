package storage

import (
	"sync"
	"time"
)

type (
	App struct {
		TTL string `env:"TTL" envDefault:"10"`
	}
	Data struct {
		Value     interface{}
		ExpiresAt *time.Time
	}
	Storage struct {
		Cfg   *App
		Store map[string]Data
		Mutex sync.RWMutex
	}
)
