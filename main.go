package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Res struct {
	// 1. responding with struct first letter has to capital or it'll return empty
	Response string `json:"response"`
	// to change key name set `json:"<your-custom-name>"`
}

func main() {
	// making a wait group so the server won't close on it's own
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // set up a go routine
		// newMux := http.ServeMux()
		// newMux.HandleFunc("")
		mux := http.NewServeMux() // set a http handler
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method)
			if r.Method != "GET" { // setting only GET method to handle
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Method not found!"))
				return
			}
			w.Header().Set("Content-Type", "application/json") // setting header
			defer fmt.Println("response is sent!!")            // defer works after function is finished
			// res := []byte(`hello world`)
			res := Res{Response: "not sending simple str"}
			fmt.Println(res)
			jsonRes, _ := json.Marshal(res) // setting up a json response

			fmt.Println("in json", string(jsonRes))
			w.Write(jsonRes) // sending response
			// io.WriteString(w, res)
		})

		// mux.ServeHTTP()

		for { // making graceful shutdown to server
			// adding localhost or network ip addr will prevent windows err from thinking it's a malware
			if err := http.ListenAndServe("localhost:8080", mux); err != nil {
				fmt.Println("Server error:", err)
				wg.Done() // if error accurs close server & print error

			}
		}
	}()
	wg.Wait() // keep the go routine working until it's done (currently when error accures)
	// time.Sleep(60 * time.Second) // was a temporary solution before wait group
}
