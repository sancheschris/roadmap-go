package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to make requests: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error to read data: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error to parse data: %v\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprint(os.Stderr, "Error to create file: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Unidade: %s, Bairro: %s, Localidade: %s, UF: %s, Estado: %s, Regiao: %s, IBGE: %s, GIA: %s, DDD: %s, SIAFI: %s\n",
			data.Cep, data.Logradouro, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.Uf, data.Estado, data.Regiao, data.Ibge, data.Gia, data.Ddd, data.Siafi))
		fmt.Println("Data saved successfully")
		fmt.Println("Cidade:", data.Localidade)

	}
}
