# **Configuração para estudo local**

## Kubectl

Download:
```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

Instalar:
```
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

Verificar instalação:
```
kubectl version --short 
```

## Kind

Instalação e criação de cluster local (Necessário ter Golang e Docker instalados):
```
go install sigs.k8s.io/kind@v0.12.0 && kind create cluster

kubectl cluster-info --context kind-kind
```

Apontando o Kubectl ao cluster criado pelo Kind (~/.kube/config):
```
kubectl cluster-info --context kind-kind
```

## Metric Server
> [Github da Aplicação](https://github.com/kubernetes-sigs/metrics-server)

Baixe manualmente:
```
wget kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```
Adicionar o seguinte argumento no arquivo baixado para que seja possível realizar a comunicação sem TLS com o container:
***- --kubelet-insecure-tls***

Aplica o serviço:
```
kubectl apply -f metrics-server.yaml
```
Verifica se a instalação ocorreu com sucesso:
```
kubectl get apiservices
```    