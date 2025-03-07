# 🛒 API em Go - CRUD para Mercadinho

## 📌 Sobre a API
Esta API foi desenvolvida em **Go (Golang)** para gerenciar um **mercadinho**, permitindo o cadastro, consulta, atualização e remoção de produtos. 

## 🛠️ Tecnologias Utilizadas
- **Go (Golang)** 🐹
- **GORM** - ORM para Go
- **PostgreSQL / MySQL / SQLite** (escolha a opção desejada)
- **Gin** - Framework web para Go
- **Docker** - Facilita a execução da API

---

## 🔧 Endpoints Disponíveis

### 🏗️ Adicionar um Produto
**POST** `/products`
```json
{
  "name": "Arroz 5kg",
  "description": "Arroz branco tipo 1",
  "price": 25.90,
  "quantity": 100
}
```

### 📖 Listar Produtos
**GET** `/products`

### 🔍 Buscar Produto por ID
**GET** `/products/{id}`

### ✏️ Atualizar Produto
**PUT** `/products/{id}`
```json
{
  "name": "Arroz 5kg",
  "description": "Arroz branco tipo 1",
  "price": 25.90,
  "quantity": 100
}
```

### ❌ Deletar Produto
**DELETE** `/products/{id}`

---

## 📦 Configuração com Docker
Para facilitar a execução da API, podemos utilizar **Docker**. 
Crie um arquivo `docker-compose.yml` com o seguinte conteúdo:

```yaml
version: '3.8'

services:
  db:
    container_name: mercadinho_db
    image: postgres:latest
    environment:
      POSTGRES_USER: db_user
      POSTGRES_PASSWORD: db_senha
      POSTGRES_DB: db_nome
    ports:
      - "5432:5432"
  api:
    container_name: mercadinho_api
    build: .
    depends_on:
      - db
    ports:
      - "8080:8080"
```

Para rodar a API com o banco de dados, execute:
```sh
docker-compose up -d
```

---

## 🧾 Script para o Banco de Dados
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INT,
    description VARCHAR(255)
);
```

---

## 🚀 Como Rodar a API
1. Clone este repositório:
   ```sh
   git clone https://github.com/seu-usuario/api-mercadinho.git
   ```
2. Entre no diretório do projeto:
   ```sh
   cd api-mercadinho
   ```
3. Instale as dependências:
   ```sh
   go mod tidy
   ```
4. Configure o banco no arquivo `config.yaml`.
5. Execute a API:
   ```sh
   go run main.go
   ```

Agora sua API está pronta para gerenciar os produtos do seu mercadinho! 🛍️🥦

---

## 📚 Documentações Úteis
- [Documentação Oficial do Go](https://golang.org/doc/)
- [Documentação do GORM](https://gorm.io/docs/)
- [Documentação do PostgreSQL](https://www.postgresql.org/docs/)
- [Documentação do Docker](https://docs.docker.com/)

---

## 🏗 Estrutura do Projeto
A API segue a seguinte estrutura de arquivos:
```
/project-root
│── main.go
│── docker-compose.yml
│── Dockerfile
│── models/
│   ├── product.go
│   ├── message.go
│── controllers/
│   ├── productController.go
│── routes/
│   ├── routes.go
│── repository/
│   ├── productRepository.go
│── usecase/
│   ├── productsUsecase.go
│── db/
│   ├── conn.go
```

---

## 📌 Exemplo de Struct para Produto
```go
type Produto struct {
    ID          uint    `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Quantity    int     `json:"quantity"`
}
```

---
