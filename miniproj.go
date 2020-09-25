package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"time"
)

type Data struct {
<<<<<<< HEAD
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
=======
	ID         int     `gorm:"primaryKey"`
	Timestamp  int     `json:"updated"`
	Coin       string  `json:"coin"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Difficulty float64 `json:"difficulty"`
	Volume     float64 `json:"volume"`
>>>>>>> 7aac8c5843332315530602e71d08167303807860
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

	messages := []string{"https://api.minerstat.com/v2/coins?list=0xBTC", "https://api.minerstat.com/v2/coins?list=BSHA3"}

	for _, message := range messages {
		resp, _ := http.Get(message)

		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)

		if err1 != nil {
			fmt.Println(err1)
		}

		body1 := body

		//fmt.Println(body1)

		data1 := []Data{}

		err2 := json.Unmarshal(body1, &data1)
		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(data1[0])

		d = append(d, data1[0])
	}

	return d
}

func main() {

	dsn := "host=localhost user=postgres password=monadmonad dbname=coin port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
