# Projeto de Conversão de Arquivo Excel para CSV

Este projeto é um programa escrito em Go (Golang) que lê um arquivo Excel chamado "cartoes.xlsx" e converte os dados em um arquivo CSV chamado "_cartoes_.csv". Ele realiza algumas operações de limpeza e formatação dos dados durante o processo de conversão.

## Requisitos

- [Go (Golang)](https://golang.org/dl/): Certifique-se de ter o Go instalado no seu sistema.
- [Excelize](https://github.com/xuri/excelize/v2): Esta biblioteca é usada para lidar com arquivos Excel no projeto. Você pode instalá-la usando `go get -u github.com/xuri/excelize/v2`.

## Funcionalidade

O programa lê o arquivo "cartoes.xlsx", que contém uma planilha chamada "Results". Ele extrai os dados dessa planilha, faz algumas transformações nos dados e, em seguida, gera um arquivo CSV "_cartoes_.csv" com os seguintes campos:

- **Número de Série**: O número de série do cartão.
- **CPF**: O CPF do titular do cartão, com formatação limpa (sem pontos ou traços).
- **Valor da Carga**: Espaço em branco (não fornecido nos dados originais).
- **Observação**: O nome do titular do cartão, limitado a 35 caracteres (caso seja maior).

## Como Usar

1. Certifique-se de que o arquivo "cartoes.xlsx" esteja presente no mesmo diretório em que você executará o programa.

2. Instale a biblioteca Excelize, se ainda não o tiver feito, com o seguinte comando:

   ```bash
   go get -u github.com/xuri/excelize/v2
   ```

3. Compile e execute o programa:

   ```bash
   go run main.go
   ```

4. Após a execução bem-sucedida, o arquivo "_cartoes_.csv" será gerado no mesmo diretório.

5. Você pode abrir o arquivo CSV em qualquer aplicativo de planilha, como o Microsoft Excel, para visualizar os dados.

## Estrutura do Projeto

- `main.go`: O código-fonte principal do programa.
- `cartoes.xlsx`: O arquivo Excel de entrada que contém os dados.
- `_cartoes_.csv`: O arquivo CSV de saída gerado pelo programa.

## Exemplo de Saída

Após a execução bem-sucedida do programa, você verá a seguinte mensagem no console:

```
Arquivo gerado com sucesso!
Pressione ENTER para sair
```

Isso indica que o arquivo CSV foi gerado com êxito.

## Observações

- Certifique-se de que o arquivo "cartoes.xlsx" está localizado no mesmo diretório em que o programa é executado.
- Os dados do arquivo de entrada são lidos a partir da planilha "Results" do arquivo Excel.
- O campo "Valor da Carga" no arquivo CSV é preenchido com um espaço em branco, pois não é fornecido nos dados originais.
