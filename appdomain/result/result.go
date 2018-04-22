package result

import (
	"fmt"
)

func init() {
	fmt.Println("package: appdomain/result has been initialized!")
}

type Result struct {
	Succes bool
	Errors []string
	Data   IEnitity
}
