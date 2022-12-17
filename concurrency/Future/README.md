### Future — реализация принципа 'fire-and-forget'

### Описание

Паттерн проектирования Future (также называемый Promise) — это быстрый и простой способ создания конкурентных структур для асинхронного программирования. Идея здесь состоит в том, чтобы реализовать принцип «fire-and-forget», который обрабатывает все возможные результаты действия.
Короче говоря, мы определим каждое возможное поведение действия перед его выполнением в разных горутинах. Здесь интересно то, что мы можем запустить новый Future внутри Future и встроить столько Future, сколько захотим, в одну и ту же горутину (или в новые).

С помощью паттерна Future мы можем запускать множество новых горутин, каждая из которых имеет действие и собственный обработчик. Это позволяет нам сделать следующее:
* Делегировать обработчик действий другой горутине
* Стекировать между собой множество асинхронных вызовов (асинхронный вызов, который в своих результатах вызывает другой асинхронный вызов)

### Пример — простой асинхронный requester
В этом примере у нас будет метод, который возвращает строку или ошибку, но мы хотим выполнить все конкурентно. Используя канал, мы можем запустить новую горутину и обработать входящий результат из канала.
Но в этом случае нам придется обрабатывать результат (строку или ошибку), а этого мы не хотим. Вместо этого мы определим, что делать в случае успеха и что делать в случае ошибки.

Требования и критерии приемлемости:
* Делегировать выполнение функции другой горутине
* Функция вернет string (maybe) или error
* Обработчики должны быть уже определены до выполнения функции
* Дизайн должен быть многоразовым

### Реализация
```go
package Future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}
```

### Тесты
```go
package Future

import (
	"errors"
	"sync"
	"testing"
)

func TestStringOrError(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})
	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", errors.New("error occurred")
		})
		wg.Wait()
	})
}
```