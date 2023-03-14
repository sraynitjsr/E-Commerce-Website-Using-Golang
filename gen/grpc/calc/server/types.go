// Code generated by goa v3.11.1, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen github.com/sraynitjsr/design

package server

import (
	calc "github.com/sraynitjsr/gen/calc"
	calcpb "github.com/sraynitjsr/gen/grpc/calc/pb"
)

// NewAddPayload builds the payload of the "add" endpoint of the "calc" service
// from the gRPC request type.
func NewAddPayload(message *calcpb.AddRequest) *calc.AddPayload {
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewProtoAddResponse builds the gRPC response type from the result of the
// "add" endpoint of the "calc" service.
func NewProtoAddResponse(result int) *calcpb.AddResponse {
	message := &calcpb.AddResponse{}
	message.Field = int32(result)
	return message
}

// NewSubPayload builds the payload of the "sub" endpoint of the "calc" service
// from the gRPC request type.
func NewSubPayload(message *calcpb.SubRequest) *calc.SubPayload {
	v := &calc.SubPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewProtoSubResponse builds the gRPC response type from the result of the
// "sub" endpoint of the "calc" service.
func NewProtoSubResponse(result int) *calcpb.SubResponse {
	message := &calcpb.SubResponse{}
	message.Field = int32(result)
	return message
}

// NewMultiplyPayload builds the payload of the "multiply" endpoint of the
// "calc" service from the gRPC request type.
func NewMultiplyPayload(message *calcpb.MultiplyRequest) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewProtoMultiplyResponse builds the gRPC response type from the result of
// the "multiply" endpoint of the "calc" service.
func NewProtoMultiplyResponse(result int) *calcpb.MultiplyResponse {
	message := &calcpb.MultiplyResponse{}
	message.Field = int32(result)
	return message
}

// NewDividePayload builds the payload of the "divide" endpoint of the "calc"
// service from the gRPC request type.
func NewDividePayload(message *calcpb.DivideRequest) *calc.DividePayload {
	v := &calc.DividePayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewProtoDivideResponse builds the gRPC response type from the result of the
// "divide" endpoint of the "calc" service.
func NewProtoDivideResponse(result int) *calcpb.DivideResponse {
	message := &calcpb.DivideResponse{}
	message.Field = int32(result)
	return message
}
