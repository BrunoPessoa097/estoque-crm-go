## Estoque CRM 
Um CRM simples sem acesso ao banco de dados para fim de mostrar simples para uma loja geral para melhor atendimento do cliente

### Informações sobre o projeto
- **Versão:** 1.0.0
- **Status:** Concluido
- **Licença:** Proprietária

### Tecnologia usadas
- [Golang(go)](https://go.dev/) - (versão: 1.25)
- [Docker](https://www.docker.com/)

### Funcionalidades
- **Cadastro** (nome, preço, estoque)
- **Listar produtos**
- **Venda**
- **Histórico** (venda, adição, alteração)

### Estrutura do Projeto

```
.
├── cmd
│   └── main.go         # ponto de entrada 
├── historico
│   └── historico.go    # Registro de histórico
├── model
│   └── produto.go      # Models de Produto e histórico
├── produto
│   └── produto.go      # Produtos integrações
├── LICENSE             # Licença
├── Dockerfile          # Docker File
├── go.mod              # Módulo go
└── README.md           # Readme
```

### Como Executar
1. Clone o projeto
```bash
    git clone https://github.com/BrunoPessoa097/estoque-crm-go.git
```
2. Acesse o diretorio
```bash
    cd estoque-crm-go/estoque-crm-go
```
3. Execute o projeto
```bash
    go run ./cmd
```

### Executar via docker
Siga os mesmos passos anterior até o passo 2, e siga os próximos passos:

1. Monte a imagem
```bash
    docker build -t estoque-crm-go .
```
2. Executar o container
```bash
    docker run -it estoque-crm-go
```

### Exemplo do projeto em execussão 
```cshap
--- Menu ---
[1]- Adicionar produto
[2]- Listar produtos
[3]- Vender produto
[4]- Adicionar Estoque ao Produto
[5]- Histórico
[0]- Sair
Qual sua escolha?


Nome do produto: Pizza
Preço do produto: 32.99
Quantidade em estoque: 5
Produto adicionado
Voltar para o menu?
[1]- Sim
[2]- Não

--- Histórico ---
Tipo: Adicionado | Nome: espeto | Preço: 1.99 | Estoque: 12 
------
Voltar para o menu?
[1]- Sim
[2]- Não

```

### Estrutura Lógica
```go
type Produto struct {
	Nome    "Pizza"
	Preco   25.99
	Estoque 15
}

type Historico struct {
	Tipo    "Venda"
	Produto *Produto
}
```


## Criado Por
* **Nome**: Bruno Pessoa
* **Área**: Desenvolvedor NodeJs|Typescript|Javascript
* **Formado**: UNIGRANDE - Centro Universitário da Grande Fortaleza.
* **Curso**: Sistemas para _Internet_.
* **Git Hub**: [github.com/BrunoPessoa097](https://github.com/BrunoPessoa097/api-agenda.git)
* **LinkedIn**: [www.linkedin.com/in/bruno-pessoa-097](https://www.linkedin.com/in/bruno-pessoa-097/)

## Licença

Este projeto é **proprietário**.  
O código está disponível **apenas para fins de visualização e demonstração em portfólio**.  

- **Não é permitido** copiar, modificar, redistribuir ou utilizar em projetos pessoais, acadêmicos ou comerciais.  
- Para obter permissão de uso, entre em contato: [bruno-pessoa009@outlook.com](mailto:bruno-pessoa009@outlook.com).  

Texto completo da licença em: [LICENSE](./LICENSE)