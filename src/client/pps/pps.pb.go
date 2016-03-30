// Code generated by protoc-gen-go.
// source: client/pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	client/pps/pps.proto

It has these top-level messages:
	Transform
	Job
	JobInput
	JobInfo
	JobInfos
	Pipeline
	PipelineInput
	PipelineInfo
	PipelineInfos
	CreateJobRequest
	InspectJobRequest
	ListJobRequest
	CreatePipelineRequest
	InspectPipelineRequest
	ListPipelineRequest
	DeletePipelineRequest
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type JobState int32

const (
	JobState_JOB_STATE_RUNNING JobState = 0
	JobState_JOB_STATE_FAILURE JobState = 1
	JobState_JOB_STATE_SUCCESS JobState = 2
)

var JobState_name = map[int32]string{
	0: "JOB_STATE_RUNNING",
	1: "JOB_STATE_FAILURE",
	2: "JOB_STATE_SUCCESS",
}
var JobState_value = map[string]int32{
	"JOB_STATE_RUNNING": 0,
	"JOB_STATE_FAILURE": 1,
	"JOB_STATE_SUCCESS": 2,
}

func (x JobState) String() string {
	return proto.EnumName(JobState_name, int32(x))
}
func (JobState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Transform struct {
	Image string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Cmd   []string `protobuf:"bytes,2,rep,name=cmd" json:"cmd,omitempty"`
	Stdin []string `protobuf:"bytes,3,rep,name=stdin" json:"stdin,omitempty"`
}

func (m *Transform) Reset()                    { *m = Transform{} }
func (m *Transform) String() string            { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()               {}
func (*Transform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Job struct {
	ID string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type JobInput struct {
	Commit *pfs.Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Reduce bool        `protobuf:"varint,2,opt,name=reduce" json:"reduce,omitempty"`
}

func (m *JobInput) Reset()                    { *m = JobInput{} }
func (m *JobInput) String() string            { return proto.CompactTextString(m) }
func (*JobInput) ProtoMessage()               {}
func (*JobInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobInput) GetCommit() *pfs.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type JobInfo struct {
	Job          *Job                        `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Transform    *Transform                  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Pipeline     *Pipeline                   `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline,omitempty"`
	Shards       uint64                      `protobuf:"varint,4,opt,name=shards" json:"shards,omitempty"`
	Inputs       []*JobInput                 `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob    *Job                        `protobuf:"bytes,6,opt,name=parent_job" json:"parent_job,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=created_at" json:"created_at,omitempty"`
	OutputCommit *pfs.Commit                 `protobuf:"bytes,8,opt,name=output_commit" json:"output_commit,omitempty"`
	State        JobState                    `protobuf:"varint,9,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
}

func (m *JobInfo) Reset()                    { *m = JobInfo{} }
func (m *JobInfo) String() string            { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()               {}
func (*JobInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JobInfo) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *JobInfo) GetInputs() []*JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobInfo) GetParentJob() *Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JobInfo) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()                    { *m = JobInfos{} }
func (m *JobInfos) String() string            { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()               {}
func (*JobInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type Pipeline struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Pipeline) Reset()                    { *m = Pipeline{} }
func (m *Pipeline) String() string            { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()               {}
func (*Pipeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PipelineInput struct {
	Repo   *pfs.Repo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	Reduce bool      `protobuf:"varint,2,opt,name=reduce" json:"reduce,omitempty"`
}

func (m *PipelineInput) Reset()                    { *m = PipelineInput{} }
func (m *PipelineInput) String() string            { return proto.CompactTextString(m) }
func (*PipelineInput) ProtoMessage()               {}
func (*PipelineInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PipelineInput) GetRepo() *pfs.Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

type PipelineInfo struct {
	Pipeline   *Pipeline                   `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform  *Transform                  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards     uint64                      `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	Inputs     []*PipelineInput            `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	OutputRepo *pfs.Repo                   `protobuf:"bytes,5,opt,name=output_repo" json:"output_repo,omitempty"`
	CreatedAt  *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *PipelineInfo) Reset()                    { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()               {}
func (*PipelineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PipelineInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *PipelineInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputs() []*PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineInfo) GetOutputRepo() *pfs.Repo {
	if m != nil {
		return m.OutputRepo
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()                    { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()               {}
func (*PipelineInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

type CreateJobRequest struct {
	Transform *Transform  `protobuf:"bytes,1,opt,name=transform" json:"transform,omitempty"`
	Pipeline  *Pipeline   `protobuf:"bytes,2,opt,name=pipeline" json:"pipeline,omitempty"`
	Shards    uint64      `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	Inputs    []*JobInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob *Job        `protobuf:"bytes,5,opt,name=parent_job" json:"parent_job,omitempty"`
}

func (m *CreateJobRequest) Reset()                    { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()               {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreateJobRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreateJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreateJobRequest) GetInputs() []*JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateJobRequest) GetParentJob() *Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

type InspectJobRequest struct {
	Job         *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	BlockOutput bool `protobuf:"varint,2,opt,name=block_output" json:"block_output,omitempty"`
	BlockState  bool `protobuf:"varint,3,opt,name=block_state" json:"block_state,omitempty"`
}

func (m *InspectJobRequest) Reset()                    { *m = InspectJobRequest{} }
func (m *InspectJobRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectJobRequest) ProtoMessage()               {}
func (*InspectJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *InspectJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type ListJobRequest struct {
	Pipeline    *Pipeline     `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	InputCommit []*pfs.Commit `protobuf:"bytes,2,rep,name=input_commit" json:"input_commit,omitempty"`
}

func (m *ListJobRequest) Reset()                    { *m = ListJobRequest{} }
func (m *ListJobRequest) String() string            { return proto.CompactTextString(m) }
func (*ListJobRequest) ProtoMessage()               {}
func (*ListJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ListJobRequest) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

type CreatePipelineRequest struct {
	Pipeline  *Pipeline        `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform *Transform       `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards    uint64           `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	Inputs    []*PipelineInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
}

func (m *CreatePipelineRequest) Reset()                    { *m = CreatePipelineRequest{} }
func (m *CreatePipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePipelineRequest) ProtoMessage()               {}
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CreatePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreatePipelineRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreatePipelineRequest) GetInputs() []*PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type InspectPipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *InspectPipelineRequest) Reset()                    { *m = InspectPipelineRequest{} }
func (m *InspectPipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectPipelineRequest) ProtoMessage()               {}
func (*InspectPipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *InspectPipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type ListPipelineRequest struct {
}

func (m *ListPipelineRequest) Reset()                    { *m = ListPipelineRequest{} }
func (m *ListPipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPipelineRequest) ProtoMessage()               {}
func (*ListPipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type DeletePipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *DeletePipelineRequest) Reset()                    { *m = DeletePipelineRequest{} }
func (m *DeletePipelineRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePipelineRequest) ProtoMessage()               {}
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeletePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func init() {
	proto.RegisterType((*Transform)(nil), "pachyderm.pps.Transform")
	proto.RegisterType((*Job)(nil), "pachyderm.pps.Job")
	proto.RegisterType((*JobInput)(nil), "pachyderm.pps.JobInput")
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.JobInfos")
	proto.RegisterType((*Pipeline)(nil), "pachyderm.pps.Pipeline")
	proto.RegisterType((*PipelineInput)(nil), "pachyderm.pps.PipelineInput")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.PipelineInfo")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.PipelineInfos")
	proto.RegisterType((*CreateJobRequest)(nil), "pachyderm.pps.CreateJobRequest")
	proto.RegisterType((*InspectJobRequest)(nil), "pachyderm.pps.InspectJobRequest")
	proto.RegisterType((*ListJobRequest)(nil), "pachyderm.pps.ListJobRequest")
	proto.RegisterType((*CreatePipelineRequest)(nil), "pachyderm.pps.CreatePipelineRequest")
	proto.RegisterType((*InspectPipelineRequest)(nil), "pachyderm.pps.InspectPipelineRequest")
	proto.RegisterType((*ListPipelineRequest)(nil), "pachyderm.pps.ListPipelineRequest")
	proto.RegisterType((*DeletePipelineRequest)(nil), "pachyderm.pps.DeletePipelineRequest")
	proto.RegisterEnum("pachyderm.pps.JobState", JobState_name, JobState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/ListJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/CreatePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/InspectPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/ListPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.API/DeletePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	InspectJob(context.Context, *InspectJobRequest) (*JobInfo, error)
	ListJob(context.Context, *ListJobRequest) (*JobInfos, error)
	CreatePipeline(context.Context, *CreatePipelineRequest) (*google_protobuf.Empty, error)
	InspectPipeline(context.Context, *InspectPipelineRequest) (*PipelineInfo, error)
	ListPipeline(context.Context, *ListPipelineRequest) (*PipelineInfos, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*google_protobuf.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).InspectJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreatePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_InspectPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).InspectPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeletePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _API_CreateJob_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _API_InspectJob_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _API_ListJob_Handler,
		},
		{
			MethodName: "CreatePipeline",
			Handler:    _API_CreatePipeline_Handler,
		},
		{
			MethodName: "InspectPipeline",
			Handler:    _API_InspectPipeline_Handler,
		},
		{
			MethodName: "ListPipeline",
			Handler:    _API_ListPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _API_DeletePipeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for JobAPI service

type JobAPIClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
}

type jobAPIClient struct {
	cc *grpc.ClientConn
}

func NewJobAPIClient(cc *grpc.ClientConn) JobAPIClient {
	return &jobAPIClient{cc}
}

func (c *jobAPIClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/ListJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobAPI service

type JobAPIServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	InspectJob(context.Context, *InspectJobRequest) (*JobInfo, error)
	ListJob(context.Context, *ListJobRequest) (*JobInfos, error)
}

func RegisterJobAPIServer(s *grpc.Server, srv JobAPIServer) {
	s.RegisterService(&_JobAPI_serviceDesc, srv)
}

func _JobAPI_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).CreateJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).InspectJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).ListJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _JobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.JobAPI",
	HandlerType: (*JobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobAPI_CreateJob_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _JobAPI_InspectJob_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _JobAPI_ListJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for PipelineAPI service

type PipelineAPIClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type pipelineAPIClient struct {
	cc *grpc.ClientConn
}

func NewPipelineAPIClient(cc *grpc.ClientConn) PipelineAPIClient {
	return &pipelineAPIClient{cc}
}

func (c *pipelineAPIClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/CreatePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/InspectPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/ListPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/DeletePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PipelineAPI service

type PipelineAPIServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*google_protobuf.Empty, error)
	InspectPipeline(context.Context, *InspectPipelineRequest) (*PipelineInfo, error)
	ListPipeline(context.Context, *ListPipelineRequest) (*PipelineInfos, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*google_protobuf.Empty, error)
}

func RegisterPipelineAPIServer(s *grpc.Server, srv PipelineAPIServer) {
	s.RegisterService(&_PipelineAPI_serviceDesc, srv)
}

func _PipelineAPI_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).CreatePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_InspectPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).InspectPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_ListPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).ListPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).DeletePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PipelineAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.PipelineAPI",
	HandlerType: (*PipelineAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineAPI_CreatePipeline_Handler,
		},
		{
			MethodName: "InspectPipeline",
			Handler:    _PipelineAPI_InspectPipeline_Handler,
		},
		{
			MethodName: "ListPipeline",
			Handler:    _PipelineAPI_ListPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineAPI_DeletePipeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x56, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0xc6, 0x18, 0x08, 0x1c, 0x7e, 0x96, 0x4c, 0x42, 0x62, 0x91, 0xec, 0x26, 0x6b, 0xed, 0x66,
	0xb3, 0x6d, 0xe5, 0x48, 0xb4, 0x6a, 0x7b, 0x1b, 0x28, 0x6d, 0x89, 0xa2, 0x24, 0x02, 0xa2, 0x4a,
	0xbd, 0xa8, 0x65, 0xcc, 0x90, 0xb8, 0xc5, 0xd8, 0xb5, 0xcd, 0x45, 0x1e, 0xa9, 0x17, 0x7d, 0x92,
	0x3e, 0x42, 0xa5, 0xf6, 0x25, 0xfa, 0x00, 0x9d, 0x19, 0xdb, 0x80, 0x07, 0x93, 0x94, 0xb4, 0xea,
	0x45, 0xd5, 0x0b, 0x2e, 0x3c, 0xe7, 0xcc, 0x39, 0xe7, 0xfb, 0x39, 0x23, 0x60, 0x5d, 0x1f, 0x1a,
	0x78, 0xe4, 0x1d, 0xd8, 0xb6, 0x4b, 0x7f, 0x8a, 0xed, 0x58, 0x9e, 0x85, 0x8a, 0xb6, 0xa6, 0x5f,
	0x5e, 0xf5, 0xb1, 0x63, 0x2a, 0xe4, 0xb0, 0xba, 0x75, 0x61, 0x59, 0x17, 0x43, 0x7c, 0xc0, 0x82,
	0xbd, 0xf1, 0xe0, 0x00, 0x9b, 0xb6, 0x77, 0xe5, 0xe7, 0x56, 0x77, 0xf8, 0xa0, 0x67, 0x98, 0xd8,
	0xf5, 0x34, 0xd3, 0x0e, 0x12, 0x26, 0x2d, 0x06, 0x2e, 0xfd, 0xf9, 0xa7, 0xf2, 0x43, 0xc8, 0x75,
	0x1d, 0x6d, 0xe4, 0x0e, 0x2c, 0xc7, 0x44, 0x45, 0x48, 0x1b, 0xa6, 0x76, 0x81, 0x25, 0x61, 0x57,
	0xd8, 0xcf, 0xa1, 0x3c, 0x88, 0xba, 0xd9, 0x97, 0x92, 0xbb, 0x22, 0xf9, 0x20, 0x31, 0xd7, 0xeb,
	0x1b, 0x23, 0x49, 0xa4, 0x9f, 0xf2, 0x2a, 0x88, 0x47, 0x56, 0x0f, 0x01, 0x24, 0x8d, 0xbe, 0x9f,
	0x2e, 0x3f, 0x82, 0x2c, 0x39, 0x6a, 0x8d, 0xec, 0xb1, 0x87, 0xb6, 0x20, 0xa3, 0x5b, 0xa6, 0x69,
	0x78, 0x2c, 0x96, 0xaf, 0xe5, 0x15, 0xda, 0xb2, 0xc1, 0x8e, 0x50, 0x09, 0x32, 0x0e, 0xee, 0x8f,
	0x75, 0x4c, 0x4a, 0x0b, 0xfb, 0x59, 0xf9, 0x73, 0x12, 0x56, 0xd8, 0xcd, 0x81, 0x85, 0x76, 0x40,
	0x7c, 0x6d, 0xf5, 0x82, 0x5b, 0x48, 0x89, 0x10, 0xa0, 0xd0, 0x8e, 0x77, 0x21, 0xe7, 0x85, 0x03,
	0xb3, 0xfb, 0xf9, 0x9a, 0xc4, 0xa5, 0x4d, 0x01, 0xfd, 0x0f, 0x59, 0xdb, 0xb0, 0xf1, 0xd0, 0x18,
	0x61, 0x32, 0x37, 0xcd, 0xdd, 0xe4, 0x72, 0xcf, 0x82, 0x30, 0x1d, 0xca, 0xbd, 0xd4, 0x9c, 0xbe,
	0x2b, 0xa5, 0x48, 0x62, 0x0a, 0xfd, 0x07, 0x19, 0x83, 0x42, 0x71, 0xa5, 0x34, 0x01, 0x3c, 0x7f,
	0x71, 0x02, 0x75, 0x0f, 0xc0, 0xd6, 0x1c, 0xc2, 0xac, 0x4a, 0x07, 0xcf, 0x2c, 0x1c, 0x5c, 0x01,
	0xd0, 0x1d, 0xac, 0x79, 0xb8, 0xaf, 0x6a, 0x9e, 0xb4, 0xc2, 0xf2, 0xaa, 0x8a, 0xaf, 0x9a, 0x12,
	0xaa, 0xa6, 0x74, 0x43, 0xd5, 0x90, 0x0c, 0x45, 0x6b, 0xec, 0x91, 0x0e, 0x6a, 0xc0, 0x64, 0x76,
	0x9e, 0xc9, 0x3d, 0x2a, 0x0a, 0x29, 0x29, 0xe5, 0x48, 0xac, 0x14, 0x37, 0x63, 0x87, 0x86, 0xe5,
	0x07, 0x81, 0x34, 0x03, 0xcb, 0x45, 0xfb, 0x90, 0x25, 0x83, 0xaa, 0x06, 0xf9, 0x20, 0x34, 0x53,
	0x68, 0x1b, 0x71, 0xd0, 0x06, 0x96, 0x2c, 0x41, 0x76, 0x42, 0x4f, 0x01, 0x52, 0x23, 0xcd, 0x0c,
	0x9c, 0x21, 0x3f, 0x86, 0x62, 0x18, 0xf1, 0x49, 0xd8, 0x84, 0x94, 0x83, 0x6d, 0x2b, 0xd0, 0x2d,
	0xc7, 0x66, 0x6c, 0x93, 0x83, 0x39, 0xad, 0xbf, 0x08, 0x50, 0x98, 0x5e, 0x25, 0x82, 0xcf, 0x4a,
	0x24, 0x5c, 0x2f, 0xd1, 0x52, 0xd2, 0x4f, 0xf5, 0x14, 0x99, 0x9e, 0xf7, 0x26, 0x7a, 0xa6, 0x18,
	0xe8, 0xed, 0x05, 0x5d, 0x7c, 0x3c, 0x7f, 0x41, 0x3e, 0x20, 0x9f, 0xc1, 0x4a, 0xf3, 0xb0, 0xa2,
	0x62, 0x66, 0x6e, 0x12, 0x53, 0x6e, 0xcc, 0x12, 0x46, 0x55, 0xa8, 0x41, 0x31, 0x84, 0x3d, 0x2b,
	0xc5, 0xd6, 0xc2, 0xa9, 0x88, 0x1e, 0x1f, 0x04, 0x28, 0x37, 0x58, 0x57, 0xa2, 0x50, 0x1b, 0xbf,
	0x1d, 0x93, 0xe2, 0x51, 0x52, 0x84, 0x25, 0xf6, 0x21, 0xf9, 0xad, 0xfb, 0x20, 0x72, 0xfb, 0x90,
	0x5a, 0x66, 0x1f, 0xd2, 0x8b, 0xf6, 0x41, 0x56, 0x61, 0xb5, 0x35, 0x72, 0x6d, 0xac, 0x7b, 0x33,
	0x68, 0x6e, 0x5c, 0xff, 0x75, 0x28, 0xf4, 0x86, 0x96, 0xfe, 0x46, 0xf5, 0xe5, 0xf1, 0x5d, 0x85,
	0xd6, 0x20, 0xef, 0x9f, 0xfa, 0xdb, 0x20, 0x32, 0xab, 0xbd, 0x82, 0xd2, 0xb1, 0xe1, 0xce, 0x56,
	0x5f, 0xc2, 0x6b, 0x7f, 0x43, 0x81, 0xc1, 0x0d, 0x97, 0x2f, 0xc9, 0x40, 0xcf, 0x2e, 0x9f, 0xfc,
	0x5e, 0x80, 0x8a, 0x2f, 0x47, 0x78, 0xeb, 0x16, 0x7d, 0x7e, 0x9e, 0xa7, 0x89, 0x07, 0x37, 0x02,
	0xc2, 0x6f, 0x3f, 0xaf, 0x5c, 0x81, 0x35, 0x4a, 0x2a, 0x57, 0x41, 0xae, 0x43, 0xe5, 0x09, 0x1e,
	0xe2, 0xef, 0xa1, 0xe2, 0xce, 0x29, 0x7b, 0xa4, 0xd8, 0x83, 0x85, 0x2a, 0xb0, 0x7a, 0x74, 0x5a,
	0x57, 0x3b, 0xdd, 0xc3, 0x6e, 0x53, 0x6d, 0x9f, 0x9f, 0x9c, 0xb4, 0x4e, 0x9e, 0x95, 0x13, 0xd1,
	0xe3, 0xa7, 0x87, 0xad, 0xe3, 0xf3, 0x76, 0xb3, 0x2c, 0x44, 0x8f, 0x3b, 0xe7, 0x8d, 0x46, 0xb3,
	0xd3, 0x29, 0x27, 0x6b, 0xef, 0x52, 0x20, 0x1e, 0x9e, 0xb5, 0x50, 0x1d, 0x72, 0x93, 0xb5, 0x41,
	0x3b, 0x5c, 0x7b, 0x7e, 0xa1, 0xaa, 0x71, 0x5e, 0x4d, 0xa0, 0xe7, 0x00, 0x53, 0xb7, 0xa2, 0x5d,
	0x2e, 0x67, 0xce, 0xc8, 0xd5, 0x45, 0x6f, 0x6a, 0x02, 0x35, 0x60, 0x25, 0xb0, 0x25, 0xfa, 0x93,
	0x4b, 0x8a, 0xda, 0xb5, 0xba, 0x19, 0x5f, 0xc3, 0x25, 0x45, 0xce, 0xa0, 0x14, 0xb5, 0x1e, 0xfa,
	0x27, 0x16, 0x17, 0x27, 0x07, 0x19, 0x8b, 0x7f, 0xa3, 0x9a, 0xf4, 0x3f, 0x04, 0xa9, 0xf8, 0x02,
	0xfe, 0xe0, 0xdc, 0x81, 0xfe, 0x8d, 0x47, 0xc9, 0xd7, 0xbc, 0xf6, 0xcd, 0x4a, 0xa0, 0x36, 0x14,
	0x66, 0x1d, 0x83, 0xe4, 0x18, 0xd0, 0x7c, 0xc9, 0xed, 0x6b, 0x4a, 0x06, 0xf0, 0xa3, 0x76, 0x9b,
	0x83, 0x1f, 0xeb, 0xc6, 0xc5, 0xf0, 0x6b, 0x1f, 0x05, 0xc8, 0x10, 0x7e, 0x7f, 0x4d, 0xbb, 0xd4,
	0x3e, 0x25, 0x21, 0x1f, 0x72, 0x41, 0x21, 0xfe, 0xb6, 0xcf, 0x8f, 0xb5, 0x4f, 0x3d, 0xfd, 0x52,
	0x24, 0xd7, 0x7a, 0x19, 0x16, 0xb8, 0xff, 0x35, 0x00, 0x00, 0xff, 0xff, 0x46, 0xb8, 0x2b, 0xb8,
	0xcf, 0x0b, 0x00, 0x00,
}