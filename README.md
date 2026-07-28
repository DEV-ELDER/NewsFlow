# NewsFlow 📰

API backend em Go que agrega notícias de múltiplas categorias, com persistência em PostgreSQL e busca concorrente.

🔗 **API ao vivo:** https://newsflow-ycdp.onrender.com/noticias

> ⚠️ Hospedado em plano gratuito — a primeira requisição pode levar ~30-60s para "acordar" o servidor.

## Sobre o projeto

O NewsFlow busca notícias de múltiplas categorias simultaneamente (usando goroutines), remove duplicatas automaticamente, e expõe os dados através de uma API REST própria.

## Funcionalidades

- 🔄 Busca concorrente de múltiplas categorias de notícias (goroutines, channels, WaitGroup)
- 🗄️ Persistência em PostgreSQL, com constraint de unicidade evitando duplicatas
- 🌐 API REST própria, com filtro por categoria via query parameter
- 🔐 Configuração segura de credenciais via variáveis de ambiente
- ✅ Testes automatizados (table-driven tests)
- ☁️ Deploy em produção (Render)

## Tecnologias

- **Go** — linguagem principal
- **PostgreSQL** — banco de dados
- **net/http** — servidor HTTP (biblioteca padrão, sem framework)
- **lib/pq** — driver PostgreSQL
- **godotenv** — variáveis de ambiente
- **Currents API** — fonte externa de notícias
- **Render** — hospedagem (aplicação + banco de dados)

## Endpoints

| Rota | Método | Descrição |
|---|---|---|
| `/noticias` | GET | Lista todas as notícias salvas |
| `/noticias?categoria=sport` | GET | Filtra notícias por categoria |
| `/atualizar` | GET | Busca novas notícias na API externa e salva no banco |

## Rodando localmente

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

Servidor disponível em `http://localhost:8080`.

## Sobre este projeto

Este projeto foi construído como parte do meu aprendizado prático de Go para desenvolvimento backend, cobrindo desde fundamentos da linguagem até concorrência, persistência de dados e deploy em produção.