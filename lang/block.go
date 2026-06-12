package lang

type Block struct {
	statements []Statement
	defering   Statement
}

func (b *Block) Execute() disruptive {
	var err disruptive
	err = nil

	NewNestedEnvironment(false)
	defer func() {
		if b.defering != nil {
			err = b.defering.Execute()
		} else {
			err = nil
		}
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
