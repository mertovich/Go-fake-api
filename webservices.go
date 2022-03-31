package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Periyot/BodyParser"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todo", indexHome)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file, _ := ioutil.ReadFile(`data.json`)
	_, err := fmt.Fprint(w, string(file))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func indexHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}

	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)
	mapsJSON, _ := json.Marshal(maps)
	timeNow := time.Now()
	fmt.Println(`[`, timeNow, `]`, `Body = `, string(mapsJSON))
	fmt.Fprint(w, string(mapsJSON))
}
