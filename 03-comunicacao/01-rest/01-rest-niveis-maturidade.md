# **REST - Represetational State of Transfer**

-   Surgiu em 2000 por Roy Fielding;
-   Simplicidade;
-   Stateless:
    -   Não mantém o estado/sessão, por conta disso em muitos casos é necessário passar formas de identificação, como Tokens.
-   Cacheável;

## **Níveis de Maturidade (Richardson Maturity Model)**

-   Nível 0: The Swamp of POX:

        Parte do princípio de que tudo que for trafegado por HTTP é realizado uma transação, como inserir um dado no banco de dado, ou executar uma função, sem uma padronização.

-   Nível 1: Utilização de resources:
    
    ![Rest Resources](../../assets/rest-resources.png)

-   Nível 2: Verbos HTTP:

    ![Rest Resources](../../assets/verbos-http.png)

-   Nível 3: HATEOAS: Hypermedia as the Engine of Application State:
        
        Sempre vai responder sua solicitação informando o que mais é possível realizar nesse serviço. Tornando assim o serviço auto explicativo.

    ![Rest Resources](../../assets/hateoas.png)

## **Uma boa API REST precisa:**

1.  Utilizar URIs únicas para serviços e itens que são expostos para esses serviços;
2.  Utiliza todos os verbos HTTP para realizar as operações em seus recursos, incluindo caching;
3.  Provê links relacionais para os recursos exemplificando o que pode ser feito;

## **HAL, Collection+JSON e Siren**

-   JSON não provê uma padrão de hipermídia para realizar a linkagem;
-   HAL (Hypermedia Application Language);
-   Siren

## **REST: HAL**
>   Além de trazer padrão para os dados no JSON, traz informações complementares e de como acessá-las.

Media type = application/hal+json

![REST: HAL](../../assets/hal.png)