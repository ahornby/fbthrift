// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package module // [[[ program thrift source path ]]]

import (
    "fmt"
    "strings"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = strings.Split
var _ = thrift.ZERO


type MyEnumAlias = MyEnum

const (
    MyEnumAlias_MyValue1 MyEnumAlias = MyEnum_MyValue1
    MyEnumAlias_MyValue2 MyEnumAlias = MyEnum_MyValue2
)

// Enum value maps for MyEnum
var (
    MyEnumAliasToName  = MyEnumToName
    MyEnumAliasToValue = MyEnumToValue
)

// Deprecated: Use MyEnumAliasToValue instead (e.g. `x, ok := MyEnumAliasToValue["name"]`).
func MyEnumAliasFromString(s string) (MyEnumAlias, error) {
    return MyEnumFromString(s)
}

func NewMyEnumAlias() MyEnumAlias {
    return 0
}

func WriteMyEnumAlias(item MyEnumAlias, p thrift.Format) error {
    if err := p.WriteI32(int32(item)); err != nil {
    return err
}
    return nil
}

func ReadMyEnumAlias(p thrift.Format) (MyEnumAlias, error) {
    var decodeResult MyEnumAlias
    decodeErr := func() error {
        enumResult, err := p.ReadI32()
if err != nil {
    return err
}
result := MyEnum(enumResult)
        decodeResult = result
        return nil
    }()
    return decodeResult, decodeErr
}

type MyDataItemAlias = MyDataItem

func NewMyDataItemAlias() *MyDataItemAlias {
    return NewMyDataItem()
}

func WriteMyDataItemAlias(item *MyDataItemAlias, p thrift.Format) error {
    if err := item.Write(p); err != nil {
    return err
}
    return nil
}

func ReadMyDataItemAlias(p thrift.Format) (MyDataItemAlias, error) {
    var decodeResult MyDataItemAlias
    decodeErr := func() error {
        result := *NewMyDataItem()
err := result.Read(p)
if err != nil {
    return err
}
        decodeResult = result
        return nil
    }()
    return decodeResult, decodeErr
}

type MyEnum int32

const (
    MyEnum_MyValue1 MyEnum = 0
    MyEnum_MyValue2 MyEnum = 1
)

// Enum value maps for MyEnum
var (
    MyEnumToName = map[MyEnum]string {
        MyEnum_MyValue1: "MyValue1",
        MyEnum_MyValue2: "MyValue2",
    }

    MyEnumToValue = map[string]MyEnum {
        "MyValue1": MyEnum_MyValue1,
        "MyValue2": MyEnum_MyValue2,
    }
)

func (x MyEnum) String() string {
    if v, ok := MyEnumToName[x]; ok {
        return v
    }
    return "<UNSET>"
}

func (x MyEnum) Ptr() *MyEnum {
    return &x
}

// Deprecated: Use MyEnumToValue instead (e.g. `x, ok := MyEnumToValue["name"]`).
func MyEnumFromString(s string) (MyEnum, error) {
    if v, ok := MyEnumToValue[s]; ok {
        return v, nil
    }
    return MyEnum(0), fmt.Errorf("not a valid MyEnum string")
}


type HackEnum int32

const (
    HackEnum_Value1 HackEnum = 0
    HackEnum_Value2 HackEnum = 1
)

// Enum value maps for HackEnum
var (
    HackEnumToName = map[HackEnum]string {
        HackEnum_Value1: "Value1",
        HackEnum_Value2: "Value2",
    }

    HackEnumToValue = map[string]HackEnum {
        "Value1": HackEnum_Value1,
        "Value2": HackEnum_Value2,
    }
)

func (x HackEnum) String() string {
    if v, ok := HackEnumToName[x]; ok {
        return v
    }
    return "<UNSET>"
}

func (x HackEnum) Ptr() *HackEnum {
    return &x
}

// Deprecated: Use HackEnumToValue instead (e.g. `x, ok := HackEnumToValue["name"]`).
func HackEnumFromString(s string) (HackEnum, error) {
    if v, ok := HackEnumToValue[s]; ok {
        return v, nil
    }
    return HackEnum(0), fmt.Errorf("not a valid HackEnum string")
}


type MyStruct struct {
    MyIntField int64 `thrift:"MyIntField,1" json:"MyIntField" db:"MyIntField"`
    MyStringField string `thrift:"MyStringField,2" json:"MyStringField" db:"MyStringField"`
    MyDataField *MyDataItemAlias `thrift:"MyDataField,3" json:"MyDataField" db:"MyDataField"`
    MyEnum MyEnum `thrift:"myEnum,4" json:"myEnum" db:"myEnum"`
    Oneway bool `thrift:"oneway,5" json:"oneway" db:"oneway"`
    Readonly bool `thrift:"readonly,6" json:"readonly" db:"readonly"`
    Idempotent bool `thrift:"idempotent,7" json:"idempotent" db:"idempotent"`
    FloatSet []float32 `thrift:"floatSet,8" json:"floatSet" db:"floatSet"`
    NoHackCodegenField string `thrift:"no_hack_codegen_field,9" json:"no_hack_codegen_field" db:"no_hack_codegen_field"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*MyStruct)(nil)

func NewMyStruct() *MyStruct {
    return (&MyStruct{}).
        SetMyIntFieldNonCompat(0).
        SetMyStringFieldNonCompat("").
        SetMyDataFieldNonCompat(*NewMyDataItemAlias()).
        SetMyEnumNonCompat(0).
        SetOnewayNonCompat(false).
        SetReadonlyNonCompat(false).
        SetIdempotentNonCompat(false).
        SetFloatSetNonCompat(make([]float32, 0)).
        SetNoHackCodegenFieldNonCompat("")
}

func (x *MyStruct) GetMyIntField() int64 {
    return x.MyIntField
}

func (x *MyStruct) GetMyStringField() string {
    return x.MyStringField
}

func (x *MyStruct) GetMyDataField() *MyDataItemAlias {
    if !x.IsSetMyDataField() {
        return nil
    }

    return x.MyDataField
}

func (x *MyStruct) GetMyEnum() MyEnum {
    return x.MyEnum
}

func (x *MyStruct) GetOneway() bool {
    return x.Oneway
}

func (x *MyStruct) GetReadonly() bool {
    return x.Readonly
}

func (x *MyStruct) GetIdempotent() bool {
    return x.Idempotent
}

func (x *MyStruct) GetFloatSet() []float32 {
    if !x.IsSetFloatSet() {
        return make([]float32, 0)
    }

    return x.FloatSet
}

func (x *MyStruct) GetNoHackCodegenField() string {
    return x.NoHackCodegenField
}

func (x *MyStruct) SetMyIntFieldNonCompat(value int64) *MyStruct {
    x.MyIntField = value
    return x
}

func (x *MyStruct) SetMyIntField(value int64) *MyStruct {
    x.MyIntField = value
    return x
}

func (x *MyStruct) SetMyStringFieldNonCompat(value string) *MyStruct {
    x.MyStringField = value
    return x
}

func (x *MyStruct) SetMyStringField(value string) *MyStruct {
    x.MyStringField = value
    return x
}

func (x *MyStruct) SetMyDataFieldNonCompat(value MyDataItemAlias) *MyStruct {
    x.MyDataField = &value
    return x
}

func (x *MyStruct) SetMyDataField(value *MyDataItemAlias) *MyStruct {
    x.MyDataField = value
    return x
}

func (x *MyStruct) SetMyEnumNonCompat(value MyEnum) *MyStruct {
    x.MyEnum = value
    return x
}

func (x *MyStruct) SetMyEnum(value MyEnum) *MyStruct {
    x.MyEnum = value
    return x
}

func (x *MyStruct) SetOnewayNonCompat(value bool) *MyStruct {
    x.Oneway = value
    return x
}

func (x *MyStruct) SetOneway(value bool) *MyStruct {
    x.Oneway = value
    return x
}

func (x *MyStruct) SetReadonlyNonCompat(value bool) *MyStruct {
    x.Readonly = value
    return x
}

func (x *MyStruct) SetReadonly(value bool) *MyStruct {
    x.Readonly = value
    return x
}

func (x *MyStruct) SetIdempotentNonCompat(value bool) *MyStruct {
    x.Idempotent = value
    return x
}

func (x *MyStruct) SetIdempotent(value bool) *MyStruct {
    x.Idempotent = value
    return x
}

func (x *MyStruct) SetFloatSetNonCompat(value []float32) *MyStruct {
    x.FloatSet = value
    return x
}

func (x *MyStruct) SetFloatSet(value []float32) *MyStruct {
    x.FloatSet = value
    return x
}

func (x *MyStruct) SetNoHackCodegenFieldNonCompat(value string) *MyStruct {
    x.NoHackCodegenField = value
    return x
}

func (x *MyStruct) SetNoHackCodegenField(value string) *MyStruct {
    x.NoHackCodegenField = value
    return x
}

func (x *MyStruct) IsSetMyDataField() bool {
    return x != nil && x.MyDataField != nil
}

func (x *MyStruct) IsSetFloatSet() bool {
    return x != nil && x.FloatSet != nil
}

func (x *MyStruct) writeField1(p thrift.Format) error {  // MyIntField
    if err := p.WriteFieldBegin("MyIntField", thrift.I64, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.MyIntField
    if err := p.WriteI64(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField2(p thrift.Format) error {  // MyStringField
    if err := p.WriteFieldBegin("MyStringField", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.MyStringField
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField3(p thrift.Format) error {  // MyDataField
    if !x.IsSetMyDataField() {
        return nil
    }

    if err := p.WriteFieldBegin("MyDataField", thrift.STRUCT, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.MyDataField
    err := WriteMyDataItemAlias(item, p)
if err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField4(p thrift.Format) error {  // MyEnum
    if err := p.WriteFieldBegin("myEnum", thrift.I32, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.MyEnum
    if err := p.WriteI32(int32(item)); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField5(p thrift.Format) error {  // Oneway
    if err := p.WriteFieldBegin("oneway", thrift.BOOL, 5); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Oneway
    if err := p.WriteBool(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField6(p thrift.Format) error {  // Readonly
    if err := p.WriteFieldBegin("readonly", thrift.BOOL, 6); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Readonly
    if err := p.WriteBool(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField7(p thrift.Format) error {  // Idempotent
    if err := p.WriteFieldBegin("idempotent", thrift.BOOL, 7); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Idempotent
    if err := p.WriteBool(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField8(p thrift.Format) error {  // FloatSet
    if err := p.WriteFieldBegin("floatSet", thrift.SET, 8); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.FloatSet
    if err := p.WriteSetBegin(thrift.FLOAT, len(item)); err != nil {
    return thrift.PrependError("error writing set begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteFloat(item); err != nil {
    return err
}
    }
}
if err := p.WriteSetEnd(); err != nil {
    return thrift.PrependError("error writing set end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField9(p thrift.Format) error {  // NoHackCodegenField
    if err := p.WriteFieldBegin("no_hack_codegen_field", thrift.STRING, 9); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.NoHackCodegenField
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) readField1(p thrift.Format) error {  // MyIntField
    result, err := p.ReadI64()
if err != nil {
    return err
}

    x.MyIntField = result
    return nil
}

func (x *MyStruct) readField2(p thrift.Format) error {  // MyStringField
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.MyStringField = result
    return nil
}

func (x *MyStruct) readField3(p thrift.Format) error {  // MyDataField
    result, err := ReadMyDataItemAlias(p)
if err != nil {
    return err
}

    x.MyDataField = &result
    return nil
}

func (x *MyStruct) readField4(p thrift.Format) error {  // MyEnum
    enumResult, err := p.ReadI32()
if err != nil {
    return err
}
result := MyEnum(enumResult)

    x.MyEnum = result
    return nil
}

func (x *MyStruct) readField5(p thrift.Format) error {  // Oneway
    result, err := p.ReadBool()
if err != nil {
    return err
}

    x.Oneway = result
    return nil
}

func (x *MyStruct) readField6(p thrift.Format) error {  // Readonly
    result, err := p.ReadBool()
if err != nil {
    return err
}

    x.Readonly = result
    return nil
}

func (x *MyStruct) readField7(p thrift.Format) error {  // Idempotent
    result, err := p.ReadBool()
if err != nil {
    return err
}

    x.Idempotent = result
    return nil
}

func (x *MyStruct) readField8(p thrift.Format) error {  // FloatSet
    _ /* elemType */, size, err := p.ReadSetBegin()
if err != nil {
    return thrift.PrependError("error reading set begin: ", err)
}

setResult := make([]float32, 0, size)
for i := 0; i < size; i++ {
    var elem float32
    {
        result, err := p.ReadFloat()
if err != nil {
    return err
}
        elem = result
    }
    setResult = append(setResult, elem)
}

if err := p.ReadSetEnd(); err != nil {
    return thrift.PrependError("error reading set end: ", err)
}
result := setResult

    x.FloatSet = result
    return nil
}

func (x *MyStruct) readField9(p thrift.Format) error {  // NoHackCodegenField
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.NoHackCodegenField = result
    return nil
}

func (x *MyStruct) toString1() string {  // MyIntField
    return fmt.Sprintf("%v", x.MyIntField)
}

func (x *MyStruct) toString2() string {  // MyStringField
    return fmt.Sprintf("%v", x.MyStringField)
}

func (x *MyStruct) toString3() string {  // MyDataField
    return fmt.Sprintf("%v", x.MyDataField)
}

func (x *MyStruct) toString4() string {  // MyEnum
    return fmt.Sprintf("%v", x.MyEnum)
}

func (x *MyStruct) toString5() string {  // Oneway
    return fmt.Sprintf("%v", x.Oneway)
}

func (x *MyStruct) toString6() string {  // Readonly
    return fmt.Sprintf("%v", x.Readonly)
}

func (x *MyStruct) toString7() string {  // Idempotent
    return fmt.Sprintf("%v", x.Idempotent)
}

func (x *MyStruct) toString8() string {  // FloatSet
    return fmt.Sprintf("%v", x.FloatSet)
}

func (x *MyStruct) toString9() string {  // NoHackCodegenField
    return fmt.Sprintf("%v", x.NoHackCodegenField)
}

// Deprecated: Use NewMyStruct().GetMyDataField() instead.
func (x *MyStruct) DefaultGetMyDataField() *MyDataItemAlias {
    if !x.IsSetMyDataField() {
        return NewMyDataItemAlias()
    }
    return x.MyDataField
}



func (x *MyStruct) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("MyStruct"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := x.writeField5(p); err != nil {
        return err
    }

    if err := x.writeField6(p); err != nil {
        return err
    }

    if err := x.writeField7(p); err != nil {
        return err
    }

    if err := x.writeField8(p); err != nil {
        return err
    }

    if err := x.writeField9(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.I64)):  // MyIntField
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.STRING)):  // MyStringField
            if err := x.readField2(p); err != nil {
                return err
            }
        case (id == 3 && wireType == thrift.Type(thrift.STRUCT)):  // MyDataField
            if err := x.readField3(p); err != nil {
                return err
            }
        case (id == 4 && wireType == thrift.Type(thrift.I32)):  // myEnum
            if err := x.readField4(p); err != nil {
                return err
            }
        case (id == 5 && wireType == thrift.Type(thrift.BOOL)):  // oneway
            if err := x.readField5(p); err != nil {
                return err
            }
        case (id == 6 && wireType == thrift.Type(thrift.BOOL)):  // readonly
            if err := x.readField6(p); err != nil {
                return err
            }
        case (id == 7 && wireType == thrift.Type(thrift.BOOL)):  // idempotent
            if err := x.readField7(p); err != nil {
                return err
            }
        case (id == 8 && wireType == thrift.Type(thrift.SET)):  // floatSet
            if err := x.readField8(p); err != nil {
                return err
            }
        case (id == 9 && wireType == thrift.Type(thrift.STRING)):  // no_hack_codegen_field
            if err := x.readField9(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *MyStruct) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("MyStruct({")
    sb.WriteString(fmt.Sprintf("MyIntField:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("MyStringField:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("MyDataField:%s ", x.toString3()))
    sb.WriteString(fmt.Sprintf("MyEnum:%s ", x.toString4()))
    sb.WriteString(fmt.Sprintf("Oneway:%s ", x.toString5()))
    sb.WriteString(fmt.Sprintf("Readonly:%s ", x.toString6()))
    sb.WriteString(fmt.Sprintf("Idempotent:%s ", x.toString7()))
    sb.WriteString(fmt.Sprintf("FloatSet:%s ", x.toString8()))
    sb.WriteString(fmt.Sprintf("NoHackCodegenField:%s", x.toString9()))
    sb.WriteString("})")

    return sb.String()
}

type MyDataItem struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*MyDataItem)(nil)

func NewMyDataItem() *MyDataItem {
    return (&MyDataItem{})
}



func (x *MyDataItem) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("MyDataItem"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyDataItem) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *MyDataItem) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("MyDataItem({")
    sb.WriteString("})")

    return sb.String()
}

type MyUnion struct {
    MyEnum *MyEnumAlias `thrift:"myEnum,1" json:"myEnum,omitempty" db:"myEnum"`
    MyStruct *MyStruct `thrift:"myStruct,2" json:"myStruct,omitempty" db:"myStruct"`
    MyDataItem *MyDataItem `thrift:"myDataItem,3" json:"myDataItem,omitempty" db:"myDataItem"`
    FloatSet []float32 `thrift:"floatSet,4" json:"floatSet,omitempty" db:"floatSet"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*MyUnion)(nil)

func NewMyUnion() *MyUnion {
    return (&MyUnion{})
}

func (x *MyUnion) GetMyEnum() MyEnumAlias {
    if !x.IsSetMyEnum() {
        return NewMyEnumAlias()
    }

    return *x.MyEnum
}

func (x *MyUnion) GetMyStruct() *MyStruct {
    if !x.IsSetMyStruct() {
        return nil
    }

    return x.MyStruct
}

func (x *MyUnion) GetMyDataItem() *MyDataItem {
    if !x.IsSetMyDataItem() {
        return nil
    }

    return x.MyDataItem
}

func (x *MyUnion) GetFloatSet() []float32 {
    if !x.IsSetFloatSet() {
        return make([]float32, 0)
    }

    return x.FloatSet
}

func (x *MyUnion) SetMyEnumNonCompat(value MyEnumAlias) *MyUnion {
    x.MyEnum = &value
    return x
}

func (x *MyUnion) SetMyEnum(value *MyEnumAlias) *MyUnion {
    x.MyEnum = value
    return x
}

func (x *MyUnion) SetMyStructNonCompat(value MyStruct) *MyUnion {
    x.MyStruct = &value
    return x
}

func (x *MyUnion) SetMyStruct(value *MyStruct) *MyUnion {
    x.MyStruct = value
    return x
}

func (x *MyUnion) SetMyDataItemNonCompat(value MyDataItem) *MyUnion {
    x.MyDataItem = &value
    return x
}

func (x *MyUnion) SetMyDataItem(value *MyDataItem) *MyUnion {
    x.MyDataItem = value
    return x
}

func (x *MyUnion) SetFloatSetNonCompat(value []float32) *MyUnion {
    x.FloatSet = value
    return x
}

func (x *MyUnion) SetFloatSet(value []float32) *MyUnion {
    x.FloatSet = value
    return x
}

func (x *MyUnion) IsSetMyEnum() bool {
    return x != nil && x.MyEnum != nil
}

func (x *MyUnion) IsSetMyStruct() bool {
    return x != nil && x.MyStruct != nil
}

func (x *MyUnion) IsSetMyDataItem() bool {
    return x != nil && x.MyDataItem != nil
}

func (x *MyUnion) IsSetFloatSet() bool {
    return x != nil && x.FloatSet != nil
}

func (x *MyUnion) writeField1(p thrift.Format) error {  // MyEnum
    if !x.IsSetMyEnum() {
        return nil
    }

    if err := p.WriteFieldBegin("myEnum", thrift.I32, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.MyEnum
    err := WriteMyEnumAlias(item, p)
if err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyUnion) writeField2(p thrift.Format) error {  // MyStruct
    if !x.IsSetMyStruct() {
        return nil
    }

    if err := p.WriteFieldBegin("myStruct", thrift.STRUCT, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.MyStruct
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyUnion) writeField3(p thrift.Format) error {  // MyDataItem
    if !x.IsSetMyDataItem() {
        return nil
    }

    if err := p.WriteFieldBegin("myDataItem", thrift.STRUCT, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.MyDataItem
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyUnion) writeField4(p thrift.Format) error {  // FloatSet
    if !x.IsSetFloatSet() {
        return nil
    }

    if err := p.WriteFieldBegin("floatSet", thrift.SET, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.FloatSet
    if err := p.WriteSetBegin(thrift.FLOAT, len(item)); err != nil {
    return thrift.PrependError("error writing set begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteFloat(item); err != nil {
    return err
}
    }
}
if err := p.WriteSetEnd(); err != nil {
    return thrift.PrependError("error writing set end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyUnion) readField1(p thrift.Format) error {  // MyEnum
    result, err := ReadMyEnumAlias(p)
if err != nil {
    return err
}

    x.MyEnum = &result
    return nil
}

func (x *MyUnion) readField2(p thrift.Format) error {  // MyStruct
    result := *NewMyStruct()
err := result.Read(p)
if err != nil {
    return err
}

    x.MyStruct = &result
    return nil
}

func (x *MyUnion) readField3(p thrift.Format) error {  // MyDataItem
    result := *NewMyDataItem()
err := result.Read(p)
if err != nil {
    return err
}

    x.MyDataItem = &result
    return nil
}

func (x *MyUnion) readField4(p thrift.Format) error {  // FloatSet
    _ /* elemType */, size, err := p.ReadSetBegin()
if err != nil {
    return thrift.PrependError("error reading set begin: ", err)
}

setResult := make([]float32, 0, size)
for i := 0; i < size; i++ {
    var elem float32
    {
        result, err := p.ReadFloat()
if err != nil {
    return err
}
        elem = result
    }
    setResult = append(setResult, elem)
}

if err := p.ReadSetEnd(); err != nil {
    return thrift.PrependError("error reading set end: ", err)
}
result := setResult

    x.FloatSet = result
    return nil
}

func (x *MyUnion) toString1() string {  // MyEnum
    if x.IsSetMyEnum() {
        return fmt.Sprintf("%v", *x.MyEnum)
    }
    return fmt.Sprintf("%v", x.MyEnum)
}

func (x *MyUnion) toString2() string {  // MyStruct
    return fmt.Sprintf("%v", x.MyStruct)
}

func (x *MyUnion) toString3() string {  // MyDataItem
    return fmt.Sprintf("%v", x.MyDataItem)
}

func (x *MyUnion) toString4() string {  // FloatSet
    return fmt.Sprintf("%v", x.FloatSet)
}


// Deprecated: Use NewMyUnion().GetMyStruct() instead.
func (x *MyUnion) DefaultGetMyStruct() *MyStruct {
    if !x.IsSetMyStruct() {
        return NewMyStruct()
    }
    return x.MyStruct
}

// Deprecated: Use NewMyUnion().GetMyDataItem() instead.
func (x *MyUnion) DefaultGetMyDataItem() *MyDataItem {
    if !x.IsSetMyDataItem() {
        return NewMyDataItem()
    }
    return x.MyDataItem
}

func (x *MyUnion) countSetFields() int {
    count := int(0)
    if (x.IsSetMyEnum()) {
        count++
    }
    if (x.IsSetMyStruct()) {
        count++
    }
    if (x.IsSetMyDataItem()) {
        count++
    }
    if (x.IsSetFloatSet()) {
        count++
    }
    return count
}

func (x *MyUnion) CountSetFieldsMyUnion() int {
    return x.countSetFields()
}



func (x *MyUnion) Write(p thrift.Format) error {
    if countSet := x.countSetFields(); countSet > 1 {
        return fmt.Errorf("%T write union: no more than one field must be set (%d set).", x, countSet)
    }
    if err := p.WriteStructBegin("MyUnion"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyUnion) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.I32)):  // myEnum
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.STRUCT)):  // myStruct
            if err := x.readField2(p); err != nil {
                return err
            }
        case (id == 3 && wireType == thrift.Type(thrift.STRUCT)):  // myDataItem
            if err := x.readField3(p); err != nil {
                return err
            }
        case (id == 4 && wireType == thrift.Type(thrift.SET)):  // floatSet
            if err := x.readField4(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *MyUnion) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("MyUnion({")
    sb.WriteString(fmt.Sprintf("MyEnum:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("MyStruct:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("MyDataItem:%s ", x.toString3()))
    sb.WriteString(fmt.Sprintf("FloatSet:%s", x.toString4()))
    sb.WriteString("})")

    return sb.String()
}

type ReservedKeyword struct {
    ReservedField int32 `thrift:"reserved_field,1" json:"reserved_field" db:"reserved_field"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*ReservedKeyword)(nil)

func NewReservedKeyword() *ReservedKeyword {
    return (&ReservedKeyword{}).
        SetReservedFieldNonCompat(0)
}

func (x *ReservedKeyword) GetReservedField() int32 {
    return x.ReservedField
}

func (x *ReservedKeyword) SetReservedFieldNonCompat(value int32) *ReservedKeyword {
    x.ReservedField = value
    return x
}

func (x *ReservedKeyword) SetReservedField(value int32) *ReservedKeyword {
    x.ReservedField = value
    return x
}

func (x *ReservedKeyword) writeField1(p thrift.Format) error {  // ReservedField
    if err := p.WriteFieldBegin("reserved_field", thrift.I32, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.ReservedField
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *ReservedKeyword) readField1(p thrift.Format) error {  // ReservedField
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.ReservedField = result
    return nil
}

func (x *ReservedKeyword) toString1() string {  // ReservedField
    return fmt.Sprintf("%v", x.ReservedField)
}



func (x *ReservedKeyword) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("ReservedKeyword"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *ReservedKeyword) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.I32)):  // reserved_field
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *ReservedKeyword) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("ReservedKeyword({")
    sb.WriteString(fmt.Sprintf("ReservedField:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type UnionToBeRenamed struct {
    ReservedField *int32 `thrift:"reserved_field,1" json:"reserved_field,omitempty" db:"reserved_field"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*UnionToBeRenamed)(nil)

func NewUnionToBeRenamed() *UnionToBeRenamed {
    return (&UnionToBeRenamed{})
}

func (x *UnionToBeRenamed) GetReservedField() int32 {
    if !x.IsSetReservedField() {
        return 0
    }

    return *x.ReservedField
}

func (x *UnionToBeRenamed) SetReservedFieldNonCompat(value int32) *UnionToBeRenamed {
    x.ReservedField = &value
    return x
}

func (x *UnionToBeRenamed) SetReservedField(value *int32) *UnionToBeRenamed {
    x.ReservedField = value
    return x
}

func (x *UnionToBeRenamed) IsSetReservedField() bool {
    return x != nil && x.ReservedField != nil
}

func (x *UnionToBeRenamed) writeField1(p thrift.Format) error {  // ReservedField
    if !x.IsSetReservedField() {
        return nil
    }

    if err := p.WriteFieldBegin("reserved_field", thrift.I32, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.ReservedField
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *UnionToBeRenamed) readField1(p thrift.Format) error {  // ReservedField
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.ReservedField = &result
    return nil
}

func (x *UnionToBeRenamed) toString1() string {  // ReservedField
    if x.IsSetReservedField() {
        return fmt.Sprintf("%v", *x.ReservedField)
    }
    return fmt.Sprintf("%v", x.ReservedField)
}


func (x *UnionToBeRenamed) countSetFields() int {
    count := int(0)
    if (x.IsSetReservedField()) {
        count++
    }
    return count
}

func (x *UnionToBeRenamed) CountSetFieldsUnionToBeRenamed() int {
    return x.countSetFields()
}



func (x *UnionToBeRenamed) Write(p thrift.Format) error {
    if countSet := x.countSetFields(); countSet > 1 {
        return fmt.Errorf("%T write union: no more than one field must be set (%d set).", x, countSet)
    }
    if err := p.WriteStructBegin("UnionToBeRenamed"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *UnionToBeRenamed) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.I32)):  // reserved_field
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *UnionToBeRenamed) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("UnionToBeRenamed({")
    sb.WriteString(fmt.Sprintf("ReservedField:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {
    registry.RegisterType("test.dev/fixtures/basic/MyStruct", func() any { return NewMyStruct() })
    registry.RegisterType("test.dev/fixtures/basic/MyDataItem", func() any { return NewMyDataItem() })
    registry.RegisterType("test.dev/fixtures/basic/MyUnion", func() any { return NewMyUnion() })
    registry.RegisterType("test.dev/fixtures/basic/ReservedKeyword", func() any { return NewReservedKeyword() })
    registry.RegisterType("test.dev/fixtures/basic/UnionToBeRenamed", func() any { return NewUnionToBeRenamed() })

    registry.RegisterType("test.dev/fixtures/basic/MyEnum", func() any { return MyEnum(0) })
    registry.RegisterType("test.dev/fixtures/basic/HackEnum", func() any { return HackEnum(0) })
}
