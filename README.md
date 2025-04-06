
# 🛒 API em GO - CRUD para Mercadinho

## 📌 Sobre a API
Esta API foi desenvolvida em **Go** para gerenciar um **mercadinho**, permitindo o cadastro, consulta, atualização e remoção de produtos. 

## 🛠️ Tecnologias Utilizadas
- **Go (Golang)** 🐹
- **GORM** - ORM para Go
- **PostgreSQL/MySQL/SQLite** (escolha sua opção)
- **Gin** - Framework web para Go

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
  "nome": "Arroz 5kg",
  "descricao": "Arroz branco premium",
  "preco": 27.50,
  "quantidade": 80
}
```

### ❌ Deletar Produto
**DELETE** `/products/{id}`

---

## 🛠 Configuração do Banco de Dados (YAML)
Crie um arquivo `config.yaml` e configure conforme seu banco de dados:

```yaml
banco_de_dados:
  tipo: "postgres"
  host: "localhost"
  usuario: "seu_usuario"
  senha: "sua_senha"
  nome: "mercadinho_db"
  porta: 5432
  sslmode: "disable"
```

---

## 📦 Docker
Para facilitar a execução da API, podemos utilizar **Docker**. 
Crie um arquivo `docker-compose.yml` e adicione:

```yaml
version: '3.8'

services:
  db:
    container_name: nome_container_db
    image: postgres:latest
    environment:
      POSTGRES_USER: db_user
      POSTGRES_PASSWORD: db_senha
      POSTGRES_DB: db_nome
    ports:
      - "5432:5432"
  api:
    container_name: nome_container_api
    image: imagem_api
    build: .
    depends_on:
      - db
    ports:
      - "8080:8080"
```

Para rodar a API com o banco de dados, basta executar:
```sh
docker-compose up -d
```

## 🧾 Script para o banco
 ```sql
 create table products(
      id serial PRIMARY KEY,
      name varchar(255) not null ,
      price numeric(10,2) not null ,
      quantity int,
      description varchar(255)
 )
 ```

---

## 🚀 Rodando a API
1. Clone este repositório
2. Instale as dependências: `go mod tidy`
3. Configure o banco em `config.yaml`
4. Execute a API: `go run main.go`

Agora sua API está pronta para gerenciar os produtos do seu mercadinho! 🛍️🥦

---

## 📚 Documentações Úteis
- [Documentação Oficial do Go](https://golang.org/doc/)
- [Documentação do GORM](https://gorm.io/docs/)
- [Documentação do PostgreSQL](https://www.postgresql.org/docs/)
- [Documentação do Docker](https://docs.docker.com/)

---

## 🏗 Estrutura do Código
A API segue a seguinte estrutura de arquivos:
```
/project-root
│── main.go
│── docker-compose.yml
│── Dockerfile
│── models/
│   ├── product.go
│── controllers/
│   ├── productController.go
│── routes/
│   ├── routes.go
│── repository/
│   ├── productRespository.go
│── usecase/
│   ├── productsUsecase.go
│── db/
│   ├── conn.go
```

### 📌 Exemplo de Struct para Produto
```go
type Produto struct {
    id            uint    `json:"id" gorm:"primaryKey"
    name          string  `json:"nome"`
    description   string  `json:"descricao"`
    Price         float64 `json:"preco"`
    Quantity      int     `json:"quantidade"`
}
```

