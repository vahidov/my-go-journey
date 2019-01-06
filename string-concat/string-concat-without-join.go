package string_concat

import (
	"fmt"
	"os"
)

func main() {
	StringConcatWithoutJoin()
}

func StringConcatWithoutJoin() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
