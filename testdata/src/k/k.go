package i

import (
	"sync"
	"fmt"
)

type Stats struct {
	mu sync.Mutex
	values [2]string
}

func (s *Stats) Snapshot() [2]string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.values
}

func main() {
	stats := &Stats{}
	snapshot := stats.Snapshot()
	fmt.Println(snapshot)
} 