package _const

type ArgPos uint8

const (
	Default ArgPos = iota
	Path
	Query
	Body
)
