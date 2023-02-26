package connection

import (
	"errors"
	"fmt"
	"net"
	"os"
)

// GetFile - Получает файл размером dataLen из conn и сохраняетего в dirToFiles c именем fileName
func GetFile(path string, dataLen int, conn net.Conn) error {
	var currentByte int64 = 0
	file, err := os.Create(path)
	if err != nil {
		return errors.New("ошибка создания файла")
	}
	countAll := 0
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
		fmt.Printf("Загрузка файла: %3.f%%\n", float32(countAll)/float32(dataLen)*float32(100))
		if countAll == dataLen {
			break
		}
	}
	fmt.Println()
	file.Close()
	return nil
}
