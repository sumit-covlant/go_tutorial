# Database Interaction in Go

## Overview

Go provides excellent support for database interactions through the `database/sql` package and various database drivers. This guide covers working with SQL databases, ORMs, connection pooling, transactions, and best practices for database operations.

## Database/SQL Package

### Basic Database Connection

The `database/sql` package provides a generic interface around SQL databases.

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
    // Connection string
    connStr := "postgres://username:password@localhost/dbname?sslmode=disable"
    
    // Open database connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Successfully connected to database")
}
```

### Basic CRUD Operations

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
)

type User struct {
    ID    int
    Name  string
    Email string
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Create user
    userID, err := createUser(db, "Alice", "alice@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created user with ID: %d\n", userID)
    
    // Get user
    user, err := getUser(db, userID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User: %+v\n", user)
    
    // Update user
    err = updateUser(db, userID, "Alice Updated", "alice.updated@example.com")
    if err != nil {
        log.Fatal(err)
    }
    
    // Delete user
    err = deleteUser(db, userID)
    if err != nil {
        log.Fatal(err)
    }
}

func createUser(db *sql.DB, name, email string) (int, error) {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    var id int
    err := db.QueryRow(query, name, email).Scan(&id)
    return id, err
}

func getUser(db *sql.DB, id int) (User, error) {
    query := `SELECT id, name, email FROM users WHERE id = $1`
    var user User
    err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
    return user, err
}

func updateUser(db *sql.DB, id int, name, email string) error {
    query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
    _, err := db.Exec(query, name, email, id)
    return err
}

func deleteUser(db *sql.DB, id int) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := db.Exec(query, id)
    return err
}
```

### Querying Multiple Rows

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
)

func getAllUsers(db *sql.DB) ([]User, error) {
    query := `SELECT id, name, email FROM users ORDER BY id`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return users, nil
}

func getUsersByEmail(db *sql.DB, emailPattern string) ([]User, error) {
    query := `SELECT id, name, email FROM users WHERE email LIKE $1`
    rows, err := db.Query(query, "%"+emailPattern+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, nil
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Get all users
    users, err := getAllUsers(db)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, user := range users {
        fmt.Printf("User: %+v\n", user)
    }
    
    // Get users by email pattern
    gmailUsers, err := getUsersByEmail(db, "gmail.com")
    if err != nil {
        log.Fatal(err)
    }
    
    for _, user := range gmailUsers {
        fmt.Printf("Gmail user: %+v\n", user)
    }
}
```

## Transactions

### Basic Transaction

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
)

func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    // Start transaction
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Defer a rollback in case anything fails
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p) // re-throw panic after Rollback
        } else if err != nil {
            tx.Rollback() // err is non-nil; don't change it
        } else {
            err = tx.Commit() // err is nil; if Commit returns error update err
        }
    }()
    
    // Deduct from source account
    _, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    // Add to destination account
    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    return nil
}

func createUserWithProfile(db *sql.DB, name, email, bio string) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()
    
    // Insert user
    var userID int
    err = tx.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", name, email).Scan(&userID)
    if err != nil {
        return err
    }
    
    // Insert profile
    _, err = tx.Exec("INSERT INTO profiles (user_id, bio) VALUES ($1, $2)", userID, bio)
    if err != nil {
        return err
    }
    
    return nil
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Transfer money
    err = transferMoney(db, 1, 2, 100.0)
    if err != nil {
        log.Fatal(err)
    }
    
    // Create user with profile
    err = createUserWithProfile(db, "Alice", "alice@example.com", "Software developer")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Transaction with Context

```go
package main

import (
    "context"
    "database/sql"
    "time"
    
    _ "github.com/lib/pq"
)

func transferMoneyWithContext(ctx context.Context, db *sql.DB, fromID, toID int, amount float64) error {
    // Start transaction with context
    tx, err := db.BeginTx(ctx, &sql.TxOptions{
        Isolation: sql.LevelReadCommitted,
        ReadOnly:  false,
    })
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()
    
    // Execute queries with context
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    _, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    return nil
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    err = transferMoneyWithContext(ctx, db, 1, 2, 100.0)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Connection Pooling

### Connection Pool Configuration

```go
package main

import (
    "database/sql"
    "time"
    
    _ "github.com/lib/pq"
)

func configureConnectionPool(db *sql.DB) {
    // Set maximum number of open connections
    db.SetMaxOpenConns(25)
    
    // Set maximum number of idle connections
    db.SetMaxIdleConns(5)
    
    // Set maximum lifetime of a connection
    db.SetConnMaxLifetime(5 * time.Minute)
    
    // Set maximum idle time of a connection
    db.SetConnMaxIdleTime(3 * time.Minute)
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    configureConnectionPool(db)
    
    // Test connection pool
    for i := 0; i < 10; i++ {
        go func(id int) {
            var result int
            err := db.QueryRow("SELECT 1").Scan(&result)
            if err != nil {
                log.Printf("Query %d failed: %v", id, err)
            } else {
                log.Printf("Query %d succeeded: %d", id, result)
            }
        }(i)
    }
    
    time.Sleep(2 * time.Second)
}
```

## Prepared Statements

### Using Prepared Statements

```go
package main

import (
    "database/sql"
    "fmt"
    
    _ "github.com/lib/pq"
)

type UserRepository struct {
    db *sql.DB
    stmts map[string]*sql.Stmt
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
    repo := &UserRepository{
        db:    db,
        stmts: make(map[string]*sql.Stmt),
    }
    
    // Prepare statements
    stmts := map[string]string{
        "create": "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
        "get":    "SELECT id, name, email FROM users WHERE id = $1",
        "update": "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        "delete": "DELETE FROM users WHERE id = $1",
        "list":   "SELECT id, name, email FROM users ORDER BY id",
    }
    
    for name, query := range stmts {
        stmt, err := db.Prepare(query)
        if err != nil {
            return nil, fmt.Errorf("failed to prepare %s statement: %v", name, err)
        }
        repo.stmts[name] = stmt
    }
    
    return repo, nil
}

func (r *UserRepository) Close() {
    for _, stmt := range r.stmts {
        stmt.Close()
    }
}

func (r *UserRepository) Create(name, email string) (int, error) {
    var id int
    err := r.stmts["create"].QueryRow(name, email).Scan(&id)
    return id, err
}

func (r *UserRepository) Get(id int) (User, error) {
    var user User
    err := r.stmts["get"].QueryRow(id).Scan(&user.ID, &user.Name, &user.Email)
    return user, err
}

func (r *UserRepository) Update(id int, name, email string) error {
    _, err := r.stmts["update"].Exec(name, email, id)
    return err
}

func (r *UserRepository) Delete(id int) error {
    _, err := r.stmts["delete"].Exec(id)
    return err
}

func (r *UserRepository) List() ([]User, error) {
    rows, err := r.stmts["list"].Query()
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, nil
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    repo, err := NewUserRepository(db)
    if err != nil {
        log.Fatal(err)
    }
    defer repo.Close()
    
    // Use repository
    id, err := repo.Create("Alice", "alice@example.com")
    if err != nil {
        log.Fatal(err)
    }
    
    user, err := repo.Get(id)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Created user: %+v\n", user)
}
```

## ORM - GORM

### Basic GORM Usage

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:255;not null"`
    Email     string    `gorm:"size:255;uniqueIndex;not null"`
    Age       int       `gorm:"default:18"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Profile struct {
    ID     uint   `gorm:"primaryKey"`
    UserID uint   `gorm:"not null"`
    Bio    string `gorm:"type:text"`
    User   User   `gorm:"foreignKey:UserID"`
}

func main() {
    // Connect to database
    dsn := "host=localhost user=username password=password dbname=dbname port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // Auto migrate
    err = db.AutoMigrate(&User{}, &Profile{})
    if err != nil {
        log.Fatal(err)
    }
    
    // Create user
    user := User{
        Name:  "Alice",
        Email: "alice@example.com",
        Age:   25,
    }
    
    result := db.Create(&user)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
    
    fmt.Printf("Created user: %+v\n", user)
    
    // Find user
    var foundUser User
    result = db.First(&foundUser, user.ID)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
    
    fmt.Printf("Found user: %+v\n", foundUser)
    
    // Update user
    result = db.Model(&foundUser).Update("Age", 26)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
    
    // Find by email
    var userByEmail User
    result = db.Where("email = ?", "alice@example.com").First(&userByEmail)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
    
    fmt.Printf("User by email: %+v\n", userByEmail)
    
    // Delete user
    result = db.Delete(&user)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
}
```

### GORM with Associations

```go
package main

import (
    "fmt"
    "log"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID       uint      `gorm:"primaryKey"`
    Name     string    `gorm:"size:255;not null"`
    Email    string    `gorm:"size:255;uniqueIndex;not null"`
    Profile  Profile   `gorm:"foreignKey:UserID"`
    Posts    []Post    `gorm:"foreignKey:UserID"`
}

type Profile struct {
    ID     uint   `gorm:"primaryKey"`
    UserID uint   `gorm:"not null"`
    Bio    string `gorm:"type:text"`
}

type Post struct {
    ID      uint   `gorm:"primaryKey"`
    UserID  uint   `gorm:"not null"`
    Title   string `gorm:"size:255;not null"`
    Content string `gorm:"type:text"`
}

func main() {
    dsn := "host=localhost user=username password=password dbname=dbname port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    
    // Auto migrate
    err = db.AutoMigrate(&User{}, &Profile{}, &Post{})
    if err != nil {
        log.Fatal(err)
    }
    
    // Create user with profile
    user := User{
        Name:  "Alice",
        Email: "alice@example.com",
        Profile: Profile{
            Bio: "Software developer",
        },
        Posts: []Post{
            {Title: "First Post", Content: "Hello World!"},
            {Title: "Second Post", Content: "Another post"},
        },
    }
    
    result := db.Create(&user)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
    
    // Load user with associations
    var loadedUser User
    result = db.Preload("Profile").Preload("Posts").First(&loadedUser, user.ID)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
    
    fmt.Printf("User with profile and posts: %+v\n", loadedUser)
    fmt.Printf("Profile: %+v\n", loadedUser.Profile)
    for _, post := range loadedUser.Posts {
        fmt.Printf("Post: %+v\n", post)
    }
}
```

## Database Drivers

### Popular Database Drivers

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
)

// PostgreSQL
func connectPostgreSQL() (*sql.DB, error) {
    return sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
}

// MySQL
func connectMySQL() (*sql.DB, error) {
    return sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname?parseTime=true")
}

// SQLite
func connectSQLite() (*sql.DB, error) {
    return sql.Open("sqlite3", "./database.db")
}

// SQL Server
func connectSQLServer() (*sql.DB, error) {
    return sql.Open("sqlserver", "server=localhost;user id=username;password=password;database=dbname")
}

func main() {
    // Import drivers
    _ "github.com/lib/pq"           // PostgreSQL
    _ "github.com/go-sql-driver/mysql" // MySQL
    _ "github.com/mattn/go-sqlite3"    // SQLite
    _ "github.com/denisenkom/go-mssqldb" // SQL Server
    
    // Connect to PostgreSQL
    db, err := connectPostgreSQL()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    fmt.Println("Connected to PostgreSQL")
}
```

## Best Practices

### Error Handling and Logging

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"
    
    _ "github.com/lib/pq"
)

type DatabaseService struct {
    db *sql.DB
}

func NewDatabaseService(dsn string) (*DatabaseService, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }
    
    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    // Test connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := db.PingContext(ctx); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }
    
    return &DatabaseService{db: db}, nil
}

func (s *DatabaseService) Close() error {
    return s.db.Close()
}

func (s *DatabaseService) GetUserWithContext(ctx context.Context, id int) (User, error) {
    query := `SELECT id, name, email FROM users WHERE id = $1`
    
    var user User
    err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, fmt.Errorf("user not found: %w", err)
        }
        return User{}, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}

func (s *DatabaseService) CreateUserWithRetry(ctx context.Context, name, email string) (int, error) {
    const maxRetries = 3
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        id, err := s.createUser(ctx, name, email)
        if err != nil {
            lastErr = err
            log.Printf("Attempt %d failed: %v", i+1, err)
            time.Sleep(time.Duration(i+1) * time.Second)
            continue
        }
        return id, nil
    }
    
    return 0, fmt.Errorf("failed after %d attempts: %w", maxRetries, lastErr)
}

func (s *DatabaseService) createUser(ctx context.Context, name, email string) (int, error) {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    
    var id int
    err := s.db.QueryRowContext(ctx, query, name, email).Scan(&id)
    if err != nil {
        return 0, fmt.Errorf("failed to create user: %w", err)
    }
    
    return id, nil
}

func main() {
    service, err := NewDatabaseService("postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer service.Close()
    
    // Use service with context
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    user, err := service.GetUserWithContext(ctx, 1)
    if err != nil {
        log.Printf("Error getting user: %v", err)
        return
    }
    
    fmt.Printf("User: %+v\n", user)
    
    // Create user with retry
    id, err := service.CreateUserWithRetry(ctx, "Bob", "bob@example.com")
    if err != nil {
        log.Printf("Error creating user: %v", err)
        return
    }
    
    fmt.Printf("Created user with ID: %d\n", id)
}
```

### Database Migrations

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
)

type Migration struct {
    ID   int
    Name string
    SQL  string
}

func runMigrations(db *sql.DB) error {
    // Create migrations table if it doesn't exist
    createMigrationsTable := `
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `
    
    _, err := db.Exec(createMigrationsTable)
    if err != nil {
        return fmt.Errorf("failed to create migrations table: %w", err)
    }
    
    // Define migrations
    migrations := []Migration{
        {
            Name: "create_users_table",
            SQL: `
                CREATE TABLE IF NOT EXISTS users (
                    id SERIAL PRIMARY KEY,
                    name VARCHAR(255) NOT NULL,
                    email VARCHAR(255) UNIQUE NOT NULL,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                )
            `,
        },
        {
            Name: "create_profiles_table",
            SQL: `
                CREATE TABLE IF NOT EXISTS profiles (
                    id SERIAL PRIMARY KEY,
                    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                    bio TEXT,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                )
            `,
        },
        {
            Name: "add_index_to_users_email",
            SQL: `CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)`,
        },
    }
    
    // Run migrations
    for _, migration := range migrations {
        // Check if migration already executed
        var count int
        err := db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = $1", migration.Name).Scan(&count)
        if err != nil {
            return fmt.Errorf("failed to check migration %s: %w", migration.Name, err)
        }
        
        if count > 0 {
            log.Printf("Migration %s already executed, skipping", migration.Name)
            continue
        }
        
        // Execute migration
        _, err = db.Exec(migration.SQL)
        if err != nil {
            return fmt.Errorf("failed to execute migration %s: %w", migration.Name, err)
        }
        
        // Record migration
        _, err = db.Exec("INSERT INTO migrations (name) VALUES ($1)", migration.Name)
        if err != nil {
            return fmt.Errorf("failed to record migration %s: %w", migration.Name, err)
        }
        
        log.Printf("Migration %s executed successfully", migration.Name)
    }
    
    return nil
}

func main() {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    err = runMigrations(db)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Migrations completed successfully")
}
```

## Summary

Database interaction in Go provides:

- **Standard interface**: `database/sql` package for SQL databases
- **Multiple drivers**: Support for PostgreSQL, MySQL, SQLite, SQL Server
- **Connection pooling**: Efficient connection management
- **Transaction support**: ACID compliance
- **ORM support**: GORM for object-relational mapping
- **Prepared statements**: Security and performance
- **Context support**: Timeout and cancellation

Key points to remember:
1. Always handle errors properly
2. Use connection pooling for production applications
3. Implement proper transaction management
4. Use prepared statements to prevent SQL injection
5. Set appropriate timeouts with context
6. Implement retry logic for transient failures
7. Use migrations for database schema management
8. Monitor and log database operations
9. Choose the right tool for the job (raw SQL vs ORM)
10. Implement proper cleanup and resource management

Understanding database interaction enables you to build robust, scalable applications with persistent data storage. 