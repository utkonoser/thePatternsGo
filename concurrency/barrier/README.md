### Barrier — ожидание всех горутин и вывод одного общего результата

### Описание

Представьте ситуацию, когда у нас есть приложение микросервисов, в котором один микросервис должна составить свой ответ путем слияния ответов трех других микросервисов. Здесь и поможет паттерн Barrier.
Barrier может быть сервисом, который будет блокировать свой ответ до тех пор, пока он не будет составлен из результатов, возвращаемых одной или несколькими различными горутинами (или сервисами).
Как следует из названия, паттерн Barrier пытается остановить выполнение, чтобы оно не завершилось до тех пор, пока все нужные результаты не будут готовы.

### Пример — HTTP GET aggregator

Для примера мы собираемся написать очень типичную ситуацию в приложении микросервисов — приложении, выполняющего два вызова HTTP GET и объединяющего их в один ответ, который будет напечатан на консоли.
Наше небольшое приложение должно выполнять каждый запрос в горутине и выводить результат на консоль, если оба ответа верны. Если какой-либо из них возвращает ошибку, мы печатаем только ошибку.

Требования и критерии приемлемости:
* Выведите на консоль объединенный результат двух вызовов URL-адресов `http://httpbin.org/headers` и `http://httpbin.org/User-Agent`. Это пара общедоступных конечных точек, которые отвечают данными о входящих соединениях
* Если какой-либо из вызовов терпит неудачу, он не должен печатать никакого результата — только сообщение об ошибке (или сообщения об ошибках, если оба вызова не удались)
* Вывод должен быть напечатан как составной результат после завершения обоих вызовов. Это означает, что мы не можем вывести результат одного вызова, а затем другого

### Реализация
```go
package barrier

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

```

### Тесты

```go
package barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || strings.Contains(result, "User-Agent") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out
	return temp
}

```