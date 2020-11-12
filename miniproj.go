package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Data is the structure save the fetched data from the API and used also as a table in the data base
type Data struct {
	Id              float64 `json:"id"`
	Name            string  `json:"name"`
	Difficulty      float64 `json:"difficulty"`
	ExchangeRateVol float64 `json:"exchange_rate_vol"`
	Timestamp       int     `json:"timestamp"`
}

type DataDB struct {
	IDKey      int     `gorm:"primaryKey"`
	Timestamp  int     `json:"updated"`    //The UNIX timestamp of the last time the coin was updated
	Coin       string  `json:"coin"`       //Coin's ticker
	Name       string  `json:"name"`       //Coin's name
	Algorithm  string  `json:"algorithm"`  //Coin's algorithm
	Price      float64 `json:"price"`      //Coin's price in USD
	Difficulty float64 `json:"difficulty"` //Coin's difficulty
	Volume     float64 `json:"volume"`     //Coin's last 24h volume in USD

}

//Environment contains the information needed to connect to the database
type Environment struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

func getCoin() []Data {

	// write the api link to get the data, the output from minerstat is a slice
	messages := "https://api.minerstat.com/v2/coins?list=BSHA3,TRB,0xBTC,KDA,DGB"

	resp, _ := http.Get(messages)

	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil {
		fmt.Println(err1)
	}

	// prepare the data
	data := []Data{}

	// unmarshall the data
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		fmt.Println(err2)
	}

	return data
}

// this function will append unique elements of the slice into datas
// first loop : get 1 data to be compared
// second loop : compare that 1 data to the others ( make sure you dont compare it with the same data), if there is no other data that is the same ( data is unique ) save the data
// additionally, since we know DGA have many algo, just check that the non unique if they have Odocrypt and append it to datas

func filterCoin(d []Data) []Data {

	datas := []Data{}
	flagUnique := 0

	for i, d1 := range d {
		flagUnique = 1
		for j, d2 := range d {
			if i != j {
				if d1.Name == d2.Name {
					flagUnique = 0
				}
			}
		}
		if flagUnique == 1 {
			datas = append(datas, d1)
		}
	}

	return datas
}

func getDatabaseInfo() Environment {

	u := os.Getenv("DATABASE_URL")
	e := Environment{}

	p, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	host, port, _ := net.SplitHostPort(p.Host)
	e.Host = host                     //"localhost"
	e.User = p.User.Username()        //"postgres"
	e.Password, _ = p.User.Password() //"monadmonad"
	e.Dbname = p.Path                 //"coin"
	e.Port = port                     //"5432"

	return e
}

func main() {
	const AWS_KEY = "***REMOVED***"
	fmt.Println(AWS_KEY)

	// get information to be able to connect to the database
	// initialize the connection to the postgreSQL server
	db, err := gorm.Open(postgres.Open("dsn"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create the table with the same structure as Data type
	db.AutoMigrate(&Data{})

	// fetch the coin data from minerstat
	datas := getCoin()

	// filter the data to make sure that each data is unique
	datas = filterCoin(datas)

	// add the filtered data to the database
	db.Create(&datas)

}
