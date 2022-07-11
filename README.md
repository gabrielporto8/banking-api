# banking-api

## Setup do projeto
Comandos docker para subir o projeto localmente:

Build da aplicação:
`docker build -t banking-api .`

Execução
`docker run -d -p 8080:8080 --name banking-api banking-api` (o parâmetro `name` é opcional, ajuda a identificar o container caso esteja usando algum programa docker no desktop)

Após subir o container, a aplicação estará rodando na porta `:8080` do seu localhost.

## Descrição da API
### Accounts
#### `POST http://localhost:8080/accounts`
Endpoint responsável pela criação de accounts.

Body da requisição:
```json
{
  "name": "Billy Butcher",
  "cpf": "11122233344",
  "secret": "password123",
  "balance": 20
}
```
Responses:
- Sucesso
  - HTTP Status 202
  - Body: 
```json
{
  "id": 0,
  "name": "Billy Butcher",
  "cpf": "11122233344",
  "secret": "$2a$14$CXJm/oYRimCUCn..ouooz.p2Mp6EJ1fz1ycmP6j7ienAEXdwssO2.",
  "balance": 20,
  "created_at": "2022-07-11T05:55:32.9788359Z"
}
```
- Erro
  - HTTP Status 409 (Conflict)
  - HTTP Status 400 (Bad Request)
  - Body:
  ```json
  {
    "Error": "Entered CPF already exists."
  }
  ```
  
#### `GET http://localhost:8080/accounts`
Endpoint responsável por listar os accounts cadastrados.

Responses:
- Sucesso
  - HTTP Status 202
  - Body:
  ```json
  {
   "0": {
		"id": 0,
		"name": "Billy Butcher",
		"cpf": "11122233344",
		"secret": "$2a$14$CXJm/oYRimCUCn..ouooz.p2Mp6EJ1fz1ycmP6j7ienAEXdwssO2.",
		"balance": 20,
		"created_at": "2022-07-11T05:55:32.9788359Z"
	},
	"1": {
		"id": 1,
		"name": "Hughie Campbell",
		"cpf": "00000000001",
		"secret": "$2a$14$zzEwixTDBHmdppES1lhipOQUCgMaMfuoiHW1un12GFuS0RuxnPemy",
		"balance": 0,
		"created_at": "2022-07-11T05:56:36.0556845Z"
	},
	"2": {
		"id": 2,
		"name": "Victoria Neuman",
		"cpf": "00000000002",
		"secret": "$2a$14$$2a$14$8sZCESrMYXZtop5MmIfHOuwv7i5yMLo2htoJq7uO5ayBwBMYc6xva",
		"balance": 0,
		"created_at": "2022-07-11T05:57:36.0556845Z"
	}
  }
  ```
  
#### `GET http://localhost:8080/accounts/{id}/balance`
Endpoint responsável por retornar o balance para o ID informado.

Responses:
- Sucesso
  - HTTP Status 202
  - Body:
  ```json
  {
    "Balance": 20
  }
  ```
  
- Erro
  - HTTP Status 404 (Not Found)
  - Body:
  ```json
  {
    "Error": "Account not found."
  }
  ```
  
### Login
#### `POST http://localhost:8080/login`
Endpoint responsável por realizar o login.

Body da requisição:
```json
{
  "cpf": "11122233344",
  "secret": "password123"
}
```

Responses:
- Sucesso
  - HTTP Status 200 (OK)
  - Body:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiIwMjgyNDQ3NDA3NiIsImV4cCI6MTY1NzUyMTc3MH0.reDkav8Cc7LDZKJOPSAOFteVjsL8wDCI3E7FaYm3Brg"
  }
  ```
  
O token retornado nessa response é o token que será utilizado nas rotas que exigem autenticação.
  
- Erro
  - HTTP Status 404 (Not Found)
  - Body:
  ```json
  {
    "Error": "Unauthorized."
  }
  ```
  
### Transfers
#### `GET http://localhost:8080/transfers`
Endpoint responsável por listar as transferências do usuário logado.

Header da requisição:
`"Authorization: Bearer token123"`
