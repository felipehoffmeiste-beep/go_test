# Deploy no Render (Alternativa ao Fly.io)

O Render não exige verificação de conta para apps gratuitos.

## Passo a Passo

1. **Acesse**: https://render.com e faça login com GitHub

2. **Criar novo Web Service**:
   - Clique em "New +" → "Web Service"
   - Conecte seu repositório GitHub
   - Selecione o repositório `go_test`

3. **Configurações**:
   - **Name**: `go-soap-test` (ou qualquer nome)
   - **Region**: `São Paulo` (ou mais próximo)
   - **Branch**: `main`
   - **Root Directory**: deixe vazio (ou `./` se o projeto estiver em subpasta)
   - **Runtime**: `Docker`
   - **Dockerfile Path**: `Dockerfile` (ou `PREPS/go_soap_test/Dockerfile` se estiver na raiz)
   - **Docker Build Command**: deixe vazio
   - **Docker Start Command**: deixe vazio
   - **Instance Type**: `Free` (512 MB RAM)
   - **Auto-Deploy**: `Yes`

4. **Environment Variables** (opcional):
   - `PORT`: `8000` (o Render define automaticamente, mas pode configurar)

5. **Clique em "Create Web Service"**

6. **Aguarde o build** (pode levar 5-10 minutos na primeira vez)

## Endpoint

Após o deploy, seu servidor estará em:
- `https://go-soap-test.onrender.com/rastro` (ou o nome que você escolheu)

## Notas

- ⚠️ **Free tier**: Apps gratuitos "dormem" após 15 minutos de inatividade
- ⚠️ **Primeira requisição**: Pode levar 30-60 segundos para "acordar"
- ✅ **Sem verificação de conta**: Funciona imediatamente
- ✅ **HTTPS automático**: Incluído

## Alternativa: Railway (se tiver upgrade)

Se você fizer upgrade no Railway, pode usar o `railway.json` que já está configurado.
