# 📚 CRUD de Autores e Livros em Go

Este é um projeto simples de CRUD (Create, Read, Update, Delete) desenvolvido em Go, com o objetivo de praticar e aprofundar os conhecimentos na linguagem. O sistema gerencia informações sobre autores e seus respectivos livros.

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM para Go
- [SQLite]

## 📦 Funcionalidades

- CRUD de Autores:
  - Criar autor
  - Listar autores
  - Buscar autor por ID
  - Atualizar autor
  - Deletar autor

- CRUD de Livros (em construção):
  - Criar livro vinculado a um autor
  - Listar livros
  - Buscar livro por ID
  - Atualizar livro
  - Deletar livro

## 🛠 Como rodar o projeto

### Pré-requisitos

- Go instalado ([Download Go](https://golang.org/dl/))

### Passos

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/rest-api-books.git

# Acesse o diretório
cd rest-api-books

# Instale as dependências (caso use go mod)
go mod tidy

# Execute a aplicação
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`.

## 📌 TODO

- [ ] Finalizar CRUD de Livros
- [ ] Testes unitários
- [ ] Autenticação (JWT)
- [ ] Paginação de resultados
- [ ] Documentação com Swagger

---

## 🧑‍💻 Autor

Desenvolvido por [Heitor Vasconcelos](https://github.com/fhva29) como forma de estudo.