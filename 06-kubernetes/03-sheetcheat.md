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
- Processo de Rollout/Rollback:
    
    Lista hitórico de versões:
    ```
    kubectl rollout history deployment goserver
    ```
    Volta para a versão anterior a atual:
    ```
    kubectl rollout undo deployment goserver
    ```
    Volta para uma versão específica:
    ```
    kubectl rollout undo deployment goserver --to-revision=<NUMERO LISTADO NO HISTORICO DE VERSOES>
    ```
- Executar um Service:
    ```
    kubectl apply -f service.yml
    ```
- Exportar a porta do Service em execução:
    ```
    kubectl port-forward service/goserver-service 3000:3000
    ```
- Acessando a API do Kubernetes:
    ```
    kubectl proxy --port=8080
    ```
- Acessando o bash de um Pod assim como se acessa um container Docker:
    ```
    kubectl exec -it goserver-7d59bcb6c8-7b449 -- bash 
    ```
- Acessando os logs do Pod:
    ```
    kubectl logs goserver-7d59bcb6c8-7b449
    ```
- Executar um Deployment e inspecionar sua inicialização:
    ```
    kubectl apply -f deployment.yml && watch -n1 kubectl get pods
    ```
- Verificar quais APIs estão em execução no Kubernetes (exemplo usado com Metrics Server):
    ```
    kubectl get apiservices
    ```
- Visualizer recursos utilizados de CPU e Memória do Pod:
    ```
    kubectl top pods <NOME DO POD>
    ```
- Teste de stress com Fortio usando a imagem Docker oficial
    ```
    kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/secret"
    ```
- Visualizer Volumes persistentes criados:
    ```
    kubectl get pvc
    ```
- Visualizer Storage Class:
    ```
    kubectl get storageclass
    ```
- Escalando um Deployment ou StatefulSet:
    ```
    kubectl scale statefulset mysql --replicas=5
    ```
- Visualizando cada Pod em um Headless Service (de outro container no mesmo cluster):
    ```
    kubectl exec -it goserver-XXXXXXXX-XXX -- bash
    ping mysql-headless
    ping mysql-0.mysql-headless
    ```
- Visualizar os namespaces:
    ```
    kubectl get namespace
    ```
- Visualizar os pods por namespaces:
    ```
    kubectl get pods -n=kube-system
    ```
- Criar um namespace:
    ```
    kubectl create ns ns-develop
    ```
- Aplicar um arquivo YAML ao namespace específico:
    ```
    kubectl apply -f <NOME DO ARQUIVO> -n=<NOME DO NAMESPACE>
    ```
- Filtrar pods por Labels:
    ```
    kubectl get pods -l <NOME DA LABEL>
    ```
- Exibir configurações do cluster:
    ```
    kubectl config view
    ```
- Exibir contexto atual do cluster:
    ```
    kubectl config current-context
    ```
- Criar um novo contexto no cluster onde:
    ```
    kubectl config set-context dev --namespace=dev --cluster=kind-second-cluster --user=kind-second-cluster
    ```
- Lista os Service Accounts existentes:
    ```
    kubectl get serviceaccounts
    ```
- Lista os recursos e suas APIs Groups:
    ```
    kubectl api-resources
    ```
