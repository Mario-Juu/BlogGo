###  Projeto Blog em Go
## Descrição
Um projeto para desenvolvimento de uma aplicação web afim de utilizar de toda a variabilidade da linguagem Go, criando um blog e hosteando um site pelo próprio computador.


## Tecnologias
- Go
- HTML
- CSS
- Bootstrap
- Docker
- MariaDB

## Como usar
1. Abra o VSCode ou sua IDE compatível de preferência
2. Clone o repositório
```sh
git clone https://github.com/Mario-Juu/WebDevGo.git
```
3. Crie o banco de dados 
```sh
docker run --name dbblog -e MYSQL_ROOT_PASSWORD=secret -e MARIADB_MSQL_LOCALHOST_USER=true -p 3306:3306 -d mariadb:latest
```
4. Abra o executável (WebDevGo.exe)
5. Visite o site na porta 8080 (http://localhost:8080)


## Objetivo 
Tendo como objetivo ser um backend que se vincula a um frontend, ambos criados na mesma aplicação, se pode fazer algumas mudanças com flags na criação da aplicação:


Vá até o caminho que clonou o repositório, execute o seguinte comando para configurações personalizadas:
```sh
./WebDevGo.exe -port={insira a porta} -env=prd (caso queira cache ativada)
```
