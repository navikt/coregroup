package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"flag"
	"os"
	"fmt"
)

var (
	coregroupFile = flag.String("file", "", "")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "", "")
)

var usage = `Usage: coregroups [options...]

Options:

  -file  		JSON-file containing your endpoints
  -cert			Server certificate
  -key			Server private key
`

type coregroup struct {
	Application string `json:"application"`
	CoregroupName string `json:"coregroupName"`
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
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

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	if flag.NFlag() < 3 {
		usageAndExit("You did not supply enough arguments")
	}

	data, err := ioutil.ReadFile(*coregroupFile)

	if err != nil {
		log.Fatal("unable to read file ", *coregroupFile)
		panic(err)
	}
		
	coregroups := []coregroup{}
	err = json.Unmarshal(data, &coregroups)

	if err != nil {
		log.Fatal("Couldn't parse JSON: ", string(data))	
		panic(err)
	}

	if _, err := os.Stat(*cert); os.IsNotExist(err) {
  		log.Fatal("Certificate %s does not exist", *cert)
	}

	if _, err := os.Stat(*key); os.IsNotExist(err) {
	 	log.Fatal("Certificate %s does not exist", *key)
	}

	mux := http.NewServeMux()
	vh := viewHandler(&coregroups)
	mux.Handle("/coregroup/", vh)
    err = http.ListenAndServeTLS(":8443", *cert, *key, mux)

	if err != nil {
		log.Fatal("Couldn't start application on :8443")	
		panic(err)
	}
}

