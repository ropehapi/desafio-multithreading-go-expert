package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	resultChannel := make(chan string)
	go fetchFromApi("https://brasilapi.com.br/api/cep/v1/86601002", resultChannel)
	go fetchFromApi("http://viacep.com.br/ws/86601002/json/", resultChannel)

	result := <-resultChannel
	if result == "" {
		result = <-resultChannel
	}

	fmt.Println(result)
}

func fetchFromApi(url string, resultChannel chan<- string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	resultChannel <- string(body)
}
