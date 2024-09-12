# DesafioPismo

# Introdução

Projeto se solução proposta para o desafio Pismo de event-processor. O sistema proposto se conecta a uma fila de eventos
em uma instância RabbitMQ, e baseado no tipo de evento recebido é realizada a persistência do mesmo em uma tabela de
dados correspondente em um banco de dados Postgres.

# Arquitetura da solução

Toda a aplicação roda baseda em uma stack de serviços utilizando o Docker. Ao iniciar a aplicação, o código será
compilado
e estará disponível dentro da stack. Ele se conectará aos serviços de mensageria e bancos de dados e estará pronto
para receber eventos.

## Executando a aplicação

Inicie o serviço e as intâncias do postgresql (banco de dados), rabbitmq (mensageria) e o pgadmin (plataforma de
administração do PostgreSQL)

```
docker-compose up -d
```

# Build and Test

TODO: Describe and show how to build your code and run the tests.

# Contribute

TODO: Explain how other users and developers can contribute to make your code better.

