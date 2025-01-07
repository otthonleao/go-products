# go-products

Este repositório contém uma API desenvolvida em Go para gerenciar produtos e cadastrar usuários com testes unitários.

### Funcionalidades
- CRUD de produtos: permite criar, ler, atualizar e deletar produtos.
- Usuários: criar um usuário e fazer login para pegar o token de acesso com JWT.

### Estrutura do Projeto
O projeto segue uma estrutura modularizada para facilitar a manutenção e escalabilidade:

- cmd/server: Contém o código para inicializar e configurar o servidor.
- configs: Arquivos de configuração da aplicação.
- docs: Documentação da API
- internal: Implementações internas, incluindo lógica de negócio e handlers.
- pkg/entity: Definições das entidades utilizadas na aplicação.
- test: Testes unitários e de integração.

### Tecnologias utilizadas
- Go 1.16
- Go Chi Framework
- Gorm
- Sqlite
- Testify
- Bcrypt
- JWT
- SwagGo

### Instalação
Clone o repositório:
```
git clone https://github.com/otthonleao/go-products.git
```
Navegue até o diretório do projeto:
```
cd go-products
```
Instale as dependências:
```
go mod download
```

### Uso
Navege até o arquivo principal:
```
cd cmd/server
```
Inicie o servidor:
```
go run main.go
```
- A API estará disponível em **http://localhost:8000**.

- Para acessar pelo swagger utilize: **http://localhost:8000/swagger/index.html#/**.

- Pegue o token de acesso no endpoint `/users/login` com o seguinte body:
```json
{
  "email": "otthon@mail.com",
  "password": "123456"
}
```

### User Endpoints
- `POST /users/login`: Cria um token de acesso
- `POST /users`: Cadastra um novo usuário

### Product Endpoints protegidos pelo JWT
- `GET /products`: Retorna a lista de produtos.
- `GET /products/{id}`: Retorna um produto específico pelo ID.
- `POST /products`: Cria um novo produto.
- `PUT /products/{id}`: Atualiza um produto existente pelo ID.
- `DELETE /products/{id}`: Deleta um produto pelo ID.`
