package workersPool

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkersPool(t *testing.T) {
	bufferSize := 100
	dispatcher := NewDispatcher(bufferSize)
	workers := 3
	for i := 1; i <= workers; i++ {
		var w WorkerLauncher = &PrefixSuffixWorker{
			id:      i,
			prefixS: fmt.Sprintf("Worker id: %d -> ", i),
			suffixS: " World",
		}
		dispatcher.LaunchWorker(w)
	}
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()
}
