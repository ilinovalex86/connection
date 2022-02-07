package connection

type Query struct {
	Method string
	Query  string
}

type Response struct {
	Response string
	DataLen  int
	Err      error
}
