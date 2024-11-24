# Oxford 5000™ Dictionary API Platform

A full-stack application providing programmatic access to Oxford 5000 word definitions, examples, and audio pronunciations. The platform includes both a RESTful API service and a modern web interface.

## Project Structure

.
├── backend/ # Go API server
├── frontend/ # Next.js web application
└── README.md

## Features

### API Service

- Word lookup with detailed information
  - Comprehensive definitions and examples
  - CEFR level indicators
  - Audio pronunciations (UK & US)
  - Part of speech classification
- User authentication system
- API key management
- Usage tracking and analytics
- Swagger documentation

### Web Interface

- Modern, responsive design
- Interactive API documentation
- User dashboard
- Authentication system
- API key management portal

## Technology Stack

### Backend

- Go 1.22+
- Gin Web Framework
- MongoDB
- JWT Authentication
- Swagger/OpenAPI

### Frontend

- Next.js 14
- TypeScript
- Tailwind CSS
- Radix UI Components
- React Hook Form

## Getting Started

### Backend Setup

See [Backend Setup Instructions](https://github.com/AkifhanIlgaz/oxford-5000-api/blob/main/backend/README.md) for detailed setup steps.
