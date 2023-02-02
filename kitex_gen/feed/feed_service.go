// Code generated by thriftgo (0.2.5). DO NOT EDIT.

package feed

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/bytecamp-galaxy/mini-tiktok/scripts/kitex_gen/rpcmodel"
)

type FeedRequest struct {
	LatestTime *int64 `thrift:"LatestTime,1,optional" frugal:"1,optional,i64" json:"LatestTime,omitempty"`
	Uid        *int64 `thrift:"uid,2,optional" frugal:"2,optional,i64" json:"uid,omitempty"`
}

func NewFeedRequest() *FeedRequest {
	return &FeedRequest{}
}

func (p *FeedRequest) InitDefault() {
	*p = FeedRequest{}
}

var FeedRequest_LatestTime_DEFAULT int64

func (p *FeedRequest) GetLatestTime() (v int64) {
	if !p.IsSetLatestTime() {
		return FeedRequest_LatestTime_DEFAULT
	}
	return *p.LatestTime
}

var FeedRequest_Uid_DEFAULT int64

func (p *FeedRequest) GetUid() (v int64) {
	if !p.IsSetUid() {
		return FeedRequest_Uid_DEFAULT
	}
	return *p.Uid
}
func (p *FeedRequest) SetLatestTime(val *int64) {
	p.LatestTime = val
}
func (p *FeedRequest) SetUid(val *int64) {
	p.Uid = val
}

var fieldIDToName_FeedRequest = map[int16]string{
	1: "LatestTime",
	2: "uid",
}

func (p *FeedRequest) IsSetLatestTime() bool {
	return p.LatestTime != nil
}

func (p *FeedRequest) IsSetUid() bool {
	return p.Uid != nil
}

func (p *FeedRequest) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
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

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FeedRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FeedRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.LatestTime = &v
	}
	return nil
}

func (p *FeedRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Uid = &v
	}
	return nil
}

func (p *FeedRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("FeedRequest"); err != nil {
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

func (p *FeedRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetLatestTime() {
		if err = oprot.WriteFieldBegin("LatestTime", thrift.I64, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.LatestTime); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FeedRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetUid() {
		if err = oprot.WriteFieldBegin("uid", thrift.I64, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.Uid); err != nil {
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

func (p *FeedRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeedRequest(%+v)", *p)
}

func (p *FeedRequest) DeepEqual(ano *FeedRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.LatestTime) {
		return false
	}
	if !p.Field2DeepEqual(ano.Uid) {
		return false
	}
	return true
}

func (p *FeedRequest) Field1DeepEqual(src *int64) bool {

	if p.LatestTime == src {
		return true
	} else if p.LatestTime == nil || src == nil {
		return false
	}
	if *p.LatestTime != *src {
		return false
	}
	return true
}
func (p *FeedRequest) Field2DeepEqual(src *int64) bool {

	if p.Uid == src {
		return true
	} else if p.Uid == nil || src == nil {
		return false
	}
	if *p.Uid != *src {
		return false
	}
	return true
}

type FeedResponse struct {
	VideoList []*rpcmodel.Video `thrift:"VideoList,3" frugal:"3,default,list<rpcmodel.Video>" json:"VideoList"`
	NextTime  *int64            `thrift:"NextTime,4,optional" frugal:"4,optional,i64" json:"NextTime,omitempty"`
}

func NewFeedResponse() *FeedResponse {
	return &FeedResponse{}
}

func (p *FeedResponse) InitDefault() {
	*p = FeedResponse{}
}

func (p *FeedResponse) GetVideoList() (v []*rpcmodel.Video) {
	return p.VideoList
}

var FeedResponse_NextTime_DEFAULT int64

func (p *FeedResponse) GetNextTime() (v int64) {
	if !p.IsSetNextTime() {
		return FeedResponse_NextTime_DEFAULT
	}
	return *p.NextTime
}
func (p *FeedResponse) SetVideoList(val []*rpcmodel.Video) {
	p.VideoList = val
}
func (p *FeedResponse) SetNextTime(val *int64) {
	p.NextTime = val
}

var fieldIDToName_FeedResponse = map[int16]string{
	3: "VideoList",
	4: "NextTime",
}

func (p *FeedResponse) IsSetNextTime() bool {
	return p.NextTime != nil
}

func (p *FeedResponse) Read(iprot thrift.TProtocol) (err error) {

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
		case 3:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FeedResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FeedResponse) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.VideoList = make([]*rpcmodel.Video, 0, size)
	for i := 0; i < size; i++ {
		_elem := rpcmodel.NewVideo()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.VideoList = append(p.VideoList, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *FeedResponse) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.NextTime = &v
	}
	return nil
}

func (p *FeedResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("FeedResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
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

func (p *FeedResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("VideoList", thrift.LIST, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.VideoList)); err != nil {
		return err
	}
	for _, v := range p.VideoList {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
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

func (p *FeedResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetNextTime() {
		if err = oprot.WriteFieldBegin("NextTime", thrift.I64, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.NextTime); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *FeedResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeedResponse(%+v)", *p)
}

func (p *FeedResponse) DeepEqual(ano *FeedResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field3DeepEqual(ano.VideoList) {
		return false
	}
	if !p.Field4DeepEqual(ano.NextTime) {
		return false
	}
	return true
}

func (p *FeedResponse) Field3DeepEqual(src []*rpcmodel.Video) bool {

	if len(p.VideoList) != len(src) {
		return false
	}
	for i, v := range p.VideoList {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}
func (p *FeedResponse) Field4DeepEqual(src *int64) bool {

	if p.NextTime == src {
		return true
	} else if p.NextTime == nil || src == nil {
		return false
	}
	if *p.NextTime != *src {
		return false
	}
	return true
}

type FeedService interface {
	GetFeed(ctx context.Context, req *FeedRequest) (r *FeedResponse, err error)
}

type FeedServiceClient struct {
	c thrift.TClient
}

func NewFeedServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FeedServiceClient {
	return &FeedServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewFeedServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FeedServiceClient {
	return &FeedServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewFeedServiceClient(c thrift.TClient) *FeedServiceClient {
	return &FeedServiceClient{
		c: c,
	}
}

func (p *FeedServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *FeedServiceClient) GetFeed(ctx context.Context, req *FeedRequest) (r *FeedResponse, err error) {
	var _args FeedServiceGetFeedArgs
	_args.Req = req
	var _result FeedServiceGetFeedResult
	if err = p.Client_().Call(ctx, "getFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type FeedServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      FeedService
}

func (p *FeedServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FeedServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FeedServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFeedServiceProcessor(handler FeedService) *FeedServiceProcessor {
	self := &FeedServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("getFeed", &feedServiceProcessorGetFeed{handler: handler})
	return self
}
func (p *FeedServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type feedServiceProcessorGetFeed struct {
	handler FeedService
}

func (p *feedServiceProcessorGetFeed) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FeedServiceGetFeedArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getFeed", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FeedServiceGetFeedResult{}
	var retval *FeedResponse
	if retval, err2 = p.handler.GetFeed(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getFeed: "+err2.Error())
		oprot.WriteMessageBegin("getFeed", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getFeed", thrift.REPLY, seqId); err2 != nil {
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

type FeedServiceGetFeedArgs struct {
	Req *FeedRequest `thrift:"req,1" frugal:"1,default,FeedRequest" json:"req"`
}

func NewFeedServiceGetFeedArgs() *FeedServiceGetFeedArgs {
	return &FeedServiceGetFeedArgs{}
}

func (p *FeedServiceGetFeedArgs) InitDefault() {
	*p = FeedServiceGetFeedArgs{}
}

var FeedServiceGetFeedArgs_Req_DEFAULT *FeedRequest

func (p *FeedServiceGetFeedArgs) GetReq() (v *FeedRequest) {
	if !p.IsSetReq() {
		return FeedServiceGetFeedArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *FeedServiceGetFeedArgs) SetReq(val *FeedRequest) {
	p.Req = val
}

var fieldIDToName_FeedServiceGetFeedArgs = map[int16]string{
	1: "req",
}

func (p *FeedServiceGetFeedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FeedServiceGetFeedArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FeedServiceGetFeedArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FeedServiceGetFeedArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewFeedRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FeedServiceGetFeedArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("getFeed_args"); err != nil {
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

func (p *FeedServiceGetFeedArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *FeedServiceGetFeedArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeedServiceGetFeedArgs(%+v)", *p)
}

func (p *FeedServiceGetFeedArgs) DeepEqual(ano *FeedServiceGetFeedArgs) bool {
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

func (p *FeedServiceGetFeedArgs) Field1DeepEqual(src *FeedRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type FeedServiceGetFeedResult struct {
	Success *FeedResponse `thrift:"success,0,optional" frugal:"0,optional,FeedResponse" json:"success,omitempty"`
}

func NewFeedServiceGetFeedResult() *FeedServiceGetFeedResult {
	return &FeedServiceGetFeedResult{}
}

func (p *FeedServiceGetFeedResult) InitDefault() {
	*p = FeedServiceGetFeedResult{}
}

var FeedServiceGetFeedResult_Success_DEFAULT *FeedResponse

func (p *FeedServiceGetFeedResult) GetSuccess() (v *FeedResponse) {
	if !p.IsSetSuccess() {
		return FeedServiceGetFeedResult_Success_DEFAULT
	}
	return p.Success
}
func (p *FeedServiceGetFeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*FeedResponse)
}

var fieldIDToName_FeedServiceGetFeedResult = map[int16]string{
	0: "success",
}

func (p *FeedServiceGetFeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FeedServiceGetFeedResult) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FeedServiceGetFeedResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FeedServiceGetFeedResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewFeedResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FeedServiceGetFeedResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("getFeed_result"); err != nil {
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

func (p *FeedServiceGetFeedResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *FeedServiceGetFeedResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeedServiceGetFeedResult(%+v)", *p)
}

func (p *FeedServiceGetFeedResult) DeepEqual(ano *FeedServiceGetFeedResult) bool {
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

func (p *FeedServiceGetFeedResult) Field0DeepEqual(src *FeedResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
