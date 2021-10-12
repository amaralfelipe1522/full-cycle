# Exemplo de REST API - Nível 3
>   Usando Laminas API Tools

1.  Inicia um Container Remote de PHP 7 + MariaDB;
2.  Instalar o Laminas API Tool:
```
    sudo composer create-project laminas-api-tools/api-tools-skeleton
```
3. Atualizar o Composer dentro do projeto:
```
    cd api-tools-skeleton
    sudo composer update
```
4.  Habilite o modo de desenvolvedor e execute a API Tools na porta 8080:
```
    sudo composer development-enable
    php -S 0.0.0.0:8080 -t public public/index.php    
```
5.  Instalar SQlite:
```
    sudo apt-get update
    sudo apt-get install sqlite3
```
6.  Crie o arquivo de banco de dados local e acesse-o:
```
    touch test.sqlite
    sqlite3 test.sqlite
```
7.  Crie uma tabela de exemplo:
```
CREATE TABLE users (id int, name varchar(255), email varchar(255));
INSERT INTO users (id, name, email) VALUES (1, "Felipe", "amaral.felipe@live.com");
```
8. No API Tools:
    1.  Vá em Databases e edite a adapter Dummy;
    2.  Selecione o driver PDO_Sqlite e indique no campo "Database" o nome do arquivo que criamos localmente: *test.sqlite*;
    3.  Crie uma Nova API;
    4.  Crie um novo serviço para a API do tipo DB Connected;
    5.  Seleciona a tabela Dummy;
