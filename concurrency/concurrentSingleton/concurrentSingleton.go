package concurrentSingleton

import "sync"

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

// with use mutexes

type singleton2 struct {
	count int
	sync.RWMutex
}

var instance2 singleton2

func GetInstance2() *singleton2 {
	return &instance2
}
func (s *singleton2) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}
func (s *singleton2) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
