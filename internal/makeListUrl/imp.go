package makeListUrl

import (
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/endApp"
	"net/url"
)

//var counter int

func MakeList(s []string, cfg *MakeConfiger.Config) {

	_, err := url.ParseRequestURI(s[1])
	if err != nil {
		fmt.Println(`Ошибка в файле импорта.`)
		fmt.Println(s[1])
		fmt.Println(`Не является URL`)
		endApp.Fin()
	}

	textBefore := s[0]
	if len(textBefore) < 1 {
		fmt.Println(`Нет дополнительного значения для URL`)
		fmt.Println(s[1])
		endApp.Fin()
	}

	URLServer := s[1]

	dataTree(URLServer, s[0], cfg)

}
