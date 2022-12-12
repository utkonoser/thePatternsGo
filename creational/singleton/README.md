### Singleton — наличие уникального экземпляра типа во всей программе

### Описание

Шаблон Singleton легко запомнить. Как следует из названия, он предоставляет единственный экземпляр объекта и гарантирует отсутствие дубликатов.
При первом вызове экземпляра он создается, а затем повторно используется всеми частями в приложении, которое должно использовать это конкретное поведение.
Шаблон Singleton используется во многих различных ситуациях. Например:
* Если вы хотите использовать одно и то же соединение с базой данных для выполнения каждого запроса.
* Когда вы открываете соединение Secure Shell (SSH) с сервером для выполнения нескольких задач, и не хотите заново открывать соединение для каждой задачи.
* Если вам нужно ограничить доступ к какой-либо переменной или пространству, вы используете Singleton как дверь к этой переменной.


### Пример — уникальный счетчик
В качестве примера объекта, будет уникальный счетчик, для которого мы должны убедиться, что существует только один такой экземпляр, счетчик будет сожержать количество вызовов во время исполнения программы. Неважно, сколько у нас экземпляров счетчика, все они будут считать одно и то же значение, и оно должно быть согласовано между экземплярами.

Требования и критерии приемлемости:
* Если счетчик ранее не создавался, создается новый со значением 0.
* Если счетчик уже создан, возвращается экземпляр, содержащий фактический
счетчик.
* Если мы вызываем метод `AddOne()`, счетчик должен быть увеличен на 1.

### Реализация

```go
package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
```

### Тесты

```go
package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("expected pointer to Singleton after " +
			"calling GetInstance(), not nil\n")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("after AddOne() the count must be" +
			" 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("expected same instance in counter2 but" +
			" it got a different instance\n")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("after AddOne() the count must be" +
			" 2 but it is %d\n", currentCount)
	}
}

```

