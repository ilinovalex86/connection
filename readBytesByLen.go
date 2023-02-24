package connection

import (
	"errors"
	"net"
)

// ReadBytesByLen - Читает lenBytes данных из conn
func ReadBytesByLen(lenBytes int, conn net.Conn) ([]byte, error) {
	var data []byte
	for {
		buffer := make([]byte, 1024)
		count, err := conn.Read(buffer)
		if err != nil {
			return nil, errors.New("error read data from conn")
		}
		data = append(data, buffer[:count]...)
		lenBytes -= count
		if lenBytes == 0 {
			break
		}
		if lenBytes < 0 {
			return nil, errors.New("error read toBig data from conn")
		}
	}
	return data, nil
}
