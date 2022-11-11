// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/Cmd.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PCommandType int32

const (
	PCommandType_NONE                     PCommandType = 0
	PCommandType_PING                     PCommandType = 100
	PCommandType_PONG                     PCommandType = 101
	PCommandType_ECHO                     PCommandType = 710
	PCommandType_ACTIVE_THREAD_COUNT      PCommandType = 730
	PCommandType_ACTIVE_THREAD_DUMP       PCommandType = 740
	PCommandType_ACTIVE_THREAD_LIGHT_DUMP PCommandType = 750
)

var PCommandType_name = map[int32]string{
	0:   "NONE",
	100: "PING",
	101: "PONG",
	710: "ECHO",
	730: "ACTIVE_THREAD_COUNT",
	740: "ACTIVE_THREAD_DUMP",
	750: "ACTIVE_THREAD_LIGHT_DUMP",
}

var PCommandType_value = map[string]int32{
	"NONE":                     0,
	"PING":                     100,
	"PONG":                     101,
	"ECHO":                     710,
	"ACTIVE_THREAD_COUNT":      730,
	"ACTIVE_THREAD_DUMP":       740,
	"ACTIVE_THREAD_LIGHT_DUMP": 750,
}

func (x PCommandType) String() string {
	return proto.EnumName(PCommandType_name, int32(x))
}

func (PCommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{0}
}

type PCmdMessage struct {
	// Types that are valid to be assigned to Message:
	//	*PCmdMessage_HandshakeMessage
	//	*PCmdMessage_FailMessage
	Message              isPCmdMessage_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PCmdMessage) Reset()         { *m = PCmdMessage{} }
func (m *PCmdMessage) String() string { return proto.CompactTextString(m) }
func (*PCmdMessage) ProtoMessage()    {}
func (*PCmdMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{0}
}

func (m *PCmdMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdMessage.Unmarshal(m, b)
}
func (m *PCmdMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdMessage.Marshal(b, m, deterministic)
}
func (m *PCmdMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdMessage.Merge(m, src)
}
func (m *PCmdMessage) XXX_Size() int {
	return xxx_messageInfo_PCmdMessage.Size(m)
}
func (m *PCmdMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdMessage proto.InternalMessageInfo

type isPCmdMessage_Message interface {
	isPCmdMessage_Message()
}

type PCmdMessage_HandshakeMessage struct {
	HandshakeMessage *PCmdServiceHandshake `protobuf:"bytes,1,opt,name=handshakeMessage,proto3,oneof"`
}

type PCmdMessage_FailMessage struct {
	FailMessage *PCmdResponse `protobuf:"bytes,2,opt,name=failMessage,proto3,oneof"`
}

func (*PCmdMessage_HandshakeMessage) isPCmdMessage_Message() {}

func (*PCmdMessage_FailMessage) isPCmdMessage_Message() {}

func (m *PCmdMessage) GetMessage() isPCmdMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PCmdMessage) GetHandshakeMessage() *PCmdServiceHandshake {
	if x, ok := m.GetMessage().(*PCmdMessage_HandshakeMessage); ok {
		return x.HandshakeMessage
	}
	return nil
}

func (m *PCmdMessage) GetFailMessage() *PCmdResponse {
	if x, ok := m.GetMessage().(*PCmdMessage_FailMessage); ok {
		return x.FailMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PCmdMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PCmdMessage_HandshakeMessage)(nil),
		(*PCmdMessage_FailMessage)(nil),
	}
}

type PCmdServiceHandshake struct {
	// initial message
	SupportCommandServiceKey []int32  `protobuf:"varint,1,rep,packed,name=supportCommandServiceKey,proto3" json:"supportCommandServiceKey,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *PCmdServiceHandshake) Reset()         { *m = PCmdServiceHandshake{} }
func (m *PCmdServiceHandshake) String() string { return proto.CompactTextString(m) }
func (*PCmdServiceHandshake) ProtoMessage()    {}
func (*PCmdServiceHandshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{1}
}

func (m *PCmdServiceHandshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdServiceHandshake.Unmarshal(m, b)
}
func (m *PCmdServiceHandshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdServiceHandshake.Marshal(b, m, deterministic)
}
func (m *PCmdServiceHandshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdServiceHandshake.Merge(m, src)
}
func (m *PCmdServiceHandshake) XXX_Size() int {
	return xxx_messageInfo_PCmdServiceHandshake.Size(m)
}
func (m *PCmdServiceHandshake) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdServiceHandshake.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdServiceHandshake proto.InternalMessageInfo

func (m *PCmdServiceHandshake) GetSupportCommandServiceKey() []int32 {
	if m != nil {
		return m.SupportCommandServiceKey
	}
	return nil
}

type PCmdResponse struct {
	ResponseId           int32                   `protobuf:"varint,1,opt,name=responseId,proto3" json:"responseId,omitempty"`
	Status               int32                   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PCmdResponse) Reset()         { *m = PCmdResponse{} }
func (m *PCmdResponse) String() string { return proto.CompactTextString(m) }
func (*PCmdResponse) ProtoMessage()    {}
func (*PCmdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{2}
}

func (m *PCmdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdResponse.Unmarshal(m, b)
}
func (m *PCmdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdResponse.Marshal(b, m, deterministic)
}
func (m *PCmdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdResponse.Merge(m, src)
}
func (m *PCmdResponse) XXX_Size() int {
	return xxx_messageInfo_PCmdResponse.Size(m)
}
func (m *PCmdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdResponse proto.InternalMessageInfo

func (m *PCmdResponse) GetResponseId() int32 {
	if m != nil {
		return m.ResponseId
	}
	return 0
}

func (m *PCmdResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PCmdResponse) GetMessage() *wrapperspb.StringValue {
	if m != nil {
		return m.Message
	}
	return nil
}

type PCmdStreamResponse struct {
	ResponseId           int32                   `protobuf:"varint,1,opt,name=responseId,proto3" json:"responseId,omitempty"`
	SequenceId           int32                   `protobuf:"varint,2,opt,name=sequenceId,proto3" json:"sequenceId,omitempty"`
	Message              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PCmdStreamResponse) Reset()         { *m = PCmdStreamResponse{} }
func (m *PCmdStreamResponse) String() string { return proto.CompactTextString(m) }
func (*PCmdStreamResponse) ProtoMessage()    {}
func (*PCmdStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{3}
}

func (m *PCmdStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdStreamResponse.Unmarshal(m, b)
}
func (m *PCmdStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdStreamResponse.Marshal(b, m, deterministic)
}
func (m *PCmdStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdStreamResponse.Merge(m, src)
}
func (m *PCmdStreamResponse) XXX_Size() int {
	return xxx_messageInfo_PCmdStreamResponse.Size(m)
}
func (m *PCmdStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdStreamResponse proto.InternalMessageInfo

func (m *PCmdStreamResponse) GetResponseId() int32 {
	if m != nil {
		return m.ResponseId
	}
	return 0
}

func (m *PCmdStreamResponse) GetSequenceId() int32 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *PCmdStreamResponse) GetMessage() *wrapperspb.StringValue {
	if m != nil {
		return m.Message
	}
	return nil
}

type PCmdRequest struct {
	RequestId int32 `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	// Types that are valid to be assigned to Command:
	//	*PCmdRequest_CommandEcho
	//	*PCmdRequest_CommandActiveThreadCount
	//	*PCmdRequest_CommandActiveThreadDump
	//	*PCmdRequest_CommandActiveThreadLightDump
	Command              isPCmdRequest_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PCmdRequest) Reset()         { *m = PCmdRequest{} }
func (m *PCmdRequest) String() string { return proto.CompactTextString(m) }
func (*PCmdRequest) ProtoMessage()    {}
func (*PCmdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{4}
}

func (m *PCmdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdRequest.Unmarshal(m, b)
}
func (m *PCmdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdRequest.Marshal(b, m, deterministic)
}
func (m *PCmdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdRequest.Merge(m, src)
}
func (m *PCmdRequest) XXX_Size() int {
	return xxx_messageInfo_PCmdRequest.Size(m)
}
func (m *PCmdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdRequest proto.InternalMessageInfo

func (m *PCmdRequest) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type isPCmdRequest_Command interface {
	isPCmdRequest_Command()
}

type PCmdRequest_CommandEcho struct {
	CommandEcho *PCmdEcho `protobuf:"bytes,710,opt,name=commandEcho,proto3,oneof"`
}

type PCmdRequest_CommandActiveThreadCount struct {
	CommandActiveThreadCount *PCmdActiveThreadCount `protobuf:"bytes,730,opt,name=commandActiveThreadCount,proto3,oneof"`
}

type PCmdRequest_CommandActiveThreadDump struct {
	CommandActiveThreadDump *PCmdActiveThreadDump `protobuf:"bytes,740,opt,name=commandActiveThreadDump,proto3,oneof"`
}

type PCmdRequest_CommandActiveThreadLightDump struct {
	CommandActiveThreadLightDump *PCmdActiveThreadLightDump `protobuf:"bytes,750,opt,name=commandActiveThreadLightDump,proto3,oneof"`
}

func (*PCmdRequest_CommandEcho) isPCmdRequest_Command() {}

func (*PCmdRequest_CommandActiveThreadCount) isPCmdRequest_Command() {}

func (*PCmdRequest_CommandActiveThreadDump) isPCmdRequest_Command() {}

func (*PCmdRequest_CommandActiveThreadLightDump) isPCmdRequest_Command() {}

func (m *PCmdRequest) GetCommand() isPCmdRequest_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *PCmdRequest) GetCommandEcho() *PCmdEcho {
	if x, ok := m.GetCommand().(*PCmdRequest_CommandEcho); ok {
		return x.CommandEcho
	}
	return nil
}

func (m *PCmdRequest) GetCommandActiveThreadCount() *PCmdActiveThreadCount {
	if x, ok := m.GetCommand().(*PCmdRequest_CommandActiveThreadCount); ok {
		return x.CommandActiveThreadCount
	}
	return nil
}

func (m *PCmdRequest) GetCommandActiveThreadDump() *PCmdActiveThreadDump {
	if x, ok := m.GetCommand().(*PCmdRequest_CommandActiveThreadDump); ok {
		return x.CommandActiveThreadDump
	}
	return nil
}

func (m *PCmdRequest) GetCommandActiveThreadLightDump() *PCmdActiveThreadLightDump {
	if x, ok := m.GetCommand().(*PCmdRequest_CommandActiveThreadLightDump); ok {
		return x.CommandActiveThreadLightDump
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PCmdRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PCmdRequest_CommandEcho)(nil),
		(*PCmdRequest_CommandActiveThreadCount)(nil),
		(*PCmdRequest_CommandActiveThreadDump)(nil),
		(*PCmdRequest_CommandActiveThreadLightDump)(nil),
	}
}

type PCmdEcho struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PCmdEcho) Reset()         { *m = PCmdEcho{} }
func (m *PCmdEcho) String() string { return proto.CompactTextString(m) }
func (*PCmdEcho) ProtoMessage()    {}
func (*PCmdEcho) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{5}
}

func (m *PCmdEcho) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdEcho.Unmarshal(m, b)
}
func (m *PCmdEcho) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdEcho.Marshal(b, m, deterministic)
}
func (m *PCmdEcho) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdEcho.Merge(m, src)
}
func (m *PCmdEcho) XXX_Size() int {
	return xxx_messageInfo_PCmdEcho.Size(m)
}
func (m *PCmdEcho) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdEcho.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdEcho proto.InternalMessageInfo

func (m *PCmdEcho) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PCmdEchoResponse struct {
	CommonResponse       *PCmdResponse `protobuf:"bytes,1,opt,name=commonResponse,proto3" json:"commonResponse,omitempty"`
	Message              string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PCmdEchoResponse) Reset()         { *m = PCmdEchoResponse{} }
func (m *PCmdEchoResponse) String() string { return proto.CompactTextString(m) }
func (*PCmdEchoResponse) ProtoMessage()    {}
func (*PCmdEchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{6}
}

func (m *PCmdEchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdEchoResponse.Unmarshal(m, b)
}
func (m *PCmdEchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdEchoResponse.Marshal(b, m, deterministic)
}
func (m *PCmdEchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdEchoResponse.Merge(m, src)
}
func (m *PCmdEchoResponse) XXX_Size() int {
	return xxx_messageInfo_PCmdEchoResponse.Size(m)
}
func (m *PCmdEchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdEchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdEchoResponse proto.InternalMessageInfo

func (m *PCmdEchoResponse) GetCommonResponse() *PCmdResponse {
	if m != nil {
		return m.CommonResponse
	}
	return nil
}

func (m *PCmdEchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PCmdActiveThreadDump struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	ThreadName           []string `protobuf:"bytes,2,rep,name=threadName,proto3" json:"threadName,omitempty"`
	LocalTraceId         []int64  `protobuf:"varint,3,rep,packed,name=localTraceId,proto3" json:"localTraceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PCmdActiveThreadDump) Reset()         { *m = PCmdActiveThreadDump{} }
func (m *PCmdActiveThreadDump) String() string { return proto.CompactTextString(m) }
func (*PCmdActiveThreadDump) ProtoMessage()    {}
func (*PCmdActiveThreadDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{7}
}

func (m *PCmdActiveThreadDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdActiveThreadDump.Unmarshal(m, b)
}
func (m *PCmdActiveThreadDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdActiveThreadDump.Marshal(b, m, deterministic)
}
func (m *PCmdActiveThreadDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdActiveThreadDump.Merge(m, src)
}
func (m *PCmdActiveThreadDump) XXX_Size() int {
	return xxx_messageInfo_PCmdActiveThreadDump.Size(m)
}
func (m *PCmdActiveThreadDump) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdActiveThreadDump.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdActiveThreadDump proto.InternalMessageInfo

func (m *PCmdActiveThreadDump) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PCmdActiveThreadDump) GetThreadName() []string {
	if m != nil {
		return m.ThreadName
	}
	return nil
}

func (m *PCmdActiveThreadDump) GetLocalTraceId() []int64 {
	if m != nil {
		return m.LocalTraceId
	}
	return nil
}

type PCmdActiveThreadDumpRes struct {
	CommonResponse       *PCmdResponse        `protobuf:"bytes,1,opt,name=commonResponse,proto3" json:"commonResponse,omitempty"`
	ThreadDump           []*PActiveThreadDump `protobuf:"bytes,2,rep,name=threadDump,proto3" json:"threadDump,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	SubType              string               `protobuf:"bytes,4,opt,name=subType,proto3" json:"subType,omitempty"`
	Version              string               `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PCmdActiveThreadDumpRes) Reset()         { *m = PCmdActiveThreadDumpRes{} }
func (m *PCmdActiveThreadDumpRes) String() string { return proto.CompactTextString(m) }
func (*PCmdActiveThreadDumpRes) ProtoMessage()    {}
func (*PCmdActiveThreadDumpRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{8}
}

func (m *PCmdActiveThreadDumpRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdActiveThreadDumpRes.Unmarshal(m, b)
}
func (m *PCmdActiveThreadDumpRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdActiveThreadDumpRes.Marshal(b, m, deterministic)
}
func (m *PCmdActiveThreadDumpRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdActiveThreadDumpRes.Merge(m, src)
}
func (m *PCmdActiveThreadDumpRes) XXX_Size() int {
	return xxx_messageInfo_PCmdActiveThreadDumpRes.Size(m)
}
func (m *PCmdActiveThreadDumpRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdActiveThreadDumpRes.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdActiveThreadDumpRes proto.InternalMessageInfo

func (m *PCmdActiveThreadDumpRes) GetCommonResponse() *PCmdResponse {
	if m != nil {
		return m.CommonResponse
	}
	return nil
}

func (m *PCmdActiveThreadDumpRes) GetThreadDump() []*PActiveThreadDump {
	if m != nil {
		return m.ThreadDump
	}
	return nil
}

func (m *PCmdActiveThreadDumpRes) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PCmdActiveThreadDumpRes) GetSubType() string {
	if m != nil {
		return m.SubType
	}
	return ""
}

func (m *PCmdActiveThreadDumpRes) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PCmdActiveThreadLightDump struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	ThreadName           []string `protobuf:"bytes,2,rep,name=threadName,proto3" json:"threadName,omitempty"`
	LocalTraceId         []int64  `protobuf:"varint,3,rep,packed,name=localTraceId,proto3" json:"localTraceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PCmdActiveThreadLightDump) Reset()         { *m = PCmdActiveThreadLightDump{} }
func (m *PCmdActiveThreadLightDump) String() string { return proto.CompactTextString(m) }
func (*PCmdActiveThreadLightDump) ProtoMessage()    {}
func (*PCmdActiveThreadLightDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{9}
}

func (m *PCmdActiveThreadLightDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdActiveThreadLightDump.Unmarshal(m, b)
}
func (m *PCmdActiveThreadLightDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdActiveThreadLightDump.Marshal(b, m, deterministic)
}
func (m *PCmdActiveThreadLightDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdActiveThreadLightDump.Merge(m, src)
}
func (m *PCmdActiveThreadLightDump) XXX_Size() int {
	return xxx_messageInfo_PCmdActiveThreadLightDump.Size(m)
}
func (m *PCmdActiveThreadLightDump) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdActiveThreadLightDump.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdActiveThreadLightDump proto.InternalMessageInfo

func (m *PCmdActiveThreadLightDump) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PCmdActiveThreadLightDump) GetThreadName() []string {
	if m != nil {
		return m.ThreadName
	}
	return nil
}

func (m *PCmdActiveThreadLightDump) GetLocalTraceId() []int64 {
	if m != nil {
		return m.LocalTraceId
	}
	return nil
}

type PCmdActiveThreadLightDumpRes struct {
	CommonResponse       *PCmdResponse             `protobuf:"bytes,1,opt,name=commonResponse,proto3" json:"commonResponse,omitempty"`
	ThreadDump           []*PActiveThreadLightDump `protobuf:"bytes,2,rep,name=threadDump,proto3" json:"threadDump,omitempty"`
	Type                 string                    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	SubType              string                    `protobuf:"bytes,4,opt,name=subType,proto3" json:"subType,omitempty"`
	Version              string                    `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PCmdActiveThreadLightDumpRes) Reset()         { *m = PCmdActiveThreadLightDumpRes{} }
func (m *PCmdActiveThreadLightDumpRes) String() string { return proto.CompactTextString(m) }
func (*PCmdActiveThreadLightDumpRes) ProtoMessage()    {}
func (*PCmdActiveThreadLightDumpRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{10}
}

func (m *PCmdActiveThreadLightDumpRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdActiveThreadLightDumpRes.Unmarshal(m, b)
}
func (m *PCmdActiveThreadLightDumpRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdActiveThreadLightDumpRes.Marshal(b, m, deterministic)
}
func (m *PCmdActiveThreadLightDumpRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdActiveThreadLightDumpRes.Merge(m, src)
}
func (m *PCmdActiveThreadLightDumpRes) XXX_Size() int {
	return xxx_messageInfo_PCmdActiveThreadLightDumpRes.Size(m)
}
func (m *PCmdActiveThreadLightDumpRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdActiveThreadLightDumpRes.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdActiveThreadLightDumpRes proto.InternalMessageInfo

func (m *PCmdActiveThreadLightDumpRes) GetCommonResponse() *PCmdResponse {
	if m != nil {
		return m.CommonResponse
	}
	return nil
}

func (m *PCmdActiveThreadLightDumpRes) GetThreadDump() []*PActiveThreadLightDump {
	if m != nil {
		return m.ThreadDump
	}
	return nil
}

func (m *PCmdActiveThreadLightDumpRes) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PCmdActiveThreadLightDumpRes) GetSubType() string {
	if m != nil {
		return m.SubType
	}
	return ""
}

func (m *PCmdActiveThreadLightDumpRes) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PCmdActiveThreadCount struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PCmdActiveThreadCount) Reset()         { *m = PCmdActiveThreadCount{} }
func (m *PCmdActiveThreadCount) String() string { return proto.CompactTextString(m) }
func (*PCmdActiveThreadCount) ProtoMessage()    {}
func (*PCmdActiveThreadCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{11}
}

func (m *PCmdActiveThreadCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdActiveThreadCount.Unmarshal(m, b)
}
func (m *PCmdActiveThreadCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdActiveThreadCount.Marshal(b, m, deterministic)
}
func (m *PCmdActiveThreadCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdActiveThreadCount.Merge(m, src)
}
func (m *PCmdActiveThreadCount) XXX_Size() int {
	return xxx_messageInfo_PCmdActiveThreadCount.Size(m)
}
func (m *PCmdActiveThreadCount) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdActiveThreadCount.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdActiveThreadCount proto.InternalMessageInfo

type PCmdActiveThreadCountRes struct {
	CommonStreamResponse *PCmdStreamResponse `protobuf:"bytes,1,opt,name=commonStreamResponse,proto3" json:"commonStreamResponse,omitempty"`
	HistogramSchemaType  int32               `protobuf:"varint,2,opt,name=histogramSchemaType,proto3" json:"histogramSchemaType,omitempty"`
	ActiveThreadCount    []int32             `protobuf:"varint,3,rep,packed,name=activeThreadCount,proto3" json:"activeThreadCount,omitempty"`
	TimeStamp            int64               `protobuf:"varint,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PCmdActiveThreadCountRes) Reset()         { *m = PCmdActiveThreadCountRes{} }
func (m *PCmdActiveThreadCountRes) String() string { return proto.CompactTextString(m) }
func (*PCmdActiveThreadCountRes) ProtoMessage()    {}
func (*PCmdActiveThreadCountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aca7d559118759c, []int{12}
}

func (m *PCmdActiveThreadCountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCmdActiveThreadCountRes.Unmarshal(m, b)
}
func (m *PCmdActiveThreadCountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCmdActiveThreadCountRes.Marshal(b, m, deterministic)
}
func (m *PCmdActiveThreadCountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCmdActiveThreadCountRes.Merge(m, src)
}
func (m *PCmdActiveThreadCountRes) XXX_Size() int {
	return xxx_messageInfo_PCmdActiveThreadCountRes.Size(m)
}
func (m *PCmdActiveThreadCountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PCmdActiveThreadCountRes.DiscardUnknown(m)
}

var xxx_messageInfo_PCmdActiveThreadCountRes proto.InternalMessageInfo

func (m *PCmdActiveThreadCountRes) GetCommonStreamResponse() *PCmdStreamResponse {
	if m != nil {
		return m.CommonStreamResponse
	}
	return nil
}

func (m *PCmdActiveThreadCountRes) GetHistogramSchemaType() int32 {
	if m != nil {
		return m.HistogramSchemaType
	}
	return 0
}

func (m *PCmdActiveThreadCountRes) GetActiveThreadCount() []int32 {
	if m != nil {
		return m.ActiveThreadCount
	}
	return nil
}

func (m *PCmdActiveThreadCountRes) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("v1.PCommandType", PCommandType_name, PCommandType_value)
	proto.RegisterType((*PCmdMessage)(nil), "v1.PCmdMessage")
	proto.RegisterType((*PCmdServiceHandshake)(nil), "v1.PCmdServiceHandshake")
	proto.RegisterType((*PCmdResponse)(nil), "v1.PCmdResponse")
	proto.RegisterType((*PCmdStreamResponse)(nil), "v1.PCmdStreamResponse")
	proto.RegisterType((*PCmdRequest)(nil), "v1.PCmdRequest")
	proto.RegisterType((*PCmdEcho)(nil), "v1.PCmdEcho")
	proto.RegisterType((*PCmdEchoResponse)(nil), "v1.PCmdEchoResponse")
	proto.RegisterType((*PCmdActiveThreadDump)(nil), "v1.PCmdActiveThreadDump")
	proto.RegisterType((*PCmdActiveThreadDumpRes)(nil), "v1.PCmdActiveThreadDumpRes")
	proto.RegisterType((*PCmdActiveThreadLightDump)(nil), "v1.PCmdActiveThreadLightDump")
	proto.RegisterType((*PCmdActiveThreadLightDumpRes)(nil), "v1.PCmdActiveThreadLightDumpRes")
	proto.RegisterType((*PCmdActiveThreadCount)(nil), "v1.PCmdActiveThreadCount")
	proto.RegisterType((*PCmdActiveThreadCountRes)(nil), "v1.PCmdActiveThreadCountRes")
}

func init() { proto.RegisterFile("v1/Cmd.proto", fileDescriptor_2aca7d559118759c) }

var fileDescriptor_2aca7d559118759c = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4f, 0x6f, 0xdb, 0x36,
	0x14, 0xb7, 0xa2, 0xa8, 0x8b, 0x9f, 0x83, 0x41, 0x63, 0xd2, 0x46, 0x0d, 0xd2, 0x20, 0x33, 0x76,
	0x08, 0x86, 0x41, 0x9e, 0xbb, 0x6e, 0x18, 0x7a, 0x4b, 0x1c, 0x2f, 0xce, 0xd6, 0x3a, 0x06, 0xed,
	0x14, 0xc3, 0x2e, 0x05, 0x23, 0x31, 0xb6, 0x30, 0x4b, 0xd4, 0x28, 0x4a, 0x43, 0x2f, 0xc3, 0xae,
	0x03, 0x76, 0xde, 0x47, 0xea, 0x07, 0xd8, 0xb1, 0xd8, 0x31, 0xe7, 0x7e, 0x86, 0x81, 0xa4, 0x64,
	0xcb, 0x95, 0x3c, 0x0c, 0x2d, 0x72, 0x23, 0xdf, 0xef, 0xf7, 0xfe, 0xfd, 0xf8, 0x48, 0xc2, 0x76,
	0xd6, 0xed, 0xf4, 0x42, 0xdf, 0x8d, 0x39, 0x13, 0x0c, 0x6d, 0x64, 0xdd, 0xfd, 0xc3, 0x29, 0x63,
	0xd3, 0x39, 0xed, 0x28, 0xcb, 0x75, 0x7a, 0xd3, 0xf9, 0x95, 0x93, 0x38, 0xa6, 0x3c, 0xd1, 0x9c,
	0xfd, 0x9d, 0xac, 0xdb, 0x99, 0xcc, 0x38, 0x25, 0xfe, 0x59, 0x1a, 0xc6, 0xda, 0xd8, 0xfe, 0xcb,
	0x80, 0xd6, 0xa8, 0x17, 0xfa, 0xcf, 0x69, 0x92, 0x90, 0x29, 0x45, 0xdf, 0x81, 0x3d, 0x23, 0x91,
	0x9f, 0xcc, 0xc8, 0xcf, 0x34, 0xb7, 0x39, 0xc6, 0x91, 0x71, 0xdc, 0x7a, 0xec, 0xb8, 0x59, 0xd7,
	0x95, 0xd4, 0x31, 0xe5, 0x59, 0xe0, 0xd1, 0x41, 0x41, 0x1b, 0x34, 0x70, 0xc5, 0x07, 0x3d, 0x81,
	0xd6, 0x0d, 0x09, 0xe6, 0x45, 0x88, 0x0d, 0x15, 0xc2, 0x2e, 0x42, 0x60, 0x9a, 0xc4, 0x2c, 0x4a,
	0xa4, 0x6b, 0x99, 0x76, 0xda, 0x84, 0x8f, 0x42, 0xbd, 0x6c, 0x63, 0xd8, 0xad, 0x4b, 0x86, 0x9e,
	0x82, 0x93, 0xa4, 0x71, 0xcc, 0xb8, 0xe8, 0xb1, 0x30, 0x24, 0x51, 0xc1, 0xf8, 0x81, 0xbe, 0x72,
	0x8c, 0x23, 0xf3, 0xd8, 0xc2, 0x6b, 0xf1, 0xf6, 0x6f, 0xb0, 0x5d, 0xce, 0x8e, 0x0e, 0x01, 0x78,
	0xbe, 0xbe, 0xf0, 0x55, 0x9b, 0x16, 0x2e, 0x59, 0xd0, 0x03, 0xb8, 0x97, 0x08, 0x22, 0xd2, 0x44,
	0xd5, 0x6f, 0xe1, 0x7c, 0x87, 0xbe, 0x59, 0x94, 0xe9, 0x98, 0xaa, 0xb1, 0x03, 0x57, 0x6b, 0xef,
	0x16, 0xda, 0xbb, 0x63, 0xc1, 0x83, 0x68, 0xfa, 0x82, 0xcc, 0x53, 0x8a, 0x17, 0x3d, 0xfd, 0x69,
	0x00, 0x52, 0x4d, 0x09, 0x4e, 0x49, 0xf8, 0xbf, 0xcb, 0x38, 0x04, 0x48, 0xe8, 0x2f, 0x29, 0x8d,
	0x3c, 0x89, 0xeb, 0x52, 0x4a, 0x96, 0xf7, 0x2e, 0xe7, 0x77, 0x53, 0x9f, 0x3d, 0x96, 0xa1, 0x12,
	0x81, 0x0e, 0xa0, 0xc9, 0xf5, 0x72, 0x51, 0xc6, 0xd2, 0x80, 0xba, 0xd0, 0xf2, 0xb4, 0xa2, 0x7d,
	0x6f, 0xc6, 0x9c, 0xd7, 0x96, 0x4a, 0xb5, 0x5d, 0x1c, 0xa9, 0x34, 0xca, 0xe3, 0x2c, 0x71, 0xd0,
	0x8f, 0xe0, 0xe4, 0xdb, 0x13, 0x4f, 0x04, 0x19, 0xd5, 0xd3, 0xd7, 0x63, 0x69, 0x24, 0x9c, 0x37,
	0xda, 0xff, 0x61, 0xe1, 0x5f, 0x61, 0x0c, 0x1a, 0x78, 0xad, 0x37, 0xba, 0x82, 0xbd, 0x1a, 0x4c,
	0xce, 0xb5, 0x73, 0x6b, 0xad, 0x8e, 0xeb, 0xbb, 0x84, 0x41, 0x03, 0xaf, 0xf3, 0x45, 0x3e, 0x1c,
	0xd4, 0x40, 0xcf, 0x82, 0xe9, 0x4c, 0xa8, 0xd8, 0x6f, 0x75, 0xec, 0x47, 0x75, 0xb1, 0x17, 0xac,
	0x41, 0x03, 0xff, 0x67, 0x14, 0x39, 0xe5, 0x39, 0xde, 0xfe, 0x0c, 0xb6, 0x0a, 0xf1, 0x90, 0xb3,
	0x3c, 0x46, 0x29, 0x7e, 0x73, 0x79, 0x50, 0x37, 0x60, 0x17, 0xac, 0xc5, 0xd0, 0x7c, 0x0b, 0x1f,
	0xcb, 0x20, 0x2c, 0x2a, 0x2c, 0xf9, 0x35, 0xad, 0xdc, 0x31, 0xfc, 0x0e, 0xaf, 0x9c, 0x67, 0x63,
	0x35, 0x4f, 0xac, 0xef, 0x5c, 0x45, 0x96, 0x5d, 0xb0, 0xe6, 0x41, 0x18, 0x88, 0x7c, 0x28, 0xf4,
	0x46, 0x8e, 0xa5, 0x50, 0x9c, 0x21, 0x09, 0x65, 0x28, 0xf3, 0xb8, 0x89, 0x4b, 0x16, 0xd4, 0x86,
	0xed, 0x39, 0xf3, 0xc8, 0x7c, 0xc2, 0x89, 0x1a, 0x5c, 0xf3, 0xc8, 0x3c, 0x36, 0xf1, 0x8a, 0xad,
	0xfd, 0xb7, 0x01, 0x7b, 0x75, 0x29, 0x31, 0x4d, 0x3e, 0xa0, 0xc3, 0xaf, 0x8b, 0xca, 0xd4, 0xa1,
	0xc9, 0xca, 0x5a, 0x8f, 0xef, 0x2b, 0xaf, 0x4a, 0x9e, 0x12, 0x11, 0x21, 0xd8, 0x14, 0xaf, 0x62,
	0x7d, 0x89, 0x9a, 0x58, 0xad, 0xa5, 0x58, 0x49, 0x7a, 0x3d, 0x91, 0xe6, 0x4d, 0x2d, 0x56, 0xbe,
	0x95, 0x48, 0x46, 0x79, 0x12, 0xb0, 0xc8, 0xb1, 0x34, 0x92, 0x6f, 0xdb, 0x29, 0x3c, 0x5c, 0x3b,
	0x1c, 0x77, 0xa8, 0xe5, 0x3f, 0x06, 0x1c, 0xac, 0xcd, 0xfb, 0x61, 0x82, 0x3e, 0xad, 0x11, 0x74,
	0xbf, 0x22, 0xe8, 0x32, 0xd9, 0x5d, 0xa8, 0xba, 0x07, 0xf7, 0x6b, 0xdf, 0x89, 0xf6, 0xad, 0x01,
	0x4e, 0x2d, 0x22, 0x7b, 0xfe, 0x1e, 0x76, 0x75, 0x2f, 0xab, 0x6f, 0x6e, 0xde, 0xf9, 0x83, 0xc5,
	0x9f, 0xb6, 0x82, 0xe2, 0x5a, 0x1f, 0xf4, 0x25, 0xec, 0xcc, 0x82, 0x44, 0xb0, 0x29, 0x27, 0xe1,
	0xd8, 0x9b, 0xd1, 0x90, 0xa8, 0x0e, 0xf4, 0x83, 0x5c, 0x07, 0xa1, 0x2f, 0xe0, 0x13, 0x52, 0x79,
	0xf9, 0x4c, 0xf5, 0x4b, 0x55, 0x01, 0xf9, 0xfe, 0x8a, 0x20, 0xa4, 0x63, 0x41, 0xc2, 0x58, 0xe9,
	0x62, 0xe2, 0xa5, 0xe1, 0xf3, 0x3f, 0x0c, 0xf9, 0x7b, 0xe9, 0x77, 0x43, 0x05, 0xdf, 0x82, 0xcd,
	0xe1, 0xe5, 0xb0, 0x6f, 0x37, 0xe4, 0x6a, 0x74, 0x31, 0x3c, 0xb7, 0x7d, 0xb5, 0xba, 0x1c, 0x9e,
	0xdb, 0x14, 0x35, 0x61, 0xb3, 0xdf, 0x1b, 0x5c, 0xda, 0xaf, 0x2d, 0xe4, 0xc0, 0xce, 0x49, 0x6f,
	0x72, 0xf1, 0xa2, 0xff, 0x72, 0x32, 0xc0, 0xfd, 0x93, 0xb3, 0x97, 0xbd, 0xcb, 0xab, 0xe1, 0xc4,
	0x7e, 0x63, 0xa1, 0x3d, 0x40, 0xab, 0xc8, 0xd9, 0xd5, 0xf3, 0x91, 0x7d, 0x6b, 0xa1, 0x47, 0xe0,
	0xac, 0x02, 0xcf, 0x2e, 0xce, 0x07, 0x13, 0x0d, 0xbf, 0xb5, 0x4e, 0x9f, 0xc0, 0xa7, 0x1e, 0x0b,
	0xdd, 0x88, 0x64, 0x94, 0x7b, 0x8c, 0xc7, 0x6e, 0x1c, 0x44, 0x31, 0x0b, 0x22, 0xe1, 0x4e, 0x79,
	0xec, 0xb9, 0x42, 0x4e, 0xe4, 0xe9, 0x56, 0x2f, 0xf4, 0x47, 0xf2, 0x03, 0x1a, 0x19, 0x3f, 0x99,
	0x9d, 0xac, 0x7b, 0x7d, 0x4f, 0x7d, 0x47, 0x5f, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x16, 0xea,
	0xe8, 0x49, 0xbb, 0x08, 0x00, 0x00,
}
