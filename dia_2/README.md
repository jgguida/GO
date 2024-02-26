# Dia 2 estudando GO

Conceitos Importantes em Go
Pacotes: Todo programa em Go é composto por pacotes. Seu programa começa a rodar no pacote main. Isso é similar ao que você já deve ter feito, definindo seu ponto de entrada.

Imports: Você já deve ter notado que para usar certas funcionalidades, como servidores HTTP ou manipulação de JSON, você precisa importar pacotes. Em Go, isso é feito com a declaração import.

Funções: Go tem suporte para funções com múltiplos retornos, o que pode ser muito útil, por exemplo, ao retornar um resultado e um erro de uma função.

Tipos: Além dos tipos básicos, Go permite a definição de tipos compostos, como structs, que são coleções de campos. Eles são úteis para agrupar dados relacionados.

Interfaces: Uma interface é um tipo que especifica um conjunto de métodos. Um tipo implementa uma interface simplesmente implementando seus métodos. Isso é usado para conseguir polimorfismo em Go.

Goroutines: São funções ou métodos que rodam concorrentemente com outras goroutines. Elas são uma forma leve de executar threads e são uma das características mais poderosas de Go.

Canais: São um meio de comunicação entre goroutines. Eles permitem sincronizar execuções e trocar informações de forma segura.