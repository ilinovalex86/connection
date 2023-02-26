package connection

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

// SendFile - Отправляет файл путь к которому filePath покилобайтно в net.Conn
func SendFile(filePath string, conn net.Conn) error {
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		return errors.New("error open")
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := info.Size()
	var currentByte int64 = 0
	for {
		fileBuffer := make([]byte, 1024)
		byteRead, err := file.ReadAt(fileBuffer, currentByte)
		if err != nil {
			if err != io.EOF {
				return errors.New("error read file")
			}
			if err == io.EOF {
				if byteRead > 0 {
					fileBuffer = fileBuffer[:byteRead]
					err = SendBytes(fileBuffer, conn)
					if err != nil {
						return errors.New("error send file")
					}
				}
				fmt.Println("Отправка файла: 100%")
				break
			}
		}
		err = SendBytes(fileBuffer, conn)
		if err != nil {
			return errors.New("error send file")
		}
		currentByte += 1024
		fmt.Printf("Отправка файла: %3.f%%\r", float32(currentByte)/float32(fileSize)*float32(100))
	}
	return nil
}
