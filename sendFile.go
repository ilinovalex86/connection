package connection

import (
	"errors"
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
				break
			}
		}
		err = SendBytes(fileBuffer, conn)
		if err != nil {
			return errors.New("error send file")
		}
		currentByte += 1024
	}
	return nil
}
