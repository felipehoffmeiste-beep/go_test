# Deploy no Railway

## Pré-requisitos

1. Conta no [Railway](https://railway.app)
2. Railway CLI instalado (opcional, pode usar o dashboard web)

## Passos para Deploy

### Opção 1: Via Dashboard Web (Recomendado)

1. Acesse [railway.app](https://railway.app) e faça login
2. Clique em "New Project"
3. Selecione "Deploy from GitHub repo" (se o código estiver no GitHub) ou "Empty Project"
4. Se escolheu "Empty Project":
   - Clique em "New" → "GitHub Repo" e conecte seu repositório
   - Ou use "Empty Project" e faça upload dos arquivos
5. Railway detectará automaticamente o `Dockerfile`
6. O deploy começará automaticamente

### Opção 2: Via Railway CLI

```bash
# Instalar Railway CLI
npm i -g @railway/cli

# Login
railway login

# Inicializar projeto
railway init

# Deploy
railway up
```

## Configurações

O servidor está configurado para:
- Usar a porta definida pela variável de ambiente `PORT` (Railway define automaticamente)
- Fallback para porta 8000 se `PORT` não estiver definida
- Criar automaticamente os arquivos `clients.json` e `pontos.json` se não existirem

## Endpoints

Após o deploy, o servidor estará disponível em:
- `https://seu-projeto.railway.app/rastro`

### Endpoints SOAP:
- **registraPontos**: POST em `/rastro` com SOAPAction `urn:RastroAction#registraPontos`
- **consultaPontos**: POST em `/rastro` com SOAPAction `urn:RastroAction#consultaPontos`

## Notas Importantes

⚠️ **Persistência de Dados**: Os arquivos JSON são armazenados no sistema de arquivos do container. 
   - Em deploys no Railway, os dados são **voláteis** e serão perdidos quando o container reiniciar
   - Para persistência real, considere usar um banco de dados (PostgreSQL, MongoDB, etc.) ou volumes persistentes

## Variáveis de Ambiente

Não são necessárias variáveis de ambiente adicionais para o funcionamento básico.

## Logs

Para ver os logs do servidor:
- No dashboard: vá em "Deployments" → selecione o deployment → "View Logs"
- Via CLI: `railway logs`
