package g

import (
	"sync"
	"fmt"
)

type Stats struct {
	mu sync.Mutex
	counters map[string]int
}
  
// Snapshot returns the current stats.
func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()
  
	return s.counters
}
  
func main() {
	stats := &Stats{}
	// snapshot は mutex で守られない
	// レースコンディションが起きる
	snapshot := stats.Snapshot() // want "WARN"
	fmt.Println(snapshot)
}