# 🚀 GoCrafter

**GoCrafter** é uma poderosa ferramenta de scaffolding e templating para projetos Go que ajuda você a criar projetos Go prontos para produção com melhores práticas, ferramentas modernas e templates personalizáveis.

[![Build](https://github.com/rafa-mori/gocrafter/actions/workflows/release.yml/badge.svg)](https://github.com/rafa-mori/gocrafter/actions/workflows/release.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E=1.21-blue)](go.mod)
[![Releases](https://img.shields.io/github/v/release/rafa-mori/gocrafter?include_prereleases)](https://github.com/rafa-mori/gocrafter/releases)

---

## ✨ Funcionalidades

- 🎯 **Criação Interativa de Projetos** - Assistente guiado para configuração de projetos
- 📦 **Múltiplos Templates** - API REST, ferramentas CLI, microsserviços, serviços gRPC e mais
- ⚙️ **Configuração Inteligente** - Integração com banco de dados, cache, autenticação e DevOps
- 🛠️ **Ferramentas Modernas** - Docker, Kubernetes, CI/CD, documentação Swagger
- 🎨 **Personalizável** - Estenda com seus próprios templates
- 🚀 **Pronto para Produção** - Melhores práticas e estrutura profissional

## 🏃‍♂️ Início Rápido

### Instalação

```bash
# Usando Go install
go install github.com/rafa-mori/gocrafter@latest

# Ou baixar das releases
curl -sSL https://github.com/rafa-mori/gocrafter/releases/latest/download/gocrafter-linux-amd64.tar.gz | tar xz
```

### Crie Seu Primeiro Projeto

```bash
# Modo interativo (recomendado para primeira vez)
gocrafter new

# Modo rápido
gocrafter new minha-api --template api-rest

# Listar templates disponíveis
gocrafter list

# Obter detalhes do template
gocrafter info api-rest
```

## 📦 Templates Disponíveis

| Template | Descrição | Funcionalidades |
|----------|-----------|-----------------|
| **api-rest** | Servidor de API REST | Framework Gin, middleware, health checks, Swagger |
| **cli-tool** | Aplicação de linha de comando | Framework Cobra, subcomandos, configuração |
| **microservice** | Arquitetura de microsserviços | gRPC + HTTP, service discovery, métricas |
| **grpc-service** | Serviço gRPC puro | Protocol buffers, streaming, service mesh |
| **worker** | Processador de jobs em background | Integração com filas, retry, monitoramento |
| **library** | Biblioteca/pacote Go | Documentação, testes, workflows CI/CD |

## 🎯 Exemplo: Criando uma API REST

```bash
$ gocrafter new minha-api-blog --template api-rest
🚀 Iniciando geração do projeto...
✅ Projeto gerado com sucesso!
📁 Localização: minha-api-blog

Próximos passos:
  cd minha-api-blog
  make run    # Iniciar a aplicação
  make test   # Executar testes
  make build  # Compilar a aplicação
```

**Estrutura do projeto gerado:**

```text
minha-api-blog/
├── cmd/main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── config/             # Gerenciamento de configuração
│   ├── handler/            # Handlers HTTP
│   ├── middleware/         # Middleware HTTP
│   ├── model/              # Modelos de dados
│   ├── repository/         # Camada de acesso a dados
│   └── service/            # Lógica de negócio
├── pkg/                    # Pacotes públicos
├── Makefile               # Automação de build
├── Dockerfile             # Configuração de container
├── docker-compose.yml     # Ambiente de desenvolvimento
├── .env.example           # Template de variáveis de ambiente
└── README.md              # Documentação do projeto
```

## ⚙️ Opções de Configuração

GoCrafter suporta configuração extensiva através de prompts interativos:

### Suporte a Bancos de Dados

- **PostgreSQL** - Pronto para produção com connection pooling
- **MySQL** - Banco relacional de alta performance
- **MongoDB** - Banco NoSQL baseado em documentos
- **SQLite** - Banco embarcado para desenvolvimento

### Cache

- **Redis** - Store de estrutura de dados em memória
- **Memcached** - Sistema de cache de alta performance
- **In-Memory** - Cache Go integrado

### Autenticação

- **JWT** - Autenticação com JSON Web Token
- **OAuth2** - Provedores de autenticação de terceiros
- **API Keys** - Autenticação simples com chaves de API

### Integração DevOps

- **Docker** - Containerização com multi-stage builds
- **Kubernetes** - Manifests de deployment para produção
- **CI/CD** - GitHub Actions, GitLab CI, Jenkins, Azure DevOps

## 🛠️ Uso Avançado

### Modo Interativo

```bash
$ gocrafter new
🚀 Bem-vindo ao GoCrafter - Gerador de Projetos Go!
Vamos criar seu projeto Go perfeito juntos...

? Qual o nome do seu projeto? minha-api-incrivel
? Qual o nome do módulo Go? github.com/usuario/minha-api-incrivel
? Que tipo de projeto você quer criar? api-rest - Servidor API REST com HTTP
? Qual banco de dados você quer usar? postgres
? Você quer adicionar uma camada de cache? redis
? Quais funcionalidades adicionais você quer incluir? [Use setas para mover, espaço para selecionar]
  ◯ Autenticação (JWT)
  ◉ Documentação da API (Swagger)
  ◉ Health Checks
  ◉ Métricas (Prometheus)
  ◯ Distributed Tracing
```

### Modo Rápido

```bash
# Criar API com funcionalidades específicas
gocrafter new blog-api \
  --template api-rest \
  --output ./projetos \
  --config api-config.json

# Criar ferramenta CLI
gocrafter new minha-cli \
  --template cli-tool \
  --quick

# Criar microsserviço
gocrafter new user-service \
  --template microservice
```

### Informações do Template

```bash
# Listar todos os templates com descrições
gocrafter list

# Obter informações detalhadas do template
gocrafter info api-rest

# Mostrar estrutura do template
gocrafter info microservice --show-structure
```

## 📚 Documentação

- 📖 [**Guia do Usuário**](user-guide.md) - Documentação completa de uso
- 🛠️ [**Desenvolvimento de Templates**](template-development.md) - Criar templates customizados
- 🏗️ [**Arquitetura**](architecture.md) - Como o GoCrafter funciona
- 🎯 [**Exemplos**](examples/) - Exemplos de projetos e tutoriais
- 🤝 [**Contribuindo**](CONTRIBUTING.md) - Como contribuir

## 🌍 Suporte a Idiomas

- [🇺🇸 English](../README.md)
- [🇧🇷 Português](README.pt-BR.md)

## 🤝 Contribuindo

Damos as boas-vindas a contribuições! Por favor, veja nosso [Guia de Contribuição](CONTRIBUTING.md) para detalhes.

1. Faça fork do repositório
2. Crie sua branch de feature (`git checkout -b feature/funcionalidade-incrivel`)
3. Commit suas mudanças (`git commit -m 'Adiciona funcionalidade incrível'`)
4. Push para a branch (`git push origin feature/funcionalidade-incrivel`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](../LICENSE) para detalhes.

## 🙏 Agradecimentos

- [Cobra](https://github.com/spf13/cobra) - Framework CLI
- [Survey](https://github.com/AlecAivazis/survey) - Prompts interativos
- [Gin](https://github.com/gin-gonic/gin) - Framework web HTTP
- [Logrus](https://github.com/sirupsen/logrus) - Logging estruturado

---

Feito com ❤️ por [@rafa-mori](https://github.com/rafa-mori)

[⭐ Nos dê uma estrela se você acha o GoCrafter útil!](https://github.com/rafa-mori/gocrafter)