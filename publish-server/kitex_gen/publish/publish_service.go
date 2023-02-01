// Code generated by thriftgo (0.2.5). DO NOT EDIT.

package publish

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type PublishRequest struct {
	Uid   int64  `thrift:"uid,1,required" frugal:"1,required,i64" json:"uid"`
	Title string `thrift:"Title,2,required" frugal:"2,required,string" json:"Title"`
	Data  []byte `thrift:"data,3,required" frugal:"3,required,binary" json:"data"`
}

func NewPublishRequest() *PublishRequest {
	return &PublishRequest{}
}

func (p *PublishRequest) InitDefault() {
	*p = PublishRequest{}
}

func (p *PublishRequest) GetUid() (v int64) {
	return p.Uid
}

func (p *PublishRequest) GetTitle() (v string) {
	return p.Title
}

func (p *PublishRequest) GetData() (v []byte) {
	return p.Data
}
func (p *PublishRequest) SetUid(val int64) {
	p.Uid = val
}
func (p *PublishRequest) SetTitle(val string) {
	p.Title = val
}
func (p *PublishRequest) SetData(val []byte) {
	p.Data = val
}

var fieldIDToName_PublishRequest = map[int16]string{
	1: "uid",
	2: "Title",
	3: "data",
}

func (p *PublishRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUid bool = false
	var issetTitle bool = false
	var issetData bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetUid = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetTitle = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetData = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetUid {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetTitle {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetData {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PublishRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_PublishRequest[fieldId]))
}

func (p *PublishRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Uid = v
	}
	return nil
}

func (p *PublishRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Title = v
	}
	return nil
}

func (p *PublishRequest) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return err
	} else {
		p.Data = []byte(v)
	}
	return nil
}

func (p *PublishRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PublishRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PublishRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("uid", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Uid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PublishRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Title", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Title); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *PublishRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBinary([]byte(p.Data)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *PublishRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PublishRequest(%+v)", *p)
}

func (p *PublishRequest) DeepEqual(ano *PublishRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Uid) {
		return false
	}
	if !p.Field2DeepEqual(ano.Title) {
		return false
	}
	if !p.Field3DeepEqual(ano.Data) {
		return false
	}
	return true
}

func (p *PublishRequest) Field1DeepEqual(src int64) bool {

	if p.Uid != src {
		return false
	}
	return true
}
func (p *PublishRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Title, src) != 0 {
		return false
	}
	return true
}
func (p *PublishRequest) Field3DeepEqual(src []byte) bool {

	if bytes.Compare(p.Data, src) != 0 {
		return false
	}
	return true
}

type PublishResponse struct {
	StatusCode int32   `thrift:"StatusCode,1,required" frugal:"1,required,i32" json:"StatusCode"`
	StatusMsg  *string `thrift:"StatusMsg,2,optional" frugal:"2,optional,string" json:"StatusMsg,omitempty"`
}

func NewPublishResponse() *PublishResponse {
	return &PublishResponse{}
}

func (p *PublishResponse) InitDefault() {
	*p = PublishResponse{}
}

func (p *PublishResponse) GetStatusCode() (v int32) {
	return p.StatusCode
}

var PublishResponse_StatusMsg_DEFAULT string

func (p *PublishResponse) GetStatusMsg() (v string) {
	if !p.IsSetStatusMsg() {
		return PublishResponse_StatusMsg_DEFAULT
	}
	return *p.StatusMsg
}
func (p *PublishResponse) SetStatusCode(val int32) {
	p.StatusCode = val
}
func (p *PublishResponse) SetStatusMsg(val *string) {
	p.StatusMsg = val
}

var fieldIDToName_PublishResponse = map[int16]string{
	1: "StatusCode",
	2: "StatusMsg",
}

func (p *PublishResponse) IsSetStatusMsg() bool {
	return p.StatusMsg != nil
}

func (p *PublishResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetStatusCode bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetStatusCode = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetStatusCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PublishResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_PublishResponse[fieldId]))
}

func (p *PublishResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *PublishResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = &v
	}
	return nil
}

func (p *PublishResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PublishResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PublishResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusCode", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PublishResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatusMsg() {
		if err = oprot.WriteFieldBegin("StatusMsg", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.StatusMsg); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *PublishResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PublishResponse(%+v)", *p)
}

func (p *PublishResponse) DeepEqual(ano *PublishResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.StatusCode) {
		return false
	}
	if !p.Field2DeepEqual(ano.StatusMsg) {
		return false
	}
	return true
}

func (p *PublishResponse) Field1DeepEqual(src int32) bool {

	if p.StatusCode != src {
		return false
	}
	return true
}
func (p *PublishResponse) Field2DeepEqual(src *string) bool {

	if p.StatusMsg == src {
		return true
	} else if p.StatusMsg == nil || src == nil {
		return false
	}
	if strings.Compare(*p.StatusMsg, *src) != 0 {
		return false
	}
	return true
}

type PublishService interface {
	PublishVideo(ctx context.Context, req *PublishRequest) (r *PublishResponse, err error)
}

type PublishServiceClient struct {
	c thrift.TClient
}

func NewPublishServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PublishServiceClient {
	return &PublishServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPublishServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PublishServiceClient {
	return &PublishServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPublishServiceClient(c thrift.TClient) *PublishServiceClient {
	return &PublishServiceClient{
		c: c,
	}
}

func (p *PublishServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PublishServiceClient) PublishVideo(ctx context.Context, req *PublishRequest) (r *PublishResponse, err error) {
	var _args PublishServicePublishVideoArgs
	_args.Req = req
	var _result PublishServicePublishVideoResult
	if err = p.Client_().Call(ctx, "publishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type PublishServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PublishService
}

func (p *PublishServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PublishServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PublishServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPublishServiceProcessor(handler PublishService) *PublishServiceProcessor {
	self := &PublishServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("publishVideo", &publishServiceProcessorPublishVideo{handler: handler})
	return self
}
func (p *PublishServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type publishServiceProcessorPublishVideo struct {
	handler PublishService
}

func (p *publishServiceProcessorPublishVideo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PublishServicePublishVideoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("publishVideo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := PublishServicePublishVideoResult{}
	var retval *PublishResponse
	if retval, err2 = p.handler.PublishVideo(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing publishVideo: "+err2.Error())
		oprot.WriteMessageBegin("publishVideo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("publishVideo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type PublishServicePublishVideoArgs struct {
	Req *PublishRequest `thrift:"req,1" frugal:"1,default,PublishRequest" json:"req"`
}

func NewPublishServicePublishVideoArgs() *PublishServicePublishVideoArgs {
	return &PublishServicePublishVideoArgs{}
}

func (p *PublishServicePublishVideoArgs) InitDefault() {
	*p = PublishServicePublishVideoArgs{}
}

var PublishServicePublishVideoArgs_Req_DEFAULT *PublishRequest

func (p *PublishServicePublishVideoArgs) GetReq() (v *PublishRequest) {
	if !p.IsSetReq() {
		return PublishServicePublishVideoArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *PublishServicePublishVideoArgs) SetReq(val *PublishRequest) {
	p.Req = val
}

var fieldIDToName_PublishServicePublishVideoArgs = map[int16]string{
	1: "req",
}

func (p *PublishServicePublishVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PublishServicePublishVideoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PublishServicePublishVideoArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PublishServicePublishVideoArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewPublishRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *PublishServicePublishVideoArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("publishVideo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PublishServicePublishVideoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PublishServicePublishVideoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PublishServicePublishVideoArgs(%+v)", *p)
}

func (p *PublishServicePublishVideoArgs) DeepEqual(ano *PublishServicePublishVideoArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *PublishServicePublishVideoArgs) Field1DeepEqual(src *PublishRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type PublishServicePublishVideoResult struct {
	Success *PublishResponse `thrift:"success,0,optional" frugal:"0,optional,PublishResponse" json:"success,omitempty"`
}

func NewPublishServicePublishVideoResult() *PublishServicePublishVideoResult {
	return &PublishServicePublishVideoResult{}
}

func (p *PublishServicePublishVideoResult) InitDefault() {
	*p = PublishServicePublishVideoResult{}
}

var PublishServicePublishVideoResult_Success_DEFAULT *PublishResponse

func (p *PublishServicePublishVideoResult) GetSuccess() (v *PublishResponse) {
	if !p.IsSetSuccess() {
		return PublishServicePublishVideoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PublishServicePublishVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*PublishResponse)
}

var fieldIDToName_PublishServicePublishVideoResult = map[int16]string{
	0: "success",
}

func (p *PublishServicePublishVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PublishServicePublishVideoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PublishServicePublishVideoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PublishServicePublishVideoResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewPublishResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *PublishServicePublishVideoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("publishVideo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PublishServicePublishVideoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *PublishServicePublishVideoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PublishServicePublishVideoResult(%+v)", *p)
}

func (p *PublishServicePublishVideoResult) DeepEqual(ano *PublishServicePublishVideoResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PublishServicePublishVideoResult) Field0DeepEqual(src *PublishResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}