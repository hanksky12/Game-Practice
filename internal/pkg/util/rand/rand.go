package rand

import (
	"math/rand"
	"sync"
)

type SafeRand struct {
	Rand *rand.Rand
	mu   sync.Mutex
}

func (sr *SafeRand) Intn(n int) int {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	return sr.Rand.Intn(n)
}
