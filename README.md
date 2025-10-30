# TODO App in Go

A simple, full-stack TODO application built with Go for the backend, MySQL for data persistence, and vanilla JavaScript for the frontend. This app allows users to create, read, update, and delete TODO items with features like progress tracking, start/end times, and completion status.

## Features

- **CRUD Operations**: Create, read, update, and delete TODO items.
- **Progress Tracking**: Manual progress input or automatic calculation based on start/end times.
- **Time Management**: Set start and end times for tasks.
- **Responsive UI**: Clean, mobile-friendly interface.
- **Real-time Updates**: Progress bars update automatically if times are set.
- **CORS Support**: Configured for development with cross-origin requests.

## Prerequisites

Before running this application, ensure you have the following installed:

- **Go**: Version 1.16 or later (tested with 1.25.1).
- **MySQL**: A running MySQL server.
- **Git**: For cloning the repository.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/nallyhoo/todo_app_go.git
   cd todo_app_go
   ```

2. **Install Dependencies**:
   The project uses Go modules. Dependencies are already listed in `go.mod`. Run:
   ```bash
   go mod tidy
   ```

3. **Set Up the Database**:
   - Create a MySQL database named `todo_db`.
   - Update the DSN in `db/db.go` with your MySQL credentials:
     ```go
     DSN := "your_username:your_password@tcp(127.0.0.1:3306)/todo_db?parseTime=true"
     ```
   - The app will automatically create the `todos` table on first run if it doesn't exist. (Note: Table creation is handled in the handlers; ensure your user has CREATE privileges.)

4. **Build and Run**:
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`.

## Usage

1. Open your browser and navigate to `http://localhost:8080`.
2. Use the form to add new TODO items:
   - **Title**: Required field.
   - **Description**: Optional.
   - **Start Time / End Time**: Optional datetime fields.
   - **Progress**: Manual percentage (0-100) or leave blank for auto-calculation.
   - **Completed**: Checkbox to mark as done.
3. View the list of TODOs below the form.
4. Click "Edit" to modify an existing TODO or "Delete" to remove it.
5. Progress bars update in real-time if start/end times are set.

## API Endpoints

The backend provides a RESTful API for TODO management:

- `GET /todos`: Retrieve all TODOs.
- `GET /todos/{id}`: Retrieve a specific TODO by ID.
- `POST /todos`: Create a new TODO.
- `PUT /todos/{id}`: Update an existing TODO.
- `DELETE /todos/{id}`: Delete a TODO by ID.

All endpoints return JSON responses. The frontend interacts with these via AJAX.

## Project Structure

```
todo-app/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum
├── db/
│   └── db.go               # Database initialization
├── models/
│   └── todo.go             # Todo model and JSON handling
├── handlers/
│   └── todo_handlers.go    # HTTP handlers for CRUD operations
├── routes/
│   └── routes.go           # Router setup and middleware
└── frontend/
    ├── index.html          # Main HTML page
    ├── styles.css          # CSS styles
    └── app.js              # Frontend JavaScript logic
```

## Database Schema

The `todos` table structure:

```sql
CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    start_time TIMESTAMP NULL,
    end_time TIMESTAMP NULL,
    progress DECIMAL(5,2) DEFAULT 0.00
);
```

## Contributing

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Make your changes and commit: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Open a pull request.

## License

This project is open-source and available under the [MIT License](LICENSE).

## Acknowledgments

- Built with [Gorilla Mux](https://github.com/gorilla/mux) for routing.
- MySQL driver: [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql).
- Inspired by simple TODO apps for learning full-stack development.
