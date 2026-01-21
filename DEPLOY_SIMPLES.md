# Deploy Simples - SEM Cartão de Crédito

## 🎯 Opção 1: Render (Mais Fácil)

### Passo a Passo:

1. **Acesse**: https://render.com
2. **Login**: Use sua conta GitHub
3. **Criar Web Service**:
   - Clique em "New +" → "Web Service"
   - Conecte seu repositório: `felipehoffmeiste-beep/go_test`
   - Clique em "Connect"

4. **Configurações** (IMPORTANTE):
   ```
   Name: go-soap-test
   Region: São Paulo (ou mais próximo)
   Branch: main
   Root Directory: (deixe VAZIO)
   Runtime: Docker
   Dockerfile Path: Dockerfile
   Docker Build Command: (deixe VAZIO)
   Docker Start Command: (deixe VAZIO)
   Instance Type: Free
   ```

5. **Environment Variables**:
   - Clique em "Advanced"
   - Adicione: `PORT` = `8000`

6. **Clique em "Create Web Service"**

7. **Aguarde** o build (5-10 minutos na primeira vez)

### ✅ Seu endpoint será:
`https://go-soap-test.onrender.com/rastro`

### ⚠️ Limitação do Plano Gratuito:
- App "dorme" após 15 minutos de inatividade
- Primeira requisição após dormir pode levar 30-60 segundos

---

## 🚀 Opção 2: Koyeb (Sempre Ativo)

### Passo a Passo:

1. **Acesse**: https://www.koyeb.com
2. **Login**: Use sua conta GitHub
3. **Criar App**:
   - Clique em "Create App"
   - Escolha "GitHub"
   - Selecione: `felipehoffmeiste-beep/go_test`

4. **Configurações**:
   ```
   App Name: go-soap-test
   Region: São Paulo
   Build: Dockerfile
   Dockerfile Path: Dockerfile
   Port: 8000
   Plan: Starter (FREE)
   ```

5. **Environment Variables**:
   - Adicione: `PORT` = `8000`

6. **Clique em "Deploy"**

### ✅ Seu endpoint será:
`https://go-soap-test-<seu-nome>.koyeb.app/rastro`

### ✅ Vantagens:
- App **sempre ativo** (não dorme)
- Sem necessidade de cartão
- HTTPS automático

---

## 📋 Checklist Antes de Deploy

- [ ] Repositório está no GitHub: `felipehoffmeiste-beep/go_test`
- [ ] `Dockerfile` está na raiz do repositório
- [ ] `main.go` está configurado para usar variável `PORT`
- [ ] `go.mod` e `go.sum` estão commitados

---

## 🐛 Troubleshooting

### Erro: "Dockerfile not found"
- Verifique se o `Dockerfile` está na raiz do repositório
- Se estiver em subpasta, ajuste o "Dockerfile Path" no painel

### Erro: "Build failed"
- Verifique os logs no painel do Render/Koyeb
- Certifique-se que `go.mod` e `go.sum` estão commitados

### App não responde
- Verifique se a porta está correta (8000)
- Verifique os logs no painel
- No Render: aguarde 30-60s se o app estava "dormindo"

---

## 🎉 Pronto!

Depois do deploy, teste seu endpoint SOAP:
```bash
curl https://seu-app.onrender.com/rastro
```

Ou use o Postman/SoapUI para testar os métodos SOAP:
- `registraPontos`
- `consultaPontos`
