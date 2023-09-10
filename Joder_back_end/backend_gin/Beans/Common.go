package beans

type ResponseWrapper struct {
	resultCode    string
	resultMessage string
	model         interface{}
}
