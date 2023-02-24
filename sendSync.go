package connection

import "net"

// SendSync - служит для синхронизации клиента и сервера
func SendSync(conn net.Conn) {
	_, _ = conn.Write([]byte("1"))
}
