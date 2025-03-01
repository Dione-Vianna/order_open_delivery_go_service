# Order open delivery go service

## Descrição
Este é um projeto desenvolvido em Go.

## Como Executar o Projeto

### 1. Instalar as Dependências
Certifique-se de ter o Go instalado e execute:
```sh
 go mod tidy
```

### 2. Rodar o Servidor
Para iniciar o servidor, rode o comando:
```sh
make run
```

## Estrutura de Pastas

```
/order_open_delivery_go_service
│── cmd/                 # Ponto de entrada do app
│   ├── server/ 
│   │   ├── main.go      # Arquivo principal para rodar a API
│── api/                 
│   ├── order-proto/     # Definições de API (gRPC, Protobuf)
│   │   ├── order_create.proto
│   │   ├── order_create_grpc.go
│   │   ├── order_create.pb.proto
│── go.mod               # Arquivo de módulo Go
│── go.sum               # Hashes das dependências
│── README.md            # Documentação do projeto
```

## Licença
Este projeto está sob a licença MIT. Para mais detalhes, consulte o arquivo `LICENSE`.

