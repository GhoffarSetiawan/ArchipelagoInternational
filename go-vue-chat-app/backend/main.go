package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

// User represents a chat user
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Conn     *websocket.Conn `json:"-"`
	IsOnline bool `json:"isOnline"`
}

// Message represents a chat message
type Message struct {
	Type    string `json:"type"`
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

// Global state
var (
	users     = make(map[string]*User)
	userMutex = sync.RWMutex{}
)

func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	// Serve static files from the frontend (for production deployment)
	app.Static("/", "../frontend/dist")

	// Setup websocket route
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Handle WebSocket connections
	app.Get("/ws/:id", websocket.New(handleWebSocket))

	// Start the server
	log.Fatal(app.Listen(":8080"))
}

func handleWebSocket(c *websocket.Conn) {
	userID := c.Params("id")
	log.Printf("New WebSocket connection: %s", userID)

	// Register the new user
	userMutex.Lock()
	user, exists := users[userID]
	if exists {
		// User reconnected
		user.Conn = c
		user.IsOnline = true
	} else {
		// Create new user with placeholder name
		user = &User{
			ID:       userID,
			Name:     userID, // Temporary name until user sets it
			Conn:     c,
			IsOnline: true,
		}
		users[userID] = user
	}
	userMutex.Unlock()

	// Broadcast updated user list to all clients
	broadcastUserList()

	var (
		msg     Message
		err     error
	)

	// WebSocket message handling loop
	for {
		if err = c.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}

		// Process the message based on type
		switch msg.Type {
		case "register":
			// Update user name
			userMutex.Lock()
			user.Name = msg.Content
			userMutex.Unlock()
			
			// Broadcast updated user list
			broadcastUserList()
			
		case "message":
			// Send message to specified recipient
			forwardMessage(msg)
			
		default:
			log.Printf("Unknown message type: %s", msg.Type)
		}
	}

	// Handle disconnection
	userMutex.Lock()
	user.IsOnline = false
	userMutex.Unlock()
	
	// Broadcast updated user list
	broadcastUserList()
	
	log.Printf("WebSocket connection closed: %s", userID)
}

// broadcastUserList sends the current user list to all connected clients
func broadcastUserList() {
	userMutex.RLock()
	defer userMutex.RUnlock()
	
	// Create a list of users with only the necessary information
	userList := make([]map[string]interface{}, 0, len(users))
	for _, user := range users {
		if user.IsOnline {
			userList = append(userList, map[string]interface{}{
				"id":   user.ID,
				"name": user.Name,
			})
		}
	}
	
	// Prepare the message
	msg := Message{
		Type:    "users",
		Content: "",
	}
	
	// Convert user list to JSON
	content, err := json.Marshal(userList)
	if err != nil {
		log.Println("Error marshaling user list:", err)
		return
	}
	msg.Content = string(content)
	
	// Send to all online users
	for _, user := range users {
		if user.IsOnline && user.Conn != nil {
			if err := user.Conn.WriteJSON(msg); err != nil {
				log.Printf("Error sending user list to %s: %v", user.ID, err)
			}
		}
	}
}

// forwardMessage sends a message to the specified recipient
func forwardMessage(msg Message) {
	userMutex.RLock()
	defer userMutex.RUnlock()
	
	// Find the recipient
	recipient, exists := users[msg.To]
	if !exists || !recipient.IsOnline || recipient.Conn == nil {
		log.Printf("Recipient not found or offline: %s", msg.To)
		return
	}
	
	// Send the message to the recipient
	if err := recipient.Conn.WriteJSON(msg); err != nil {
		log.Printf("Error sending message to %s: %v", msg.To, err)
	}
}
