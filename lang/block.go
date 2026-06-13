package lang

type Block struct {
	statements []Statement
}

func (b *Block) Execute() disruptive {
	var err disruptive
	err = nil

	NewNestedEnvironment(false)
	defer func() {
		RetractEnvironment()
	}()

	if b.statements != nil {
		for _, stmt := range b.statements {
			err = stmt.Execute()
			if err != nil {
				return err
			}
		}
	}
	return err
}
