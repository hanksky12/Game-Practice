package file

import (
	"encoding/csv"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"unicode/utf8"
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

	if err := handleBOM(file); err != nil {
		log.Fatalf("處理 BOM 時出錯: %v", err)
		return records, err
	}

	reader := csv.NewReader(file)

	records, err = reader.ReadAll()
	if err != nil {
		log.Fatalf("無法讀取 CSV: %v", err)
	}
	return records, err
}

func handleBOM(file *os.File) error {
	// 讀取文件內容
	content := make([]byte, 3)
	_, err := file.Read(content)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}

	// 檢查是否存在 BOM(在UTF-8編碼的檔案的開首加入一段位元組串EF BB BF)
	if utf8.Valid(content) && content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		// 存在，跳過三個字節
		_, err = file.Seek(3, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking file: %s", err)
		}
	} else {
		// 不存在，將文件重置到開頭
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking file: %s", err)
		}
	}

	return nil
}
