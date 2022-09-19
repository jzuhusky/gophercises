package main

import (
	"flag"
	"fmt"
	"net/http"

	"urlshort/handlers"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func main() {

	jsonFileName := flag.String("json-file", "", "Json file of short url mappings")
	yamlFileName := flag.String("yaml-file", "", "Yaml file of short url mappings")
	flag.Parse()

	if *jsonFileName == "" && *yamlFileName == "" {
		fmt.Println("must specify at least one of --json-file or --yaml-file")
		return
	}

	// Defaults
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathToUrls, mux)
	lastHandler := &mapHandler

	if *jsonFileName != "" {
		jsnBytes := handlers.GetFileBytes(*jsonFileName)
		jsonHandler, err := handlers.JsonHandler(jsnBytes, *lastHandler)
		handlers.CheckErr(err)
		lastHandler = &jsonHandler
	}

	if *yamlFileName != "" {
		ymlBytes := handlers.GetFileBytes(*yamlFileName)
		ymlHandler, err := handlers.YamlHandler(ymlBytes, *lastHandler)
		handlers.CheckErr(err)
		lastHandler = &ymlHandler
	}

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", *lastHandler)

}
