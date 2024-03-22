package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func serve(){

	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", netWelcome)
	mux.HandleFunc("/rmcalc", netRMCalculateBufferSeries)
	http.ListenAndServe(":5000", mux)
}

func netWelcome(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Welcome to a calcbuf istance\n")
	fmt.Fprintf(w, "\t check /rmcalc to calculate buffer series\n")
}

func netRMCalculateBufferSeries(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w, "Please send a POST request with the correct json formatting", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// populate json
	var cb CalcBuf
	err = json.Unmarshal(body, &cb)
	if err != nil {
		http.Error(w, "Error Parsing user data", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(cb.RMBufferSeries())
	if err != nil {
		http.Error(w, "An error occurred!", http.StatusInternalServerError)
		return
	}

}

