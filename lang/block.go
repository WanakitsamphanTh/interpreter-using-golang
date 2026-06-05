package lang

type Block struct {
	statements []Statement
	env *Environment
}

func NewBlock(stmts []Statement, env *Environment) Statement {
	return &Block{stmts, env}
}

func (b *Block) Execute() error {
	for _, stmt := range b.statements {
		err := stmt.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}