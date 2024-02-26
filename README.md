# go-cms

Creating a web application that simplifies backend development by allowing users to define API endpoints, request methods, and database schemas through a web interface is an ambitious and exciting project. Using Go for the backend, HTMX for dynamic HTML content without writing JavaScript, and SQLite as the database is a solid choice for building a lightweight, efficient, and easy-to-deploy application.

This project involves creating a user-friendly interface where users can define new API endpoints, specify the request methods (GET, POST, etc.), and outline the database schema (fields and types) for each endpoint. Essentially, you're building a CMS-like system tailored for API management and data storage, which could significantly speed up development processes for small to medium projects or prototyping phases.

### Project Structure
```plaintext
Go-cms/
├── cmd/
│   └── server/
│       └── main.go        # Entry point to start the server
├── internal/
│   ├── api/
│   │   ├── handler.go     # Request handlers for the API and dynamic endpoints
│   │   └── router.go      # Setup routes and dynamic endpoint registration
│   ├── db/
│   │   ├── sqlite.go      # SQLite database operations
│   │   └── schema.go      # Schema definition and management
│   └── ui/
│       ├── htmx.go        # HTMX dynamic content handlers
│       └── templates/     # HTML templates for the UI
├── static/
│   ├── css/
│   └── js/
├── go.mod
├── go.sum
└── README.md
```

### Project Overview

#### Technologies:
- **Backend**: Go (Golang)
- **Frontend**: HTML enhanced with HTMX for dynamic interactions
- **Database**: SQLite

#### Core Features:
- GUI for creating and managing API endpoints.
- Functionality to define request methods (GET, POST) for each endpoint.
- Ability to specify database schema directly from the GUI.
- Auto-generation of database tables based on the defined schema.
- CRUD operations on the defined endpoints and associated database records.

### Suggested Architecture

#### Backend (Go):
- **API Server**: Handles HTTP requests, serves the frontend, and interacts with the SQLite database.
- **Dynamic Endpoint Creation**: Dynamically register user-defined endpoints and their methods at runtime.
- **Database Management**: Interface for executing schema changes and performing CRUD operations based on user input.

#### Frontend (HTML + HTMX):
- **Endpoint Management UI**: Allows users to create, modify, and delete API endpoints.
- **Schema Definition UI**: Enables users to define and edit the schema (fields and data types) for each endpoint's database table.
- **Dynamic Content Loading**: Use HTMX for partial page updates when creating or modifying endpoints and schemas, providing a smooth user experience without full page reloads.

#### Database (SQLite):
- **Schema Storage**: Stores the schema definitions for each user-created endpoint.
- **Data Storage**: Maintains data for each endpoint according to its defined schema.
