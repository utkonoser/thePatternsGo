### Concurrent Singleton — используя мьютексы и каналы

### Описание

В Creational паттернах есть паттерн Singleton — это некая структура или переменная, которая существует в коде только один раз. Весь доступ к этой структуре должен осуществляться с использованием описанного паттерна, но на самом деле он не безопасен с параллельной точки зрения.
Concurrent Singleton будет описан с учетом параллелизма.

### Пример — уникальный счетчик с помощью каналов и мьютексов
Чтобы ограничить одновременный доступ к экземпляру Singleton, только одна горутина сможет получить к нему доступ. Мы получим доступ к нему, используя каналы — первый для добавления единицы в счетчик, второй для получения текущего счетчика и третий для остановки горутины.
Мы добавим единицу в счетчик 10 000 раз, используя 10 000 различных горутин, запущенных из двух разных экземпляров Singleton. Затем мы введем цикл для проверки количества Singleton до тех пор, пока оно не станет равным 5000, и напишем значение счетчика перед запуском цикла.
Как только счетчик достигнет 5000, цикл завершится и закроет запущенную горутину.

### Реализация с помощью каналов

```go
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
```

### Реализация с помощью мьютексов
```go
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
```

### Тесты
```go
package concurrentSingleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()
}

func TestStartInstanceMutex(t *testing.T) {
	singleton := GetInstance2()
	singleton2 := GetInstance2()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
}
```