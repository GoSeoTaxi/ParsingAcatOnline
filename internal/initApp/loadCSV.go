package initApp

import (
	"encoding/csv"
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"log"
	"os"
)

func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(`Ошибка чтения входного файла. Нет доступов`)
		return nil, err
		//	log.Fatal("Unable to read input file "+filePath, err)
	}

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(`Ошибка структуры входного файла. Проверить файл ` + constData.InputCSV)
		return nil, err
		//	log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	err = f.Close()
	if err != nil {
		fmt.Println(`Ошибка закрытия файла на запись.`)
		log.Fatal(err)
	}

	return records, nil
}
