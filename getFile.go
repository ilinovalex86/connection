package connection

import (
	"errors"
	"net"
	"os"
)

// GetFile - Получает файл размером dataLen из conn и сохраняетего в dirToFiles c именем fileName
func GetFile(fileName string, dataLen int, dirToFiles string, conn net.Conn) error {
	var currentByte int64 = 0
	file, err := os.Create(dirToFiles + "/" + fileName)
	if err != nil {
		return errors.New("ошибка создания файла")
	}
	defer file.Close()
	countAll := 0
	for {
		count, data, err := ReadBytes(conn)
		if err != nil {
			return err
		}
		_, err = file.WriteAt(data, currentByte)
		if err != nil {
			return errors.New("ошибка записи файла")
		}
		currentByte += int64(count)
		countAll += count
		if countAll == dataLen {
			break
		}
	}
	return nil
}
