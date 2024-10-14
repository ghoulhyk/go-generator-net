package arg

import (
	"github.com/ghoulhyk/go-generator-net/types/const"
)

func Custom(name string, customType any) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeCustom, CustomType: customType}
}
func Bool(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeBool}
}
func Time(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeTime}
}
func Byte(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeByte}
}
func String(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeString}
}
func Int8(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeInt8}
}
func Int16(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeInt16}
}
func Int32(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeInt32}
}
func Int(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeInt}
}
func Int64(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeInt64}
}
func Uint8(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeUint8}
}
func Uint16(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeUint16}
}
func Uint32(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeUint32}
}
func Uint(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeUint}
}
func Uint64(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeUint64}
}
func Float32(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeFloat32}
}
func Float64(name string) DynamicArg {
	return DynamicArg{ParaName: name, ReqName: name, Type: _const.TypeFloat64}
}

func Static(name string, val any) StaticArg {
	return StaticArg{
		ReqName: name,
		Value:   val,
	}
}
