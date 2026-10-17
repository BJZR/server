# Go API Server

A minimal, production-ready REST API written in Go, ready for continuous deployment on [Render](https://render.com) via GitHub.

## Project Structure

```
.
├── main.go
├── config/
│   └── config.go
├── handlers/
│   └── handlers.go
├── go.mod
├── .gitignore
└── README.md
```

## Endpoints

| Method | Path      | Description                  |
|--------|-----------|------------------------------|
| GET    | `/`       | Welcome/status JSON message  |
| GET    | `/health` | Health check — `{"status":"ok"}` |

## Run Locally

```bash
# Clone the repo
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>

# (Optional) set a custom port
export PORT=9000

# Run
go run .
```

The server starts on `http://localhost:8080` (or the port you set).

## Render Deployment

### Build Command
```
go build -o main .
```

### Start Command
```
./main
```

> Render automatically injects the `PORT` environment variable. The server listens on `0.0.0.0:PORT` as required by Render.

---

## Push to GitHub & Connect to Render

### 1. Initialize the local repository

```bash
git init
git add .
git commit -m "feat: initial Go API"
```

### 2. Create a repository on GitHub and push

```bash
# Replace <your-username> and <your-repo> with your actual values
git remote add origin https://github.com/<your-username>/<your-repo>.git
git branch -M main
git push -u origin main
```

### 3. Connect to Render

1. Log in to [render.com](https://render.com) and click **New → Web Service**.
2. Select your GitHub repository.
3. Set the following fields:
   - **Environment**: `Go`
   - **Build Command**: `go build -o main .`
   - **Start Command**: `./main`
4. Click **Create Web Service**.

Render will build and deploy automatically on every push to `main`.
