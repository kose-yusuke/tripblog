# gocrud
Goを用いたWebアプリ作成

## ディレクトリ構成

.
├── README.md
├── api
│   ├── build
│   │   └── Dockerfile
│   └── cmd
│       ├── calendar
│       │   ├── controllers
│       │   │   ├── calendar_controller.go
│       │   │   ├── plan_controller.go
│       │   │   └── user_controller.go
│       │   ├── model
│       │   │   ├── calendar.go
│       │   │   ├── color.go
│       │   │   ├── plan.go
│       │   │   └── user.go
│       │   └── service
│       │       ├── calendar_psql.go
│       │       ├── plan_psql.go
│       │       └── user_psql.go
│       ├── db
│       │   └── db.go
│       ├── endpoint
│       │   └── calendar
│       ├── form
│       │   └── api
│       │       ├── post.go
│       │       └── user.go
│       ├── go.mod
│       ├── go.sum
│       ├── main.go
│       └── server
│           └── server.go
├── db
│   ├── pgadmin4

## 技術
- React/NextJS
- Go/Gin/GORM
- Postgresql
- Docker
- AWS ECS

## 構成
<img width="623" alt="スクリーンショット 2022-09-02 11 03 53" src="https://user-images.githubusercontent.com/61424757/188043149-b22ad2e0-d676-494f-bebd-3af69d885c17.png">
