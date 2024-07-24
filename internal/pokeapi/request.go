package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err := fmt.Errorf("Response failed with status code: %d and body: %s", res.StatusCode, body)
		log.Fatal(err)
		return nil, err
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return body, nil
}
