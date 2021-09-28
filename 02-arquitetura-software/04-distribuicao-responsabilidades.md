# **Distribuição de Responsabilidades**

Exemplo de responsabilidades distribuídas:
-   Aplicação:
    1. Instância de Aplicação 1;
    2. Instância de Aplicação 2;
    3. Instância de Aplicação 3;
    4. Instância de Aplicação 4;
-   Proxy Reverso;
-   Elastic Search;
-   Static (servidor de assets);
-   Cache;
-   Banco de Dados.

> Obs.: Esse cenário ainda se trata de uma aplicação monolítica, com escala horizontal.

## **Escala Horizontal: Aplicação Monolítica**

-   Ter imagens (Ex.: AWS) / Containers;
-   Ser facilmente reconstruída;
-   Ter suas responsabilidades, incluindo assets.

## **Até que ponto vale a pena aplicações monolíticas?**

-   Times grandes;
-   Necessidade de escalar todo o sistema pelo fato de uma área em específico esteja em pico de utilização;
-   Risco de um deploy completo começa a se elevar;
-   Necessidade de usar tecnologias diferentes;
