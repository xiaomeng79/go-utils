package path

import (
	"fmt"
)

func ExamplePathExists() {

	fmt.Println(PathExists("example.file"))
	fmt.Println(PathExists("example.json"))

	//Output:
	//true <nil>
	//false <nil>
}
