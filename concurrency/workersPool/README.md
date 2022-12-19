### Workers pool — пул обработчиков

### Описание

Workers pool — это множество потоков, предназначенных для обработки назначаемых им заданий. Веб-сервер Apache и Go-пакет net/http работают приблизительно так: основной процесс принимает все входящие запросы, которые затем перенаправляются рабочим процессам для обработки. Как только рабочий процесс завершает свою работу, он готов к обслуживанию нового клиента.
Однако здесь есть главное различие: пул обработчиков использует не потоки,
а горутины. Кроме того, потоки обычно не умирают после обработки запросов, потому что затраты на завершение потока и создание нового слишком высоки, тогда как горутина прекращает существовать после завершения работы. Пулы обработчиков в Go реализованы с помощью буферизованных каналов, поскольку они позволяют ограничить число одновременно выполняемых горутин.

Создание Workers pool связано с управлением ресурсами: ЦП, ОЗУ, временем, соединениями и так далее. Паттерн проектирования Workers pool помогает сделать следующее:
* Контролировать доступ к общим ресурсам
* Создавать ограниченное количество горутин для каждого приложения
* Обеспечить больше возможностей параллелизма для других конкурентных структур

### Пример — pool of pipelines
В примере мы запустим ограниченное количество Pipeline, чтобы планировщик Go мог попытаться обрабатывать запросы параллельно.
Идея здесь состоит в том, чтобы контролировать количество горутин, изящно останавливать их, когда приложение завершает работу, и максимизировать параллелизм. В примере мы будем передавать строки, к которым будем добавлять данные и префиксы.

### Реализация
```go
package workersPool

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// workers pipeline

type Request struct {
	Data    any
	Handler RequestHandler
}

type RequestHandler func(any)

func NewStringRequest(data string, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: data, Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}

// worker

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type PrefixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercaseStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercaseStringWithSuffix))
		}
	}()
}

// dispatcher

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(r Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{inCh: make(chan Request, b)}
}
```

### Тесты
```go
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

```
