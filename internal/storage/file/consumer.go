package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func newConsumer(filename string) (*consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &consumer{
		file: file,
		// создаём новый scanner
		scanner: bufio.NewScanner(file),
	}, nil
}

// поиск JSON строки в файле по маске "{param}":"{value}"
func (c *consumer) findJSONByParam(param string, value string) ([]byte, error) {
	for true {
		// одиночное сканирование до следующей строки
		if !c.scanner.Scan() {
			return nil, c.scanner.Err()
		}
		// читаем данные из scanner
		data := c.scanner.Bytes()
		// проверяем наличие подстроки "{param}":"{id}"
		if strings.Contains(string(data), "\""+param+"\":\""+value+"\"") {
			return data, nil
		}
	}

	return nil, fmt.Errorf("%s not found", value)
}

func (c *consumer) close() error {
	return c.file.Close()
}
