# Solução para Erro de Autenticação GitHub

## Problema
```
remote: Permission to felipehoffmeiste-beep/go_test.git denied to hoffera.
fatal: unable to access 'https://github.com/felipehoffmeiste-beep/go_test.git/': The requested URL returned error: 403
```

## Solução Aplicada
✅ Credenciais antigas do Windows Credential Manager foram removidas.

## Próximos Passos

### Opção 1: Usar Personal Access Token (PAT) - Recomendado

1. **Criar um PAT no GitHub:**
   - Acesse: https://github.com/settings/tokens
   - Clique em "Generate new token" → "Generate new token (classic)"
   - Dê um nome (ex: "Railway Deploy")
   - Selecione o escopo `repo` (acesso completo aos repositórios)
   - Clique em "Generate token"
   - **COPIE O TOKEN** (você só verá ele uma vez!)

2. **Fazer push usando o token:**
   ```bash

   ```
   - Quando pedir usuário: `felipehoffmeiste-beep`
   - Quando pedir senha: **cole o token** (não sua senha do GitHub)

### Opção 2: Usar SSH (Mais Seguro)

1. **Gerar chave SSH (se ainda não tiver):**
   ```bash
   ssh-keygen -t ed25519 -C "felipe.hoffmeiste@edu.univali.br"
   ```

2. **Adicionar chave pública ao GitHub:**
   - Copie o conteúdo de `~/.ssh/id_ed25519.pub`
   - Acesse: https://github.com/settings/keys
   - Clique em "New SSH key"
   - Cole a chave e salve

3. **Alterar remote para SSH:**
   ```bash
   git remote set-url origin git@github.com:felipehoffmeiste-beep/go_test.git
   git push -u origin main
   ```

### Opção 3: Usar GitHub CLI

```bash
# Instalar GitHub CLI (se não tiver)
winget install GitHub.cli

# Login
gh auth login

# Push
git push -u origin main
```

## Verificar Configuração

```bash
# Ver remote atual
git remote -v

# Ver credenciais salvas
cmdkey /list | Select-String -Pattern "github"
```
