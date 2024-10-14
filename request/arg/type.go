package arg

import (
	"ghoulhyk/go-generator-net/types/const"
)

type Arg interface {
	GetReqName() string
}

// region DynamicArg

type DynamicArg struct {
	ParaName   string // 参数名
	ReqName    string // 请求参数名
	Type       _const.DefType
	CustomType any
	PtrType    bool
}

func (receiver DynamicArg) SetReqName(reqName string) DynamicArg {
	receiver.ReqName = reqName
	return receiver
}

func (receiver DynamicArg) GetReqName() string {
	return receiver.ReqName
}

func (receiver DynamicArg) Ptr() DynamicArg {
	receiver.PtrType = true
	return receiver
}

func (receiver DynamicArg) IsPtr() bool {
	return receiver.PtrType
}

// endregion

// region StaticArg

type StaticArg struct {
	ReqName string // 请求参数名
	Value   any
}

func (receiver StaticArg) SetReqName(reqName string) StaticArg {
	receiver.ReqName = reqName
	return receiver
}

func (receiver StaticArg) GetReqName() string {
	return receiver.ReqName
}

func (receiver StaticArg) SetValue(Value any) StaticArg {
	receiver.Value = Value
	return receiver
}

// endregion
