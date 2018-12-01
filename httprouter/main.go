package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getHello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lang := p.ByName("lang")
	fmt.Fprintf(w, "hello %s\n", lang)
}

func postHello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type SampleParameter struct {
		ID   int
		Name string
	}
	var param SampleParameter

	err = json.Unmarshal(bodyBytes, &param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%+v\n", param))
}

func main() {
	router := httprouter.New()

	// リクエスト例
	// curl http://localhost:8080/hello/world
	router.GET("/hello/:lang", getHello)

	// リクエスト例
	// curl -X POST -d "{\"id\": 1, \"name\": \"ktny\"}" http://localhost:8080/hello
	router.POST("/hello", postHello)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
