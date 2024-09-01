package main

/*
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadCast = make(chan Message)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all connections
}

type SignalingMessage struct {
	Type      string `json:"type"`
	SDP       string `json:sdp`
	Candidate string `json:candidate`
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func reader(conn *websocket.Conn) {
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients, conn)
			return
		}

		log.Println("Received:", msg.Content)

		// Forward the message to other clients
		for client := range clients {
			if client != conn {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Println("Error writing message:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
*/
/*
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return // Return early if there's an error
	}
	defer ws.Close() // Ensure the connection is closed when we're done

	clients[ws] = true
	log.Println("Client Connected")

	// Send a welcome message to the client
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hi Client!"))
	if err != nil {
		log.Println("Error writing message:", err)
		return
	}

	// Start the reader to listen for messages from the client
	go reader(ws)
}
*/
/*
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Server started on :8080")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	messageType, p, err := conn.ReadMessage()

	for {
		// Read in a message
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Print out that message for clarity
	log.Println(string(p))
	if err := conn.WriteMessage(messageType, p); err != nil {
		log.Println(err)
		return
	}
}*/
/*
func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	//server := socketio.NewServer(nil)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))

	reader(ws)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, Http!\n")
}*/
