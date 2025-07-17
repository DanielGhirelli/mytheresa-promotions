# mytheresa Promotions API

## Description
A REST API that applies discounts to fashion products based on specific rules and provides filtering capabilities. The API processes a catalog of product and applies the following discount logic:

- `30% discount` for all products in the `"boots"` category
- `15% discount` for the product with SKU `"000003"`
- When multiple discounts apply, the `highest discount` is selected

## Technical Stack
- **Language**: Go 1.23
- **Containerization**: Docker
- **Testing**: Native Go testing package

## Prerequisites
Before running this project, ensure you have installed:

- [**Go 1.23+**](https://go.dev/dl/) - Required to build and run the application
- [**Docker**](https://www.docker.com/) - For containerized deployment (optional)

## Setup

### 1. Environment Configuration
Create an `.env` file in the project root with these variables:

```env
SERVER_PORT=8080
DATA_FILE_PATH=data/products.json
API_KEY=4075ea28e50c84f01937b136fc81f182
```

### 2. Install Dependencies
```go
go mod download
go mod vendor
```

### 3. Prepare Sample Data
Create a `data/products.json` file with your product catalog:
```json
{
  "products": [
    {
      "sku": "000001",
      "name": "Leather Boots",
      "category": "boots",
      "price": 99900
    },
    {
      "sku": "000003",
      "name": "Special Sandals",
      "category": "sandals",
      "price": 79500
    }
    ...
  ]
}
```

## Running the Application
### Option 1: Local Execution
```go
go run main.go
```

### Option 2: Docker Compose
```go
docker-compose up
```

The API will be available at http://localhost:8080 (for example)


## API Endpoints
### GET /products
Retrieve products with applied discounts

#### Query Parameters

| Parameter      | Type    | Description                                     |
|---------------|---------|--------------------------------------------------|
| `category`    | string  | Filter by product category (e.g. "boots")        |
| `priceLessThan` | integer | Filter by original price (before discounts) in cents (e.g. 80000 → €800) |

#### Header Parameters

| Parameter      | Type    | Description                                                                       |
|---------------|---------|------------------------------------------------------------------------------------|
| `x-api-key`    | string  | Simple Authentication Validation. Value must be set under the `.env` → `API_KEY` |


#### Examples: 

```
# All products
GET /products
Headers: { x-api-key: 4075ea28e50c84f01937b136fc81f182 }

# Boots category
GET /products?category=boots
Headers: { x-api-key: 4075ea28e50c84f01937b136fc81f182 }

# Products under €800
GET /products?priceLessThan=80000
Headers: { x-api-key: 4075ea28e50c84f01937b136fc81f182 }
```

#### Response:

```json
{
  "products": [
    {
        "sku": "000001",
        "name": "BV Lean leather ankle boots",
        "category": "boots",
        "price": {
            "original": 89000,
            "final": 62300,
            "discount_percentage": "30%",
            "currency": "EUR"
        }
    },
    {
        "sku": "000003",
        "name": "Ashlington leather ankle sneakers",
        "category": "sneakers",
        "price": {
            "original": 71000,
            "final": 60350,
            "discount_percentage": "15%",
            "currency": "EUR"
        }
    },
    {
        "sku": "000004",
        "name": "Naima embellished suede sandals",
        "category": "sandals",
        "price": {
            "original": 79500,
            "final": 79500,
            "discount_percentage": null,
            "currency": "EUR"
        }
    },
    ...
  ]
}
```

## Project Structure
```
├── config/             # Configuration management
├── data/               # Product data file
├── http/               # Server and Middleware utilities
├── pkg/                # Reusable packages
│   └── application/    # Use cases
│   └── domain/         # Entities, Value Objects, Domain Services
│   └── repo/           # Product repository implementation
│   └── service/        # Product service implementation
├── .env                # Environmental file
├── go.mod              # Dependency management
├── Dockerfile          # Service orchestration
└── docker-compose.yml  # Service orchestration
```

| Layer                    | Folder                                   | Purpose                                                  |
| ------------------------ | ---------------------------------------- | -------------------------------------------------------- |
| **Domain Layer**         | `pkg/domain/product/`                    | Core business rules and logic: `Product`, `Price`, `Discount` |
| **Application Layer**    | `pkg/application/`                       | Input orchestration (HTTP calls) |
| **Infrastructure Layer** | `pkg/repo/`, `http/`, `config/`, `data/` | File system repo, HTTP server, middleware, environment loading  |
| **Service Layer**        | `pkg/service/`                           | Application logic       |


## Testing
Run unit tests with coverage:

```go
go test -cover ./...
```
