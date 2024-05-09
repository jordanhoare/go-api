# Go App

## Getting Started

## MakeFile

```
go-app
├─ .github
│  └─ workflows
│     ├─ backend.yaml
│     └─ frontend.yaml
├─ .gitignore
├─ Makefile
├─ README.md
├─ docker-compose.yml
└─ src
   ├─ backend
   │  ├─ .air.toml
   │  ├─ cmd
   │  │  └─ api
   │  ├─ go.mod
   │  ├─ go.sum
   │  ├─ internal
   │  │  ├─ database
   │  │  │  └─ database.go
   │  │  └─ server
   │  │     ├─ routes.go
   │  │     └─ server.go
   │  └─ tests
   │     └─ handler_test.go
   └─ frontend
      ├─ .prettierrc.json
      ├─ index.html
      ├─ package-lock.json
      ├─ package.json
      ├─ public
      │  └─ vite.svg
      ├─ src
      │  ├─ assets
      │  │  └─ react.svg
      │  ├─ index.css
      │  ├─ index.tsx
      │  ├─ pages
      │  │  └─ app
      │  │     ├─ App.css
      │  │     └─ App.tsx
      │  └─ vite-env.d.ts
      ├─ tsconfig.json
      └─ vite.config.ts

```
```
go-app
├─ .github
│  └─ workflows
│     ├─ backend.yaml
│     └─ frontend.yaml
├─ .gitignore
├─ Makefile
├─ README.md
├─ docker-compose.yml
└─ src
   ├─ backend
   │  ├─ .air.toml
   │  ├─ cmd
   │  │  └─ api
   │  ├─ go.mod
   │  ├─ go.sum
   │  ├─ internal
   │  │  ├─ database
   │  │  │  └─ database.go
   │  │  └─ server
   │  │     ├─ routes.go
   │  │     └─ server.go
   │  └─ tests
   │     └─ handler_test.go
   └─ frontend
      ├─ ...
```








   ├─ backend
   │  ├─ .air.toml
   │  ├─ cmd
   │  │  └─ api
   │  ├─ go.mod
   │  ├─ go.sum
   │  ├─ api
   │  │  ├─ handler
   │  │  │  ├─ user_handler.go
   │  │  │  ├─ product_handler.go
   │  │  │  └─ order_handler.go
   │  │  ├─ middleware
   │  │  └─ routes
   │  │     └─ router.go
   │  ├─ pkg
   │  │  ├─ config
   │  │  │  └─ config.go
   │  │  ├─ model
   │  │  │  ├─ user.go
   │  │  │  ├─ product.go
   │  │  │  └─ order.go
   │  │  ├─ repository
   │  │  │  ├─ user_repo.go
   │  │  │  ├─ product_repo.go
   │  │  │  └─ order_repo.go
   │  │  └─ service
   │  │     ├─ user_service.go
   │  │     ├─ product_service.go
   │  │     └─ order_service.go
   │  ├─ internal
   │  │  ├─ auth
   │  │  └─ database
   │  │     └─ database.go
   │  └─ tests
   │     ├─ handler_test.go
   │     ├─ service_test.go
   │     └─ repo_test.go








   ├─ backend
   │  ├─ .air.toml
   │  ├─ api
   │  │  ├─ handler
   │  │  │  ├─ user_handler.go
   │  │  │  ├─ product_handler.go
   │  │  │  └─ order_handler.go
   │  │  ├─ middleware
   │  │  │  ├─ auth_middleware.go
   │  │  │  └─ logging_middleware.go
   │  │  └─ routes
   │  │     └─ router.go
   │  ├─ cmd
   │  │  └─ api
   │  │     └─ main.go
   │  ├─ go.mod
   │  ├─ go.sum
   │  ├─ internal
   │  │  ├─ auth
   │  │  │  └─ jwt.go
   │  │  ├─ config
   │  │  │  └─ config.go
   │  │  ├─ database
   │  │  │  └─ database.go
   │  │  └─ server
   │  │     ├─ routes.go
   │  │     └─ server.go
   │  └─ tests
   │     ├─ handler_test.go
   │     └─ middleware_test.go
```
go-app
├─ Makefile
├─ README.md
├─ docker-compose.yml
└─ src
   ├─ backend
   │  ├─ .air.toml
   │  ├─ api
   │  │  ├─ handler
   │  │  │  ├─ health_handler.go
   │  │  │  └─ hello_handler.go
   │  │  └─ routes
   │  │     └─ routes.go
   │  ├─ cmd
   │  │  └─ api
   │  ├─ go.mod
   │  ├─ go.sum
   │  ├─ internal
   │  │  ├─ database
   │  │  │  └─ database.go
   │  │  ├─ model
   │  │  │  ├─ health.go
   │  │  │  └─ hello.go
   │  │  └─ server
   │  │     └─ server.go
   │  └─ tests
   │     └─ handler_test.go
   └─ frontend
      └─ ...
```