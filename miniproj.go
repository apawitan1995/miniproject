package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"time"
)

type Data struct {
	Id              float64 `json:"id"`
	Name            string  `json:"name"`
	Difficulty      float64 `json:"difficulty"`
	ExchangeRateVol float64 `json:"exchange_rate_vol"`
	Timestamp       int     `json:"timestamp"`
}

type DataDB struct {
	Id              float64
	Name            string
	Difficulty      float64
	ExchangeRateVol float64
	Timestamp       int
}

func doEvery(d time.Duration, f func() []Data) []Data {
	data := []Data{}
	for _ = range time.Tick(d) {
		data = f()
	}

	return data
}

func printCoin() []Data {
	d := []Data{}

	messages := []string{"https://whattomine.com/coins/315.json", "https://whattomine.com/coins/334.json"}

	for _, message := range messages {
		resp, _ := http.Get(message)

		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)

		if err1 != nil {
			fmt.Println(err1)
		}

		body1 := body

		data1 := Data{}

		err2 := json.Unmarshal(body1, &data1)

		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(data1)

		d = append(d, data1)
	}

	return d
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&DataDB{})

	datas := []Data{}

	datas = printCoin()

	dataprint := fmt.Sprintf("%v", datas)

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", dataprint)
	})

	// Create
	db.Create(&DataDB{Id: datas[0].Id, Name: datas[0].Name, Difficulty: datas[0].Difficulty, ExchangeRateVol: datas[0].ExchangeRateVol, Timestamp: datas[0].Timestamp})

	log.Fatal(http.ListenAndServe(":8081", nil))


}
