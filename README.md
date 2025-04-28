# Go MySQL Learning Project

A learning project that demonstrates the integration of MySQL with three different Go API frameworks: HTTP, Echo, and Gin. This project includes a web interface to test all three APIs and showcases secure database connections, CORS implementation, and proper error handling.

## Features

- Three API implementations (HTTP, Echo, Gin) for the same functionality
- Secure database connection using environment variables
- CORS implementation for all APIs
- Modern web interface for testing
- Comprehensive test coverage
- Environment variable configuration

## Prerequisites

- Go 1.16 or higher
- MySQL 8.0 or higher
- Access to the `world` database (MySQL sample database)

## Project Structure

```
.
├── db/                 # Database connection and queries
├── http_api/          # Standard HTTP API implementation
├── echo_api/          # Echo framework API implementation
├── gin_api/           # Gin framework API implementation
├── static/            # Static files (HTML, CSS, JS)
├── static_server/     # Static file server
└── .env               # Environment variables (not committed)
```

## Setup

1. **Clone the repository**
   ```bash
   git clone [https://github.com/AkshDesai04/GoMySQLLearning/](https://github.com/AkshDesai04/GoMySQLLearning.git)
   cd GoMySQLLearning
   ```

2. **Initialize Go module**
   ```bash
   go mod init github.com/AkshDesai04/GoMySQLLearning
   ```

3. **Install required packages**
   ```bash
   go get github.com/go-sql-driver/mysql
   go get github.com/gin-gonic/gin
   go get github.com/labstack/echo/v4
   go get github.com/joho/godotenv
   ```

4. **Install dependencies**
   ```bash
   go mod download
   go mod tidy
   ```

5. **Create .env file**
   Create a `.env` file in the project root with the following content:
   ```
   DB_USER=your_id
   DB_PASSWORD=your_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=world
   ```

6. **Set up MySQL**
   - Install MySQL if not already installed
   - Import the `world` database
   - Ensure the database user has proper permissions

## Running the Project

1. **Start the HTTP API Server** (Port 8080)
   ```bash
   go run http_api/main.go
   ```

2. **Start the Echo API Server** (Port 8081)
   ```bash
   go run echo_api/main.go
   ```

3. **Start the Gin API Server** (Port 8082)
   ```bash
   go run gin_api/main.go
   ```

4. **Start the Static File Server** (Port 8083)
   ```bash
   go run static_server/main.go
   ```

5. **Access the Web Interface**
   Open your browser and navigate to:
   ```
   http://localhost:8083
   ```

## Testing

### Running Tests

To run all tests:
```bash
go test ./http_api ./echo_api ./gin_api -v
```

To run tests for a specific API:
```bash
go test ./http_api -v
go test ./echo_api -v
go test ./gin_api -v
```

### API Testing with Postman

All APIs support the following endpoint:
- **URL:** `http://localhost:{port}/cities`
- **Method:** GET
- **Query Parameter:** `search` (optional)

Port numbers:
- HTTP API: 8080
- Echo API: 8081
- Gin API: 8082

Example request:
```
GET http://localhost:8080/cities?search=London
```

### Web Interface Testing

1. Open `http://localhost:8083` in your browser
2. Enter a search term in the text field
3. Select an API using the radio buttons
4. Click "Run" to execute the query
5. Results will be displayed below
6. Use "Clear" to reset the form

## Security Considerations

- Database credentials are stored in `.env` file
- `.env` is included in `.gitignore` to prevent accidental commits
- CORS is properly configured for all APIs
- Input validation is implemented for search queries

## Contributing

Feel free to submit issues and enhancement requests!
