## 💼 MinhasCriptos – API Rest (WIP)

## 📝 Sumário

- [Sobre](#about)
- [Estrutura](#pattern)
- [Como baixar e iniciar](#comobaixar)
- [Executar as operações com Swagger](#swagger)
- [Futuras atualizações](#update)

## 💻 Sobre: <a name="about"></a>

MinhasCriptos será uma aplicação web para gerência e controle pessoal de criptomoedas. Será possível acompanhar a valorização das moedas adicionadas, o quanto já foi gasto, qual o retorno atual, o histórico de moedas e outras funcionalidades.

## 🏛 Estrutura: <a name="pattern"></a>

MinhasCriptos – API é uma API Rest desenvolvida em GO para atender a aplicação web <a href="https://minhascriptos.netlify.app/">MinhasCriptos</a>.

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

## 🐳 Como baixar e iniciar: <a name="comobaixar"></a>
⚠ Ter o <a href="https://www.docker.com/products/docker-desktop/">Docker</a> e Docker Compose instalado.
- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/MinhasCriptos-API
```
⚠ No diretório clonado:
- Fazer a build da aplicação.
```bash
docker-compose build
```
- Iniciar a aplicação.
```bash
docker-compose up
```
## 🧵 Executar as operações com Swagger: <a name="swagger"></a>

👉 Com a aplicação iniciada, é possível executar operações clicando <a href="http://localhost:8080/swagger/index.html">aqui</a>.

## ⚡ Futuras atualizações: <a name="update"></a>

- Autenticação com JWT:
    - Será implementado autenticação com JWT quando a aplicação atingir a meta de usuários. No momento a aplicação será local.
- Documentação em inglês: 
    - Documentação do código em inglês para melhor entendimento.
