# Build en 2 étapes
# - 1: creation d'un environnement de compilation pour le serveur go
# - 2: utilisation du produit de la compilation précédente pour créer une autre image plus légère qui ne contient que l'executable

# 1: Compilation du server
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go build -o=/app/server ./cmd/server

# 2: Creation de l'environnement d'execution
FROM debian:stable
WORKDIR /app
# RUN apk add vim \
#     && apk add curl \
#     && apk add git 
COPY --from=builder /app/server .
COPY ./ui /app/ui
CMD ["./server"]
