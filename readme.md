# Voting System
## en
voting system

Introducing one of the solutions, of the problems that most applications that suffer large amounts of request have.

With this, a module of a voting system was built.

As an example the social networks based on the voting of the best social networks.

**MicroServices**

two microservices were created

user and votes

User registers his data to perform the votes.

**RabbitMQ**

We use *RabbitMQ's* messaging system to publish polls.

Then we will read the voting data and store it in the *sqlite3* database.

## pt
Sistema de voto

Apresentando uma das soluções, dos problemas que maior parte das apliacações que sofre grandes quantidades de requisição têm.

Com isto construiu-se um o modulo de um sistema de votação.

Como exemplo as redes socias baseando na votação das melhores redes sociais.

**Microserviços**

foi criado dois microserviços 

usuário e votos

usuário registra seu dados para efetuar os votos.

**RabbitMQ**

em usamos o sistema de Mensageria o *RabbitMQ* para publicar as votações .

Depois vamos  ler os dados das votações e armazenamos na base de dados *sqlite3*


## Installation

```
docker-compose up -d
```
