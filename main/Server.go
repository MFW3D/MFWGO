//main service entry
package main

import (
	"flag"
	"log"
	"../Hub"
	"net/http"
)

var addr = flag.String("addr", ":8081", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "../html/home.html")
}

func main() {
	flag.Parse()
	//ini chat room
	hub := Hub.NewHub()
	//run room
	go hub.Run()
	//http service init
	http.HandleFunc("/", serveHome)
	//ws request handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		Hub.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
