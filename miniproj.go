package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Id                float64 `json:"id"`
	Name              string  `json:"name"`
	Difficulty        float64 `json:"difficulty"`
	Exchange_rate_vol float64 `json:"exchange_rate_vol"`
	Timestamp         int     `json:"timestamp"`
}

func main() {

	message := "https://whattomine.com/coins/315.json"

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

	databaru := fmt.Sprintf("%v", data1)
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", databaru)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
