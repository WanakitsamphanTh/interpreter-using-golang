# Lox Interpreter
I'm currently studying formal languages and decided to build an interpreter for  Lox programming language by following the book *Crafting Interpreters* by Robert Nystrom. While the tutorial used Java, I implemented the interpreter in Golang. 

**Milestones**
- Implemented the tokenizer (scanner)
- Implemented the parser using operator-precedence hierarchy
- Implemented expression evaluation
- Implementing syxtax & runtime error
- Implementing statement evaluation

## Grammar
***Types of expressions***
**expression** &rarr; literal | unary | binary | grouping \
**literal** &rarr; NUMBER | STRING | `true` | `false` | `nil` \
**unary** &rarr; (`-`|`!`) expression \
**binary** &rarr; expression operator expression \
**operator** &rarr; `+` | `-` | `*` | `/` | `>` | `>=` | `<` | `<=` | `==` | `!=` \
**grouping** &rarr; `(` expression `)`

***Grammar of expressions in which the order of operations is applied.*** \
**expression** &rarr; equality \
**equality** &rarr; comparison [(`==` | `!=`) comparison]* \
**comparison** &rarr; term [(`>` | `>=` | `<` | `<=`) term]* \
**term** &rarr; factor [(`+` | `-`) factor]* \
**factor** &rarr; unary [(`*` | `/`) unary]* \
**unary** &rarr; (`-`|`!`) unary | primary \
**primary** &rarr; literal | grouping \

## Reference
*Crafting Interpreters* by Robert Nystrom 