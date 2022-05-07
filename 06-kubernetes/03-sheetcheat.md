# Cheatsheet - Kubectl && Kind

/``````

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

> A extensão Kubernetes do VSCode torna a alteração do contexto mais prática

