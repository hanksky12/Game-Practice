package file

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
)

type CSV struct{}

func (c *CSV) Read(filePath string) ([][]string, error) {
	var records [][]string
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("無法打開檔案: %v", err)
		return records, err
	}

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatalf("無法讀取 CSV: %v", err)
	}
	return records, err
}
