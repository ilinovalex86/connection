Package connection содержит функции для работы с net.Conn

Основные функции: ReadQRStruct, SendQRStruct
работают со структурой QueryResponse
 type QueryResponse struct {
   Method   string
   Query    string
   Response string
   DataLen  int
   FileName string
   Err      error
 }
