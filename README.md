# 🚀 Task Tracker API

REST API для управления задачами, написанный на Go.

Проект демонстрирует архитектуру backend-приложения с использованием:

- Go
- PostgreSQL
- Docker
- Swagger (OpenAPI)
- Echo framework

API позволяет создавать, получать, обновлять и удалять задачи.

---

# ✨ Features

- CRUD API для задач
- PostgreSQL база данных
- Swagger документация
- Docker контейнеризация
- SQL migrations
- REST архитектура
- разделение на handler / service / repository
- простой frontend для тестирования API

---

# 🧰 Tech Stack

- Go
- Echo
- PostgreSQL
- pgx
- Docker
- Swagger (swaggo)

---

# 📁 Project Structure
task-tracker
│
├── cmd
│ └── main.go
│
├── docs
│ ├── docs.go
│ ├── swagger.json
│ └── swagger.yaml
│
├── internal
│
│ ├── handler
│ │ └── task_handler.go
│ │
│ ├── model
│ │ └── task.go
│ │
│ ├── repository
│ │ ├── db.go
│ │ └── task.go
│ │
│ └── service
│ └── task_service.go
│
├── migrations
│ └── 001_create_tasks_table.sql
│
├── web
│ ├── index.html
│ ├── about.html
│ ├── script.js
│ ├── style.css
│ └── logo.png
│
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── task-tracker
└── README.md

---

# ⚙️ Run project

## 1 Clone repository
git clone https://github.com/yourusername/task-tracker.git
cd task-tracker

---

## 2 Run with Docker
docker compose up --build

API будет доступен:
    http://localhost:8080

Swagger:
    http://localhost:8080/swagger/index.html

---

# 🗄 Database

Используется PostgreSQL.

Миграция:
    migrations/001_create_tasks_table.sql

Пример структуры таблицы:

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT DEFAULT 'todo',
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

# 📡 API Endpoints

Get tasks:

    GET /tasks

Get tasks:

    GET /tasks

Create task:

    POST /tasks

    Example body:
        {
        "title": "Learn Go",
        "description": "Build REST API"
        }

Update task:

    PUT /tasks/{id}

Update task:

    PUT /tasks/{id}

---

# 📚 Swagger

Swagger UI доступен:

http://localhost:8080/swagger/index.html

Документация автоматически генерируется через:

    swag init

---

# 🌐 Frontend

В проекте есть простой frontend для тестирования API.

web/

    Страницы:
        index.html
        about.html
    Файлы:
        script.js
        style.css

---

# Architecture

Проект использует разделение слоёв:

    Handler → Service → Repository → Database
    Handler

    HTTP обработчики.

    Service

    Бизнес-логика.

    Repository

    Работа с базой данных.

---

# 🐳 Docker

Запуск:

    docker compose up --build

    Поднимаются сервисы:
        API
        PostgreSQL
---

# 🎰 Author

    &Andrew
    &Go Backend Developer
    &Contacts: andrecentov.space, or TG: @hannaXmontana