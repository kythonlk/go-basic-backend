# Go CockroachDB Backend Project Structure

## Technologies
- **Backend**: Go, Chi
- **Database**: CockroachDB, pgx

## Database Models

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    last_login TIMESTAMP,
    reset_token TEXT,
    reset_token_expiry TIMESTAMP,
    mfa_secret TEXT,
    role VARCHAR(50) DEFAULT 'user'
);
```

### JWT Token Management Table
```sql
CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    is_valid BOOLEAN DEFAULT TRUE,
    user_agent TEXT,
    ip_address VARCHAR(45)
);
```

### Roles and Permissions
```sql
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE user_roles (
    user_id INT REFERENCES users(id),
    role_id INT REFERENCES roles(id),
    PRIMARY KEY(user_id, role_id)
);
```

### Audit Logs
```sql
CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    action VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    ip_address VARCHAR(45),
    user_agent TEXT
);
```

## JWT Authentication Workflow
1. **User Registration**
   - User submits email and password.
   - Store hashed password and send verification email.

2. **User Login**
   - Validate credentials and generate JWT access and refresh tokens.
   - Store refresh token in the database.

3. **Token Flow**
   - Use short-lived access tokens; refresh tokens for obtaining new access tokens.

4. **Password Reset**
   - Generate and email reset token; update password upon successful reset.

## Project Structure
```
/root
│
├── /cmd                    # Main application and migration commands
│   ├── /app                # Entry point for the application
│   │   └── main.go        
│   └── /migrate            # Database migration tool
│       └── migrate.go      
│
├── /config                 # Configuration handling
│   └── config.go           
│
├── /controllers            # HTTP request handlers
│   ├── user_controller.go  
│   ├── auth_controller.go  
│   └── health_controller.go
│
├── /models                 # Database models
│   ├── user.go             
│   └── base.go             
│
├── /routes                 # API route definitions
│   └── routes.go           
│
├── /services               # Business logic
│   ├── auth_service.go     
│   └── user_service.go     
│
├── /middleware             # Middleware for handling requests
│   ├── auth_middleware.go  
│   └── logging.go          
│
├── /utils                  # Utility functions
│   ├── jwt.go              
│   └── hash.go             
│
├── /migrations             # Database migration scripts
│   └── 20230928_create_user_table.sql
│
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── README.md               # Project documentation
```
