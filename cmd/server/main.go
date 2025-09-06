package main

import (
    "log"
    "os"

    "gofr.dev/pkg/gofr"
    "github.com/KraisuN-1010/student-rooms-backend/db"
    "github.com/KraisuN-1010/student-rooms-backend/internal/handlers"
    "github.com/KraisuN-1010/student-rooms-backend/services"
)

func main() {
    // Connect to database
    if err := db.Connect(); err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }

    // Create GoFr app
    app := gofr.New()

    // Initialize services
    roomService := services.NewRoomService()
    authService := services.NewAuthService()
    noteService := services.NewNoteService()
    commentService := services.NewCommentService()
    doubtService := services.NewDoubtService()

    // Initialize WebSocket service (for future integration)
    _ = services.NewWebSocketService()

    // Initialize handlers
    roomHandler := handlers.NewRoomHandler(roomService)
    authHandler := handlers.NewAuthHandler(authService)
    noteHandler := handlers.NewNoteHandler(noteService)
    commentHandler := handlers.NewCommentHandler(commentService)
    doubtHandler := handlers.NewDoubtHandler(doubtService)
    uploadHandler := handlers.NewFileUploadHandler("/tmp/uploads")

    // Health check
    app.GET("/", func(c *gofr.Context) (interface{}, error) {
        return "Student Rooms Backend is running!", nil
    })

    // Auth routes
    app.POST("/auth/signup", authHandler.SignUp)
    app.POST("/auth/login", authHandler.Login)

    // Rooms routes
    app.GET("/rooms", roomHandler.GetRooms)
    app.POST("/rooms", roomHandler.CreateRoom)

    // Posts routes (using posts table)
    app.GET("/rooms/:roomId/posts", noteHandler.GetNotesByRoom)
    app.POST("/rooms/:roomId/posts", noteHandler.CreateNote)
    
    // Notes routes (alias for backward compatibility)
    app.GET("/rooms/:roomId/notes", noteHandler.GetNotesByRoom)
    app.POST("/rooms/:roomId/notes", noteHandler.CreateNote)

    // Comments routes
    app.GET("/comments/:parentId", commentHandler.GetCommentsByParent)
    app.POST("/comments", commentHandler.CreateComment)

    // Doubts routes
    app.GET("/rooms/:roomId/doubts", doubtHandler.GetDoubtsByRoom)
    app.POST("/rooms/:roomId/doubts", doubtHandler.CreateDoubt)

    // File upload routes
    app.POST("/upload", uploadHandler.UploadFile)
    app.POST("/upload/multiple", uploadHandler.UploadMultipleFiles)
    app.GET("/files/:fileId", uploadHandler.GetFileInfo)

    // WebSocket integration endpoint
    app.GET("/ws-info", func(c *gofr.Context) (interface{}, error) {
        // For production, use environment variables or configuration
        websocketURL := os.Getenv("WEBSOCKET_URL")
        if websocketURL == "" {
            websocketURL = "ws://localhost:8001/ws"
        }
        
        return map[string]string{
            "websocket_url": websocketURL,
            "status": "WebSocket server running",
            "test_client": "/static/websocket_client.html",
        }, nil
    })

    // Port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    log.Println("Starting server on port", port)
    app.Run()
}
