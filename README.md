# Tech Stack

- **Go**: Programming language ([https://golang.org/](https://golang.org/))
- **Fiber**: Web framework ([https://gofiber.io/](https://gofiber.io/))
- **GORM**: ORM library ([https://gorm.io/](https://gorm.io/))
- **PostgreSQL**: Database ([https://www.postgresql.org/](https://www.postgresql.org/))
- **Swagger**: API documentation ([https://swagger.io/](https://swagger.io/))
- **Docker**: Containerization ([https://www.docker.com/](https://www.docker.com/))

## Project Structure

The project follows Clean Architecture principles:

- **client**: External service clients
- **configuration**: Application and database configuration
- **controller**: HTTP request handlers
- **docs**: Swagger documentation
- **exception**: Custom error handling
- **repository**: Data access layer
- **service**: Business logic layer

## Getting Started

### Prerequisites

- **Go 1.16+**
- **Docker** and **Docker Compose**

### Running the Application

1. Clone the repository:

```bash
git clone https://github.com/RianAsmara/personal-finance-advisor-api.git
cd family-tree-be
```

2. Build and run the application using Docker Compose:

```bash
make run
```

3. The API will be available at `http://localhost:9999`

### Running Tests

To run the tests:

```bash
make test
```
