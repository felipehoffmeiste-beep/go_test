# 🚀 Deploy no Koyeb - Passo a Passo

## ✅ Pré-requisitos
- [x] Repositório no GitHub: `felipehoffmeiste-beep/go_test`
- [x] Dockerfile configurado
- [x] main.go usando variável PORT
- [x] Sem necessidade de cartão de crédito

---

## 📋 Passo a Passo

### 1. Acesse o Koyeb
👉 https://www.koyeb.com

### 2. Faça Login
- Clique em "Sign Up" ou "Log In"
- Escolha **"Continue with GitHub"**
- Autorize o acesso ao GitHub

### 3. Criar Novo App
- No dashboard, clique no botão **"Create App"** (canto superior direito)
- Escolha **"GitHub"** como fonte

### 4. Conectar Repositório
- Selecione seu repositório: **`felipehoffmeiste-beep/go_test`**
- Clique em **"Connect"**

### 5. Configurar o App

#### Aba "Overview":
```
App Name: go-soap-test
Region: São Paulo (ou mais próximo disponível)
```

#### Aba "Build & Deploy":
```
Source: GitHub
Repository: felipehoffmeiste-beep/go_test
Branch: main
Build: Dockerfile
Dockerfile Path: Dockerfile
```

#### Aba "Settings":
```
Port: 8000
Plan: Starter (FREE)
```

### 6. Environment Variables (Opcional mas Recomendado)
- Clique em **"Environment Variables"**
- Adicione:
  ```
  Key: PORT
  Value: 8000
  ```
- Clique em **"Add"**

### 7. Deploy
- Clique no botão **"Deploy"** (canto inferior direito)
- Aguarde o build (5-10 minutos na primeira vez)

---

## 🎉 Pronto!

Seu servidor SOAP estará disponível em:
```
https://go-soap-test-<seu-nome>.koyeb.app/rastro
```

### Endpoints SOAP:
- **registraPontos**: `POST https://seu-app.koyeb.app/rastro`
  - SOAPAction: `urn:RastroAction#registraPontos`

- **consultaPontos**: `POST https://seu-app.koyeb.app/rastro`
  - SOAPAction: `urn:RastroAction#consultaPontos`

---

## ✅ Vantagens do Koyeb

- ✅ **Sempre ativo** - App não "dorme" (diferente do Render)
- ✅ **Sem cartão de crédito** - Plano Starter é gratuito
- ✅ **HTTPS automático** - Certificado SSL incluído
- ✅ **Deploy automático** - Atualiza ao fazer push no GitHub
- ✅ **Logs em tempo real** - Fácil debug

---

## 🐛 Troubleshooting

### Build falha
- Verifique os logs no painel do Koyeb
- Certifique-se que `go.mod` e `go.sum` estão commitados
- Verifique se o Dockerfile está na raiz do repositório

### App não responde
- Verifique se a porta está configurada como `8000`
- Verifique os logs: Dashboard → Seu App → Logs
- Certifique-se que a variável `PORT` está definida

### Erro de autenticação SOAP
- Verifique se você tem clientes cadastrados em `clients.json`
- O app cria o arquivo automaticamente se não existir

---

## 📝 Notas Importantes

⚠️ **Persistência de Dados**: 
- Os arquivos JSON (`clients.json`, `pontos.json`) são armazenados no sistema de arquivos do container
- **Dados são voláteis** e serão perdidos quando o container reiniciar
- Para persistência real, considere usar um banco de dados (PostgreSQL, MongoDB, etc.)

---

## 🔄 Atualizações Futuras

Para atualizar o app:
1. Faça alterações no código
2. Commit e push para o GitHub
3. O Koyeb detecta automaticamente e faz redeploy

Ou manualmente:
- Dashboard → Seu App → "Redeploy"

---

## 📞 Suporte

Se tiver problemas:
- Logs do Koyeb: Dashboard → Seu App → Logs
- Documentação: https://www.koyeb.com/docs
