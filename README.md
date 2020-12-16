How to Install
Clone o projeto
git clone
Instale as dependências
npm install
Execute o banco no Docker
docker run --name db_postgres  -p 5432:5432 -e POSTGRES_DB=thanos -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -d postgres:10.11
