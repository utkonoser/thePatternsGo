### Proxy — оборачивание объекта для сокрытия характеристик

### Описание

Шаблон Proxy обычно оборачивает объект, чтобы скрыть некоторые его характеристики. Эти характеристики могут заключаться в том, что это удаленный объект (remote proxy), очень тяжелый объект, такой как дамп терабайтной базы данных (virtual proxy), или объект с ограниченным доступом (protection proxy).

Возможностей паттерна Proxy много, но в целом все они пытаются обеспечить одни и те же следующие функции:
* Скрыть объект за прокси-сервером для того, чтобы возможные функции можно было скрыть или ограничить
* Обеспечить новый уровень абстракции, с которым легко работать и можно легко изменить

### Пример

В примере мы собираемся создать удаленный прокси, который будет кэшировать объекты перед доступом к базе данных. Давайте представим, что у нас есть база данных со многими пользователями, но вместо того, чтобы обращаться к базе данных каждый раз, когда нам нужна информация о пользователе, у нас будет стек пользователей в порядке поступления (FIFO) в шаблоне Proxy.

Требования и критерии приемлемости:
* Весь доступ к базе данных пользователей будет осуществляться через тип Proxy
* Стек из n последних пользователей будет храниться в Proxy
* Если пользователь уже существует в стеке, запроса в базу данных не будет, вернется кешируемое значение
* Если запрошенный пользователь не существует в стеке, будет сделан запрос в базу данных, если стек полон, то удалим самого старого пользователя в стеке, далее сохраним нового пользователя и вернем его

### Реализация
```go
package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found\n", id)
}

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}
	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)
	fmt.Println("returning user from database")
	u.DidLastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

```

### Тесты

```go
package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCache:    UserList{},
		StackCapacity: 2,
	}

	knownIDs := [3]int32{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("after one successful search empty cache, the size of it must be one")
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("no user can be returned from empty cache")
		}
	})

	t.Run("FindUser - one user, ask fo the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("cache must not grow if we asked for an object that is stored on it")
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("the user should have been returned from the cache")
		}
	})

	user1, err := proxy.FindUser(knownIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	user2, _ := proxy.FindUser(knownIDs[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("the user wasn't stored on the proxy cache yet")
	}

	user3, _ := proxy.FindUser(knownIDs[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("the user wasn't stored on the proxy cache yet")
	}

	for i := 0; i < len(proxy.StackCache); i++ {
		if proxy.StackCache[i].ID == user1.ID {
			t.Error("user that should be gone was found")
		}
	}

	if len(proxy.StackCache) != 2 {
		t.Error("after inserting 3 users the cache should not grow more than to two")
	}

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Error("a non expected user was found on the cache")
		}
	}

}

```