package connection

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

// GetFile - Получает файл размером dataLen из conn и сохраняетего в dirToFiles c именем fileName
func GetFile(path string, dataLen int, conn net.Conn) error {
	var currentByte int64 = 0
	file, err := os.Create(path)
	if err != nil {
		return errors.New("ошибка создания файла")
	}
	countAll := 0
	t := time.Now()
	for {
		count, data, err := ReadBytes(conn)
		if err != nil {
			file.Close()
			os.Remove(path)
			return err
		}
		_, err = file.WriteAt(data, currentByte)
		if err != nil {
			file.Close()
			os.Remove(path)
			return errors.New("ошибка записи файла")
		}
		currentByte += int64(count)
		countAll += count
		p := float32(currentByte) / float32(dataLen) * float32(100)
		s := float32(currentByte) / float32(time.Now().Sub(t).Seconds()) / float32(1024)
		m, sec := secToMinSec(int(float32(dataLen-int(currentByte)) / float32(1024) / s))
		fmt.Printf("Отправка файла: %8d/%d %3.f%% %10s %6dm %02ds %10s\r", currentByte/1024, dataLen/1024, p, speedCount(s), m, sec, "")
		if countAll == dataLen {
			break
		}
	}
	fmt.Println()
	file.Close()
	return nil
}
