package connection

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
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
	t := time.Now()
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
				s := float32(currentByte) / float32(time.Now().Sub(t).Seconds()) / float32(1024)
				fmt.Printf("Отправка файла: %8d/%d 100%% %10s      0m 00s %10s\n", fileSize/1024, fileSize/1024, speedCount(s), "")
				break
			}
		}
		err = SendBytes(fileBuffer, conn)
		if err != nil {
			return errors.New("error send file")
		}
		currentByte += 1024
		p := float32(currentByte) / float32(fileSize) * float32(100)
		s := float32(currentByte) / float32(time.Now().Sub(t).Seconds()) / float32(1024)
		m, sec := secToMinSec(int(float32(fileSize-currentByte) / float32(1024) / s))
		fmt.Printf("Отправка файла: %8d/%d %3.f%% %10s %6dm %02ds %10s\r", currentByte/1024, fileSize/1024, p, speedCount(s), m, sec, "")
	}
	return nil
}
