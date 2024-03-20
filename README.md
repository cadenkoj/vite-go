# React + Vite + Echo + Go

This is a full-stack web application built with React, Vite, Echo, and Go. It combines the frontend framework React with the lightning-fast build tool Vite for efficient development. The backend is powered by Echo, a high-performance web framework for Go.

## Features

- **React**: A popular JavaScript library for building user interfaces.
- **Vite**: A fast, opinionated web dev build tool that serves your code via native ES Module imports during development.
- **Echo**: A high-performance, minimalist web framework for Go.
- **Go**: A statically typed, compiled programming language designed for building simple, reliable, and efficient software.

## Prerequisites

Before running this project, ensure you have the following installed:

- Node.js
- Go
- PNPM (Package Manager)
- Vite

## Getting Started

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/cadenkoj/vite-go
   ```

2. Navigate to the project directory:

   ```bash
   cd vite-go
   ```

3. Install dependencies:

   ```bash
   cd frontend
   pnpm install
   cd ../backend
   go mod tidy
   ```

### Development

To start the development servers for both frontend and backend, follow these steps:

```bash
cd backend
air
```

This will start the Vite development alongside the Echo server.

Now you should be able to access your application at `http://localhost:8080` in your web browser.

## Deployment

### Building

To build the production-ready version of the frontend, run:

```bash
cd frontend
pnpm build
```

This will create an optimized build of your React application in the `backend/web/static` directory.

### Deployment Options

You can deploy your application using various platforms like Railway, AWS, or Vercel. Ensure that you configure your deployment environment according to your specific needs.

## Contributing

Contributions are welcome! Please fork this repository and create a pull request with your changes.

## License

This project is licensed under the [Apache 2.0](LICENSE).
