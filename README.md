PROJETO DESENVOLVIDO NA IDE GOLAND

HOW TO INSTALL:

1 - Clonar o projeto

2 - Rodar o comando docker-compose up -d para criar o banco postgres no docker

3 - Rodar o comando go run main.go para executar o projeto.

* As migrações são rodadas automaticamente ao se executar o projeto.
** Todas as vezes que o projeto é executado ele apaga e cria novamente as tabelas.

FUNCIONAMENTO

1 - Criação de usuário
2 - Criação de livro
3 - Emprestar livro
4 - Devolver livro
5 - Buscar usuário por ID

TESTES

Para a realização dos testes é necessario a alteração das variáveis no .env
e posteriormente caminhar até a pasta tests/modeltests e executar o comando go test
