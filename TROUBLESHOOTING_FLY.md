# Troubleshooting - Erro ao Criar App no Fly.io

## Problemas Comuns e Soluções

### 1. Nome do App Já Existe
**Erro**: "App name already taken" ou similar

**Solução**: 
- Escolha um nome único no `fly.toml`
- Exemplo: `app = "go-test-seu-nome-unico-123"`

### 2. Erro de Região
**Erro**: "Region not available" ou similar

**Solução**:
- Tente outra região no `fly.toml`:
  - `primary_region = "iad"` (Washington, EUA)
  - `primary_region = "sjc"` (San Jose, EUA)
  - `primary_region = "scl"` (Santiago, Chile)

### 3. Erro de Build/Dockerfile
**Erro**: "Build failed" ou "Dockerfile not found"

**Solução**:
- Certifique-se que o `Dockerfile` está na raiz do projeto
- No painel do Fly, verifique:
  - **Working directory**: deixe vazio ou `./`
  - **Config path**: deixe vazio ou `./fly.toml`

### 4. Erro de Repositório
**Erro**: "Repository not found" ou "Access denied"

**Solução**:
- Certifique-se que o repositório está público OU
- Conecte sua conta GitHub ao Fly.io nas configurações

### 5. Erro de Sintaxe no fly.toml
**Erro**: "Invalid configuration"

**Solução**:
- Verifique se não há caracteres especiais no nome do app
- Nome do app deve conter apenas letras minúsculas, números e hífens
- Exemplo válido: `go-test-ids2cq`
- Exemplo inválido: `go_test` (underscore não permitido)

## Passo a Passo Alternativo

Se continuar dando erro, tente criar o app **sem** o fly.toml primeiro:

1. **No painel do Fly.io**:
   - Clique em "New App"
   - Escolha "Deploy from GitHub"
   - Selecione seu repositório
   - **NÃO** selecione "Use existing fly.toml"
   - Configure manualmente:
     - **Region**: gru (São Paulo)
     - **Internal port**: 8000
     - **Machine size**: shared-cpu-1x / 256MB
   - Clique em "Deploy"

2. **Depois do deploy**, o Fly criará um `fly.toml` automaticamente

3. **Faça pull** do `fly.toml` gerado e ajuste conforme necessário

## Verificar Logs

Se o app foi criado mas não está funcionando:
- No painel do Fly: vá em "Monitoring" → "Logs"
- Procure por erros de build ou runtime

## Contato

Se nenhuma solução funcionar, compartilhe:
- A mensagem de erro exata
- Screenshot do erro (se possível)
- O conteúdo do `fly.toml` atual
