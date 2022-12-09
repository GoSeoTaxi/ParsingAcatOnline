package writeOutFiles

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"io/ioutil"
	"log"
	"os"
)

type ImpDataArr struct {
	ArrayStrings []string
	TimeScraping string
}

func (data ImpDataArr) WriteToFile() {

	err := writeLines(data.ArrayStrings, data.TimeScraping, constData.FileNameOutWriter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v %v \n", data.TimeScraping, constData.WriterSucces)

	return
}

func writeLines(lines []string, time string, path string) error {

	var abcd []byte
	lineFiles, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) || lineFiles.Size() < 0 {
		ioutil.WriteFile(path, abcd, 0644)
		log.Fatal(`Нет доступа к файлу`)
	}

	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
			return err
		}
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
