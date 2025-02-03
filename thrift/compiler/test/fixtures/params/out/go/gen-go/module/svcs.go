// Autogenerated by Thrift for thrift/compiler/test/fixtures/params/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module


import (
    "context"
    "fmt"
    "reflect"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = context.Background
var _ = fmt.Printf
var _ = reflect.Ptr
var _ = thrift.ZERO
var _ = metadata.GoUnusedProtection__

type NestedContainers interface {
    MapList(ctx context.Context, foo map[int32][]int32) (error)
    MapSet(ctx context.Context, foo map[int32][]int32) (error)
    ListMap(ctx context.Context, foo []map[int32]int32) (error)
    ListSet(ctx context.Context, foo [][]int32) (error)
    Turtles(ctx context.Context, foo [][]map[int32]map[int32][]int32) (error)
}

type NestedContainersChannelClientInterface interface {
    thrift.ClientInterface
    NestedContainers
}

type NestedContainersClientInterface interface {
    thrift.ClientInterface
    NestedContainers
}

type NestedContainersChannelClient struct {
    ch thrift.RequestChannel
}
// Compile time interface enforcer
var _ NestedContainersChannelClientInterface = (*NestedContainersChannelClient)(nil)

func NewNestedContainersChannelClient(channel thrift.RequestChannel) *NestedContainersChannelClient {
    return &NestedContainersChannelClient{
        ch: channel,
    }
}

func (c *NestedContainersChannelClient) Close() error {
    return c.ch.Close()
}

type NestedContainersClient struct {
    chClient *NestedContainersChannelClient
}
// Compile time interface enforcer
var _ NestedContainersClientInterface = (*NestedContainersClient)(nil)

func NewNestedContainersClient(prot thrift.Protocol) *NestedContainersClient {
    return &NestedContainersClient{
        chClient: NewNestedContainersChannelClient(
            thrift.NewSerialChannel(prot),
        ),
    }
}

func (c *NestedContainersClient) Close() error {
    return c.chClient.Close()
}

func (c *NestedContainersChannelClient) MapList(ctx context.Context, foo map[int32][]int32) (error) {
    in := &reqNestedContainersMapList{
        Foo: foo,
    }
    out := newRespNestedContainersMapList()
    err := c.ch.Call(ctx, "mapList", in, out)
    if err != nil {
        return err
    }
    return nil
}

func (c *NestedContainersClient) MapList(ctx context.Context, foo map[int32][]int32) (error) {
    return c.chClient.MapList(ctx, foo)
}

func (c *NestedContainersChannelClient) MapSet(ctx context.Context, foo map[int32][]int32) (error) {
    in := &reqNestedContainersMapSet{
        Foo: foo,
    }
    out := newRespNestedContainersMapSet()
    err := c.ch.Call(ctx, "mapSet", in, out)
    if err != nil {
        return err
    }
    return nil
}

func (c *NestedContainersClient) MapSet(ctx context.Context, foo map[int32][]int32) (error) {
    return c.chClient.MapSet(ctx, foo)
}

func (c *NestedContainersChannelClient) ListMap(ctx context.Context, foo []map[int32]int32) (error) {
    in := &reqNestedContainersListMap{
        Foo: foo,
    }
    out := newRespNestedContainersListMap()
    err := c.ch.Call(ctx, "listMap", in, out)
    if err != nil {
        return err
    }
    return nil
}

func (c *NestedContainersClient) ListMap(ctx context.Context, foo []map[int32]int32) (error) {
    return c.chClient.ListMap(ctx, foo)
}

func (c *NestedContainersChannelClient) ListSet(ctx context.Context, foo [][]int32) (error) {
    in := &reqNestedContainersListSet{
        Foo: foo,
    }
    out := newRespNestedContainersListSet()
    err := c.ch.Call(ctx, "listSet", in, out)
    if err != nil {
        return err
    }
    return nil
}

func (c *NestedContainersClient) ListSet(ctx context.Context, foo [][]int32) (error) {
    return c.chClient.ListSet(ctx, foo)
}

func (c *NestedContainersChannelClient) Turtles(ctx context.Context, foo [][]map[int32]map[int32][]int32) (error) {
    in := &reqNestedContainersTurtles{
        Foo: foo,
    }
    out := newRespNestedContainersTurtles()
    err := c.ch.Call(ctx, "turtles", in, out)
    if err != nil {
        return err
    }
    return nil
}

func (c *NestedContainersClient) Turtles(ctx context.Context, foo [][]map[int32]map[int32][]int32) (error) {
    return c.chClient.Turtles(ctx, foo)
}


type NestedContainersProcessor struct {
    processorFunctionMap map[string]thrift.ProcessorFunction
    functionServiceMap   map[string]string
    handler            NestedContainers
}

func NewNestedContainersProcessor(handler NestedContainers) *NestedContainersProcessor {
    p := &NestedContainersProcessor{
        handler:              handler,
        processorFunctionMap: make(map[string]thrift.ProcessorFunction),
        functionServiceMap:   make(map[string]string),
    }
    p.AddToProcessorFunctionMap("mapList", &procFuncNestedContainersMapList{handler: handler})
    p.AddToProcessorFunctionMap("mapSet", &procFuncNestedContainersMapSet{handler: handler})
    p.AddToProcessorFunctionMap("listMap", &procFuncNestedContainersListMap{handler: handler})
    p.AddToProcessorFunctionMap("listSet", &procFuncNestedContainersListSet{handler: handler})
    p.AddToProcessorFunctionMap("turtles", &procFuncNestedContainersTurtles{handler: handler})
    p.AddToFunctionServiceMap("mapList", "NestedContainers")
    p.AddToFunctionServiceMap("mapSet", "NestedContainers")
    p.AddToFunctionServiceMap("listMap", "NestedContainers")
    p.AddToFunctionServiceMap("listSet", "NestedContainers")
    p.AddToFunctionServiceMap("turtles", "NestedContainers")

    return p
}

func (p *NestedContainersProcessor) AddToProcessorFunctionMap(key string, processorFunction thrift.ProcessorFunction) {
    p.processorFunctionMap[key] = processorFunction
}

func (p *NestedContainersProcessor) AddToFunctionServiceMap(key, service string) {
    p.functionServiceMap[key] = service
}

func (p *NestedContainersProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction) {
    return p.processorFunctionMap[key]
}

func (p *NestedContainersProcessor) ProcessorFunctionMap() map[string]thrift.ProcessorFunction {
    return p.processorFunctionMap
}

func (p *NestedContainersProcessor) FunctionServiceMap() map[string]string {
    return p.functionServiceMap
}

func (p *NestedContainersProcessor) PackageName() string {
    return "module"
}

func (p *NestedContainersProcessor) GetThriftMetadata() *metadata.ThriftMetadata {
    return GetThriftMetadataForService("module.NestedContainers")
}


type procFuncNestedContainersMapList struct {
    handler NestedContainers
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncNestedContainersMapList)(nil)

func (p *procFuncNestedContainersMapList) Read(iprot thrift.Decoder) (thrift.Struct, thrift.Exception) {
    args := newReqNestedContainersMapList()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncNestedContainersMapList) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("mapList", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncNestedContainersMapList) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqNestedContainersMapList)
    result := newRespNestedContainersMapList()
    err := p.handler.MapList(ctx, args.Foo)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing MapList: " + err.Error(), err)
        return x, x
    }

    return result, nil
}


type procFuncNestedContainersMapSet struct {
    handler NestedContainers
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncNestedContainersMapSet)(nil)

func (p *procFuncNestedContainersMapSet) Read(iprot thrift.Decoder) (thrift.Struct, thrift.Exception) {
    args := newReqNestedContainersMapSet()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncNestedContainersMapSet) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("mapSet", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncNestedContainersMapSet) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqNestedContainersMapSet)
    result := newRespNestedContainersMapSet()
    err := p.handler.MapSet(ctx, args.Foo)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing MapSet: " + err.Error(), err)
        return x, x
    }

    return result, nil
}


type procFuncNestedContainersListMap struct {
    handler NestedContainers
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncNestedContainersListMap)(nil)

func (p *procFuncNestedContainersListMap) Read(iprot thrift.Decoder) (thrift.Struct, thrift.Exception) {
    args := newReqNestedContainersListMap()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncNestedContainersListMap) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("listMap", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncNestedContainersListMap) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqNestedContainersListMap)
    result := newRespNestedContainersListMap()
    err := p.handler.ListMap(ctx, args.Foo)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing ListMap: " + err.Error(), err)
        return x, x
    }

    return result, nil
}


type procFuncNestedContainersListSet struct {
    handler NestedContainers
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncNestedContainersListSet)(nil)

func (p *procFuncNestedContainersListSet) Read(iprot thrift.Decoder) (thrift.Struct, thrift.Exception) {
    args := newReqNestedContainersListSet()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncNestedContainersListSet) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("listSet", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncNestedContainersListSet) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqNestedContainersListSet)
    result := newRespNestedContainersListSet()
    err := p.handler.ListSet(ctx, args.Foo)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing ListSet: " + err.Error(), err)
        return x, x
    }

    return result, nil
}


type procFuncNestedContainersTurtles struct {
    handler NestedContainers
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncNestedContainersTurtles)(nil)

func (p *procFuncNestedContainersTurtles) Read(iprot thrift.Decoder) (thrift.Struct, thrift.Exception) {
    args := newReqNestedContainersTurtles()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncNestedContainersTurtles) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("turtles", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncNestedContainersTurtles) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqNestedContainersTurtles)
    result := newRespNestedContainersTurtles()
    err := p.handler.Turtles(ctx, args.Foo)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing Turtles: " + err.Error(), err)
        return x, x
    }

    return result, nil
}


