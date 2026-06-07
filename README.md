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
- While the tutorial treats environment objects (which store its own variable-value hashmap) as a field within the interpreter object, the environment objects in my interpreter are treated as a global variable.
- The tutorial implements the return statement with throw/catch statement. This is not possible in Golang, hence I implemented using addition fields of the environment object instead. The field `@ret_val` is used to store return value of a function. The boolean value whether a function has terminated is stored in `@terminated` which is  true after a function call or a return statement (false by default). Since environments can be nested, the environment created after calling a function will have the field `functionBound` set to true (false by default). If a return statement is executed within any scope inside a function, it will track up to the nearest enclosing environment where the field is true before altering `@terminated`.

*Should you have any suggestions, please feel free to reach me*

## Reference
*Crafting Interpreters* by Robert Nystrom 
