// Autogenerated by Thrift for thrift/compiler/test/fixtures/interactions/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module

import (
    "fmt"
    "reflect"

    shared "shared"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

var _ = shared.GoUnusedProtection__
// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = reflect.Ptr
var _ = thrift.ZERO

type CustomException struct {
    Message string `thrift:"message,1" json:"message" db:"message"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*CustomException)(nil)

func NewCustomException() *CustomException {
    return (&CustomException{}).setDefaults()
}

func (x *CustomException) GetMessage() string {
    return x.Message
}

func (x *CustomException) SetMessageNonCompat(value string) *CustomException {
    x.Message = value
    return x
}

func (x *CustomException) SetMessage(value string) *CustomException {
    x.Message = value
    return x
}

func (x *CustomException) writeField1(p thrift.Encoder) error {  // Message
    if err := p.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Message
    if err := p.WriteString(item); err != nil {
        return err
    }

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *CustomException) readField1(p thrift.Decoder) error {  // Message
    result, err := p.ReadString()
    if err != nil {
        return err
    }

    x.Message = result
    return nil
}



func (x *CustomException) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("CustomException"); err != nil {
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

func (x *CustomException) Read(p thrift.Decoder) error {
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

        var fieldReadErr error
        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // message
            fieldReadErr = x.readField1(p)
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

func (x *CustomException) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *CustomException) setDefaults() *CustomException {
    return x.
        SetMessageNonCompat("")
}

func (x *CustomException) Error() string {
    return x.String()
}

// Service req/resp structs (below)
type reqMyServiceFoo struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*reqMyServiceFoo)(nil)

// Deprecated: MyServiceFooArgsDeprecated is deprecated, since it is supposed to be internal.
type MyServiceFooArgsDeprecated = reqMyServiceFoo

func newReqMyServiceFoo() *reqMyServiceFoo {
    return (&reqMyServiceFoo{}).setDefaults()
}



func (x *reqMyServiceFoo) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("reqMyServiceFoo"); err != nil {
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

func (x *reqMyServiceFoo) Read(p thrift.Decoder) error {
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

func (x *reqMyServiceFoo) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *reqMyServiceFoo) setDefaults() *reqMyServiceFoo {
    return x
}

type respMyServiceFoo struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*respMyServiceFoo)(nil)
var _ thrift.WritableResult = (*respMyServiceFoo)(nil)

// Deprecated: MyServiceFooResultDeprecated is deprecated, since it is supposed to be internal.
type MyServiceFooResultDeprecated = respMyServiceFoo

func newRespMyServiceFoo() *respMyServiceFoo {
    return (&respMyServiceFoo{}).setDefaults()
}



func (x *respMyServiceFoo) Exception() thrift.WritableException {
    return nil
}

func (x *respMyServiceFoo) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("respMyServiceFoo"); err != nil {
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

func (x *respMyServiceFoo) Read(p thrift.Decoder) error {
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

func (x *respMyServiceFoo) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *respMyServiceFoo) setDefaults() *respMyServiceFoo {
    return x
}

type reqFactoriesFoo struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*reqFactoriesFoo)(nil)

// Deprecated: FactoriesFooArgsDeprecated is deprecated, since it is supposed to be internal.
type FactoriesFooArgsDeprecated = reqFactoriesFoo

func newReqFactoriesFoo() *reqFactoriesFoo {
    return (&reqFactoriesFoo{}).setDefaults()
}



func (x *reqFactoriesFoo) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("reqFactoriesFoo"); err != nil {
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

func (x *reqFactoriesFoo) Read(p thrift.Decoder) error {
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

func (x *reqFactoriesFoo) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *reqFactoriesFoo) setDefaults() *reqFactoriesFoo {
    return x
}

type respFactoriesFoo struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*respFactoriesFoo)(nil)
var _ thrift.WritableResult = (*respFactoriesFoo)(nil)

// Deprecated: FactoriesFooResultDeprecated is deprecated, since it is supposed to be internal.
type FactoriesFooResultDeprecated = respFactoriesFoo

func newRespFactoriesFoo() *respFactoriesFoo {
    return (&respFactoriesFoo{}).setDefaults()
}



func (x *respFactoriesFoo) Exception() thrift.WritableException {
    return nil
}

func (x *respFactoriesFoo) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("respFactoriesFoo"); err != nil {
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

func (x *respFactoriesFoo) Read(p thrift.Decoder) error {
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

func (x *respFactoriesFoo) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *respFactoriesFoo) setDefaults() *respFactoriesFoo {
    return x
}

type reqPerformFoo struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*reqPerformFoo)(nil)

// Deprecated: PerformFooArgsDeprecated is deprecated, since it is supposed to be internal.
type PerformFooArgsDeprecated = reqPerformFoo

func newReqPerformFoo() *reqPerformFoo {
    return (&reqPerformFoo{}).setDefaults()
}



func (x *reqPerformFoo) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("reqPerformFoo"); err != nil {
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

func (x *reqPerformFoo) Read(p thrift.Decoder) error {
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

func (x *reqPerformFoo) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *reqPerformFoo) setDefaults() *reqPerformFoo {
    return x
}

type respPerformFoo struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*respPerformFoo)(nil)
var _ thrift.WritableResult = (*respPerformFoo)(nil)

// Deprecated: PerformFooResultDeprecated is deprecated, since it is supposed to be internal.
type PerformFooResultDeprecated = respPerformFoo

func newRespPerformFoo() *respPerformFoo {
    return (&respPerformFoo{}).setDefaults()
}



func (x *respPerformFoo) Exception() thrift.WritableException {
    return nil
}

func (x *respPerformFoo) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("respPerformFoo"); err != nil {
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

func (x *respPerformFoo) Read(p thrift.Decoder) error {
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

func (x *respPerformFoo) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *respPerformFoo) setDefaults() *respPerformFoo {
    return x
}

type reqInteractWithSharedDoSomeSimilarThings struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*reqInteractWithSharedDoSomeSimilarThings)(nil)

// Deprecated: InteractWithSharedDoSomeSimilarThingsArgsDeprecated is deprecated, since it is supposed to be internal.
type InteractWithSharedDoSomeSimilarThingsArgsDeprecated = reqInteractWithSharedDoSomeSimilarThings

func newReqInteractWithSharedDoSomeSimilarThings() *reqInteractWithSharedDoSomeSimilarThings {
    return (&reqInteractWithSharedDoSomeSimilarThings{}).setDefaults()
}



func (x *reqInteractWithSharedDoSomeSimilarThings) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("reqInteractWithSharedDoSomeSimilarThings"); err != nil {
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

func (x *reqInteractWithSharedDoSomeSimilarThings) Read(p thrift.Decoder) error {
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

func (x *reqInteractWithSharedDoSomeSimilarThings) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *reqInteractWithSharedDoSomeSimilarThings) setDefaults() *reqInteractWithSharedDoSomeSimilarThings {
    return x
}

type respInteractWithSharedDoSomeSimilarThings struct {
    Success *shared.DoSomethingResult `thrift:"success,0,optional" json:"success,omitempty" db:"success"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*respInteractWithSharedDoSomeSimilarThings)(nil)
var _ thrift.WritableResult = (*respInteractWithSharedDoSomeSimilarThings)(nil)

// Deprecated: InteractWithSharedDoSomeSimilarThingsResultDeprecated is deprecated, since it is supposed to be internal.
type InteractWithSharedDoSomeSimilarThingsResultDeprecated = respInteractWithSharedDoSomeSimilarThings

func newRespInteractWithSharedDoSomeSimilarThings() *respInteractWithSharedDoSomeSimilarThings {
    return (&respInteractWithSharedDoSomeSimilarThings{}).setDefaults()
}

func (x *respInteractWithSharedDoSomeSimilarThings) GetSuccess() *shared.DoSomethingResult {
    if !x.IsSetSuccess() {
        return nil
    }
    return x.Success
}

func (x *respInteractWithSharedDoSomeSimilarThings) SetSuccessNonCompat(value *shared.DoSomethingResult) *respInteractWithSharedDoSomeSimilarThings {
    x.Success = value
    return x
}

func (x *respInteractWithSharedDoSomeSimilarThings) SetSuccess(value *shared.DoSomethingResult) *respInteractWithSharedDoSomeSimilarThings {
    x.Success = value
    return x
}

func (x *respInteractWithSharedDoSomeSimilarThings) IsSetSuccess() bool {
    return x != nil && x.Success != nil
}

func (x *respInteractWithSharedDoSomeSimilarThings) writeField0(p thrift.Encoder) error {  // Success
    if !x.IsSetSuccess() {
        return nil
    }

    if err := p.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Success
    if err := item.Write(p); err != nil {
        return err
    }

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *respInteractWithSharedDoSomeSimilarThings) readField0(p thrift.Decoder) error {  // Success
    result := shared.NewDoSomethingResult()
    err := result.Read(p)
    if err != nil {
        return err
    }

    x.Success = result
    return nil
}

// Deprecated: Use newRespInteractWithSharedDoSomeSimilarThings().GetSuccess() instead.
func (x *respInteractWithSharedDoSomeSimilarThings) DefaultGetSuccess() *shared.DoSomethingResult {
    if !x.IsSetSuccess() {
        return shared.NewDoSomethingResult()
    }
    return x.Success
}



func (x *respInteractWithSharedDoSomeSimilarThings) Exception() thrift.WritableException {
    return nil
}

func (x *respInteractWithSharedDoSomeSimilarThings) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("respInteractWithSharedDoSomeSimilarThings"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField0(p); err != nil {
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

func (x *respInteractWithSharedDoSomeSimilarThings) Read(p thrift.Decoder) error {
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

        var fieldReadErr error
        switch {
        case (id == 0 && wireType == thrift.Type(thrift.STRUCT)):  // success
            fieldReadErr = x.readField0(p)
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

func (x *respInteractWithSharedDoSomeSimilarThings) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *respInteractWithSharedDoSomeSimilarThings) setDefaults() *respInteractWithSharedDoSomeSimilarThings {
    return x
}


// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {

}
