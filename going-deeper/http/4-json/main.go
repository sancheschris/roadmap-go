package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"number"` //tags 
	Balance int `json:"balance"`
}


func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account) // when I want to control it, I use Marshal
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account) // when I want to store directly into a webservice, file and so on.
	if err != nil {
		println(err)
	}

	jsonRaw := []byte(`{"n":2, "s":200}`)
	var accountX Account
	err = json.Unmarshal(jsonRaw, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)

}