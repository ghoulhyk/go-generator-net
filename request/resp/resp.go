package resp

type IResp interface {
	IsSuccess() bool
	ErrInfo() error
}
