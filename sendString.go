package connection

import (
	"errors"
	"net"
)

// SendString - Отправляет сообщение(string) в conn
func SendString(mes string, conn net.Conn) error {
	_, err := conn.Write([]byte(mes + "\n"))
	if err != nil {
		return errors.New("error write mes to conn")
	}
	return nil
}
