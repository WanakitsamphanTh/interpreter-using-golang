package lang

type Block struct {
	statements []Statement
	shared bool
}

func (b *Block) Execute() error {
	for _, stmt := range b.statements {
		if !b.shared {
			NewNestedEnvironment()
			defer RetractEnvironment()
		}
		err := stmt.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
