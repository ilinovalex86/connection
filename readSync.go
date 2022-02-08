package connection

import (
	"bufio"
	"net"
)

// ReadSync - служит для синхронизации клиента и сервера
func ReadSync(conn net.Conn) (string, error) {
	b, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
