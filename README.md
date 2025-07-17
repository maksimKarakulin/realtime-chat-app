# **Real-Time Chat App**
**A Go (Backend) + Vue.js (Frontend) Application**

---

## **📖 Overview**
This project is a **real-time chat application** with:
- **Backend**: A Go server using WebSockets for live messaging.
- **Frontend**: A Vue.js SPA with reactive UI.

### **Key Features**
✔ **Real-time messaging** (WebSockets)  
✔ **User sessions** (JWT optional)  
✔ **Scalable structure** (separate backend/frontend)  
✔ **Environment-based config** (`.env` files)

---

## **📂 Project Structure**
```bash
real-time-chat-app/
├── backend/           # Go server (WebSocket + API)
│   ├── .env          # Backend config (DB, PORT, JWT_SECRET)
│   ├── src/
│   │   ├── handlers/ # WebSocket + HTTP handlers
│   │   ├── models/   # Data structures
│   │   └── main.go   # Entry point
│   └── go.mod        # Go dependencies
│
├── frontend/         # Vue.js app
│   ├── src/
│   │   ├── assets/   # Images/fonts
│   │   ├── components/ # Reusable UI
│   │   ├── views/    # Pages (e.g., ChatView.vue)
│   │   ├── router/   # Vue Router setup
│   │   ├── store/    # Vuex state management
│   │   ├── App.vue   # Root component
│   │   └── main.js   # Vue entry point
│   ├── .env          # Frontend config (API_URL)
│   └── package.json  # JS dependencies
│
├── .gitignore        # Ignores node_modules, .env, etc.
└── README.md         # This file!
```

---

## **⚙️ Configuration**
### **Backend (Go)**
1. Rename `.env.example` to `.env` in `/backend`:
   ```env
   PORT=3000                 # Go server port
   JWT_SECRET=your_secret   # For auth (if used)
   DB_URL=postgres://...    # Optional database
   ```
2. Install Go dependencies:
   ```bash
   cd backend
   go mod download
   ```

### **Frontend (Vue.js)**
1. Rename `.env.example` to `.env` in `/frontend`:
   ```env
   VUE_APP_API_URL=http://localhost:3000  # Points to Go server
   ```
2. Install JS dependencies:
   ```bash
   cd frontend
   npm install
   ```

---

## **🚀 How to Run**
### **Option 1: Local Development**
1. **Start Backend (Go)**:
   ```bash
   cd backend
   go run src/main.go
   ```
   → Server runs at `http://localhost:3000`.

2. **Start Frontend (Vue.js)**:
   ```bash
   cd frontend
   npm run serve
   ```
   → App runs at `http://localhost:8080`.

3. Open `http://localhost:8080` in your browser.

### **Option 2: Docker (Production)**
```bash
docker-compose up --build  # Requires docker-compose.yml
```

---

## **🔧 Troubleshooting**
| Issue | Solution |  
|-------|----------|  
| Go server won’t start | Check `.env` and run `go mod tidy` |  
| Vue app can’t connect to backend | Verify `VUE_APP_API_URL` in `/frontend/.env` |  
| WebSocket errors | Ensure CORS is configured in Go backend |  

---

## **📜 License**
MIT License - Free for personal/commercial use.

---

## **📌 Notes**
- **Backend-first**: The Go server must run before the frontend.
- **No database?** The app works in memory by default (add DB if scaling).
- **Security**: For production, use HTTPS and secure `JWT_SECRET`.

**Happy coding!** ✨

--- 
