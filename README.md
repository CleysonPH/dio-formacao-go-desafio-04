# Desafio 04 - Formação Go Developer - DIO

Código com a solução do desafio 04 da formação Go Developer da Digital Innovation One. Construir uma API REST completa com os verbos que permitirão que sejam executadas diversas operações.

## Como executar

Para executar o programa, primeiro é necessario baixar as dependências do projeto. Para isso, execute o comando abaixo:

```bash
go mod download
```

Após baixar as dependências, execute o comando abaixo para executar o programa:

```bash
go run main.go
```

Agora, basta acessar a URL http://localhost:8000 para acessar a API.

## Rotas

| Método | Rota            | Descrição                           |
| ------ | --------------- | ----------------------------------- |
| GET    | /api/books      | Retorna todos os livros cadastrados |
| GET    | /api/books/{id} | Retorna o livro com o ID informado  |
| POST   | /api/books      | Cadastra um novo livro              |
| PUT    | /api/books/{id} | Atualiza o livro com o ID informado |
| DELETE | /api/books/{id} | Deleta o livro com o ID informado   |
