# Flight-Ticketing Go API

Welcome to the Flight Ticketing Go API Assignment! In this assignment, you will be building a RESTful API for a Flight Ticketing System using the Go programming language. Your API will allow users to perform the following operations:

# Flight Ticketing Go API Assignment

Features:
- Register an account
- Login to an existing account
- Search for available flights
- Book a flight
- Retrieve a list of booked flights
- Cancel a booked flight

You will need to use the following technologies to build your API:

- Go programming language
- Gorilla/mux package for HTTP routing
- PostgreSQL for data storage
- JWT for authentication

## Project Structure

Your project should be structured as follows:

```go
FlightTicketingGoAPI/
  ├── api/
  │   ├── handlers/
  │   ├── middleware/
  │   └── routes/
  ├── config/
  ├── db/
  ├── models/
  ├── utils/
  ├── main.go
  └── README.md
```

- `api/handlers`: This directory should contain the HTTP handlers for each API endpoint.
- `api/middleware`: This directory should contain any middleware functions that you use to process requests before they reach the handlers.
- `api/routes`: This directory should contain the route definitions for your API endpoints.
- `config`: This directory should contain your application configuration files.
- `db`: This directory should contain your database migration scripts and any other database-related code.
- `models`: This directory should contain your database models.
- `utils`: This directory should contain any utility functions that you use throughout your application.
- `main.go`: This is the main entry point for your application.
- `README.md`: This is where you should document your API and provide instructions on how to run and test it.

## API Specification

Your API should conform to the following specifications:

- All requests and responses should be in JSON format.
- All endpoints should be versioned using a `/v1` prefix in the URL path.
- All requests that require authentication should include a **JWT** token in the Authorization header.
- Responses should include appropriate HTTP status codes and error messages.

Your API should provide the following endpoints:

### User Registration

- **Endpoint**: `/v1/register`
- **Method**: `POST`
- **Request Body**:

```json
{
  "email": "user@example.com",
  "password": "password"
}
```

- **Response Body**:

```json
{
  "id": 1,
  "email": "user@example.com"
}
```

### User Login

- **Endpoint**: `/v1/login`
- **Method**: `POST`
- **Request Body**:

```json
{
  "email": "user@example.com",
  "password": "password"
}
```

- **Response Body**:

```json
{
  "token": "<JWT token>"
}
```

### Flight Search

- **Endpoint**: `/v1/flights`
- **Method**: `GET`
- **Request Query Parameters**:

- origin (required): The origin airport code (e.g. LAX)
- destination (required): The destination airport code (e.g. JFK)
- departure_date (required): The departure date in YYYY-MM-DD format

- **Response Body**:

```json
[
  {
    "id": 1,
    "origin": "LAX",
    "destination": "JFK",
    "departure_date": "2023-05-01",
    "departure_time": "09:00",
    "arrival_date": "2023-05-01",
    "arrival_time": "17:00",
    "price": 100.00
  },
  {
    "id": 2,
    "origin": "LAX",
    "destination": "JFK",
    "departure_date": "2023-05-01",
    "departure_time": "13:00",
    "arrival_date": "2023-05-01",
    "arrival_time": "21:00",
    "price": 150.00
  }
]
```

#### Flight Booking

- **Endpoint**: `/v1/bookings`
- **Method**: `POST`
- **Request Body**:

```json
{
  "flight_id": 1,
  "passengers": [
    {
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com"
    }
  ]
}
```

- **Response Body**:

```json
{
  "id": 1,
  "user_id": 1,
  "flight_id": 1,
  "passengers": [
    {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com"
    }
  ],
  "total_price": 100.00,
  "booking_date": "2023-04-17T14:25:00Z"
}
```

#### List Booked Flights

- **Endpoint**: `/v1/bookings`
- **Method**: `GET`
- **Request Query Parameters**:
- `user_id` (optional): The ID of the user whose bookings should be retrieved
- **Response Body**:

```json
[
  {
    "id": 1,
    "user_id": 1,
    "flight_id": 1,
    "passengers": [
      {
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com"
      }
    ],
    "total_price": 100.00,
    "booking_date": "2023-04-17T14:25:00Z"
  }
]
```

#### Cancel Booking

- **Endpoint**: `/v1/bookings/:id`
- **Method**: `DELETE`
- **Request Parameters**:
- `id` (required): The ID of the booking to cancel
- **Response Body**:

```json
{
  "message": "Booking cancelled successfully"
}
```

### Authentication

Your API should use JWT for authentication. You should implement a middleware function that verifies the JWT token in the `Authorization` header before allowing access to protected endpoints.

### Data Storage

Your API should use PostgreSQL for data storage. You should use the `github.com/jackc/pgx` package for interacting with the database.

### Testing

You should write unit tests and integration tests for your API. You can use the `net/http/httptest` package for testing your API endpoints.

### Authors

- Alireza Arzehgar
- Max Base

Copyright 2023, Max Base
