//go:generate file2gostr -src "index.html" -dst "index.html.go" -var "webpageData" -pkg "main"
//go:generate go fmt

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/pkg/browser"
)

var (
	upgrader     = websocket.Upgrader{}
	messages     = make(chan []byte)
	closed       = make(chan bool)
	ready        = make(chan bool)
	address      = "localhost:8080"
	serverLocked = false
)

func lockServer() {
	serverLocked = true
}

func unlockServer() {
	serverLocked = false
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, webpageData)
	return
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if serverLocked {
		return
	}
	lockServer()
	defer unlockServer()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		select {
		case msg := <-messages:
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err)
				return
			}
		case <-closed:
			log.Println("closing connection to ", r.RemoteAddr)
			return
		}
	}
}

func runInputLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		messages <- []byte(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/ws", wsHandler)
	s := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}

func main() {
	go runServer()
	browser.OpenURL("http://" + address)
	runInputLoop()
}
