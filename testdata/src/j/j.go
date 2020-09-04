package h

import (
	"sync"
	"fmt"
)

type Stats struct {
	mu sync.Mutex
	values []string
}

// Snapshot returns the current stats.
func (s *Stats) Snapshot() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.values
}

func main() {
	stats := &Stats{}
	// snapshot は mutex で守られない
	// レースコンディションが起きる
	snapshot := stats.Snapshot() // want "WARN: Slices or maps that are kept internally without being made public may be changed."
	fmt.Println(snapshot)
} 