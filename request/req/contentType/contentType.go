package contentType

type Type string

const (
	FormUrlencoded Type = "application/x-www-form-urlencoded"
	FormData       Type = "multipart/form-data"
	Json           Type = "application/json"
	//Plain          Type = "text/plain"
)
