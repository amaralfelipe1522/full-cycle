# **Comunicação entre Microsserviços**

A comunicação entre Microsserviços podem ocorrer de duas formas:

## **Síncrona**
![Comunicação Síncrona](../assets/sincrona.png)
> Porém em alguns casos o processo pode ficar travado.
![Comunicação Síncrona](../assets/sincrona-travado.png)

## **Assíncrona**
![Comunicação Assíncrona](../assets/assincrona.png)
> Mesmo que o Microsserviço B esteja fora do ar, a requisição não vai se perder, pois estará armazenado no sistema de Mensageria.