package connection

import (
	"net"
)

// SendBytesWithDelim - Отправляет []bytes + '\n' в conn
func SendBytesWithDelim(data []byte, conn net.Conn) error {
	data = append(data, '\n')
	err := SendBytes(data, conn)
	if err != nil {
		return err
	}
	return nil
}
