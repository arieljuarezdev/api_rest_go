# API REST em Go (MVC, PostgreSQL)

Projeto simples de API REST em Go, usando a arquitetura em camadas (Controllers, Models, Repositories). O foco é CRUD de clientes.

## 🧩 Estrutura do projeto

- `cmd/api/main.go`: ponto de entrada, rotas HTTP (Gorilla Mux)
- `internal/controllers/controller.go`: handlers HTTP + lógica de fluxo
- `internal/models/model.go`: definição do struct `Customer`
- `internal/repositories/database.go`: conexão com PostgreSQL

## 📦 Dependências

- Go (1.20+ recomendado)
- `github.com/gorilla/mux`
- `github.com/lib/pq`

## 🔧 Configuração do banco de dados

A função `OpenConnection()` em `internal/repositories/database.go` usa os dados:

- host: `localhost`
- port: `5434`
- user: `ariel`
- password: `bacon`
- dbname: `api_rest_go`
- sslmode: `disable`

### Exemplo SQL de tabela

```sql
CREATE TABLE costumer (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  phone TEXT,
  adress TEXT
);
```

> Observação: há inconsistência no nome (`costumer` em vez de `customer`). Mantenha igual ao existente.

## 🚀 Endpoints

1. `GET /customers`
   - Lista todos clientes

2. `GET /customers/{id}`
   - Busca cliente por ID

3. `POST /newCustomer`
   - Inserir cliente
   - Payload JSON:
     ```json
     {
       "name": "João",
       "phone": "(11) 99999-9999",
       "adress": "Rua X, 123"
     }
     ```

4. `PATCH /updateCustomer/{id}`
   - Atualiza cliente por ID
   - JSON:
     ```json
     {
       "phone": "(11) 98888-8888",
       "adress": "Rua Y, 456"
     }
     ```

5. `DELETE /deleteCustomer/{id}`
   - Remove cliente por ID

## 📌 Model

`internal/models/model.go`:

```go
type Customer struct {
  Id     int    `json:"id"`
  Name   string `json:"name"`
  Phone  string `json:"phone"`
  Adress string `json:"adress"`
}
```

## 💡 Pontos que preciso ajustar.

- Mudar `adress` para `address` e corrigir em DB/JSON
- TRatar e retornar erros HTTP com código adequado, em vez de `return`
- Usar `defer conn.Close()` imediatamente após a conexão aberta
- Validar _scan_ e presença de linha no `GetCustumerById`
- Separar camada `repository` com `func GetAllCustomers()`, etc., em vez de uso direto em Controller

## ▶️ Execução

```bash
go run cmd/api/main.go
```

A API ficará em `http://localhost:8080`.

## 🛠️ Teste rápido com curl

```bash
curl -X POST http://localhost:8080/newCustomer \
  -H "Content-Type: application/json" \
  -d '{"name":"Ana","phone":"123","adress":"Rua A"}'

curl http://localhost:8080/customers
```
