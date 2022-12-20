### Facade — создание библиотеки

### Описание

Если прочитать про паттерн Proxy, то можно узнать, что это способ обернуть тип, чтобы скрыть от пользователя некоторые его сложные особенности. А если представить, что можно сгруппировать множество прокси в одной точке, например, в файле или библиотеке. Это может быть паттерн Facade. Facade в архитектурном смысле – это фасадная стена, скрывающая помещения и коридоры здания. Он защищает своих обитателей от холода и дождя и обеспечивает им уединение.
Шаблон проектирования Facade делает то же самое, но в нашем коде. Он защищает код от нежелательного доступа, упорядочивает некоторые вызовы и скрывает сложные области от пользователя.

Самый яркий пример паттерна Facade — это библиотека, где кто-то должен предоставить разработчику некоторые методы для выполнения определенных действий в дружественной манере. Таким образом, если разработчику нужно использовать вашу библиотеку, ему не нужно знать все внутренние задачи, чтобы получить желаемый результат.

Итак, вы паттерн проектирования Facade используется в следующих сценариях:
* Если нужно уменьшить сложность некоторых частей кода, то можно эту сложность скрыть с помощью Facade, предоставляя более простой в использовании метод
* Если нужно сгруппировать взаимосвязанные действия в одном месте
* Когда нужно создать библиотеку, чтобы другие могли использовать ваши продукты, не беспокоясь о том, как все это работает

### Пример — HTTP REST API для OpenWeatherMaps

В качестве примера попробуем написать часть библиотеки, которая обращается к сервису OpenWeatherMaps. Если вы не знакомы с сервисом OpenWeatherMap, это HTTP-сервис, который предоставляет вам оперативную информацию о погоде. HTTP REST API очень прост в использовании и станет хорошим примером того, как реализовать паттерн Facade для сокрытия сложности сетевых подключений за службой REST.

API OpenWeatherMap предоставляет много информации, поэтому мы сосредоточимся на получении данных о погоде в реальном времени в одном городе в каком-либо географическом месте, используя его значения широты и долготы. Ниже приведены требования и критерии приемлемости для этого шаблона проектирования:
* Нужно реализовать единый тип для доступа к данным. Вся информация, полученная из сервиса OpenWeatherMap, будет проходить через него
* Создать способ получения данных о погоде для какого-либо города какой-либо страны
* Создать способ получения данных о погоде для определенной широты и долготы
* Снаружи пакета должно быть видно не все, только самое важное, остальное должно быть скрыто (включая все данные, связанные с соединением)

### Реализация
```go
package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

type CurrentWeatherData struct {
	APIkey string
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v,%v&APPID=%s", lat, lon, c.APIkey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&APPID=%s", city, countryCode, c.APIkey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := io.ReadAll(resp.Body)

		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n", resp.StatusCode, errMsg)
		return
	}
	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()
	return
}

// getMockData - mock data for our example
func getMockData() io.Reader {
	response := `{
"coord":{"lon":-3.7,"lat":40.42},"weather":
[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],
"base":"stations","main":{"temp":303.56,"pressure":1016.46,
"humidity":26.8,"temp_min":300.95,"temp_max":305.93},"wind":{"speed":3.17,"deg":151.001},
"rain":{"3h":0.0075},"clouds":{"all":68},"dt":1471295823,"sys":{"type":3,"id":1442829648,
"message":0.0278,"country":"ES","sunrise":1471238808,"sunset":1471288232},"id":3117735,
"name":"Madrid","cod":200}`
	r := bytes.NewReader([]byte(response))
	return r
}

// Weather struct from http://openweathermap.org/current#current_JSON.
type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}
```

### Тесты

```go
package facade

import (
	"fmt"
	"testing"
)

// test with mock data
func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}

// if there is api then use this test
func TestWithApi(t *testing.T) {
	weatherMap := CurrentWeatherData{"*Apikey"}
	weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Temperature in Madrid is %f celsius\n", weather.Main.Temp-273.15)
}

```