# syntax=docker/dockerfile:1

# Versão escolhida do golang
FROM golang:1.18

# Diretório da imagem
WORKDIR /app

# Copiar os arquivos que contém as dependências
COPY go.mod ./
COPY go.sum ./

# Baixar dependências
RUN go mod download

# Copiar todos os arquivos da api
COPY . .

# Fazer a build da api
RUN go build -o /docker-minhascriptos-api

# Porta da api
EXPOSE 8080:8080

# Permissão para executar comandos
CMD [ "/docker-minhascriptos-api" ]