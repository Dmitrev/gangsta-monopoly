package main

import (
	"math/rand"
	"time"

	"net/http"

	"log"

	"flag"

	"github.com/Dmitrev/gangsta-monopoly/game"
)

var addr = flag.String("addr", ":8080", "http service address")

func init() {
	// Make sure that we can generate random numbers
	rand.Seed(time.Now().UnixNano())
}

var g *game.Game

func main() {
	flag.Parse()

	server := newServer()
	go server.run()
	// Create simple file server to serve static files from the public directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Handle the websocket connection
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(server, w, r)
	})

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Started server on port %s", addr)

	//g = game.NewGame()
	//startServer()
}
