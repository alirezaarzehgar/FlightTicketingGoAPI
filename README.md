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
  ├── database/
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

### Fetch a user 

- **Endpoint**: `/v1/users/:id`
- **Method**: `GET`

**Request Body**:

**Response Body**:
```json
{
  "email": "user@example.com",
  "role": "passengar"
}
```

### Get all users

- **Endpoint**: `/v1/users`
- **Method**: `GET`
- **Admin Only**

**Response Body**:
```json
{
  "users": [
    {
      "email": "user@example.com",
      "role": "passengar"
    },
    // more users
  ]
}
```

### Update user

- **Endpoint**: `/v1/users/:id`
- **Method**: `PUT`
- **Admin Only**

**Request Body**:

```json
{
  "email": "user@example.com",
  "password": "password",
  "role": "employee"
}
```

**Response Body**:
```json
{
  "message": "Updated successfully"
}
```

### Delete user

- **Endpoint**: `/v1/users/:id`
- **Method**: `DELETE`
- **Admin Only**

**Response Body**:
```json
{
  "message": "Deleted successfully"
}
```

### Register new flight

- **Endpoint**: `/v1/flights`
- **Method**: `POST`
- **Admin Only**

**Request Body**:

```json
{
  "origin_id": "1",
  "destination_id": "2",
  "departure_date": "2023-04-17T7:00:00Z",
  "arrival_date": "2023-04-17T21:00:00Z",
  "price": 100.0
}
```

**Response Body**:
```json
{
  "id": 1,
  "origin_id": "1",
  "destination_id": "2",
  "departure_date": "2023-04-17T7:00:00Z",
  "arrival_date": "2023-04-17T21:00:00Z",
  "price": 100.0
}
```


### Flight Search

- **Endpoint**: `/v1/flights`
- **Method**: `GET`
- **Request Query Parameters**:

- origin_id (required): The origin_id airport code (e.g. 1)
- destination_id (required): The destination_id airport code (e.g. 2)
- departure_date (required): The departure date in YYYY-MM-DD format

- **Response Body**:

```json
[
  {
    "id": 1,
    "origin_id": "1",
    "destination_id": "2",
    "departure_date": "2023-05-01",
    "departure_time": "09:00",
    "arrival_date": "2023-05-01",
    "arrival_time": "17:00",
    "price": 100.00
  },
  {
    "id": 2,
    "origin_id": "1",
    "destination_id": "2",
    "departure_date": "2023-05-01",
    "departure_time": "13:00",
    "arrival_date": "2023-05-01",
    "arrival_time": "21:00",
    "price": 150.00
  }
]
```

### Get flight by id

- **Endpoint**: `/v1/flights/:id`
- **Method**: `GET`
- **Admin Only**

**Response Body**:
```json
  {
    "id": 1,
    "origin_id": "1",
    "destination_id": "2",
    "departure_date": "2023-05-01",
    "departure_time": "09:00",
    "arrival_date": "2023-05-01",
    "arrival_time": "17:00",
    "price": 100.00
  },
```

### Update flight

- **Endpoint**: `/v1/`
- **Method**: ``
- **Admin Only**

**Request Body**:

```json
  {
    "id": 1,
    "origin_id": "1",
    "destination_id": "2",
    "departure_date": "2023-05-01",
    "departure_time": "09:00",
    "arrival_date": "2023-05-01",
    "arrival_time": "17:00",
    "price": 100.00
  },
```

**Response Body**:
```json
{
  "message": "Updated successfully"
}
```

### Delete flight by id

- **Endpoint**: `/v1/flights/:id`
- **Method**: `DELETE`
- **Admin Only**

**Response Body**:
```json
{
  "message": "Deleted successfully"
}
```

#### Flight Booking

- **Endpoint**: `/v1/flights/:flight_id/booking`
- **Method**: `POST`
- **Request Body**:

```json
{
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

### Get all tickets of logged in user

- **Endpoint**: `/v1/tickets`
- **Method**: `GET`

**Response Body**:
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

### Get ticket by id

- **Endpoint**: `/v1/ticket/:id`
- **Method**: `GET`

**Response Body**:
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

### Update ticket

- **Endpoint**: `/v1/ticket/:id`
- **Method**: `PUT`

**Request Body**:

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

**Response Body**:
```json
{
  "message": "Updated successfully"
}
```

### Delete ticket by id

- **Endpoint**: `/v1/tickets/:id/cancel`
- **Method**: `DELETE`
- **Admin Only**

**Response Body**:
```json
{
  "message": "Deleted successfully"
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

Copyright 2023, Alireza Arzehgar
