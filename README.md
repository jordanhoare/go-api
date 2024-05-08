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