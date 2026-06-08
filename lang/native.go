package lang

import (
	"time"
	"fmt"
)

// Native functions
func InitNativeFunctions(){
	global := current_env
	global.Define("clock", &NativeFn{0, func([]any) (any, disruptive) {return float64(time.Now().Unix()), nil }})
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