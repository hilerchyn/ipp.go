package ipp

import (
	"fmt"
	"sync/atomic"
)

type IppRequest struct {
	data        string
	debug       bool
	parsed      bool
	operationId int16
	RequestId   int32
}

func (r *IppRequest) GetMessage() *IppMessage {
	m := NewIppMessage()
	m.SetOperationIdStatusCode(r.operationId)
	m.SetRequestId(r.RequestId)
	return m
}

func (r *IppRequest) SetDebug(debug bool) {
	r.debug = debug
}
func (r *IppRequest) SetRequestId(requestId int32) {
	r.RequestId = requestId
}
func (r *IppRequest) SetOperationId(operationId int16) {
	r.operationId = operationId
}

type CupsServer struct {
	uri            string
	username       string
	password       string
	debug          bool
	requestCounter int32
}

func (c *CupsServer) SetServer(server string) {
	c.uri = server
}

func (c *CupsServer) DoRequest(request IppRequest) string {
	return fmt.Sprintf("ok %d", request.RequestId)
}

func (c *CupsServer) getRequestId() int32 {
	return atomic.AddInt32(&c.requestCounter, 1)
}

func (c *CupsServer) CreateRequest(operationId int16) IppRequest {
	var r IppRequest
	r.SetDebug(c.debug)
	r.SetRequestId(c.getRequestId())
	r.SetOperationId(operationId)
	return r
}
