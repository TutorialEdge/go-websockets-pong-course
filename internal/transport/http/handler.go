package http

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Router   *mux.Router
	upgrader websocket.Upgrader
}

// New - returns a new http handler
func New() Handler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return Handler{
		upgrader: upgrader,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/ws", h.ServeWs).Methods("GET")
	h.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
}

func (h *Handler) ServeWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New Connection established")
	ws, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	go writer(ws)
	reader(ws)
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func writer(conn *websocket.Conn) {
	for {
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", r)
		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
