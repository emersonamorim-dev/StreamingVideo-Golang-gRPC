# Plataforma de Streaming de Vídeo - Microserviços

Codificação em Golang usando gRPC para implementação que consiste em vários microserviços que juntos formam uma plataforma de streaming de vídeo. Os microserviços são implementados em Go, utilizam MongoDB para armazenamento de dados, Kafka para streaming de eventos e se comunicam via gRPC.

## Microserviços

- `auth-service`: Gerencia autenticação e autorização.
- `user-management-service`: Gerencia dados do usuário.
- `video-streaming-service`: Transmite conteúdo de vídeo para os usuários.
- `video-upload-service`: Gerencia o upload de novo conteúdo de vídeo.

## Pré-requisitos

- Go 1.16 ou superior
- MongoDB 4.4 ou superior
- Kafka
- Docker
- Kubernetes (k8s)

## Instalação

1. Clone este repositório para o seu ambiente local.
2. Navegue até o diretório do projeto.
3. Execute `docker-compose up` para iniciar todos os serviços.

## Configuração

Os serviços se conectam a um servidor MongoDB na URI `mongodb://localhost:27017` e usam o banco de dados `streamingvideo`. Certifique-se de que o MongoDB esteja rodando e acessível nesta URI.

Os serviços gRPC escutam em várias portas. Certifique-se de que estas portas estejam disponíveis.

O Kafka é usado para streaming de eventos entre serviços. Certifique-se de que uma instância do Kafka esteja rodando e acessível.

## Execução

Após iniciar os contêineres Docker com `docker-compose up`, você pode iniciar os serviços executando o script `start-services.sh`:

bash:
./start-services.sh

## Autor
Emerson Amorim
