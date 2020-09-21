package main

//this is a test

// I am editing something

// ni coba test lagi ya

// ni coba test lagi  makin banyak

// ni coba test lagi  makin banyak biar konflik

// ngedit lagi buat di pull

// edit lagi buat tes di vscode

// edit lagi buat di github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"Result"`
}

func main() {

	apietherscan := "RBFCU93DGD1CGQMK5MHD7XAWXW1RT1QNZ7"

	wallet := "0x63FDB5011c66C686f88E836c38C874142014f0A2"

	message := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%v&tag=latest&apikey=%v", wallet, apietherscan)

	fmt.Println(message)

	resp, _ := http.Get(message)

	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(body)

	body1 := body

	data1 := Data{}

	err2 := json.Unmarshal(body1, &data1)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(data1)

}
