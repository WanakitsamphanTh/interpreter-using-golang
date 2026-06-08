package lang

import (
	"time"
	"fmt"
)

// Native functions
func InitNativeFunctions(){
	global := current_env
	global.Define("clock", &NativeFn{0, func([]any) (any, disruptive) {return float64(time.Now().Unix()), nil }})
	global.Define("print", &NativeFn{-1, func(params []any) (any,disruptive){
		if len(params) != 0 {
			for _, param := range params {
				fmt.Printf("%v", param)
			}
		}
		return nil, nil
	}})
	global.Define("printLn", &NativeFn{-1, func(params []any) (any,disruptive){
		if len(params) != 0 {
			for _, param := range params {
				fmt.Printf("%v", param)
			}
		}
		fmt.Printf("\n")
		return nil, nil
	}})
	global.Define("scan", &NativeFn{0, func([]any) (any, disruptive) {
		var s string
		_, err := fmt.Scanln(&s)
		return s, err
	}})	
	global.Define("scanNumber", &NativeFn{0, func([]any) (any, disruptive) {
		var s float64
		_, err := fmt.Scanln(&s)
		return s, err
	}})
}