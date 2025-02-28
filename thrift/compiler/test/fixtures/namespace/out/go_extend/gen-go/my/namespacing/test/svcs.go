// Autogenerated by Thrift for hsmodule.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package test


import (
    "context"
    "fmt"
    "io"
    "reflect"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = context.Background
var _ = fmt.Printf
var _ = io.EOF
var _ = reflect.Ptr
var _ = thrift.VOID
var _ = metadata.GoUnusedProtection__

type HsTestService interface {
    Init(ctx context.Context, int1 int64) (int64, error)
}

type HsTestServiceClientInterface interface {
    io.Closer
    HsTestService
}

type HsTestServiceClient struct {
    ch thrift.RequestChannel
}
// Compile time interface enforcer
var _ HsTestServiceClientInterface = (*HsTestServiceClient)(nil)

func NewHsTestServiceChannelClient(channel thrift.RequestChannel) *HsTestServiceClient {
    return &HsTestServiceClient{
        ch: channel,
    }
}

func NewHsTestServiceClient(prot thrift.DO_NOT_USE_ChannelWrapper) *HsTestServiceClient {
    var channel thrift.RequestChannel
    if prot != nil {
        channel = prot.DO_NOT_USE_WrapChannel()
    }
    return NewHsTestServiceChannelClient(channel)
}

func (c *HsTestServiceClient) Close() error {
    return c.ch.Close()
}

func (c *HsTestServiceClient) Init(ctx context.Context, int1 int64) (int64, error) {
    in := &reqHsTestServiceInit{
        Int1: int1,
    }
    out := newRespHsTestServiceInit()
    err := c.ch.Call(ctx, "init", in, out)
    if err != nil {
        return 0, err
    }
    return out.GetSuccess(), nil
}


type HsTestServiceProcessor struct {
    processorFunctionMap map[string]thrift.ProcessorFunction
    functionServiceMap   map[string]string
    handler              HsTestService
}

func NewHsTestServiceProcessor(handler HsTestService) *HsTestServiceProcessor {
    p := &HsTestServiceProcessor{
        handler:              handler,
        processorFunctionMap: make(map[string]thrift.ProcessorFunction),
        functionServiceMap:   make(map[string]string),
    }
    p.AddToProcessorFunctionMap("init", &procFuncHsTestServiceInit{handler: handler})
    p.AddToFunctionServiceMap("init", "HsTestService")

    return p
}

func (p *HsTestServiceProcessor) AddToProcessorFunctionMap(key string, processorFunction thrift.ProcessorFunction) {
    p.processorFunctionMap[key] = processorFunction
}

func (p *HsTestServiceProcessor) AddToFunctionServiceMap(key, service string) {
    p.functionServiceMap[key] = service
}

func (p *HsTestServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction) {
    return p.processorFunctionMap[key]
}

func (p *HsTestServiceProcessor) ProcessorFunctionMap() map[string]thrift.ProcessorFunction {
    return p.processorFunctionMap
}

func (p *HsTestServiceProcessor) FunctionServiceMap() map[string]string {
    return p.functionServiceMap
}

func (p *HsTestServiceProcessor) PackageName() string {
    return "test"
}

func (p *HsTestServiceProcessor) GetThriftMetadata() *metadata.ThriftMetadata {
    return GetThriftMetadataForService("hsmodule.HsTestService")
}


type procFuncHsTestServiceInit struct {
    handler HsTestService
}
// Compile time interface enforcer
var _ thrift.ProcessorFunction = (*procFuncHsTestServiceInit)(nil)

func (p *procFuncHsTestServiceInit) Read(iprot thrift.Decoder) (thrift.Struct, error) {
    args := newReqHsTestServiceInit()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncHsTestServiceInit) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Encoder) (err error) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("init", messageType, seqId); err2 != nil {
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

func (p *procFuncHsTestServiceInit) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqHsTestServiceInit)
    result := newRespHsTestServiceInit()
    retval, err := p.handler.Init(ctx, args.Int1)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing Init: " + err.Error(), err)
        return x, x
    }

    result.Success = &retval
    return result, nil
}


