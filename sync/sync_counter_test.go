package sync_counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T){
	t.Run("Increment counter linearly", func(t *testing.T){
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})
	t.Run("Increment counter concurrently", func(t *testing.T){
		counter := NewCounter()
		wantedCount := 1000
		var wg sync.WaitGroup
		for range wantedCount {
			wg.Go(func(){
				counter.Inc()
			})
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, val int){
	t.Helper()
	if counter.Value() != val {
		t.Errorf("got %d want %d", counter.Value(), val)
	}
}