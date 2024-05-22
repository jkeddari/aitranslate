package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jkeddari/aitranslate/model"
)

func main() {
	translateRequest := model.Request{
		Language: "english",
		Source:   "Hello Sir !",
	}

	jsonValue, _ := json.Marshal(translateRequest)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/translate", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var translatorResponse model.Response

	json.NewDecoder(resp.Body).Decode(&translatorResponse)

	fmt.Println(translatorResponse)
}
