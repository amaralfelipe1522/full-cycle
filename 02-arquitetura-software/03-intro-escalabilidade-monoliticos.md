# **Introdução a Escabilidade e Sistemas Monolíticos**

-   Tudo dentro do mesmo sistema;
-   Alto acoplamento;
-   Processo de Deploy "completo" a cada mudança;
-   Normalmente usa uma única tecnologia;
-   Um problema afeta todo o sistema;
-   Maior complexidade para times;
-   Não é crime usar sistema monolítico;
-   Na maioria dos casos vai atender;
-   Menos complexidade na maioria dos casos.

## **Escalando Software**

1.  **Escala Vertical:** Aumento nos recursos computacionais de uma única máquina;
2.  **Escala Horizontal:** Adiona mais máquinas a essa aplicação, compondo o sistema. Se usa um load balance para distribuir as cargas entre as aplicações.

## **Detalhes Sobre a Arquitetura da Aplicação**

-   Disco efêmero;
-   Servidor de Aplicação Vs Servidor de Assets;
-   Cache centralizado;
-   Sessões centralizadas;
-   Upload / Gravação de Arquivos;

>   Tudo pode ter que ser destruído e criado novamente.