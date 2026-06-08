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
- Implemented function
- implemented closure
- Implemented return statement
- Implemented break & skip statement
- Implementing syntax & runtime error

## Expression Grammar
***Types of expressions***
**expression** &rarr; literal | unary | binary | grouping \
**literal** &rarr; NUMBER | STRING | `true` | `false` | `nil` \
**unary** &rarr; (`-`|`!`) expression | call \
**binary** &rarr; expression operator expression \
**operator** &rarr; `+` | `-` | `*` | `/` | `>` | `>=` | `<` | `<=` | `==` | `!=` \
**grouping** &rarr; `(` expression `)`

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
**logic_or** &rarr; logic_and [`or` logic_and]* 

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
**block** &rarr; `{` statement* `}` \
**arguments** &rarr; identifier [`,` identifier]*

## Key differences
- While the tutorial implements interpreter with Visitor pattern, which means statement executions and expression evaluations are done via visitor objects, expression objects implement Exp(ression) interface with Eval() and statement objects implement Statement interface with Execute().
- While the tutorial treats environment objects (which store its own variable-value hashmap) as a field within the interpreter object, the environment objects in my interpreter are represented by a global variable.
- The tutorial implements the return statement by runtimeException subclassing and throw/catch statement. This is not possible in Golang, hence I took an alternative Golang-native approach. Semantically, I did'nt want break, continue, and return statement to implement the error interface. I alised `any` type to `disruptive` with error and such statements as intended underlying. An error propogates to the outermost scope, a return to the nearest function node, and a break/skip to the nearest loop.

*Should you have any suggestions, please feel free to reach me*

## Reference
*Crafting Interpreters* by Robert Nystrom 
