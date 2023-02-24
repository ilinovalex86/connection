package connection

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

// ReadString - Получает сообщение(string) из conn
func ReadString(conn net.Conn) (string, error) {
	mes, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", errors.New("error read mes from conn")
	}
	mes = strings.TrimSpace(mes)
	return mes, nil
}
