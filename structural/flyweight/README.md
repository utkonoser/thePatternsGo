### Flyweight — паттерн Приспособленец

### Описание

Паттерн проектирования Flyweight. Он очень часто используется в компьютерной графике и индустрии видеоигр, но не так часто в корпоративных приложениях.
Flyweight — это паттерн, который позволяет разделить состояние тяжелого объекта между многими экземплярами одного типа. Представьте, что вам нужно создать и хранить слишком много принципиально одинаковых объектов, занимающих довольно много памяти. У вас быстро закончится память. Эта проблема может быть легко решена с помощью паттерна Flyweight с дополнительной помощью паттерна Factory. Благодаря шаблону Flyweight мы можем разделить все возможные состояния объектов в одном общем объекте и, таким образом, свести к минимуму создание объектов, используя указатели на уже созданные объекты.

### Пример — букмекерская контора
Чтобы привести пример, мы собираемся смоделировать то, что можно найти на странице букмекерской конторы. Представьте себе финальный матч чемпионата Европы, который смотрят миллионы людей по всему континенту. Теперь представьте, что у нас есть веб-страница для ставок, на которой мы публикуем историческую информацию о каждой команде в Европе. Это очень много информации, которая обычно хранится в какой-то распределенной базе данных, и у каждой команды буквально мегабайты информации о своих игроках, матчах, чемпионатах и так далее. Если миллион пользователей получат доступ к информации о команде и для каждого пользователя, запрашивающего исторические данные, будет создан новый экземпляр информации, у нас в мгновение ока закончится память. Можно попробовать Proxy паттерн, где можно кэшировать n самых последних поисков, чтобы ускорить запросы, но если мы будем возвращать клон для каждой команды, нам все равно будет не хватать памяти.
Вместо этого мы будем хранить информацию о каждой команде только один раз, и будем предоставлять ссылки на них пользователям. Таким образом, если мы столкнемся с миллионом пользователей, пытающихся получить доступ к информации о матче, у нас фактически будет просто две команды в памяти с миллионом указателей на одно и то же направление памяти.

Критерии приемлемости паттерна Flyweight всегда должны уменьшать объем используемой памяти и должны быть сосредоточены в первую очередь на этой цели:
*  Мы создадим структуру Team с некоторой базовой информацией, такой как название команды, игроки, прошлые результаты матчей, изображение их эмблемы
* Мы должны обеспечить правильное создание команды и отсутствие дубликатов
* При создании одной и той же команды дважды у нас должно быть два указателя, указывающих на один и тот же адрес памяти

### Реализация
```go
package flyweight

import "time"

type Team struct {
	Id             uint64
	Name           int
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Lear          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			Id:   2,
			Name: TEAM_B,
		}
	default:
		return Team{
			Id:   1,
			Name: TEAM_A,
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{createdTeams: make(map[int]*Team)}
}

```
### Тесты
```go
package flyweight

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("the pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}
	if teamA1 != teamA2 {
		t.Error("TEAM_A pointers weren't the same")
	}
	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()
	teams := make([]*Team, 500000*2)
	for i := 0; i < 500000; i++ {
		teams[i] = factory.GetTeam(TEAM_A)
	}
	for i := 500000; i < 2*500000; i++ {
		teams[i] = factory.GetTeam(TEAM_B)
	}
	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}

```