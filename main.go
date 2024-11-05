package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please include the word or phrase you want to translate")
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")

	req := TranslateRequest{
		Source:       "sv",
		Target:       "en",
		Query:        input,
		Alternatives: 3,
		Format:       "text",
	}

	payload, err := json.Marshal(req)
	if err != nil {
		panic("error when marshalling json")
	}

	fmt.Println(string(payload))

	resp, err := http.Post("http://localhost:5000/translate", "application/json", bytes.NewReader(payload))
	if err != nil {
		panic("error when posting request")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("error when reading body")
	}

	fmt.Println(string(bodyBytes))
}

type TranslateRequest struct {
	Source       string `json:"source"`
	Target       string `json:"target"`
	Query        string `json:"q"`
	Format       string `json:"format"`
	Alternatives int    `json:"alternatives"`
}
