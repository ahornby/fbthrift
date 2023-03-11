// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package module // [[[ program thrift source path ]]]

import (
  "fmt"

  "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)


// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type GetEntityRequest struct {
    Id string `thrift:"id,1" json:"id" db:"id"`
}
// Compile time interface enforcer
var _ thrift.Struct = &GetEntityRequest{}

func NewGetEntityRequest() *GetEntityRequest {
    return (&GetEntityRequest{})
}

func (x *GetEntityRequest) GetIdNonCompat() string {
    return x.Id
}

func (x *GetEntityRequest) GetId() string {
    return x.Id
}

func (x *GetEntityRequest) SetId(value string) *GetEntityRequest {
    x.Id = value
    return x
}


func (x *GetEntityRequest) writeField1(p thrift.Protocol) error {  // Id
    if err := p.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetIdNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *GetEntityRequest) readField1(p thrift.Protocol) error {  // Id
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetId(result)
    return nil
}

func (x *GetEntityRequest) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use GetEntityRequest.Set* methods instead or set the fields directly.
type GetEntityRequestBuilder struct {
    obj *GetEntityRequest
}

func NewGetEntityRequestBuilder() *GetEntityRequestBuilder {
    return &GetEntityRequestBuilder{
        obj: NewGetEntityRequest(),
    }
}

func (x *GetEntityRequestBuilder) Id(value string) *GetEntityRequestBuilder {
    x.obj.Id = value
    return x
}

func (x *GetEntityRequestBuilder) Emit() *GetEntityRequest {
    var objCopy GetEntityRequest = *x.obj
    return &objCopy
}
func (x *GetEntityRequest) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("GetEntityRequest"); err != nil {
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

func (x *GetEntityRequest) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // id
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
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

type GetEntityResponse struct {
    Entity string `thrift:"entity,1" json:"entity" db:"entity"`
}
// Compile time interface enforcer
var _ thrift.Struct = &GetEntityResponse{}

func NewGetEntityResponse() *GetEntityResponse {
    return (&GetEntityResponse{})
}

func (x *GetEntityResponse) GetEntityNonCompat() string {
    return x.Entity
}

func (x *GetEntityResponse) GetEntity() string {
    return x.Entity
}

func (x *GetEntityResponse) SetEntity(value string) *GetEntityResponse {
    x.Entity = value
    return x
}


func (x *GetEntityResponse) writeField1(p thrift.Protocol) error {  // Entity
    if err := p.WriteFieldBegin("entity", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetEntityNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *GetEntityResponse) readField1(p thrift.Protocol) error {  // Entity
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetEntity(result)
    return nil
}

func (x *GetEntityResponse) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use GetEntityResponse.Set* methods instead or set the fields directly.
type GetEntityResponseBuilder struct {
    obj *GetEntityResponse
}

func NewGetEntityResponseBuilder() *GetEntityResponseBuilder {
    return &GetEntityResponseBuilder{
        obj: NewGetEntityResponse(),
    }
}

func (x *GetEntityResponseBuilder) Entity(value string) *GetEntityResponseBuilder {
    x.obj.Entity = value
    return x
}

func (x *GetEntityResponseBuilder) Emit() *GetEntityResponse {
    var objCopy GetEntityResponse = *x.obj
    return &objCopy
}
func (x *GetEntityResponse) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("GetEntityResponse"); err != nil {
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

func (x *GetEntityResponse) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // entity
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
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
