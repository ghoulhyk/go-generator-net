package _const

type ReqMethod string

const (
	GET     ReqMethod = "GET"
	POST    ReqMethod = "POST"
	PUT     ReqMethod = "PUT"
	DELETE  ReqMethod = "DELETE"
	PATCH   ReqMethod = "PATCH"
	HEAD    ReqMethod = "HEAD"
	OPTIONS ReqMethod = "OPTIONS"
)
