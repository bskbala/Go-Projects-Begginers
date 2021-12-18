package examples

import (
	"encoding/json"
	"net/http"
	"path"
)

type (
	fruitMap map[string]interface{}
)

// FruitHandler Creates http.Handler for the Fruits Server
//Routes
//GET/fruits  ===> Get Fruit list
//Get/fruits(name) ===> get fruit
//put/fruits(name) ===> add or update fruit

func FruitsHandler() http.Handler {
	fruits := fruitMap{}

	mux := http.NewServeMux()

	mux.HandleFunc("/fruits/", func(w http.ResponseWriter, r *http.Request) {
		handleFruitlist(fruits, w, r)
	})
	mux.HandleFunc("/fruits/", func(w http.ResponseWriter, r *http.Request) {
		handleFruit(fruits, w, r)
	})
	return mux

}

func handleFruitlist(fruits fruitMap, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ret := []string{}
		for k := range fruits {
			ret = append(ret, k)
		}
		b, err := json.Marshal(ret)
		if err != nil {
			panic(err)

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleFruit(fruits fruitMap, w http.ResponseWriter, r *http.Request) {
	_, name := path.Split(r.URL.Path)
	switch r.Method {
	case "GET":
		if data, ok := fruits[name]; ok {
			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)

		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	case "PUT":
		var data map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			panic(err)
		}
		fruits[name] = data
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}
