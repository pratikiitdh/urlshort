package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"urlshort"
)

func main() {

	mux := defaultMux()

	jsonFileName := flag.String("json", "./example.json", "json file which provides redirect routes")
	ymlFileName := flag.String("yml", "./example.yml", "yml file which provides redirect routes")
	flag.Parse()

	jsonBytes, err := os.ReadFile(*jsonFileName)
	if err != nil {
		panic(err)
	}

	ymlBytes, err := os.ReadFile(*ymlFileName)
	if err != nil {
		panic(err)
	}
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/insta": "https://instagram.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler(ymlBytes, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JsonHandler([]byte(jsonBytes), yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
