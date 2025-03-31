# API de Gerenciamento de Usuários em Go com Gin Gonic (MVC)

Este projeto é uma API RESTful desenvolvida em Go (Golang) utilizando o framework Gin Gonic. Ela oferece funcionalidades de autenticação de usuários (login e registro) e um CRUD básico para gerenciamento de usuários, com rotas protegidas por JWT. A persistência dos dados é feita em um banco de dados PostgreSQL e as senhas são armazenadas de forma segura com criptografia. A arquitetura do projeto segue o padrão MVC (Model-View-Controller) para melhor organização e manutenibilidade do código.

## Tecnologias Utilizadas

* **Go (Golang):** Linguagem de programação utilizada no backend.
* **Gin Gonic:** Framework web de alta performance para Go.
* **JWT (JSON Web Tokens):** Padrão para criação de tokens de acesso seguros.
* **PostgreSQL:** Banco de dados relacional utilizado para persistência dos dados.
* **Pacotes Go:** (Liste aqui os principais pacotes utilizados, por exemplo:)
    * `gorm.io/gorm`: ORM para interagir com o PostgreSQL.
    * `github.com/gin-gonic/gin`: Framework web Gin Gonic.
    * `github.com/golang-jwt/jwt/v5`: Biblioteca para geração e verificação de JWT.
    * `golang.org/x/crypto/bcrypt`: Para criptografia de senhas.
    * `github.com/joho/godotenv`: Para carregar variáveis de ambiente.

## Estrutura do Projeto (MVC)

A arquitetura do projeto segue o padrão MVC (Model-View-Controller):

* **`model/`:** Contém as definições das estruturas de dados (Models) e a lógica de interação com o banco de dados (usualmente utilizando um ORM como GORM).
* **`controller/`:** Contém os handlers (controladores) que recebem as requisições HTTP, interagem com os models para processar os dados e retornam as respostas.
* **`repository/`:** Contém a lógica de acesso aos dados (Data Access Layer). As interfaces e implementações para interagir com o banco de dados (PostgreSQL, utilizando um ORM como GORM) residem aqui. Essa camada abstrai a forma como os dados são persistidos.
* **`services/`:** Contém a lógica de negócios da aplicação. Os services orquestram as chamadas aos repositories, aplicam regras de negócio e preparam os dados para serem utilizados pelos controllers.
* **`router/`:** Define as rotas da API e associa os endpoints aos respectivos handlers.
* **`middleware/`:** Contém middlewares utilizados na aplicação, como o middleware de autenticação JWT.
* **`config/`:** Arquivos de configuração da aplicação, como a conexão com o banco de dados.
* **`main.go`:** Ponto de entrada da aplicação.

## Funcionalidades

* **[Registro de Usuário](#registro-de-usuário):** Permite a criação de novos usuários no sistema.
* **[Login de Usuário](#login-de-usuário):** Permite que usuários existentes façam login e recebam um token JWT para acesso às rotas protegidas.
* **CRUD de Usuários (Protegido por JWT):**
    * **[Buscar Usuário por ID](#buscar-usuário-por-id):** Retorna os dados de um usuário específico com base no seu ID.
    * **[Atualizar Usuário](#atualizar-usuário):** Permite a modificação dos dados de um usuário existente.
    * **[Deletar Usuário](#deletar-usuário):** Permite a remoção de um usuário do sistema.

## Pré-requisitos

* Go instalado (versão >= 1.18 recomendada).
* PostgreSQL instalado e configurado.
* `make` (opcional, para executar comandos do Makefile).

## Configuração

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/Murilojms7/Login-System-MVC
    cd Login-System-MVC
    ```

2.  **Crie um arquivo `.env`:**
    Copie o arquivo `.env.example` (se existir) ou crie um novo arquivo `.env` na raiz do projeto e configure as seguintes variáveis de ambiente:
    ```env
    DB_HOST=host
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=name
    DB_PORT=port
    SECRETKEY=secretkey
    ```
    **Importante:** Substitua os valores de exemplo pelas suas configurações reais do banco de dados e uma chave secreta forte para o JWT.

3.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

## Execução

Você pode executar a API de duas maneiras:

**Usando `go run`:**

```bash
go run main.go
```

### **Registro de Usuário**
---
`POST` / `localhost:8080/auth/register`

**Request Body**

|Field|Type|required|Description
|-----|----|:-----------:|---------
|Name|String|✅|Nome do usuário
|Email|String|✅|Email do usuário
|Password|String|✅|Senha do usuário


```js
{   
    "name": "murilo",
    "email": "murilo@gmail.com",
    "password": "123",
}
```
**Exemplo de Resposta**

```js
{
    "data": "user: murilo@gmail.com created successfully",
    "message": "operation from handler: Register-User success"
}
```

### **Login de Usuário**
---
`POST` / `localhost:8080/auth/login`

**Request Body**

|Field|Type|required|Description
|-----|----|:-----------:|---------
|Email|String|✅|Email do usuário
|Password|String|✅|Senha do usuário

```js
{
    "email":"murilo@gmail.com",
    "password": "123"
}
```

**Exemplo de Resposta**

```js
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM1MTcwMDQsIm5hbWUiOiJqbXMiLCJ1c2VySWQiOiJjMTAyMzNlYS1iOWYzLTRiMGYtYWJiMS05MWE4ODIxMzNiZWYifQ.9OBCtLrkf3C4gv1wZ9Yxc1agYs6smuHD_D9ErR1XOH8"
    },
    "message": "operation from handler: Login-user success"
}
```

### **Buscar Usuário por ID**
---
`GET` / `localhost:8080/user/:id`

**Request Header**
|Key|Value|
|---|----|
|Authorization|Bearer "token"|

**Exemplo de Resposta**

```js
{
    "data": {
        "email": "murilo@gmail.com",
        "id":"00e11cf0-1129-43f2-971c-7f1a1118947c",
        "name": "murilo"
    },
    "message": "operation from handler: show-user-by-id success"
}
```

### **Atualizar Usuário**
---
`PUT` / `localhost:8080/user/:id`

**Request Header**
|Key|Value|
|---|----|
|Authorization|Bearer "token"|

**Request Body**

|Field|Type|required|Description
|-----|----|:-----------:|---------
|Name|String|🚫|Nome do usuário
|Email|String|🚫|Email do usuário
|Password|String|🚫|Senha do usuário

```js
{   
    "name": "murilo",
    "email": "murilo@gmail.com",
    "password": "123",
}
```

**Exemplo de Resposta**

```js
{
    "data": "User: murilo was updated",
    "message": "operation from handler: update-user success"
}
```

### **Deletar Usuário**
---
`DELETE` / `localhost:8080/user/:id`

**Request Header**
|Key|Value|
|---|----|
|Authorization|Bearer "token"|

**Exemplo de Resposta**

```js
{
    "data": "Delete user with id c7090e11-1a3b-41fa-aaa3-786ab0952ef6 successfully",
    "message": "operation from handler: delete-user success"
}
```
