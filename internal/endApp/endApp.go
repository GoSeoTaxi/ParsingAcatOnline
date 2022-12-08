package endApp

import (
	"fmt"
	"os"
)

func Fin() {
	var end string
	fmt.Println("Нажмите Enter")
	fmt.Scanf("%s\n", &end)
	os.Exit(1)
}
