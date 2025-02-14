// Autogenerated by Thrift for foo.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package foo

import (
    "fmt"
    "reflect"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = reflect.Ptr
var _ = thrift.VOID

type Fields struct {
    InjectedField string `thrift:"injected_field,100" json:"injected_field" db:"injected_field"`
    InjectedStructuredAnnotationField *string `thrift:"injected_structured_annotation_field,101,optional" json:"injected_structured_annotation_field,omitempty" db:"injected_structured_annotation_field"`
    InjectedUnstructuredAnnotationField *string `thrift:"injected_unstructured_annotation_field,102,optional" json:"injected_unstructured_annotation_field,omitempty" db:"injected_unstructured_annotation_field"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Fields)(nil)

func NewFields() *Fields {
    return (&Fields{}).setDefaults()
}

func (x *Fields) GetInjectedField() string {
    return x.InjectedField
}

func (x *Fields) GetInjectedStructuredAnnotationField() string {
    if !x.IsSetInjectedStructuredAnnotationField() {
        return ""
    }
    return *x.InjectedStructuredAnnotationField
}

func (x *Fields) GetInjectedUnstructuredAnnotationField() string {
    if !x.IsSetInjectedUnstructuredAnnotationField() {
        return ""
    }
    return *x.InjectedUnstructuredAnnotationField
}

func (x *Fields) SetInjectedFieldNonCompat(value string) *Fields {
    x.InjectedField = value
    return x
}

func (x *Fields) SetInjectedField(value string) *Fields {
    x.InjectedField = value
    return x
}

func (x *Fields) SetInjectedStructuredAnnotationFieldNonCompat(value string) *Fields {
    x.InjectedStructuredAnnotationField = &value
    return x
}

func (x *Fields) SetInjectedStructuredAnnotationField(value *string) *Fields {
    x.InjectedStructuredAnnotationField = value
    return x
}

func (x *Fields) SetInjectedUnstructuredAnnotationFieldNonCompat(value string) *Fields {
    x.InjectedUnstructuredAnnotationField = &value
    return x
}

func (x *Fields) SetInjectedUnstructuredAnnotationField(value *string) *Fields {
    x.InjectedUnstructuredAnnotationField = value
    return x
}

func (x *Fields) IsSetInjectedStructuredAnnotationField() bool {
    return x != nil && x.InjectedStructuredAnnotationField != nil
}

func (x *Fields) IsSetInjectedUnstructuredAnnotationField() bool {
    return x != nil && x.InjectedUnstructuredAnnotationField != nil
}

func (x *Fields) writeField100(p thrift.Encoder) error {  // InjectedField
    if err := p.WriteFieldBegin("injected_field", thrift.STRING, 100); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.InjectedField
    if err := p.WriteString(item); err != nil {
        return err
    }

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Fields) writeField101(p thrift.Encoder) error {  // InjectedStructuredAnnotationField
    if !x.IsSetInjectedStructuredAnnotationField() {
        return nil
    }

    if err := p.WriteFieldBegin("injected_structured_annotation_field", thrift.STRING, 101); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.InjectedStructuredAnnotationField
    if err := p.WriteString(item); err != nil {
        return err
    }

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Fields) writeField102(p thrift.Encoder) error {  // InjectedUnstructuredAnnotationField
    if !x.IsSetInjectedUnstructuredAnnotationField() {
        return nil
    }

    if err := p.WriteFieldBegin("injected_unstructured_annotation_field", thrift.STRING, 102); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.InjectedUnstructuredAnnotationField
    if err := p.WriteString(item); err != nil {
        return err
    }

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Fields) readField100(p thrift.Decoder) error {  // InjectedField
    result, err := p.ReadString()
    if err != nil {
        return err
    }

    x.InjectedField = result
    return nil
}

func (x *Fields) readField101(p thrift.Decoder) error {  // InjectedStructuredAnnotationField
    result, err := p.ReadString()
    if err != nil {
        return err
    }

    x.InjectedStructuredAnnotationField = &result
    return nil
}

func (x *Fields) readField102(p thrift.Decoder) error {  // InjectedUnstructuredAnnotationField
    result, err := p.ReadString()
    if err != nil {
        return err
    }

    x.InjectedUnstructuredAnnotationField = &result
    return nil
}





func (x *Fields) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Fields"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField100(p); err != nil {
        return err
    }
    if err := x.writeField101(p); err != nil {
        return err
    }
    if err := x.writeField102(p); err != nil {
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

func (x *Fields) Read(p thrift.Decoder) error {
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
        case ((id == 100 && wireType == thrift.STRING) || (id == thrift.NO_FIELD_ID && fieldName == "injected_field")):  // injected_field
            fieldReadErr = x.readField100(p)
        case ((id == 101 && wireType == thrift.STRING) || (id == thrift.NO_FIELD_ID && fieldName == "injected_structured_annotation_field")):  // injected_structured_annotation_field
            fieldReadErr = x.readField101(p)
        case ((id == 102 && wireType == thrift.STRING) || (id == thrift.NO_FIELD_ID && fieldName == "injected_unstructured_annotation_field")):  // injected_unstructured_annotation_field
            fieldReadErr = x.readField102(p)
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

func (x *Fields) String() string {
    return thrift.StructToString(reflect.ValueOf(x))
}

func (x *Fields) setDefaults() *Fields {
    return x.
        SetInjectedFieldNonCompat("")
}



// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {

}
