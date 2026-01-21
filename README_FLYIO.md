# Deploy no Fly.io (pelo web via GitHub Actions)

## 1) Pré-requisitos
- Repositório no GitHub com este projeto
- Conta no Fly.io

## 2) Criar app no Fly.io (painel web)
1. Entre no painel do Fly.io
2. Crie um novo App (anote o **nome do app**)

## 3) Ajustar `fly.toml`
No arquivo `fly.toml`, troque:
- `app = "SEU-APP-NAME-AQUI"` pelo nome do app criado no Fly.io

## 4) Criar token no Fly.io e cadastrar no GitHub
1. No Fly.io, gere um **API Token**
2. No GitHub (seu repo): **Settings → Secrets and variables → Actions**
3. Crie um secret chamado **`FLY_API_TOKEN`** com o token do Fly.io

## 5) Push no `main`
Quando você der push no branch `main`, o GitHub Actions vai rodar o deploy automaticamente.

## Endpoint
- SOAP: `https://<seu-app>.fly.dev/rastro`

