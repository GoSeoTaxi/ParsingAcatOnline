package main

import (
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/endApp"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/initApp"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/makeListUrl"
	"log"
	"time"
)

func main() {

	cfg, err := MakeConfiger.InitConfig()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}

	fmt.Println(`Запуск`)
	time.Sleep(constData.TimeSleepStart * time.Second)
	records, err := initApp.ReadCsvFile(constData.InputCSV)
	if err != nil {
		endApp.Fin()
	}

	StartingParsint(records, cfg)
	endApp.Fin()
}

func StartingParsint(lines [][]string, cfg *MakeConfiger.Config) {
	for value := range lines {
		makeListUrl.MakeList(lines[value], cfg)
	}
}
