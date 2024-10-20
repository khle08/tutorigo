////////////////////////////////////////////////////////////////////////

package tools

import (
    // "github.com/gin-gonic/gin"
    // "github.com/gorilla/websocket"
    // "log"
    // "net/http"
    // "sync"
)

////////////////////////////////////////////////////////////////////////


// // Room represents a chat room
// type Room struct {
//     clients   []*websocket.Conn // Directly use *websocket.Conn
//     broadcast chan []byte
//     mu        sync.RWMutex
// }

// var rooms = make(map[string]*Room)
// var upgrader = websocket.Upgrader{
//     CheckOrigin: func(r *http.Request) bool {
//         return true // Allow connections from any origin
//     },
// }

// func main() {
//     router := gin.Default()
//     router.GET("/ws/:room", handleWebSocket)

//     if err := router.Run(":8080"); err != nil {
//         log.Fatalf("Server failed to start: %v", err)
//     }
// }

// // handleWebSocket handles incoming WebSocket connections for a specific room
// func handleWebSocket(c *gin.Context) {
//     roomName := c.Param("room")
//     conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//     if err != nil {
//         log.Printf("Failed to upgrade connection: %v", err)
//         return
//     }
//     defer conn.Close()

//     room := joinRoom(roomName, conn)
//     log.Printf("Client joined room: %s", roomName)

//     // Handle client messages
//     for {
//         _, msg, err := conn.ReadMessage()
//         if err != nil {
//             log.Printf("Client in room %s disconnected: %v", roomName, err)
//             break
//         }
//         log.Printf("Message from room %s: %s", roomName, msg)

//         // Broadcast message to the room
//         room.broadcast <- msg
//     }

//     // Leave the room
//     leaveRoom(room, conn)
// }

// // joinRoom adds a WebSocket connection to the specified room, creating the room if necessary
// func joinRoom(roomName string, conn *websocket.Conn) *Room {
//     room := getRoom(roomName)

//     room.mu.Lock()
//     defer room.mu.Unlock()
//     room.clients = append(room.clients, conn)
//     return room
// }

// // leaveRoom removes a WebSocket connection from the room
// func leaveRoom(room *Room, conn *websocket.Conn) {
//     room.mu.Lock()
//     defer room.mu.Unlock()

//     for i, c := range room.clients {
//         if c == conn {
//             room.clients = append(room.clients[:i], room.clients[i+1:]...)
//             break
//         }
//     }
// }

// // getRoom retrieves or creates a room
// func getRoom(name string) *Room {
//     if room, exists := rooms[name]; exists {
//         return room
//     }

//     room := &Room{
//         clients:   []*websocket.Conn{},
//         broadcast: make(chan []byte),
//     }
//     rooms[name] = room

//     go handleMessages(room)
//     return room
// }

// // handleMessages listens for messages in a room and broadcasts them to all clients
// func handleMessages(room *Room) {
//     for msg := range room.broadcast {
//         room.mu.RLock()
//         for _, conn := range room.clients {
//             if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
//                 log.Printf("Error sending message to client: %v", err)
//                 conn.Close()
//             }
//         }
//         room.mu.RUnlock()
//     }
// }

