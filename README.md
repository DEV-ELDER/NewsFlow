# NewsFlow 📰

API backend em Go que agrega notícias de múltiplas categorias, com persistência em PostgreSQL e busca concorrente.

🔗 **API ao vivo:** https://newsflow-ycdp.onrender.com/noticias

> ⚠️ Hospedado em plano gratuito — a primeira requisição pode levar ~30-60s para "acordar" o servidor.

## Sobre o projeto

O NewsFlow busca notícias de múltiplas categorias simultaneamente (usando goroutines), remove duplicatas automaticamente, e expõe os dados através de uma API REST própria.

## Funcionalidades

- 🔄 Busca concorrente de múltiplas categorias de notícias (goroutines, channels, WaitGroup)
- 🔒 Proteção contra travamento com `context.Context` (timeout em chamadas externas)
- 🗄️ Persistência em PostgreSQL, com constraint de unicidade evitando duplicatas
- 🌐 API REST própria, com filtro por categoria e por palavra-chave via query parameters
- 🧩 Persistência abstraída por interface (`ArmazenamentoNoticia`), permitindo testes sem depender de banco real
- 🔐 Configuração segura de credenciais via variáveis de ambiente
- ✅ Testes automatizados (table-driven tests + testes com fake/mock)
- 🐳 Containerizado com Docker (build multi-stage)
- ⚙️ CI configurado com GitHub Actions, rodando testes a cada push
- ☁️ Deploy em produção (Render), com CD automático

## Tecnologias

- **Go** — linguagem principal
- **PostgreSQL** — banco de dados
- **net/http** — servidor HTTP (biblioteca padrão, sem framework)
- **lib/pq** — driver PostgreSQL
- **godotenv** — variáveis de ambiente
- **Docker** — containerização
- **GitHub Actions** — integração contínua (CI)
- **Currents API** — fonte externa de notícias
- **Render** — hospedagem (aplicação + banco de dados)

## Endpoints

| Rota | Método | Descrição |
|---|---|---|
| `/noticias` | GET | Lista todas as notícias salvas |
| `/noticias?categoria=sport` | GET | Filtra notícias por categoria |
| `/noticias?palavraChave=IA` | GET | Filtra notícias por palavra-chave no título |
| `/atualizar` | GET | Busca novas notícias na API externa e salva no banco |

## Rodando localmente

### Opção 1 — Direto com Go

\`\`\`bash
git clone https://github.com/DEV-ELDER/NewsFlow.git
cd NewsFlow
go mod download
\`\`\`

Crie um arquivo `.env` na raiz com:

\`\`\`
CURRENTS_API_KEY=sua_chave_aqui
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=newsflow
\`\`\`

\`\`\`bash
go run .
\`\`\`

### Opção 2 — Com Docker

\`\`\`bash
docker build -t newsflow .
docker run -p 8080:8080 --env-file .env -e DB_HOST=host.docker.internal newsflow
\`\`\`

Servidor disponível em `http://localhost:8080`.

## Testes

\`\`\`bash
go test .
\`\`\`

## Sobre este projeto

Este projeto foi construído como parte do meu aprendizado prático de Go para desenvolvimento backend, cobrindo desde fundamentos da linguagem até concorrência avançada (Goroutines, Channels, Mutex, Context), interfaces, testes, containerização e deploy em produção.