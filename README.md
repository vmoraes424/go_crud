# 🏋️‍♂️ Go CRUD API - Workouts

Este é um projeto de API REST em Go feita com foco em simplicidade e boas práticas. Ele gerencia treinos (workouts) e utiliza:

- 🔥 [Gin](https://github.com/gin-gonic/gin) para criar rotas e lidar com requisições
- 🧠 [GORM](https://gorm.io/) para integração com banco de dados (ORM)
- ⚡ [Air](https://github.com/cosmtrek/air) para hot reload durante o desenvolvimento

---

## 📦 Tecnologias

- Go
- Gin
- GORM (PostgreSQL)
- Air (dev tool)

---

## 🚀 Como rodar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/vmoraes424/go_crud.git
cd go_crud
```

### 2. Copie o arquivo .env

Crie um arquivo .env na raiz do projeto e copie os dados que estão em .env.example

### 3. Instale as dependências

```bash
go mod tidy
```

### 4. Inicie com hot reload (Air)

Se ainda não tem o Air:
```bash
go install github.com/cosmtrek/air@latest
```
Depois, é só rodar:

```bash
air
```
