# Go Finance API

Uma API RESTful para controle financeiro pessoal, desenvolvida em **Go (Golang)**, com autenticação JWT, organização em camadas e testes automatizados.

## 🚀 Funcionalidades

- Registro e autenticação de usuários (JWT).
- CRUD completo de transações financeiras (entradas e saídas).
- Relatórios simples:
  - **Saldo do mês**
  - **Despesas por categoria**
- Estrutura escalável com organização em camadas:
  - `domain` → modelos de domínio
  - `repository` → acesso ao banco de dados
  - `service` → regras de negócio
  - `handler` → controladores HTTP
  - `router` → configuração das rotas
  - `middleware` → autenticação e segurança

## 📂 Estrutura do Projeto

```
cmd/api/
  main.go               # Ponto de entrada da aplicação

internal/
  app/                  # Inicialização do app
  domain/               # Modelos de domínio (User, Transaction)
  handler/              # Handlers HTTP (UserHandler, TransactionHandler)
  middleware/           # Middleware (Auth)
  repository/           # Repositórios (UserRepo, TransactionRepo)
  router/               # Rotas (User, Transaction, Reports)
  service/              # Serviços com regras de negócio
  test/                 # Testes automatizados (api_test.go)

migrations/             # Scripts SQL de criação e rollback
```

## 🛠 Tecnologias Utilizadas

- [Go](https://go.dev/) (Golang)
- [Gin](https://github.com/gin-gonic/gin) (framework web)
- [GORM](https://gorm.io/) (ORM para banco de dados)
- PostgreSQL
- JWT (JSON Web Token) para autenticação

## ▶️ Como Rodar o Projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/fiorellizz/go-finance-api.git
   cd go-finance-api
   ```

2. Configure o `.env` com suas credenciais do banco de dados:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=finance
   JWT_SECRET=seuSegredoAqui
   ```

3. Rode as migrações:
   ```bash
   go run cmd/api/main.go
   ```

4. Teste a API:
   ```bash
   go test ./internal/test -run TestAPI
   ```

## ✅ Rotas Disponíveis

- `POST /api/register` → Registrar usuário
- `POST /api/login` → Login e gerar token JWT
- `GET /api/users` → Listar usuários
- `POST /api/transactions` → Criar transação
- `GET /api/transactions` → Listar transações
- `PUT /api/transactions/:id` → Atualizar transação
- `DELETE /api/transactions/:id` → Deletar transação
- `GET /api/reports/balance` → Relatório de saldo
- `GET /api/reports/expenses-by-category` → Relatório de despesas por categoria

## 🧪 Testes Automatizados

O projeto possui testes de ponta a ponta em `internal/test/api_test.go`, cobrindo:

- Registro e login
- CRUD de transações
- Relatórios

Rodar os testes:
```bash
go test ./internal/test -v
```
