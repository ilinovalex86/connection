package connection

import (
	"bufio"
	"net"
)

// ReadSync - служит для синхронизации клиента и сервера
func ReadSync(conn net.Conn) {
	_, _ = bufio.NewReader(conn).ReadByte()
}
