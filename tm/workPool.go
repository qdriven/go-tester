package tm

import (
	"log"
	"math/rand"
)

// workerResetThreshold defines how often the stack must be reset. Every N
// requests, by spawning a new Goroutine in its place, a worker can reset its
// stack so that large stacks don't live in memory forever. 2^16 should allow
// each Goroutine stack to live for at least a few seconds in a typical
// workload (assuming a QPS of a few thousand requests/sec).
const workerResetThreshold = 1 << 16

type WorkReq struct {
	// whatever
}

type server struct {
	workerChannels []chan *WorkReq
}

func (s *server) worker(ch chan *WorkReq) {
	// To make sure all server workers don't reset at the same time, choose a
	// random number of iterations before resetting.
	threshold := workerResetThreshold + rand.Intn(workerResetThreshold)
	for completed := 0; completed < threshold; completed++ {
		req, ok := <-ch
		if !ok {
			return
		}
		s.handleSingleRequest(req)
	}
	// Restart in a new Goroutine.
	go s.worker(ch)
}

func (s *server) initWorkers(numWorkers int) {
	s.workerChannels = make([]chan *WorkReq, numWorkers)
	for i := 0; i < numWorkers; i++ {
		// One channel per worker reduces contention.
		s.workerChannels[i] = make(chan *WorkReq)
		go s.worker(s.workerChannels[i])
	}
}

func (s *server) stopWorkers() {
	for _, ch := range s.workerChannels {
		close(ch)
	}
}

func (s *server) handleSingleRequest(req *WorkReq) {
	log.Printf("processing req=%v\n", req)
}

func (s *server) listenAndHandleForever() {
	for counter := 0; ; counter++ {
		req := listenForRequest()
		select {
		case s.workerChannels[counter%len(s.workerChannels)] <- req:
		default:
			// TODO: If this workers is busy, fall back to spawning a Goroutine. Or
			// find a different worker. Or dynamically increase the number of workers.
			// Or just reject the WorkReq.
		}
	}
}

func listenForRequest() *WorkReq {
	return &WorkReq{}
}

func newServer() *server {
	s := &server{}
	s.initWorkers(16)
	return s
}

func main() {
	s := newServer()
	s.listenAndHandleForever()
}
