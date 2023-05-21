# Dockerfile para rodar projeto utilizando a versão 1.19 do Go

# Imagem base
FROM golang:1.19

# Diretório de trabalho
WORKDIR /go/src/app

# Copia o arquivo go.mod para o diretório de trabalho
COPY . .

# Executa o comando go mod download
RUN go mod download

# Executa o comando go build -o main .
RUN go build -o main .
