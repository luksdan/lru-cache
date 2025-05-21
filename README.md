# 🔄 LRU Cache em Go

Este projeto implementa uma estrutura de cache LRU (Least Recently Used) simples em Go, com suporte à execução via Docker.

## 📦 O que é um LRU Cache?

Um **LRU Cache** (Least Recently Used) é uma estrutura que armazena um número limitado de pares `chave:valor`. Quando o limite de capacidade é atingido, ele remove o item menos recentemente acessado para dar lugar a um novo.

---

## 📁 Estrutura do Projeto

- `main.go`: Implementação da estrutura LRU em Go.
- `Dockerfile`: Arquivo para build da imagem Docker do projeto.

---

## 🚀 Como rodar

### ✅ Pré-requisitos

- [Go](https://golang.org/doc/install) instalado (opcional, se usar Docker)
- [Docker](https://www.docker.com/) instalado

### 🔧 Rodando localmente (sem Docker)

```bash
go run main.go
```
## 🐳 Rodando com Docker
-Build da imagem:
```bash
docker run --rm luksdan/lru-cache:1.0
```
-Executar o container:
```bash
docker build -t lru-cache .
docker run --rm lru-cache
```
## 📌 Exemplo de uso
```go
cache := NewLRUCache(3)

cache.Set("a", 1)
cache.Set("b", 2)
fmt.Println(cache.Get("b")) // 2

cache.Set("c", 3)
cache.Set("d", 4)

fmt.Println(cache.Get("a")) // -1 (removido por ser o menos usado)
```