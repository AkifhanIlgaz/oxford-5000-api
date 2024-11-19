# Oxford-5000 API

A RESTful API service providing access to Oxford 5000 word definitions, including meanings, examples, and audio pronunciations. The API features user authentication and API key management for secure access.

## Features

- Word lookup with detailed information
- User authentication system
- API key management
- Swagger documentation
- CEFR level indicators
- Audio pronunciations (UK & US)

## API Endpoints

### Authentication

#### Register
- **POST** `/api/auth/register`
- Register a new user account
- Body:
```
{
    "email": "user@example.com",
    "password": "password123"
}
```

#### Login
- **POST** `/api/auth/login`
- Login with existing credentials
- Body:
```
{
    "email": "user@example.com",
    "password": "password123"
}
```

#### Refresh Token
- **POST** `/api/auth/refresh`
- Get new access token using refresh token
- Body:
```
{
    "refreshToken": "your-refresh-token"
}
```

### API Keys

#### Create API Key
- **POST** `/api/user/api-key`
- Creates new API key for authenticated user
- Required: Bearer token authentication
- Query Parameters:
  - `name`: Name for the API key

#### Get API Key
- **GET** `/api/user/api-key`
- Retrieves user's API key
- Required: Bearer token authentication

#### Delete API Key
- **DELETE** `/api/user/api-key`
- Deletes user's API key
- Required: Bearer token authentication

### Word Information

#### Get Word Details
- **GET** `/api/word/{word}`
- Retrieves detailed information about a word
- Optional Query Parameters:
  - `part_of_speech`: Filter by grammatical category (noun, verb, adjective, etc.)
- Response includes:
  - Definitions
  - Examples
  - CEFR Level
  - Part of speech
  - Audio pronunciations (UK & US)
  - Related idioms

## Response Format

All API responses follow a standard format:

```json
{
    "success": boolean,
    "message": "Human-readable message",
    "error": "Error details if success is false",
    "data": {
        // Response payload
    }
}
```

## Authentication

The API uses JWT (JSON Web Token) authentication:

1. Register/Login to get access and refresh tokens
2. Include access token in requests:
   ```
   Authorization: Bearer <your-access-token>
   ```
3. Use refresh token to get new access token when expired

## Development Setup

1. Clone the repository
2. Create `dev.env` and `prod.env` files with required configurations:
   ```env
   MONGO_URI=your_mongodb_uri
   PORT=8080
   ACCESS_TOKEN_PUBLIC_KEY=
   ACCESS_TOKEN_PRIVATE_KEY=
   ACCESS_TOKEN_EXPIRY_HOUR=1
   REFRESH_TOKEN_PUBLIC_KEY=
   REFRESH_TOKEN_PRIVATE_KEY=
   REFRESH_TOKEN_EXPIRY_HOUR=24
   ```
3. Run the application:
   ```bash
   # Development mode
   go run main.go -mode=dev

   # Production mode
   go run main.go -mode=prod
   ```

## API Documentation

Swagger documentation is available at `/api/swagger/*` when running the server.

## Technologies

- Go 1.22+
- Gin Web Framework
- MongoDB
- JWT Authentication
- Swagger/OpenAPI
```
