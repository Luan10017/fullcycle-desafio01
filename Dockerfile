FROM golang:1.25.6

# instalar compilador C para sqlite3 (caso use go-sqlite3)
RUN apt-get update && apt-get install -y build-essential && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# copiar módulos e baixar antes para cache
COPY go.mod go.sum ./
# use tidy to ensure go.mod/go.sum are consistent
RUN go mod tidy

# copiar código
COPY . .

# habilitar CGO (por padrão já está 1 em imagens linux, mas definimos explicitamente)
ENV CGO_ENABLED=1

# compilação de exemplo (apenas para gerar binário)
RUN go build -o /usr/local/bin/app .

CMD ["/usr/local/bin/app"]
