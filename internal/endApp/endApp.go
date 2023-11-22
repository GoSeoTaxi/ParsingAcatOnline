package endApp

import (
	"fmt"
	"os"
)

func Fin() {
	var end string
	fmt.Println("Я закончил")
	fmt.Println("Нажмите Enter")
	fmt.Scanf("%s\n", &end)
	os.Exit(1)
}
