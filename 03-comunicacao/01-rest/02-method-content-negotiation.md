# **Method e Content Negotiation**

## **REST: HTTP Method Negotiation**

Ao passar OPTIONS na requisição, este nos informa quais métodos são permitidos ou não em um determinado recurso.

A resposta pode ser:
    
    HTTP/1.1 200 OK
    Allow: GET, POST

Caso envie a requisição em outro formato em que o server não aceite:

    HTTP/1.1 405 Not Allowed
    Allow: GET, POST

## **REST: Accept Negotiation**

Baseada na requisição que o cliente está fazendo para o server, passando **O QUE** e **COMO** ele quer a resposta.

O cliente solicita a informação no formato desejado:
    
    GET /product
    Accept: application/json

E caso o server não aceite, retornará informando:

    HTTP/1.1 406 Not Acceptable

## **REST: Content-type Negotiation**
Através de um content-type no header da request, o servidor consegue verificar se ele irá conseguir processar a informação enviada pelo cliente.

    POST /product HTTP/1.1
    Accept: application/json
    Content-Type: application/json

    {
        "name":"Product 1"
    }

Do contrário:

    HTTP/1.1 415 Unsupported Media Type

