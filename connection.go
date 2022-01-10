// Package connection содержит функции для работы с net.Conn
// Написан для взаимодействия сервера и клиента:
// сервер github.com/ilinovalex86/tcpserver
// клиент github.com/ilinovalex86/tcpclient
// Основные функции: ReadQRStruct, SendQRStruct
// работают со структурой QueryResponse
//  type QueryResponse struct {
//    Method   string
//    Query    string
//    Response string
//    DataLen  int
//    FileName string
//    Err      error
//  }
package connection
