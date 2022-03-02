package connection

type Query struct {
	Method  string
	Query   string
	DataLen int
}

type Response struct {
	Response string
	DataLen  int
	Err      error
}
