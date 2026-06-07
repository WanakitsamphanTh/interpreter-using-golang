# Lox Interpreter
I'm currently studying formal languages and decided to build an interpreter for  Lox programming language by following the book *Crafting Interpreters* by Robert Nystrom. While the tutorial used Java, I implemented the interpreter in Golang. 

**Milestones**
- Implemented the tokenizer (scanner)
- Implemented the parser using operator-precedence hierarchy
- Implemented expression evaluation
- Implemented variable declaration and referencing
- Implemented expression statement & print statement execution
- Implemented scope
- Implemented control flow (if-else, while and for loop)
- Implementing syntax & runtime error

## Expression Grammar
***Types of expressions***
**expression** &rarr; literal | unary | binary | grouping \
**literal** &rarr; NUMBER | STRING | `true` | `false` | `nil` \
**unary** &rarr; (`-`|`!`) expression | call \
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
**unary** &rarr; (`-`|`!`) unary | call \
**call** &rarr; primary [`(` arguments?`)`]* \
**primary** &rarr; literal | grouping \
**logic_and** &rarr; equality [`and` equality]* \
**logic_or** &rarr; logic_and [`or` logic_and]* \

## Syntax Grammar
**program** &rarr; declaration* EOF \
**declaration** &rarr; varDecl | funcDecl | statement \
**statement** &rarr; exprStmt | printStmt | ifStmt | forStmt | whileStmt | returnStmt | block \
**exprStmt** &rarr; expression `;` \
**printStmt** &rarr; `print` expression `;`  \
**returnStmt** &rarr; `return` expression? `;` \
**ifStmt** &rarr; `if` `(` expression `)` statement [`else` statement]? \
**forStmt** &rarr; `for` `(` [varDecl | exprStmt | `;`] exprStmt `;` expression? expression? `)` statement \
**whileStmt** &rarr; `while` `(` expression `)` statement \
**varDecl** &rarr; `var` identifier [`=` expression]? `;` \
**funcDecl** &rarr; `func` identifier `(` parameters? `)` block \
**block** &rarr; `{` statement* `}`
**arguments** &rarr; identifier [`,` identifier]*

## Reference
*Crafting Interpreters* by Robert Nystrom 
