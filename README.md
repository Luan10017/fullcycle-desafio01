# 🚀 Desafio 01 - Go Expert

> Desafio 01 do curso Go Expert.

---

## ⚙️ Como executar esse projeto

Primeiro crie a tabela do banco SQLite com o comando:
```
sqlite3 ./data/cotacao.db < schema.sql
```
Após criar a tabela com sucesso, execute o server com o comando:
```
go run ./Server/server.go
```
Por fim execute em outro terminal o client da aplicação com o comando:
```
go run ./Client/client.go
```
---
O resultado será um arquivo cotacao.txt na pasta Client com a cotação do Dólar do dia da execução. Caso algum erro aconteça logs para troubleshooting serão exibidos no terminal tanto do server quanto no client.
