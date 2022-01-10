package connection

// QueryResponse - Структура общения клиента и сервера
type QueryResponse struct {
	Method   string
	Query    string
	Response string
	DataLen  int
	FileName string
	Err      error
}
