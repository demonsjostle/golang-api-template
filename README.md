 # 🚀 CPRU GraphQL API

GraphQL API สำหรับระบบจัดการข้อมูล ด้วย Go และออกแบบตามแนวทาง **Hexagonal Architecture** 
ระบบ RBAC (Role-Based Access Control)

---

## 📦 Tech Stack

| Component         | Technology                |
|------------------|---------------------------|
| Language         | [Go](https://golang.org/) |
| API              | [GraphQL](https://graphql.org/) via [gqlgen](https://gqlgen.com) |
| Web Framework    | [Echo](https://echo.labstack.com/) |
| ORM              | [Ent](https://entgo.io/) |
| Database         | PostgreSQL                |
| Architecture     | Hexagonal (Ports & Adapters) |

---

## 🧱 Project Structure

```
cpru-api/
├── cmd/                      # Bootstrap / Compose dependency
│   └── server/               # Entry point (main.go)
│
├── internal/
│   ├── domain/               # ห้าม import framework ใด ๆ (Secondary Port)
│   │   └── example/             # 💡 Domain model (entity) + interface (core entity, no logic)
│   │
│   ├── application/          # Service ห้ามรู้จัก Ent, GraphQL (Primary port && Business logic)
│   │   └── example/             # 💡 Business logic / Usecases (pure logic, testable)
│   │
│   ├── infrastructure/       # เท่านั้นที่ import DB (Secondary Adapter)
│   │   ├── db/               # Adapter to Ent (implements Repository)
│   │   └── config/           # (Optional) env, flags, config loader
│   │
│   ├── interface/            # Controller / Resolver เชื่อม request ↔ usecase  (Primary adapter) 
│   │   ├── graphql/          # ✅ Resolver Adapter + Playground
│
├── graph/
│   ├── generated/            # gqlgen generated code
│   ├── model/                # gqlgen model (map to/from domain)
│
├── ent/                      # Ent schema (generate ORM)
│   └── schema/
│
├── pkg/                      # Optional: shared utils/helpers/logger
│
├── .air.toml                 # Live reload dev config
├── gqlgen.yml                # gqlgen config
├── go.mod / go.sum           # Go module
├── Dockerfile
└── docker-compose.yml
```

---

## 🚀 Getting Started

### 1. ติดตั้งเครื่องมือ
- Go >= 1.24
- PostgreSQL
- (แนะนำ) [air](https://github.com/cosmtrek/air) สำหรับ hot reload

### 2. ติดตั้ง dependencies

```bash
go mod tidy
```

### 3. สร้าง code จาก Ent
#### create schema 
```bash 
go run -mod=mod entgo.io/ent/cmd/ent new [schemaname]

```
#### 
```bash
go generate ./ent

```

### 4. สร้าง gqlgen
```bash 
go run github.com/99designs/gqlgen generate
```

### 5. รัน Server

```bash
go run cmd/server/main.go
```

Server จะเปิดที่:  
📍 [http://localhost:8080](http://localhost:8080) (GraphQL Playground)


---

## 💡 Concept: Hexagonal Architecture

แยกระบบเป็นชั้น ๆ เพื่อให้โค้ดสามารถ:

- เปลี่ยน DB ได้ง่าย
- ทดสอบแต่ละส่วนแยกได้ (service, domain, adapter)
- ไม่ผูกกับ web framework หรือ ORM โดยตรง

```txt
GraphQL → Resolver → Service → Repository Interface → Ent (Adapter) → PostgreSQL
```

---

## 🧡 ผู้ดูแลโปรเจกต์

พัฒนาโดย [Minerta](https://minertatech.com/) 

