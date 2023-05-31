# Projeto de Estudos da Linguagem Go
Este é um projeto de estudos da linguagem Go, que tem como objetivo aprender os conceitos básicos e avançados dessa linguagem.

![Go](go1.png)

## Introdução
Go é uma linguagem de programação criada pelo Google em 2007. Ela é compilada, estaticamente tipada, com coleta de lixo e suporte a concorrência. Ela foi projetada para ser simples, eficiente e produtiva.

Go tem uma sintaxe clara e concisa, que facilita a leitura e a escrita do código. Ela também tem uma biblioteca padrão abrangente, que oferece diversas funcionalidades úteis para o desenvolvimento de aplicações.

Go é uma linguagem multiplataforma, que pode ser executada em diferentes sistemas operacionais, como Windows, Linux e MacOS. Ela também pode ser usada para desenvolver diferentes tipos de aplicações, como web, desktop, móveis e embarcadas.

Go é uma linguagem em constante evolução, que recebe atualizações periódicas com novas funcionalidades e melhorias. A versão mais recente da linguagem é a Go 1.17, lançada em agosto de 2021.

## Instalação
Para instalar a linguagem Go no seu computador, você pode seguir os passos descritos na página oficial: https://golang.org/doc/install.

Basicamente, você precisa baixar o arquivo binário correspondente ao seu sistema operacional e arquitetura, extrair o conteúdo em uma pasta de sua preferência e configurar as variáveis de ambiente ```GOROOT``` (que indica onde o Go está instalado) e ```GOPATH``` (que indica onde os seus projetos em Go serão armazenados).

Você também pode verificar se a instalação foi bem sucedida executando o comando ```go version``` no terminal ou prompt de comando. Ele deve retornar a versão do Go instalada no seu computador.

## Sintaxe
A sintaxe da linguagem Go é baseada na linguagem C, mas com algumas diferenças e simplificações. Por exemplo:

* Go não usa ponto-e-vírgula para terminar as instruções. Em vez disso, ele usa as quebras de linha para determinar o fim das instruções.
* Go não usa parênteses para delimitar os blocos de código. Em vez disso, ele usa chaves {}.
* Go não usa palavras-chave como public, private ou protected para definir o nível de acesso dos elementos do código. Em vez disso, ele usa a letra inicial maiúscula ou minúscula para indicar se o elemento é exportado (visível fora do pacote) ou não.
* Go não usa herança ou classes para definir tipos personalizados. Em vez disso, ele usa structs (estruturas de dados) e interfaces (conjuntos de métodos) para modelar os conceitos do domínio do problema.
* Go não usa exceções ou try/catch para tratar erros. Em vez disso, ele usa valores de retorno múltiplos e verificações explícitas dos erros.
* Um programa em Go consiste em um ou mais arquivos com a extensão ```.go```, que são organizados em pacotes (packages). Cada arquivo deve começar com a declaração do pacote ao qual ele pertence, usando a palavra-chave ```package```. Por exemplo:
    
        package main

O pacote ```main``` é especial, pois indica que o arquivo contém uma função main, que é o ponto de entrada do programa. A função main deve ter a seguinte assinatura:

        func main() {
            // código do programa
        }

Para executar um programa em Go, você pode usar o comando go run seguido do nome do arquivo. Por exemplo:

        go run main.go

Isso irá compilar e executar o programa em um único passo. Se você quiser apenas compilar o programa sem executá-lo, você pode usar o comando ```go build```, que irá gerar um arquivo executável com o mesmo nome do pacote. Por exemplo:

        go build main.go

Isso irá gerar um arquivo chamado ```main``` , que você pode executar diretamente no terminal ou prompt de comando.




