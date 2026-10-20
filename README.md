# StoryCraft AI

AI-powered comic and drama creation platform.

This repository contains a minimal runnable foundation for a Go + Vue application, designed to be extended into a full AI-assisted comic / drama / storyboard platform.

## Stack

- Backend: Go
- Frontend: Vue 3 + Vite + TypeScript
- Containerization: Docker Compose
- Storage: MinIO (planned)
- AI layer: LLM + Seedance integration (planned)

## Quick start

### 1) Backend

```bash
cd backend
go run .
```

Backend health check:

```bash
curl http://localhost:8080/api/v1/health
```

### 2) Frontend

```bash
cd frontend
npm install
npm run dev -- --host 0.0.0.0
```

Open:

```text
http://localhost:5173
```

### 3) Docker Compose

```bash
docker compose up --build
```

## Repo structure

```text
backend/
frontend/
config/
docker-compose.yml
.env.example
README.md
```

## Planned features

- Project management
- Chapter and scene management
- Shot / storyboard editor
- Character library
- Media asset management
- AI story generation
- AI shot generation
- Seedance-powered visual generation
- Review and publishing flow

## Notes

This is the initial scaffolding for the project. The backend is intentionally minimal and can be extended with Gin, GORM, JWT, and AI service layers.
