# Go Social Media Project

This project is a simple social media application implemented in Go using the Gin framework and follows the principles of clean architecture.

## Project Structure

The project is organized using the following directory structure:

- `cmd/`: Contains the main application entry point.
- `internal/`: Contains the internal application code.
  - `app/`: Houses the core business logic and use cases.
  - `model/`: Defines the data models for users, posts, and comments.
  - `persistence/`: Manages the interaction with the data storage (in-memory database).
  - `router/`: Configures the application's HTTP routes using the Gin framework.
- `pkg/`: Houses any shared packages or utilities used across the application.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-social-media.git
   cd go-social-media
