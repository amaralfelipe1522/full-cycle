# **Service Discovery**

> É responsável por prover mecanismos de identificação dos serviços disponíveis e suas instâncias.

## **Client Side**

![Client Side](../assets/service-discovery-client.png)

## **Server Side**

![Server Side](../assets/service-discovery-server.png)

## **Ferramentas Populares**

-   Netflix Eureka;
-   Consul;
-   Etcd;
-   ZooKeeper.

> O processo de Service Discovery já estará sendo aplicado automaticamente ao utilizar o Kubernetes, pois ao acionar o Load Balance, este por sua vez vai identificar o Pod disponível do serviço.