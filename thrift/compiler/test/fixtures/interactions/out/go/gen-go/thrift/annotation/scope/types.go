// Autogenerated by Thrift for thrift/annotation/scope.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package scope

import (
    "fmt"
    "reflect"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = reflect.Ptr
var _ = thrift.VOID

type Transitive struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Transitive)(nil)

func NewTransitive() *Transitive {
    return (&Transitive{}).setDefaults()
}



func (x *Transitive) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Transitive"); err != nil {
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

func (x *Transitive) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Transitive) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Transitive) setDefaults() *Transitive {
    return x
}

type Program struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Program)(nil)

func NewProgram() *Program {
    return (&Program{}).setDefaults()
}



func (x *Program) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Program"); err != nil {
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

func (x *Program) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Program) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Program) setDefaults() *Program {
    return x
}

type Struct struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Struct)(nil)

func NewStruct() *Struct {
    return (&Struct{}).setDefaults()
}



func (x *Struct) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Struct"); err != nil {
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

func (x *Struct) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Struct) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Struct) setDefaults() *Struct {
    return x
}

type Union struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Union)(nil)

func NewUnion() *Union {
    return (&Union{}).setDefaults()
}



func (x *Union) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Union"); err != nil {
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

func (x *Union) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Union) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Union) setDefaults() *Union {
    return x
}

type Exception struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Exception)(nil)

func NewException() *Exception {
    return (&Exception{}).setDefaults()
}



func (x *Exception) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Exception"); err != nil {
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

func (x *Exception) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Exception) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Exception) setDefaults() *Exception {
    return x
}

type Field struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Field)(nil)

func NewField() *Field {
    return (&Field{}).setDefaults()
}



func (x *Field) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Field"); err != nil {
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

func (x *Field) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Field) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Field) setDefaults() *Field {
    return x
}

type Typedef struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Typedef)(nil)

func NewTypedef() *Typedef {
    return (&Typedef{}).setDefaults()
}



func (x *Typedef) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Typedef"); err != nil {
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

func (x *Typedef) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Typedef) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Typedef) setDefaults() *Typedef {
    return x
}

type Service struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Service)(nil)

func NewService() *Service {
    return (&Service{}).setDefaults()
}



func (x *Service) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Service"); err != nil {
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

func (x *Service) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Service) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Service) setDefaults() *Service {
    return x
}

type Interaction struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Interaction)(nil)

func NewInteraction() *Interaction {
    return (&Interaction{}).setDefaults()
}



func (x *Interaction) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Interaction"); err != nil {
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

func (x *Interaction) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Interaction) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Interaction) setDefaults() *Interaction {
    return x
}

type Function struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Function)(nil)

func NewFunction() *Function {
    return (&Function{}).setDefaults()
}



func (x *Function) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Function"); err != nil {
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

func (x *Function) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Function) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Function) setDefaults() *Function {
    return x
}

type EnumValue struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*EnumValue)(nil)

func NewEnumValue() *EnumValue {
    return (&EnumValue{}).setDefaults()
}



func (x *EnumValue) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("EnumValue"); err != nil {
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

func (x *EnumValue) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *EnumValue) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *EnumValue) setDefaults() *EnumValue {
    return x
}

type Const struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Const)(nil)

func NewConst() *Const {
    return (&Const{}).setDefaults()
}



func (x *Const) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Const"); err != nil {
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

func (x *Const) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Const) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Const) setDefaults() *Const {
    return x
}

type Enum struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Enum)(nil)

func NewEnum() *Enum {
    return (&Enum{}).setDefaults()
}



func (x *Enum) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Enum"); err != nil {
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

func (x *Enum) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Enum) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Enum) setDefaults() *Enum {
    return x
}

type Structured struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Structured)(nil)

func NewStructured() *Structured {
    return (&Structured{}).setDefaults()
}



func (x *Structured) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Structured"); err != nil {
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

func (x *Structured) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Structured) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Structured) setDefaults() *Structured {
    return x
}

type Interface struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Interface)(nil)

func NewInterface() *Interface {
    return (&Interface{}).setDefaults()
}



func (x *Interface) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Interface"); err != nil {
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

func (x *Interface) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Interface) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Interface) setDefaults() *Interface {
    return x
}

type RootDefinition struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*RootDefinition)(nil)

func NewRootDefinition() *RootDefinition {
    return (&RootDefinition{}).setDefaults()
}



func (x *RootDefinition) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("RootDefinition"); err != nil {
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

func (x *RootDefinition) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *RootDefinition) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *RootDefinition) setDefaults() *RootDefinition {
    return x
}

type Definition struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Definition)(nil)

func NewDefinition() *Definition {
    return (&Definition{}).setDefaults()
}



func (x *Definition) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Definition"); err != nil {
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

func (x *Definition) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *Definition) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Definition) setDefaults() *Definition {
    return x
}

type DisableSchemaConst struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*DisableSchemaConst)(nil)

func NewDisableSchemaConst() *DisableSchemaConst {
    return (&DisableSchemaConst{}).setDefaults()
}



func (x *DisableSchemaConst) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("DisableSchemaConst"); err != nil {
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

func (x *DisableSchemaConst) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        fieldName, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d ('%s') read error: ", x, id, fieldName), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
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

func (x *DisableSchemaConst) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *DisableSchemaConst) setDefaults() *DisableSchemaConst {
    return x
}



// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {
    registry.RegisterType("facebook.com/thrift/annotation/Transitive", func() any { return NewTransitive() })
    registry.RegisterType("facebook.com/thrift/annotation/Program", func() any { return NewProgram() })
    registry.RegisterType("facebook.com/thrift/annotation/Struct", func() any { return NewStruct() })
    registry.RegisterType("facebook.com/thrift/annotation/Union", func() any { return NewUnion() })
    registry.RegisterType("facebook.com/thrift/annotation/Exception", func() any { return NewException() })
    registry.RegisterType("facebook.com/thrift/annotation/Field", func() any { return NewField() })
    registry.RegisterType("facebook.com/thrift/annotation/Typedef", func() any { return NewTypedef() })
    registry.RegisterType("facebook.com/thrift/annotation/Service", func() any { return NewService() })
    registry.RegisterType("facebook.com/thrift/annotation/Interaction", func() any { return NewInteraction() })
    registry.RegisterType("facebook.com/thrift/annotation/Function", func() any { return NewFunction() })
    registry.RegisterType("facebook.com/thrift/annotation/EnumValue", func() any { return NewEnumValue() })
    registry.RegisterType("facebook.com/thrift/annotation/Const", func() any { return NewConst() })
    registry.RegisterType("facebook.com/thrift/annotation/Enum", func() any { return NewEnum() })
    registry.RegisterType("facebook.com/thrift/annotation/Structured", func() any { return NewStructured() })
    registry.RegisterType("facebook.com/thrift/annotation/Interface", func() any { return NewInterface() })
    registry.RegisterType("facebook.com/thrift/annotation/RootDefinition", func() any { return NewRootDefinition() })
    registry.RegisterType("facebook.com/thrift/annotation/Definition", func() any { return NewDefinition() })
    registry.RegisterType("facebook.com/thrift/annotation/DisableSchemaConst", func() any { return NewDisableSchemaConst() })

}
