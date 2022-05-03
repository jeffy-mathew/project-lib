package project_lib

import (
	"fmt"
	"github.com/jeffy-mathew/project-lib/lib"
)

func main() {
	fmt.Println(string(lib.GetData()))
}

func Data() string {
	return lib.GetData()
}
