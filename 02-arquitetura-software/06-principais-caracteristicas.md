# **Microsserviços - Principais Características**

## Componentização via serviços
-   Services dos microsserviços != Services da OO;
-   Componente é uma unidade de software independente que pode ser substituída ou atualizada;
-   Desvantagens:
    1.  Chamadas externas são mais custosas do que chamadas locais;
    2.  Cruzamento entre componentes pode se tornar complexo;
    3.  Transações entre serviços são grandes desafios;
    4.  Mudanças bruscas em regras de negócio podem afetar diversos serviços.

## Organização em torno do negócio
>   Um produto é baseado em um ou mais produtos que trabalham em diferentes contextos.
-   Time de desenvolvedores por produto;
-   Muitas empresas tratam os times como "squads";
-   Cada squad é responsável por um ou mais produtos;
-   Cada produto pode ter um ou mais serviços envolvidos.

## Smart endpoints & dumb pipes
-   Exposição de APIs (Ex.: Rest, Graphql, etc.);
-   Comunicação entre serviços;
-   Comunicação síncrona e assíncrona;
-   Utilização de sistemas de mensageria (ex.: RabitMQ, Apache Kafka, etc.);
-   Garantia de que um serviço foi ou será executado.

## Governança descentralizada
-   Ferramenta certa para o trabalho certo. Tecnologias podem ser definidas baseadas na necessidade do produto;
-   Diferentes padrões de squads;
-   Contratos de interface de forma independente;

## Descentralização no gerenciamento de dados

## Automação de Infraestrutura
-   Cloud computing;
-   Testes automatizados;
-   CI/CD;
-   Load balancer / Autoscaling.

## Desenhado para falhar
-   Tolerâcia a falha;
-   Serviços precisam de fallback;
-   Logging;
-   Monitoramento em tempo real;
-   Alarmes.

## Design evolutivo
-   Produtos bem definidos podem evoluir ou serem extintos por razões de negócio;
-   Gerenciamento de versões;
-   Replacement and upgradeability.