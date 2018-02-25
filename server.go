// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "log"

type Message struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

type Server struct {
	// Registered clients
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (s *Server) run() {
	for {
		select {
		case client := <-s.register:
			log.Println("[server.run()] s.register")
			s.clients[client] = true
			sendRegisterRequest(client)
		case client := <-s.unregister:
			log.Println("[server.run()] s.unregister")
			if _, ok := s.clients[client]; ok {
				log.Println("[server.run()] s.unregister - deleting client")
				delete(s.clients, client)
				close(client.send)
			}
		case message := <-s.broadcast:
			log.Println("[server.run()] s.broadcast")
			for client := range s.clients {
				select {
				case client.send <- message:
					log.Println("[server.run()] s.register > message")
				default:
					log.Println("[server.run()] s.register > [default]")
					close(client.send)
					delete(s.clients, client)
				}
			}

		}
	}
}

//func websocketHandler(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//
//	if err != nil {
//		log.Fatalf("[websocketHandler] Failed to upgrade connection (%s)\n", err)
//	}
//
//	defer conn.Close()
//
//	// Once the user gets here, that means we have a new connection
//	// Add client to the clients slice
//	clients[conn] = true
//	log.Printf("New user connected\n")
//	// Send register request to client
//	sendRegisterRequest(conn)
//
//	var msg Message
//	// Connection loop
//	for {
//		// Incoming JSON from the client
//		err := conn.ReadJSON(&msg)
//
//		// If we cannot read the message anymore that means that the user is disconnected
//		if err != nil {
//			// Handle disconnect login in game
//			g.Disconnect(conn)
//			// Remove connection from clients
//			delete(clients, conn)
//			break
//		}
//
//		handleEvent(conn, msg)
//	}
//
//}
