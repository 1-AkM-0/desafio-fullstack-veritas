# Mini Kanban Fullstack (React + Go)

Este é um projeto de desafio fullstack que implementa um simples quadro Kanban. O frontend é construído com React e o backend é uma API REST em Go.

---

### Pré-requisitos

Antes de começar, garanta que você tenha as seguintes ferramentas instaladas em sua máquina:

* [**Go**](https://go.dev/doc/install) (versão 1.21 ou superior)
* [**Node.js**](https://nodejs.org/en) (versão 18 ou superior)

---

### Começando

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/1-AkM-0/desafio-fullstack-veritas.git
    cd desafio-fullstack-veritas
    ```

2.  **Instale as dependências do Frontend:**
    ```bash
    # Navegue até a pasta do frontend
    cd ./frontend

    # Instale as dependências
    npm install
    ```

---

### Instruções para Rodar o Projeto

Você precisará de **dois terminais** abertos, um para o backend e outro para o frontend.

#### 1. Terminal 1: Backend 

```bash
# 1. Navegue até a pasta do backend 
cd ./backend

# 2. Rode o servidor Go
go run .

# O servidor estará rodando em http://localhost:5000
```

#### 2. Terminal 2: Frontend
```bash
# 1. Navegue até a pasta do frontend
cd ./frontend

# 2. Inicie o servidor de desenvolvimento
npm run dev

# O servidor estará rodando em http://localhost:5173
```

### Rodando os Testes do Backend

Para verificar a integridade da API, você pode rodar os testes com:

```bash
# 1. Navegue até a pasta do backend
cd ./backend

# 2. Execute os testes
go test -v 
```

### Decisões Técnicas


- Backend:
  - Go Puro (net/http)
    - Em vez de usar um framework (como Gin ou Echo), optei pela biblioteca padrão net/http do Go.
    - Motivo: Para um CRUD simples, a biblioteca padrão é mais que o necessário e não introduz dependências externas.
  - Armazenamento em Memória

      - Os dados das tarefas são armazenados em um slice `[]Task` em memória dentro da struct `TaskStore`.

      - Motivo: Esta abordagem simplifica o setup do projeto (sem necessidade de banco de dados) e é suficiente para os requisitos do desafio. A segurança de concorrência é            garantida pelo uso de sync.RWMutex, que protege o slice contra leituras e escritas simultâneas.

  - Testes de Integração (httptest)

    - O projeto inclui uma suíte de testes para os endpoints da API usando a biblioteca padrão httptest.

    - Motivo: Isso garante que o CRUD funciona como esperado, validando códigos de status, content-type e os dados retornados. Isso previne regressões e documenta o comportamento da API.

      
### Limitações e Melhorias Futuras

Apesar de cumprir os requisitos básicos, exitem melhoriam que poderiam ser implementas em um cenário de produção:

- Persistência de dados:
  - Atualmente os dados estão sendo salvos em memória, o que significa que uma vez que o servidor reinicia, os dados são perdidos. Uma solução seria integrar um banco de dados SQL ou NoSQL para isso, ou salvar em um arquivo JSON mesmo.
- Validação de entradas:
  - Atualmente as mensagens retornadas são strings simples. Uma opção melhor seria um JSON (ex: `{"error": "titulo é obrigatorio"}`), o que facilita o tratamento de erros no frontend
- Cobertura de testes:
  - Os testes atuais são focados nos handlers da API, cobrindo os "caminhos felizes" e falhas de ID. O que garante apenas 65% de cobertura de testes
  - Seria interessante expandir os testes para cobrir mais cenários de falha e adicionar testes de unidade puros para a lógica de negócio como as funções `ValidateStatus` e `ValidateTitle`
