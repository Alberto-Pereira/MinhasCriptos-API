## 💼 MinhasCriptos – API Rest (WIP)

## 📝 Sumário

- [Sobre](#about)
- [Estrutura](#pattern)
- [Como baixar e executar:](#comobaixar)
    - [Localmente](#executarlocal)
    - [Docker](#executardocker)
- [Executar as operações com:](#op)
    - [Swagger](#swagger)
    - [API Clients(Postman, Insomnia, etc.)](#opapiclient)
- [Banco de Dados - PostgreSQL](#dbpostgres)
- [Futuras atualizações](#update)

## 💻 Sobre: <a name="about"></a>

MinhasCriptos – API é uma API Rest desenvolvida com GO para atender a aplicação web <a href="https://github.com/Alberto-Pereira/MinhasCriptos">MinhasCriptos</a>.

## 🏛 Estrutura: <a name="pattern"></a>

A API é dividida em duas partes:
  
  - Usuário:
    - Cadastrar usuário.
    - Obter usuário.
    - Obter dinheiro do usuário.
  
  - Cripto (Moeda): 
    - Adicionar moeda.
    - Editar moeda.
    - Deletar moeda.
    - Obter moedas.
    - Obter moedas com parâmetros personalizados.

## 🔥 Como baixar e executar: <a name="comobaixar"></a>
## 🏠 Localmente: <a name="executarlocal"></a>

- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/MinhasCriptos-API
```
- No arquivo <a href="https://github.com/Alberto-Pereira/MinhasCriptos-API/blob/main/repository/db.go">db.go</a>, linha 24, trocar o parâmetro "user" para "host".
- Antes:
```bash
psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, user, port, dbname)
```
- Depois:
```bash
psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
```
- Fazer a build da API:
```bash
go build
```
- É preciso criar o [banco de dados](#dbpostgres).
- Alterar as constantes no arquivo <a href="https://github.com/Alberto-Pereira/MinhasCriptos-API/blob/main/repository/db.go">db.go</a> para combinar com as suas preferências:
```bash
const (
	host     = "localhost"     // host
	port     = 5432            // porta
	user     = "postgres"      // usuário
	password = 123             // senha
	dbname   = "minhascriptos" // nome do banco de dados
)
```
- Executar o arquivo minhascriptos.exe.
- Escolher como [executar as operações](#op).

## 🐳 Docker: <a name="executardocker"></a>
- Ter o <a href="https://www.docker.com/products/docker-desktop/">Docker</a> e Docker Compose instalado.
- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/MinhasCriptos-API
```

## 🐘 Banco de Dados - PostgreSQL: <a name="dbpostgres"></a>
- Criar banco de dados com o nome:
```bash
minhascriptos
```
- Copiar e colar todas as linhas do arquivo de <a href="https://github.com/Alberto-Pereira/MinhasCriptos-API/blob/main/script/minhascriptos.sql">script</a> no seu SGBD para criar o restante do banco de dados.
- Para começar, você deve cadastrar um usuário usando seu SGBD preferido ou utilizando umas das [operações](#op).

## 🧵 Executar as operações com: <a name="op"></a>
