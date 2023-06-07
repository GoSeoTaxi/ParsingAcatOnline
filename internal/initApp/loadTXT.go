package initApp

import (
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"log"
	"os"
	"strings"
)

func LoadTOmapTXT(m map[int]string) {

	// Открываем файл для чтения
	data, err := os.ReadFile(constData.InputURLFileTXT)
	if err != nil {
		log.Fatal("Ошибка при открытии файла c адресами:", err)
		return
	}

	// Разделяем файл на строки
	lines := strings.Split(string(data), "\n")

	for k, line := range lines {
		m[k+1] = line
	}

	return
}
