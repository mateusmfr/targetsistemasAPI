# Projeto Go - Target Sistemas API

Este projeto contém algumas funções implementadas em Go para o teste do processo seletivo da empresa Target Sistemas.

## Funcionalidades

- **N1 - Soma**: Calcula a soma de 1 até um índice específico.
- **N2 - Sequência de Fibonacci**: Verifica se um número pertence à sequência de Fibonacci.
- **N3 - Faturamento Diário**: Calcula o menor, maior valor de faturamento e o número de dias com faturamento acima da média.
- **N4 - Percentual de Faturamento por Estado**: Calcula o percentual de faturamento de cada estado em relação ao total.
- **N5 - Inversão de String**: Inverte os caracteres de uma string.
  
## Como Rodar

Antes de rodar o projeto, verifique se o Go está instalado na sua máquina. Caso não tenha o Go, você pode instalá-lo a partir do site oficial: [Go Install](https://golang.org/dl/).

1. Clone este repositório:
    ```bash
    git clone git@github.com:mateusmfr/targetsistemasAPI.git
    cd targetsistemasAPI
    ```

2. Instale as dependências necessárias:
    ```bash
    go mod tidy
    ```

3. Compile e execute o código:
    ```bash
    go run main.go
    ```

4. A API estará disponível em `http://localhost:8080`.

## Testando a API

Você pode testar a API usando ferramentas como [Postman](https://www.postman.com/) ou direto pelo navegador:

1. **Soma**:
    ```bash
    http://localhost:8080/n1
    ```

2. **Sequência de Fibonacci**:
    ```bash
    http://localhost:8080/n2
    ```

3. **Faturamento Diário**:
    ```bash
    http://localhost:8080/n3
    ```

4. **Percentual de Faturamento por Estado**:
    ```bash
    http://localhost:8080/n4
    ```

5. **Inversão de String**:
    ```bash
    http://localhost:8080/n5
    ```

## Observação:

Durante o desenvolvimento desta API, optei por usar o método w.Write para exibir as respostas diretamente na tela. Essa decisão foi tomada devido ao tempo limitado para o desenvolvimento. Caso eu tivesse optado por retornar as respostas no formato JSON, talvez eu não conseguiria concluir o projeto dentro do prazo estipulado.

## 🗣 Me dê um feedback!

Ficarei muito grato se você puder compartilhar seu feedback sobre este projeto. Seu retorno é essencial para o meu aprendizado e aprimoramento contínuo. Se você tiver qualquer sugestão, crítica construtiva ou comentário, por favor, não hesite em entrar em contato comigo.

Meu email: `mateus.mfr10@gmail.com`