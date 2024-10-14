package _const

type DefType uint8

const (
	TypeInvalid DefType = iota
	TypeCustom
	TypeBool
	TypeTime
	TypeByte
	TypeString
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint
	TypeUint64
	TypeFloat32
	TypeFloat64
	endTypes
)
