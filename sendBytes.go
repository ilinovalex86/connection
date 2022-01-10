package connection

import (
	"errors"
	"net"
)

// SendBytes - Отправляет бинарные данные в net.Conn
func SendBytes(data []byte, conn net.Conn) error {
	_, err := conn.Write(data)
	if err != nil {
		return errors.New("error write data to conn")
	}
	return nil
}
