
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
**POST** `/produtos`
```json
{
  "nome": "Arroz 5kg",
  "descricao": "Arroz branco tipo 1",
  "preco": 25.90,
  "quantidade": 100
}
```

### 📖 Listar Produtos
**GET** `/produtos`

### 🔍 Buscar Produto por ID
**GET** `/produtos/{id}`

### ✏️ Atualizar Produto
**PUT** `/produtos/{id}`
```json
{
  "nome": "Arroz 5kg",
  "descricao": "Arroz branco premium",
  "preco": 27.50,
  "quantidade": 80
}
```

### ❌ Deletar Produto
**DELETE** `/produtos/{id}`

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
    image: postgres:latest
    restart: always
    environment:
      POSTGRES_USER: seu_usuario
      POSTGRES_PASSWORD: sua_senha
      POSTGRES_DB: mercadinho_db
    ports:
      - "5432:5432"

  api:
    build: .
    depends_on:
      - db
    ports:
      - "8080:8080"
    environment:
      DB_HOST: db
      DB_USER: seu_usuario
      DB_PASSWORD: sua_senha
      DB_NAME: mercadinho_db
```

Para rodar a API com o banco de dados, basta executar:
```sh
docker-compose up -d
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
│── models/
│   ├── products.go
│── controllers/
│   ├── productsController.go
│── routes/
│   ├── routes.go
│── database/
│   ├── connection.go
```

### 📌 Exemplo de Struct para Produto
```go
type Produto struct {
    ID          uint    `json:"id" gorm:"primaryKey"
    Nome        string  `json:"nome"`
    Descricao   string  `json:"descricao"`
    Preco       float64 `json:"preco"`
    Quantidade  int     `json:"quantidade"`
}
```

