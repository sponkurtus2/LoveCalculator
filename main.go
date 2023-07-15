package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var nombre1, nombre2 string
	fmt.Println("Ingrese el nombre de la persona 1")
	fmt.Scanln(&nombre1)
	fmt.Println("Ingrese el nombre de la persona 2")
	fmt.Scanln(&nombre2)

	url := fmt.Sprintf("https://love-calculator.p.rapidapi.com/getPercentage?sname=%s&fname=%s", nombre1, nombre2)


	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "bc97482e5amsha7b095f66e9c6a4p19d731jsn0cdfa6d84b35")
	req.Header.Add("X-RapidAPI-Host", "love-calculator.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	//check status api fmt.Println(res.Status)

	// Extraer el valor del JSON de respuesta
	type LoveCalculatorResponse struct {
		Fname string `json:"fname"`
		Sname string `json:"sname"`
		Percentage string `json:"percentage"`
	}

	var response LoveCalculatorResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return
	}

	// Imprimir el valor de 'fname'

	fmt.Printf("Nombre 1: %s, Nombre 2: %s, su porcentaje de amor es de: %s%%", response.Fname, response.Sname, response.Percentage)
}
