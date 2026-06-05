# Lox Interpreter
I'm currently studying formal languages and decided to build an interpreter for  Lox programming language by following the book *Crafting Interpreters* by Robert Nystrom. While the tutorial used Java, I implemented the interpreter in Golang. 

**Milestones**
- Implemented the tokenizer (scanner)
- Implemented the parser using operator-precedence hierarchy
- Implemented expression evaluation
- Implemented variable declaration and referencing
- Implemented expression statement & print statement execution
- Implemented scope
- Implementing syntax & runtime error

## Expression Grammar
***Types of expressions***
**expression** &rarr; literal | unary | binary | grouping \
**literal** &rarr; NUMBER | STRING | `true` | `false` | `nil` \
**unary** &rarr; (`-`|`!`) expression \
**binary** &rarr; expression operator expression \
**operator** &rarr; `+` | `-` | `*` | `/` | `>` | `>=` | `<` | `<=` | `==` | `!=` \
**grouping** &rarr; `(` expression `)`

***Grammar of expressions in which the order of operations is applied.*** \
**expression** &rarr; assignment \
**assignment** &rarr; IDENTIFIER `=` assignment | equality \
**equality** &rarr; comparison [(`==` | `!=`) comparison]* \
**comparison** &rarr; term [(`>` | `>=` | `<` | `<=`) term]* \
**term** &rarr; factor [(`+` | `-`) factor]* \
**factor** &rarr; unary [(`*` | `/`) unary]* \
**unary** &rarr; (`-`|`!`) unary | primary \
**primary** &rarr; literal | grouping \

## Syntax Grammar
**program** &rarr; declaration* EOF \
**declaration** &rarr; varDecl | statement \
**statement** &rarr; exprStmt | printStmt | block \
**exprStmt** &rarr; expression `;` \
**printStmt** &rarr; `print` expression `;`  \
**varDecl** &rarr; `var` identifier (`=` expression)? `;`
**block** &rarr; `{` statement* `}`

## Reference
*Crafting Interpreters* by Robert Nystrom 
