# Cheatsheet - Kubectl && Kind

Visualizar cluster em execução no Kubectl:
```
kubectl get nodes
```

Listar clusters criados no Kind:
```
kind get clusters
```

Deletar clusters criados no Kind:
```
kind delete clusters kind
```

Criar cluster com mais de um node (diretório: 04-cluster-nodes):
```
kind create cluster --config=<CAMINHO DO ARQUIVO YAML> --name=<NOME DO CLUSTER>
```

Listar todos os cluster no arquivo kube/config, independente se estão local ou na Cloud:
```
kubectl config get-clusters
```

Alterar o contexto para outro cluster em execução:
```
kubectl config use-context <NOME DO CLUSTER>
```

Consultar os replicasets
```
kubectl get replicasets
```

Consultar os deployments
```
kubectl get deployments
```

> A extensão Kubernetes do VSCode torna a alteração do contexto mais prática

## Resumo do primeiro exemplo
- Criado uma aplicação em Golang;
- Dockerfile e geração da imagem;
- Teste de execução do container;
- Push da imagem no Dockerhub;
- Criação do pod.yml;
- Verifica se os nodes do cluster estão de pé:
    ```
    kubectl get nodes
    ```
- Aplica o arquivo pod.yml para a criação do Pod:
    ```
    kubectl apply -f pod.yml
    ```
- Consulta o Pod em execução:
    ```
    kubectl get pods
    ```
- Exportar a porta do pod em execução:
    ```
    kubectl port-forward pod/goserver 3000:3000
    ```
