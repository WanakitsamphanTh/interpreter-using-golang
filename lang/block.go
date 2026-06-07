package lang

//import "fmt";

type Block struct {
	statements []Statement
	shared bool
}

func (b *Block) Execute() error {
	if !b.shared {
		NewNestedEnvironment(false)
		defer RetractEnvironment()
	}
	for _, stmt := range b.statements {
		err := stmt.Execute()
		if err != nil {
			return err
		}
		terminated := current_env.GetValue("@terminated").(bool)
		if terminated {
			return nil
		}
	}
	return nil
}
