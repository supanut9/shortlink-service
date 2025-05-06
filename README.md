# 🔗 shortlink-service

A simple and scalable URL shortener built with [Go](https://golang.org/) and [Fiber](https://gofiber.io/). Supports custom or randomly generated short codes, redirection, and optional link management.

---

## 🚀 Features

- Create short links for long URLs
- Fast redirects using Fiber
- Unique short code generation
- RESTful API design
- Optional link listing and deletion
- Easy to deploy and extend

---

## 🧱 Tech Stack

- **Go** – core language
- **Fiber** – fast HTTP web framework
- **PostgreSQL / Redis** – for storing URL mappings _(optional)_
- **Docker** – containerization (optional)

---

## 📦 Installation

```bash
git clone https://github.com/yourusername/shortlink-service.git
cd shortlink-service
go mod tidy
go run cmd/server/main.go
```
