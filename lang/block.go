package lang

type Block struct {
	statements []Statement
	defering Statement
	shared bool
}

func (b *Block) Execute() disruptive {
	if !b.shared {
		NewNestedEnvironment(false)
		defer RetractEnvironment()
	}
	defer func() {
		if b.defering != nil {
			b.defering.Execute()
		}
	}()
	if b.statements != nil {
		for _, stmt := range b.statements {
			err := stmt.Execute()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
