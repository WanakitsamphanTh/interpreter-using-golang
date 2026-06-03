package main
import(
	"fmt"
	"os"
	. "interpreter/lang"
)

func main(){
	if(len(os.Args) > 2){
		fmt.Println("Usage: interpreter <script>")
		return
	} else if(len(os.Args) == 2){
		path := os.Args[1]
		fmt.Println("Run program: ", path)
		err := RunScript(path)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		err := RunREPL()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}