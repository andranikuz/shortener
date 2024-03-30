package file

import (
	"bufio"
	"os"
)

type producer struct {
	file *os.File
	// добавляем writer в producer
	writer *bufio.Writer
}

func newProducer(filename string) (*producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &producer{
		file: file,
		// создаём новый writer
		writer: bufio.NewWriter(file),
	}, nil
}

func (p *producer) write(data []byte) error {
	// записываем строку в буфер
	if _, err := p.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := p.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return p.writer.Flush()
}
