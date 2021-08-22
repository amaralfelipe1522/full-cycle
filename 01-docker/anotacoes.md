# Anotações

## Container
- Namespaces = Processos isolados;
- Cgroups = Controla os recursos de cada processo;
- File System - OFS (Overlay File System) = Gerencia o sistema em camadas, onde em uma nova versão do mesmo, apenas as camadas diferentes serão criadas. Sem precisar copiar todo o sistema operacional novamente por exemplo.

## Imagem
- Uma fotografia do estado do meu sistema.
- São imutáveis, porém é criado uma camada READ/WRITE onde podemos realizar alterações no CONTAINER.
- Uma imagem é criada após realizar o Build de um Dockerfile ou após um COMMIT quando for realizada alguma alteração na camada READ/WRITE do CONTAINER.
- PULL & PUSH = Para buscar ou atualizar uma imagem armazenada no Registry.

## Dockerfile
- Arquivo declarativo onde se descreve como será a imagem que queremos consumir.

## Dockerhost
- Daemon - API = Serviço que executa os comandos informados para o Docker Client (docker build, docker ps, etc.);
- Cache = Armazena minhas imagens originadas, ou não, do Registry;
- Volume = Responsável pela persistência de dados. Quando eu compartilho um diretório da minha máquina com a imagem, toda a informação armazenada nela não será perdida quando a imagem for destruída;
- Network = Garante a comunicação entre os containers;

## Volume x Bind
- Volume é uma forma mais antiga de persistir seus dados mesmo após o encerramento do container;
```
podman run -d --rm -p 8000:80 -v $(pwd):/usr/share/nginx/html --name nginx docker.io/nginx
```
- Mount especifica o que será feito, tendo como maior diferencial o fato de não criar diretórios em relação ao caminho informado no comando.
```
podman run -d --rm -p 8000:80 --mount type=bind,source=$(pwd)/html,target=/usr/share/nginx/html --name nginx docker.io/nginx
```
Exemplo de erro por não existir o caminho especificado:
```
podman run -d --rm -p 8000:80 --mount type=bind,source=$(pwd)/html/xxx,target=/usr/share/nginx/html --name nginx docker.io/nginx
Error: statfs /home/felipe/Documentos/workspace/full-cycle/01-docker/html/xxx: no such file or directory
```

### Exec x Attach
- O docker exec é utilizado para enviar um comando para o container em questão, quando você faz **docker exec -it container /bin/bash** é enviada uma instrução para o container abrir um processo do bash e mantê-lo aberto até que feche. Enquanto que o docker attach anexa o Stdio do container ao seu terminal.
- Ao sair da função Attach, o container é encerrado.

###Quando:
- Precisar persistir os dados do seu container no computador;
- Tiver que manter o mesmo tipo de File System do Linux (Linux Virtual Machine);
- Querer ter mais performance;
- Você não sabe ou não tem controle sobre qual é o caminho dos diretórios (criando esse volume no container e posteriormente apontando para o seu computador);
- For necessário mapear e apontar esse volume para mais de um container;

*Listar Volumes*
```
podman volume ls
```
*Crir Volumes*
```
podman volume create meu-volume 
```
*Detalhes do Volume*
```
podman volume inspect meu-volume 
```
> Como todos os volumes serão criados no mesmo caminho, facilita posteriormente um backup.

*Criando o container com o volume já criado*
```
podman run -d --rm -p 8000:80 --mount type=volume,source=meu-volume,target=/usr/share/nginx/html --name nginx docker.io/nginx
```

*Criando um novo container, usando o mesmo volume* 
```
podman run -d --rm -p 8000:80 --mount type=volume,source=meu-volume,target=/usr/share/nginx/html --name nginx2 docker.io/nginx
```

*É possivel também utilizar o volume usando a forma antiga de chamada de volumes*
```
podman run -d --rm -p 8000:80 -v meu-volume:/usr/share/nginx/html --name nginx2 docker.io/nginx
```
*Quando for necessário esvaziar os volumes que não estão sendo usados*
```
podman volume prune
```
## Dicas
```
podman rm $(podman ps -a -q) -f
```

## CMD x Entrypoint
- CMD descrito no Dockerfile será substituído se no comando de rodar o container for incluído alguma ação, como por exemplo bash, sh, echo, etc.
- Entrypoint sempre é executado e nunca substituído;

Exemplo:

    FROM ubuntu:latest as img-ubuntu
    ENTRYPOINT ["echo","Hello"]
    CMD ["World"]
```
docker run --rm --name ubuntu ubuntu:1.0        
Hello World
docker run --rm --name ubuntu ubuntu:1.0 Felipe
Hello Felipe
```
- Quando no Entrypoint estiver executando um arquivo shell script, geralmente no final dele terá o comando **exec "$@"** responsável por permitir que um comando seja executado após a execução do shell script;

## Network

### Bridge
- Quando se cria uma network no Docker sem informar nada, esta será do tipo Bridge.
- Usada para conectar um ou mais containers na mesma rede;

### Host
- Conecta na mesma rede a network do container com o host do Docker na máquina local;
- Não funciona no MacOs, apenas no Linux ou WSL2.

### Overlay
- Quando se é necessário conectar na mesma rede vários containers e maquinas físicas;
- Usado com geralmente com Docker Swarm, criando algo como um cluster de containers, para escalar o sistema;

### MacVLan
- Conecta na mesma rede um container com um mac adress específico;

### None
- Isola o container em uma network.

### Comandos

- Exibe todas as networks:
```
podman network ls
```

- Remove todas as networks que não estão sendo usadas:
```
podman network prune
```

- Remove um network que não está sendo usada:
```
podman network rm
```

### Exemplo criando network Bridge "simples"
> Aparentemente essa forma simplificada não funciona no Podman
- Criando dois containers de Bash:
```
docker run -d -it --name ubuntu1 bash
docker run -d -it --name ubuntu2 bash
```
- Inspecionando as conexões do tipo Bridge:
```
docker network inspect bridge
```
- Acessando o ubuntu1 é possível dar um ping no ubuntu2:
```
docker attach ubuntu1
ping <IP DO CONTAINER>
```
> Porém nesse tipo de conexão não é possível executar ping <NOME DO CONTAINER>

### Criando uma rede Bridge

- Criar rede do tipo Bridge:
```
podman network create --driver bridge minha-rede
```
- Inicia os containers apontando para a rede criada:
```
docker run -d -it --name ubuntu1 --network minha-rede bash
docker run -d -it --name ubuntu2 --network minha-rede bash
```
> Nesse cenário é possível dar um ping no container apenas informando o nome dele: ping ubuntu2

- Conectar um container já criado em uma rede
```
docker network connect minha-rede ubuntu3
```

### Criando uma rede Host
- Criar rede do tipo Host:
```
docker run --rm -d --name nginx --network host nginx
```
> Nesse exemplo o Nginx estará em execução no localhost:80 de minha máquina.

### Container acessando minha rede local
> No exemplo em questão rodei uma aplicação Node na porta 3000 sem Docker.

- Acessando o localhost 3000 de minha máquina dentro do container:
```
curl http://host.docker.internal:3000
```

## Multistage Builder
> Exemplo no arquivo /php/Dockerfile.prod
```
docker build -t amaralfelipe1522/laravel . -f Dockerfile.prod
```
- Usado para otimizar imagens, mantendo-os fáceis de ler e manter.

## Exemplo de proxy reverse com Nginx com Laravel
- Criar a conexão:
```
docker network create --driver bridge laranet
```

- Build and Run:
```
docker build -t amaralfelipe1522/laravel . -f Dockerfile.prod
docker build -t amaralfelipe1522/nginx-laravel .

docker run -d --network laranet --name laravel amaralfelipe1522/laravel
docker run -d --network laranet --name nginx-laravel -p 8080:80 amaralfelipe1522/nginx-laravel
```

## Docker Compose

- Rebuildar imagens todas as vezes que fizer uma alteração no Dockerfile:
```
docker-compose up -d --build
```
- Comandos usados no exemplo da aula
```
docker exec -it db bash
mysql -uroot -p
use nodedb
CREATE TABLE people (id INT NOT NULL auto_increment, name VARCHAR(250), PRIMARY KEY (id));
```
> Tabela PEOPLE do container DB será populada após rodar a aplicação index.js no container NODE-APP


## Dependencias entre Containers

### depends_on
- Garante que um container será iniciado antes de outro, porém não garante que irá aguardar o primeiro container estar completo;

### Dockerize ou Wait-for-it
> Exemplo de Dockerize no exemplo de Docker-compose de Node com MySql.
- Garante que um container só irá subir após outro container ser completamente iniciado.