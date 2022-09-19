package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v3"
)

type ShortUrl struct {
	ShortPath string `yaml:"path" json:"path"`
	TargetUrl string `yaml:"url" json:"url"`
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetFileBytes(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	CheckErr(err)
	return bytes
}

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if long_url, ok := pathToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, long_url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func shortUrlsToMap(shortUrls []ShortUrl) map[string]string {
	pathsToUrls := map[string]string{}
	for _, shortUrl := range shortUrls {
		pathsToUrls[shortUrl.ShortPath] = shortUrl.TargetUrl
	}
	return pathsToUrls
}

func YamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var shortUrls []ShortUrl
	err := yaml.Unmarshal(yml, &shortUrls)
	CheckErr(err)
	pathsToUrls := shortUrlsToMap(shortUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func JsonHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var shortUrls []ShortUrl
	err := json.Unmarshal(jsn, &shortUrls)
	CheckErr(err)
	pathsToUrls := shortUrlsToMap(shortUrls)
	return MapHandler(pathsToUrls, fallback), nil
}
