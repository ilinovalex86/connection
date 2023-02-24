package connection

import (
	"errors"
	"net"
)

// ReadBytes - Читает бинарные данные из conn
func ReadBytes(conn net.Conn) (int, []byte, error) {
	buffer := make([]byte, 1024)
	count, err := conn.Read(buffer)
	if err != nil {
		return count, nil, errors.New("error read data from conn")
	}
	return count, buffer[:count], nil
}
