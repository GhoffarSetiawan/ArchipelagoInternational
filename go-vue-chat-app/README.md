# Real-Time Chat Application

A real-time chat application built with Go Fiber backend and Vue.js 3 frontend, using WebSockets for communication.

## Features

- User registration with custom names
- Real-time list of online users
- Direct messaging between users
- Modern, responsive UI with two-panel interface

## Project Structure

```
go-vue-chat-app/
├── backend/               # Go Fiber backend
│   ├── main.go            # Main Go application
│   └── go.mod             # Go modules file
├── frontend/              # Vue.js 3 frontend
│   ├── public/            # Static assets
│   ├── src/               # Source code
│   │   ├── components/    # Vue components
│   │   ├── store/         # Pinia store
│   │   ├── App.vue        # Main Vue component
│   │   └── main.js        # Vue application entry point
│   ├── index.html         # HTML entry point
│   ├── package.json       # NPM dependencies
│   └── vite.config.js     # Vite configuration
└── README.md              # This file
```

## Prerequisites

- Go (1.16 or higher)
- Node.js (14.x or higher)
- npm (6.x or higher)

## Getting Started

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Run the Go application:
   ```bash
   go run main.go
   ```

   The backend server will start on port 8080.

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

   The frontend development server will start on port 5173 by default.

4. Open your browser and navigate to `http://localhost:5173`

## How to Use

1. Enter your name in the registration form and click "Join Chat"
2. View the list of online users in the left panel
3. Click on a user to start chatting with them
4. Type your message in the input field at the bottom of the right panel and click "Send"

## Building for Production

### Frontend

```bash
cd frontend
npm run build
```

The built files will be in the `frontend/dist` directory and can be served by the Go backend.

### Backend

The Go backend is configured to serve the frontend static files in production. After building the frontend, you can start the backend and it will serve the frontend at the root URL.

```bash
cd backend
go build
./go-vue-chat
```

## License

MIT
