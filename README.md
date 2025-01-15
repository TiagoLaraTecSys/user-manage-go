# Projeto Go com Docker e MySQL

Este repositório contém uma aplicação Go que se conecta a um banco de dados MySQL, tudo orquestrado via Docker Compose.
É uma aplicação de gestão(CRUD) de usuários. Através de um Client como Insomnia, Postman ou outros, é possível acessar recursos como:

### POST: /v1/user
### GET:  /v1/user?userId=?
### GET:  /v1/users
### PUT:  /v1/user?userId=
### DELETE: /v1/user?userId=

## Pré-requisitos

Antes de começar, você precisa ter instalado em sua máquina:

- **Docker**: [Instruções de instalação do Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Instruções de instalação do Docker Compose](https://docs.docker.com/compose/install/)
- **Git** (opcional, para clonar o repositório): [Instalação do Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Passo a passo para rodar a aplicação

### 1. Clonar o repositório

Se você ainda não clonou o repositório, faça isso com o comando:

```bash
git clone https://github.com/TiagoLaraTecSys/user-manage-go.git
cd user-manage-go
```

### 2. Rodar o comando:
```bash
docker-compose up --build
```

Pronto! Se tudo ocorrer bem a sua aplicação estará de pé em http://localhost:8080