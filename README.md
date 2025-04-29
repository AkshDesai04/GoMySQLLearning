# Go MySQL Learning Project

A learning project that demonstrates the integration of MySQL with three different Go API frameworks: HTTP, Echo, and Gin. This project includes a web interface to test all three APIs and showcases secure database connections, CORS implementation, and proper error handling.

## Features

- Three API implementations (HTTP, Echo, Gin) for the same functionality
- Secure database connection using environment variables
- CORS implementation for all APIs
- Modern web interface for testing
- Comprehensive test coverage
- Environment variable configuration
- Interactive table sorting with merge sort implementation
  - Sort by any column (ID, Name, Country Code, District, Population)
  - Toggle between ascending and descending order
  - Visual indicators for current sort column and direction

## Prerequisites

- Go 1.16 or higher
- MySQL 8.0 or higher
- Access to the `world` database (MySQL sample database)

## Project Structure

```
.
├── db/                 # Database connection and queries
│   ├── db.go          # Database operations
│   └── sort.go        # Merge sort implementation
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
   git clone <repository-url>
   cd GoMySQLLearning
   ```

2. **Initialize Go module**
   ```bash
   go mod init github.com/yourusername/GoMySQLLearning
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
go test ./... -v
```

To run tests for specific components:
```bash
# Database and sorting tests
go test ./db -v

# HTTP API tests
go test ./http_api -v

# Echo API tests
go test ./echo_api -v

# Gin API tests
go test ./gin_api -v
```

### Test Coverage

To generate test coverage reports:
```bash
# Generate coverage profile
go test ./... -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

### Sorting Tests

The project includes comprehensive tests for the merge sort implementation:
- Tests for all sortable fields (ID, Name, Country Code, District, Population)
- Tests for both ascending and descending order
- Tests for edge cases (empty slice, single element)
- Tests for stability (maintaining order of equal elements)

### API Testing with Postman

All APIs support the following endpoint:
- **URL:** `http://localhost:{port}/cities`
- **Method:** GET
- **Query Parameters:**
  - `search` (optional): Filter cities by name, country code, or district
  - `sort_field` (optional): Field to sort by (id, name, country_code, district, population)
  - `sort_dir` (optional): Sort direction (asc, desc)

Port numbers:
- HTTP API: 8080
- Echo API: 8081
- Gin API: 8082

Example requests:
```
# Basic request
GET http://localhost:8080/cities

# With search
GET http://localhost:8080/cities?search=London

# With sorting
GET http://localhost:8080/cities?sort_field=population&sort_dir=desc

# Combined search and sort
GET http://localhost:8080/cities?search=London&sort_field=population&sort_dir=desc
```

### Web Interface Testing

1. Open `http://localhost:8083` in your browser
2. Enter a search term in the text field
3. Select an API using the radio buttons
4. Click "Run" to execute the query
5. Click on any column header to sort by that field
6. Click again to toggle between ascending and descending order
7. Use "Clear" to reset the form

## Security Considerations

- Database credentials are stored in `.env` file
- `.env` is included in `.gitignore` to prevent accidental commits
- CORS is properly configured for all APIs
- Input validation is implemented for search queries and sort parameters

## Contributing

Feel free to submit issues and enhancement requests!
