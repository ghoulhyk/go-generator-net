package types

import (
	"encoding/json"
	"fmt"
	"ghoulhyk/go-generator-net/types/const"
	"github.com/pkg/errors"
)

type Arg interface {
	GetReqName() string
	GetValWay() string
}

type Args []Arg

func (s *Args) UnmarshalJSON(data []byte) error {
	var rawSlice []map[string]any
	if err := json.Unmarshal(data, &rawSlice); err != nil {
		return err
	}

	for _, raw := range rawSlice {
		argType := raw["_getValWay"].(string)

		bytes, _ := json.Marshal(raw)
		if argType == (DynamicArg{}).GetValWay() {
			arg := DynamicArg{}
			err := json.Unmarshal(bytes, &arg)
			if err != nil {
				return err
			}
			*s = append(*s, arg)
		} else if argType == (StaticArg{}).GetValWay() {
			arg := StaticArg{}
			err := json.Unmarshal(bytes, &arg)
			if err != nil {
				return err
			}
			*s = append(*s, arg)
		} else {
			return errors.Errorf("unknown ValWay[%v]", argType)
		}
	}

	return nil
}

// region DynamicArg

type DynamicArg struct {
	ParaName string // 参数名
	Type     _const.DefType

	ReqName    string // 请求参数名
	CustomType *RType
	PtrType    bool
}

func (receiver DynamicArg) GetReqName() string {
	return receiver.ReqName
}

func (receiver DynamicArg) GetValWay() string {
	return "Dynamic"
}

func (receiver DynamicArg) IsPtr() bool {
	return receiver.PtrType
}

func (receiver DynamicArg) TypeStr() string {
	typeStr := ""
	switch receiver.Type {
	case _const.TypeCustom:
		typeStr = receiver.CustomType.Ident
	case _const.TypeBool:
		typeStr = "bool"
	case _const.TypeTime:
		typeStr = "time.Time"
	case _const.TypeByte:
		typeStr = "byte"
	case _const.TypeString:
		typeStr = "string"
	case _const.TypeInt8:
		typeStr = "int8"
	case _const.TypeInt16:
		typeStr = "int16"
	case _const.TypeInt32:
		typeStr = "int32"
	case _const.TypeInt:
		typeStr = "int"
	case _const.TypeInt64:
		typeStr = "int64"
	case _const.TypeUint8:
		typeStr = "uint8"
	case _const.TypeUint16:
		typeStr = "uint16"
	case _const.TypeUint32:
		typeStr = "uint32"
	case _const.TypeUint:
		typeStr = "uint"
	case _const.TypeUint64:
		typeStr = "uint64"
	case _const.TypeFloat32:
		typeStr = "float32"
	case _const.TypeFloat64:
		typeStr = "float64"
	default:
		panic("类型错误！")
	}
	if receiver.PtrType {
		typeStr = "*" + typeStr
	}

	return typeStr
}

func (receiver DynamicArg) PkgPathList() []string {
	switch receiver.Type {
	case _const.TypeCustom:
		return receiver.CustomType.PkgPathList
	case _const.TypeTime:
		return []string{"time"}
	}
	return nil
}

func (receiver DynamicArg) ValueFormatter() string {
	switch receiver.Type {
	case _const.TypeFloat32, _const.TypeFloat64:
		return "%f"
	case _const.TypeInt, _const.TypeInt8, _const.TypeInt16, _const.TypeInt32, _const.TypeInt64, _const.TypeUint, _const.TypeUint8, _const.TypeUint16, _const.TypeUint32, _const.TypeUint64:
		return "%d"
	case _const.TypeString:
		return "%s"
	default:
		return "%v"
	}
}

func (receiver DynamicArg) MarshalJSON() ([]byte, error) {
	//return json.Marshal(struct {
	//	ParaName   string // 参数名
	//	Type       _const.DefType
	//	ReqName    string // 请求参数名
	//	CustomType *RType
	//	PtrType    bool
	//	ArgType    uint8 `json:"_argType"`
	//}{
	//	ParaName:   receiver.ParaName,
	//	Type:       receiver.Type,
	//	ReqName:    receiver.ReqName,
	//	CustomType: receiver.CustomType,
	//	PtrType:    receiver.PtrType,
	//	ArgType:    1,
	//})

	type jsonArg DynamicArg

	return json.Marshal(struct {
		jsonArg
		GetValWay string `json:"_getValWay"`
	}{
		jsonArg:   jsonArg(receiver),
		GetValWay: receiver.GetValWay(),
	})
}

// endregion

// region StaticArg

type StaticArg struct {
	ReqName string // 请求参数名
	Value   any
}

func (receiver StaticArg) GetReqName() string {
	return receiver.ReqName
}

func (receiver StaticArg) GetValWay() string {
	return "Static"
}

func (receiver StaticArg) ValueStr() string {
	switch receiver.Value.(type) {
	case float32, float64:
		return fmt.Sprintf("%f", receiver.Value)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", receiver.Value)
	case string:
		return receiver.Value.(string)
	default:
		return fmt.Sprintf("%v", receiver.Value)
	}
}

func (receiver StaticArg) ValueForBody() any {
	return fmt.Sprintf("%#v", receiver.Value)
}

func (receiver StaticArg) MarshalJSON() ([]byte, error) {
	//return json.Marshal(struct {
	//	ReqName     string // 请求参数名
	//	Value any
	//
	//	ArgType uint8 `json:"_argType"`
	//}{
	//	ReqName:     receiver.ReqName,
	//	Value: receiver.Value,
	//	ArgType:     2,
	//})

	type jsonArg StaticArg

	return json.Marshal(struct {
		jsonArg
		GetValWay string `json:"_getValWay"`
	}{
		jsonArg:   jsonArg(receiver),
		GetValWay: receiver.GetValWay(),
	})
}

// endregion
