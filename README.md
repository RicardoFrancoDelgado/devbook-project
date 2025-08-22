# DevBook API

Documentação da API DevBook — descrição dos arquivos, funções e métodos presentes no projeto.

## Sumário
- Visão geral
- Como rodar
- Variáveis de ambiente
- Estrutura do projeto
- Documentação dos pacotes e funções
- Endpoints principais
- Banco de dados / seed
- Exemplos de requisição

---

## Visão geral
API em Go (Golang) para gerenciamento de usuários, publicações e relacionamento de seguidores. Utiliza MySQL, bcrypt para senhas e JWT para autenticação.

---

## Como rodar
1. Configure as variáveis de ambiente (veja seção Variáveis de ambiente).
2. Inicie o banco MySQL e importe o arquivo `sql/dados.sql`.
3. No diretório da API:
   - `go run main.go` ou `go build` e execute o binário.
4. A API escuta na porta definida em `config.Porta`.

---

## Variáveis de ambiente (exemplo)
- DB_HOST
- DB_PORT
- DB_USER
- DB_PASS
- DB_NAME
- PORTA (ou similar, conforme `config/config.go`)
- CHAVE_JWT (segredo para tokens)

(As chaves exatas e nomes estão em `src/config/config.go`.)

---

## Estrutura do projeto
src/
- autenticacao/
  - token.go
- banco/
  - banco.go
- config/
  - config.go
- controllers/
  - login.go
  - publicacoes.go
  - usuarios.go
- middlewares/
  - middlewares.go
- modelos/
  - publicacao.go
  - senha.go
  - usuario.go
- repositorios/
  - publicacoes.go
  - usuario.go
- respostas/
  - respostas.go
- router/
  - router.go
  - rotas/
    - login.go
    - publicacoes.go
    - rotas.go
    - usuarios.go
- seguranca/
  - seguranca.go
main.go

---

## Documentação dos pacotes e funções

### main.go
- `main()`
  - Carrega configurações (`config.Carregar()`), gera o router (`router.Gerar()`) e inicia o servidor HTTP com `http.ListenAndServe`.

---

### src/config/config.go
- `Carregar()`
  - Lê variáveis de ambiente e configura valores globais (porta do servidor, credenciais do banco etc).
- Variáveis exportadas (ex.: `Porta`, configurações DB) — usadas por outros pacotes.

---

### src/router/router.go
- `Gerar() *mux.Router`
  - Cria e configura o router principal, registra rotas e middlewares, e retorna o router pronto para servir.

Arquivos em `router/rotas/` registram handlers específicos:
- `rotas/login.go` — rotas de autenticação (POST /login).
- `rotas/publicacoes.go` — rotas de publicações.
- `rotas/usuarios.go` — rotas de usuários (CRUD, seguir, seguidores).

---

### src/banco/banco.go
- `Conectar() (*sql.DB, error)`
  - Abre conexão com MySQL usando as configurações e retorna o `*sql.DB`. Gerencia pool e opções de conexão.

---

### src/autenticacao/token.go
- `CriarToken(usuarioID uint64) (string, error)`
  - Gera token JWT assinando com a chave/segredo do projeto. Contém o ID do usuário no payload.
- `ValidarToken(token string) error`
  - Valida o token JWT (assinatura e expiração).
- `ExtrairUsuarioID(r *http.Request) (uint64, error)`
  - Extrai o ID do usuário do token presente no header Authorization da requisição.

---

### src/seguranca/seguranca.go
- `Hash(senha string) ([]byte, error)`
  - Gera o hash bcrypt da senha informada.
- `VerificarSenha(senhaComHash, senhaString string) error`
  - Compara senha em texto com o hash; retorna nil se bater, erro caso contrário.

---

### src/modelos/
- usuario.go
  - `Preparar(etapa string)` — valida e prepara um `Usuario` antes de persistir (ex.: valida campos obrigatórios, sanitiza).
  - Métodos de validação (ex.: checar email, tamanho de campos).
- publicacao.go
  - `Preparar()` — valida/normaliza uma `Publicacao` antes de persistir.
- senha.go
  - `ValidarSenha()` — valida critérios da nova senha (força, confirmação etc).

(As implementações concretas estão nos arquivos correspondentes.)

---

### src/repositorios/usuario.go
Tipo:
- `type Usuarios struct { db *sql.DB }`

Funções/metodos:
- `NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios`
  - Construtor do repositório.
- `Criar(usuario modelos.Usuario) (uint64, error)`
  - Insere um novo usuário e retorna o ID inserido.
- `Buscar(nomeOuNick string) ([]modelos.Usuario, error)`
  - Busca usuários cujo nome ou nick contenha o filtro.
- `BuscarPorID(ID uint64) (modelos.Usuario, error)`
  - Recupera um usuário por seu ID.
- `Atualizar(ID uint64, usuario modelos.Usuario) error`
  - Atualiza nome, nick e email do usuário.
- `Deletar(ID uint64) error`
  - Remove usuário do banco.
- `BuscarPorEmail(email string) (modelos.Usuario, error)`
  - Retorna ID e hash da senha do usuário com dado email.
- `Seguir(usuarioID, seguidorID uint64) error`
  - Insere relação de seguimento (usuario_id, seguidor_id).
- `PararDeSeguir(usuarioID, seguidorID uint64) error`
  - Remove relação de seguimento.
- `BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error)`
  - Lista usuários que seguem o `usuarioID`.
- `BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error)`
  - Lista usuários que `usuarioID` está seguindo.
- `BuscarSenha(usuarioID uint64) (string, error)`
  - Retorna o hash de senha do usuário.
- `AtualizarSenha(usuarioID uint64, senha string) error`
  - Atualiza o hash de senha do usuário no banco.

---

### src/repositorios/publicacoes.go
- `NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes`
- `Criar(publicacao modelos.Publicacao) (uint64, error)`
- `Buscar(usuarioID uint64) ([]modelos.Publicacao, error)`
- `BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error)`
- `Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error`
- `Deletar(publicacaoID uint64) error`

Descrição: operações CRUD para publicações, com queries que respeitam autor e permissões.

---

### src/controllers/login.go
- `Login(w http.ResponseWriter, r *http.Request)`
  - Recebe email e senha, busca hash via repositório, verifica senha com `seguranca.VerificarSenha`, cria JWT com `autenticacao.CriarToken` e retorna token.

---

### src/controllers/usuarios.go
Handlers:
- `CriarUsuario(w http.ResponseWriter, r *http.Request)`
  - Recebe payload, valida (model.Preparar), gera hash (seguranca.Hash), persiste via repositório.
- `BuscarUsuarios(w http.ResponseWriter, r *http.Request)`
  - Retorna lista filtrada de usuários.
- `BuscarUsuario(w http.ResponseWriter, r *http.Request)`
  - Retorna dados de um usuário por ID.
- `AtualizarUsuario(w http.ResponseWriter, r *http.Request)`
  - Atualiza campos do usuário (validação e persistência).
- `DeletarUsuario(w http.ResponseWriter, r *http.Request)`
  - Exclui usuário por ID.
- `SeguirUsuario(w http.ResponseWriter, r *http.Request)`
  - Cria relação de seguimento entre usuários.
- `PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request)`
  - Remove relação de seguimento.
- `BuscarSeguidores(w http.ResponseWriter, r *http.Request)`
  - Lista seguidores de um usuário.
- `BuscarSeguindo(w http.ResponseWriter, r *http.Request)`
  - Lista quem um usuário está seguindo.
- `AtualizarSenha(w http.ResponseWriter, r *http.Request)`
  - Valida senhas, gera novo hash e persiste com `AtualizarSenha`.

---

### src/controllers/publicacoes.go
Handlers:
- `CriarPublicacao(w http.ResponseWriter, r *http.Request)`
- `BuscarPublicacoes(w http.ResponseWriter, r *http.Request)`
- `BuscarPublicacao(w http.ResponseWriter, r *http.Request)`
- `AtualizarPublicacao(w http.ResponseWriter, r *http.Request)`
- `DeletarPublicacao(w http.ResponseWriter, r *http.Request)`

Descrição: CRUD de publicações com controle de autor/autenticação.

---

### src/middlewares/middlewares.go
Middlewares típicos:
- `Autenticacao(next http.Handler) http.Handler`
  - Verifica JWT e injeta ID do usuário no contexto da requisição.
- `Logger(next http.Handler) http.Handler`
  - Log de requisições.
- Outras funções de tratamento de headers/respostas (ex.: content-type JSON).

---

### src/respostas/respostas.go
Funções auxiliares para enviar respostas HTTP:
- `JSON(w http.ResponseWriter, status int, dados interface{})`
- `Erro(w http.ResponseWriter, status int, mensagem string)`

---

## Endpoints principais (resumo)
- POST /login — autenticação (email + senha) → token JWT
- POST /usuarios — criar usuário
- GET /usuarios — buscar usuários (filtro)
- GET /usuarios/{usuarioId} — obter usuário por ID
- PUT /usuarios/{usuarioId} — atualizar usuário
- DELETE /usuarios/{usuarioId} — deletar usuário
- POST /usuarios/{usuarioId}/seguir — seguir usuário
- POST /usuarios/{usuarioId}/parar-de-seguir — parar de seguir
- GET /usuarios/{usuarioId}/seguidores — listar seguidores
- GET /usuarios/{usuarioId}/seguindo — listar seguindo
- POST /usuarios/{usuarioId}/atualizar-senha — atualizar senha
- POST /publicacoes — criar publicação
- GET /publicacoes — listar publicações
- GET /publicacoes/{publicacaoId} — obter publicação por ID
- PUT /publicacoes/{publicacaoId} — atualizar publicação
- DELETE /publicacoes/{publicacaoId} — deletar publicação

(As rotas e verbos estão definidas em `src/router/rotas/*.go`.)

---

## Banco de dados / seed
Arquivo: `sql/dados.sql` contém inserts de exemplo:
- Usuários:
  - usuario1@gmail.com
  - usuario2@gmail.com
  - usuario3@gmail.com
- Tabela `seguidores` com algumas relações.
- Tabela `publicacoes` com publicações de cada usuário.

Observação: as senhas no seed estão em forma de hash bcrypt:
- Exemplo de hash presente: `$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy`
- Para autenticar, é necessário saber a senha original correspondente ao hash ou recriar usuários com senhas conhecidas.

---

## Exemplos de requisição

Login (curl):
```bash
curl -X POST http://localhost:5000/login \
  -H "Content-Type: application/json" \
  -d '{"email":"usuario1@gmail.com","senha":"SENHA_AQUI"}'
```

Criar usuário:
```bash
curl -X POST http://localhost:5000/usuarios \
  -H "Content-Type: application/json" \
  -d '{"nome":"Novo","nick":"novo","email":"novo@email","senha":"senha123"}'
```

Criar publicação (com token):
```bash
curl -X POST http://localhost:5000/publicacoes \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"titulo":"Olá","conteudo":"Meu post"}'
```

---

## Observações finais
- Consulte os arquivos em `src/` para detalhes de implementação e mensagens de erro.
- Para testes locais, é útil importar `sql/dados.sql` e ajustar variáveis de ambiente conforme `src/config/config.go`.
- Para segurança em produção, use variáveis de ambiente seguras e chave JWT forte.
