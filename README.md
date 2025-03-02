# Order Open Delivery Go Service

## Descrição
Este é um projeto desenvolvido em Go.

## Como Executar o Projeto

### 1. Instalar as Dependências
Certifique-se de ter o Go instalado e execute:
```sh
 go mod tidy
```

### 2. Para usar o RabbitMQ, execute o comando:

```sh
docker-compose up -d
```

A interface de gerenciamento do RabbitMQ pode ser acessada em:
```
http://localhost:15672
```
Usuário: `neverknow`  
Senha: `neverknow`

### 3. Rodar o Servidor
Para iniciar o servidor, rode o comando:
```sh
make run
```

## Estrutura de Pastas

```
/order_open_delivery_go_service
│── cmd/                         
│   ├── server/                   
│   │   ├── main.go               
│── internal/                     
│   ├── order/                    
│   │   ├── repository/           
│   │   │   ├── order_repository.go 
│   │   ├── service/              
│   │   │   ├── order_service.go   
│   │   ├── validation/           
│   │   │   ├── order_validation.go 
│   ├── queue/                    
│   │   ├── queue_client.go       
│── api/                          
│   ├── order-proto/              
│   │   ├── order_create.proto     
│   │   ├── order_create_grpc.go   
│   │   ├── order_create.pb.proto  
│── .air.toml                      
│── .env                           
│── .env.example                   
│── .gitignore                      
│── docker-compose.yml              
│── go.mod                          
│── go.sum                          
│── Makefile                        
│── README.md                       
```

## Licença
Este projeto está sob a licença MIT. Para mais detalhes, consulte o arquivo `LICENSE`.

