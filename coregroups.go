package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

type coregroup struct {
	Application string `json:"application"`
	CoregroupName string `json:"coregroupName"`
}

func viewHandler(coregroups *[]coregroup) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
		var applicationName string
		if r.URL.Query().Get("application") == "" {
			w.Write([]byte("Request parameter application missing "))
			return
		} else {
			applicationName = r.URL.Query().Get("application")
			for _,correctCoregroup := range *coregroups {
				if applicationName == correctCoregroup.Application {
					w.Write([]byte(correctCoregroup.CoregroupName))
					return
				}
			}
			w.Write([]byte("DefaultCoreGroup"))
		}
	})
}

func main() {

	data, err := ioutil.ReadFile("coregroups.json")

	if err != nil {
		log.Fatal("unable to read file coregroups.json")
		panic(err)
	}
		
	coregroups := []coregroup{}
	err = json.Unmarshal(data, &coregroups)

	if err != nil {
		log.Fatal("Couldn't parse JSON: ", string(data))	
		panic(err)
	}

	mux := http.NewServeMux()
	vh := viewHandler(&coregroups)
	mux.Handle("/coregroup/", vh)
    http.ListenAndServe(":8080", mux)
}

