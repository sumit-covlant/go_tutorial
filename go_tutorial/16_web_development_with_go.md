# Web Development with Go

## Overview

Go is excellent for web development, offering built-in HTTP support, excellent performance, and a rich ecosystem of web frameworks and libraries. This guide covers building web applications, APIs, and services using Go's standard library and popular frameworks.

## Basic HTTP Server

### Simple HTTP Server

Go's `net/http` package provides everything needed to build web servers.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Handle root path
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    // Handle specific path
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "About Us")
    })
    
    // Start server
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

### HTTP Methods and Request Handling

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        // Handle GET request
        user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
        
    case "POST":
        // Handle POST request
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        // Process user data
        fmt.Printf("Received user: %+v\n", user)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
        
    case "PUT":
        // Handle PUT request
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "User updated"})
        
    case "DELETE":
        // Handle DELETE request
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
        
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/user", handleUser)
    http.ListenAndServe(":8080", nil)
}
```

## Routing and URL Parameters

### Custom Router

```go
package main

import (
    "fmt"
    "net/http"
    "regexp"
    "strings"
)

type Route struct {
    Pattern *regexp.Regexp
    Handler http.HandlerFunc
    Methods []string
}

type Router struct {
    routes []Route
}

func NewRouter() *Router {
    return &Router{}
}

func (r *Router) AddRoute(pattern string, handler http.HandlerFunc, methods ...string) {
    regex := regexp.MustCompile("^" + pattern + "$")
    r.routes = append(r.routes, Route{
        Pattern: regex,
        Handler: handler,
        Methods: methods,
    })
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.routes {
        if route.Pattern.MatchString(req.URL.Path) {
            // Check if method is allowed
            if len(route.Methods) == 0 || contains(route.Methods, req.Method) {
                route.Handler(w, req)
                return
            }
        }
    }
    http.NotFound(w, req)
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func main() {
    router := NewRouter()
    
    router.AddRoute("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Home Page")
    })
    
    router.AddRoute("/users/([0-9]+)", func(w http.ResponseWriter, r *http.Request) {
        // Extract user ID from URL
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) >= 3 {
            userID := parts[2]
            fmt.Fprintf(w, "User ID: %s", userID)
        }
    }, "GET")
    
    router.AddRoute("/api/users", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "API Users")
    }, "GET", "POST")
    
    http.ListenAndServe(":8080", router)
}
```

### Using Gorilla Mux

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]
    
    user := User{ID: userID, Name: "Alice", Email: "alice@example.com"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Process user creation
    user.ID = "generated-id"
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    r := mux.NewRouter()
    
    // Routes with parameters
    r.HandleFunc("/users/{id}", getUser).Methods("GET")
    r.HandleFunc("/users", createUser).Methods("POST")
    
    // Subrouting
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/users/{id}", getUser).Methods("GET")
    api.HandleFunc("/users", createUser).Methods("POST")
    
    // Query parameters
    r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("q")
        fmt.Fprintf(w, "Search query: %s", query)
    }).Methods("GET")
    
    http.ListenAndServe(":8080", r)
}
```

## Middleware

### Custom Middleware

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging middleware
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Call the next handler
        next(w, r)
        
        // Log the request details
        log.Printf(
            "%s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(start),
        )
    }
}

// Authentication middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Validate token (simplified)
        if token != "Bearer valid-token" {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        next(w, r)
    }
}

// CORS middleware
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next(w, r)
    }
}

// Chain multiple middleware
func Chain(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
    for _, m := range middleware {
        h = m(h)
    }
    return h
}

func main() {
    // Handler function
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    }
    
    // Apply middleware
    finalHandler := Chain(handler, LoggingMiddleware, CORSMiddleware)
    
    http.HandleFunc("/", finalHandler)
    http.HandleFunc("/protected", Chain(handler, LoggingMiddleware, AuthMiddleware))
    
    http.ListenAndServe(":8080", nil)
}
```

### Using Negroni Middleware

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/urfave/negroni"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    // Create negroni instance with middleware
    n := negroni.New()
    n.Use(negroni.NewLogger())
    n.Use(negroni.NewRecovery())
    n.UseHandler(mux)
    
    http.ListenAndServe(":8080", n)
}
```

## Templates and HTML

### HTML Templates

```go
package main

import (
    "html/template"
    "net/http"
)

type PageData struct {
    Title   string
    Content string
    Users   []User
}

type User struct {
    Name  string
    Email string
}

func main() {
    // Parse templates
    tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title:   "Home Page",
            Content: "Welcome to our website!",
            Users: []User{
                {Name: "Alice", Email: "alice@example.com"},
                {Name: "Bob", Email: "bob@example.com"},
            },
        }
        
        tmpl.ExecuteTemplate(w, "layout", data)
    })
    
    http.ListenAndServe(":8080", nil)
}
```

### Template Files

**templates/layout.html:**
```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <header>
        <nav>
            <a href="/">Home</a>
            <a href="/about">About</a>
            <a href="/users">Users</a>
        </nav>
    </header>
    
    <main>
        {{template "content" .}}
    </main>
    
    <footer>
        <p>&copy; 2023 My Website</p>
    </footer>
</body>
</html>
```

**templates/home.html:**
```html
{{define "content"}}
<h1>{{.Title}}</h1>
<p>{{.Content}}</p>

<h2>Users</h2>
<ul>
    {{range .Users}}
    <li>{{.Name}} ({{.Email}})</li>
    {{end}}
</ul>
{{end}}
```

### Template Functions

```go
package main

import (
    "html/template"
    "net/http"
    "strings"
    "time"
)

func main() {
    // Create template functions
    funcMap := template.FuncMap{
        "uppercase": strings.ToUpper,
        "formatDate": func(t time.Time) string {
            return t.Format("2006-01-02")
        },
        "add": func(a, b int) int {
            return a + b
        },
    }
    
    // Parse template with functions
    tmpl := template.Must(template.New("").Funcs(funcMap).Parse(`
        <h1>{{.Title | uppercase}}</h1>
        <p>Date: {{.Date | formatDate}}</p>
        <p>Sum: {{add .A .B}}</p>
    `))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := struct {
            Title string
            Date  time.Time
            A, B  int
        }{
            Title: "Hello World",
            Date:  time.Now(),
            A:     5,
            B:     3,
        }
        
        tmpl.Execute(w, data)
    })
    
    http.ListenAndServe(":8080", nil)
}
```

## REST API Development

### Complete REST API

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    
    "github.com/gorilla/mux"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserService struct {
    users map[int]User
    nextID int
}

func NewUserService() *UserService {
    return &UserService{
        users: make(map[int]User),
        nextID: 1,
    }
}

func (s *UserService) CreateUser(user User) User {
    user.ID = s.nextID
    s.users[user.ID] = user
    s.nextID++
    return user
}

func (s *UserService) GetUser(id int) (User, bool) {
    user, exists := s.users[id]
    return user, exists
}

func (s *UserService) GetAllUsers() []User {
    users := make([]User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return users
}

func (s *UserService) UpdateUser(id int, user User) (User, bool) {
    if _, exists := s.users[id]; !exists {
        return User{}, false
    }
    user.ID = id
    s.users[id] = user
    return user, true
}

func (s *UserService) DeleteUser(id int) bool {
    if _, exists := s.users[id]; !exists {
        return false
    }
    delete(s.users, id)
    return true
}

type UserHandler struct {
    service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    createdUser := h.service.CreateUser(user)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    user, exists := h.service.GetUser(id)
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users := h.service.GetAllUsers()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    updatedUser, exists := h.service.UpdateUser(id, user)
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    if !h.service.DeleteUser(id) {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    service := NewUserService()
    handler := NewUserHandler(service)
    
    r := mux.NewRouter()
    
    // API routes
    r.HandleFunc("/api/users", handler.CreateUser).Methods("POST")
    r.HandleFunc("/api/users", handler.GetAllUsers).Methods("GET")
    r.HandleFunc("/api/users/{id}", handler.GetUser).Methods("GET")
    r.HandleFunc("/api/users/{id}", handler.UpdateUser).Methods("PUT")
    r.HandleFunc("/api/users/{id}", handler.DeleteUser).Methods("DELETE")
    
    http.ListenAndServe(":8080", r)
}
```

## Web Frameworks

### Gin Framework

```go
package main

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

func main() {
    r := gin.Default()
    
    // Middleware
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    
    // Static files
    r.Static("/static", "./static")
    r.LoadHTMLGlob("templates/*")
    
    // Routes
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Home Page",
        })
    })
    
    // API routes
    api := r.Group("/api")
    {
        api.GET("/users", func(c *gin.Context) {
            users := []User{
                {ID: 1, Name: "Alice", Email: "alice@example.com"},
                {ID: 2, Name: "Bob", Email: "bob@example.com"},
            }
            c.JSON(http.StatusOK, users)
        })
        
        api.POST("/users", func(c *gin.Context) {
            var user User
            if err := c.ShouldBindJSON(&user); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            
            user.ID = 3 // In real app, generate ID
            c.JSON(http.StatusCreated, user)
        })
        
        api.GET("/users/:id", func(c *gin.Context) {
            id := c.Param("id")
            user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
            c.JSON(http.StatusOK, user)
        })
    }
    
    r.Run(":8080")
}
```

### Echo Framework

```go
package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    e := echo.New()
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())
    
    // Routes
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    
    // API group
    api := e.Group("/api")
    {
        api.GET("/users", getUsers)
        api.POST("/users", createUser)
        api.GET("/users/:id", getUser)
        api.PUT("/users/:id", updateUser)
        api.DELETE("/users/:id", deleteUser)
    }
    
    e.Logger.Fatal(e.Start(":8080"))
}

func getUsers(c echo.Context) error {
    users := []User{
        {ID: 1, Name: "Alice", Email: "alice@example.com"},
        {ID: 2, Name: "Bob", Email: "bob@example.com"},
    }
    return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
    user := new(User)
    if err := c.Bind(user); err != nil {
        return err
    }
    
    user.ID = 3 // Generate ID
    return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
    id := c.Param("id")
    user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
    return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
    id := c.Param("id")
    user := new(User)
    if err := c.Bind(user); err != nil {
        return err
    }
    
    return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
    id := c.Param("id")
    return c.NoContent(http.StatusNoContent)
}
```

## WebSocket Support

### Basic WebSocket Server

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    
    "github.com/gorilla/websocket"
)

type Message struct {
    Type    string `json:"type"`
    Content string `json:"content"`
    User    string `json:"user"`
}

type Client struct {
    conn *websocket.Conn
    send chan []byte
}

type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
            
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
            
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}

func (c *Client) readPump(hub *Hub) {
    defer func() {
        hub.unregister <- c
        c.conn.Close()
    }()
    
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
        
        // Parse message
        var msg Message
        if err := json.Unmarshal(message, &msg); err != nil {
            log.Printf("error parsing message: %v", err)
            continue
        }
        
        // Broadcast message
        hub.broadcast <- message
    }
}

func (c *Client) writePump() {
    defer func() {
        c.conn.Close()
    }()
    
    for {
        select {
        case message, ok := <-c.send:
            if !ok {
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            
            w, err := c.conn.NextWriter(websocket.TextMessage)
            if err != nil {
                return
            }
            w.Write(message)
            
            if err := w.Close(); err != nil {
                return
            }
        }
    }
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins
    },
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    
    client := &Client{conn: conn, send: make(chan []byte, 256)}
    client.conn.SetReadLimit(512)
    
    hub.register <- client
    
    go client.writePump()
    go client.readPump(hub)
}

func main() {
    hub := NewHub()
    go hub.run()
    
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })
    
    log.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

## Static File Serving

### File Server with Security

```go
package main

import (
    "net/http"
    "path/filepath"
    "strings"
)

func main() {
    // Serve static files with security headers
    fs := http.FileServer(http.Dir("static"))
    
    http.Handle("/static/", http.StripPrefix("/static/", secureFileServer(fs)))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })
    
    http.ListenAndServe(":8080", nil)
}

func secureFileServer(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Security headers
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        
        // Prevent directory traversal
        if strings.Contains(r.URL.Path, "..") {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        
        // Only allow certain file types
        ext := filepath.Ext(r.URL.Path)
        allowedExts := map[string]bool{
            ".css":  true,
            ".js":   true,
            ".png":  true,
            ".jpg":  true,
            ".jpeg": true,
            ".gif":  true,
            ".ico":  true,
            ".svg":  true,
        }
        
        if !allowedExts[ext] {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

## Error Handling and Logging

### Structured Error Handling

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "runtime/debug"
)

type AppError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

func (e AppError) Error() string {
    return e.Message
}

func handleError(w http.ResponseWriter, err error) {
    var appErr AppError
    
    switch e := err.(type) {
    case AppError:
        appErr = e
    default:
        appErr = AppError{
            Code:    http.StatusInternalServerError,
            Message: "Internal Server Error",
            Details: err.Error(),
        }
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(appErr.Code)
    json.NewEncoder(w).Encode(appErr)
}

func errorMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic: %v\n%s", err, debug.Stack())
                handleError(w, AppError{
                    Code:    http.StatusInternalServerError,
                    Message: "Internal Server Error",
                })
            }
        }()
        
        next(w, r)
    }
}

func main() {
    http.HandleFunc("/", errorMiddleware(func(w http.ResponseWriter, r *http.Request) {
        // Simulate error
        if r.URL.Query().Get("error") == "true" {
            panic("Something went wrong!")
        }
        
        fmt.Fprintf(w, "Hello, World!")
    }))
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Summary

Go web development provides:

- **Built-in HTTP support**: No external dependencies for basic servers
- **Excellent performance**: Fast and efficient web applications
- **Rich ecosystem**: Many frameworks and libraries available
- **Concurrent by design**: Handle many requests efficiently
- **Type safety**: Compile-time error checking
- **Cross-platform**: Run on any platform

Key points to remember:
1. Use the standard library for simple applications
2. Choose frameworks (Gin, Echo) for complex applications
3. Implement proper middleware for cross-cutting concerns
4. Use templates for server-side rendering
5. Implement proper error handling and logging
6. Secure your applications with proper headers and validation
7. Use WebSockets for real-time communication
8. Structure your code with proper separation of concerns

Understanding web development with Go enables you to build scalable, performant web applications and APIs. 