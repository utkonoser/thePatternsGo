### Pipeline — конвейер передачи данных

### Описание

Pipeline — это паттерн, предназначенный для соединения горутин
и каналов, так что выходные данные одной горутины становятся входными данными для другой горутины, а для передачи данных используются каналы.
Одним из преимуществ использования конвейеров является наличие постоянного потока данных, так что никакие горутины и каналы не должны ожидать, пока завершится все остальное, чтобы можно было начать выполнение. Кроме того, мы используем меньше переменных и, следовательно, меньше памяти, потому что не приходится сохранять все данные в виде переменных. Наконец, использование конвейеров упрощает разработку программ и делает их удобнее для поддержки.

Возможности паттерна Pipeline:
* Можно создать параллельную структуру многошагового алгоритма
* Можно использовать параллелизм многоядерных машин, разложив алгоритм на разные горутины

### Пример — конвейер математических операций

 Мы собираемся сгенерировать список чисел, начиная с 1 и заканчивая некоторым произвольным числом N. Затем мы возьмем каждое число, возведем его в степень 2 и суммируем полученные числа с уникальным результатом. Итак, если `N = 3`, наш список будет `[1,2,3]`. После включения их в 2 наш список становится `[1,4,9]`. Если мы суммируем полученный список, результирующее значение равно 14.

План реализации:
* Нужно создать список от 1 до N, где N может быть любым целым числом
* Взять каждое число из этого сгенерированного списка и возвести его в степень 2
* Суммировать каждое полученное число в окончательный результат и вернуть его

### Реализация
```go
package pipeline

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)
	result := <-thirdCh
	return result
}

// LaunchPipeline function doesn't need to allocate every channel, and can be rewritten like this:

func LaunchPipelineSecondVar(amount int) int {
	return <-sum(power(generator(amount)))
}
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()
	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

```
### Тесты
```go
package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}

func TestLaunchPipelineSecondVar(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipelineSecondVar(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}

```