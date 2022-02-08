package connection

import "net"

// SendSync - служит для синхронизации клиента и сервера
func SendSync(conn net.Conn) (int, error) {
	n, err := conn.Write([]byte("1"))
	if err != nil {
		return 0, err
	}
	return n, nil
}
