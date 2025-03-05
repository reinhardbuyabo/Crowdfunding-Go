# Crowdfunding Microservice API

## Table of Contents
- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Database Setup](#database-setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Example Requests](#example-requests)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Overview

This Crowdfunding Microservice is a Go-based API for managing crowdfunding campaigns. It provides endpoints to create and retrieve campaigns using PostgreSQL as the database.

## Prerequisites

- Go 1.20 or higher
- PostgreSQL 12+ 
- Git
- Postman or cURL for API testing

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/crowdfunding-microservice.git
cd crowdfunding-microservice
```

2. Install dependencies:
```bash
go mod download
```

## Database Setup

1. Create PostgreSQL Database:
```sql
CREATE DATABASE crowdfunding;
USE crowdfunding;

CREATE TABLE campaigns (
    id SERIAL PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    target DECIMAL(18,2) NOT NULL,
    deadline TIMESTAMP NOT NULL,
    amount_collected DECIMAL(18,2) DEFAULT 0,
    image VARCHAR(512)
);
```

2. Update Connection String:
In `main.go`, modify the connection string:
```go
connectionString := "postgresql://username:password@localhost:5432/crowdfunding?sslmode=disable"
```

## Running the Application

1. Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Create Campaign
- **URL:** `/campaigns`
- **Method:** `POST`
- **Content-Type:** `application/json`

#### Request Body
```json
{
    "owner": "0x1234567890abcdef",
    "title": "Community Playground",
    "description": "Building a new playground for local children",
    "target": 50000.00,
    "deadline": "2024-12-31T23:59:59Z",
    "amountCollected": 0,
    "image": "https://example.com/playground.jpg"
}
```

#### Success Response
- **Code:** `201 Created`
- **Content:** Created Campaign Object with assigned ID

### Get Campaigns
- **URL:** `/campaigns`
- **Method:** `GET`

#### Success Response
- **Code:** `200 OK`
- **Content:** Array of Campaign Objects

## Example Requests

### Create Campaign (cURL)
```bash
curl -X POST http://localhost:8080/campaigns \
     -H "Content-Type: application/json" \
     -d '{
         "owner": "0x1234567890abcdef",
         "title": "Community Playground",
         "description": "Building a new playground for local children",
         "target": 50000.00,
         "deadline": "2024-12-31T23:59:59Z",
         "amountCollected": 0,
         "image": "https://example.com/playground.jpg"
     }'
```

### Get Campaigns (cURL)
```bash
curl http://localhost:8080/campaigns
```

## Error Handling

Possible Error Responses:
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database or server issues

Example Error Response:
```json
{
    "error": "Invalid campaign data"
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.