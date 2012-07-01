package ipp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync/atomic"
)

type CupsServer struct {
	uri            string
	username       string
	password       string
	requestCounter int32
}

func (c *CupsServer) SetServer(server string) {
	c.uri = server
}

func (c *CupsServer) getRequestId() int32 {
	return atomic.AddInt32(&c.requestCounter, 1)
}

func (c *CupsServer) CreateRequest(operationId int16) *IppMessage {
	m := NewIppMessage()
	m.SetOperationIdStatusCode(operationId)
	m.SetRequestId(c.getRequestId())
	return m
}

func (c *CupsServer) GetPrinters() {
	m := c.CreateRequest(CUPS_GET_PRINTERS)
	g1 := NewIppAttributeGroup(IPP_TAG_OPERATION)
	o := NewIppAttributeWithOneValue(IPP_TAG_KEYWORD)
	o.SetName("requested-attributes")
	o.SetValueString("printer-name")
	g1.AddAttribute(*o)
	p := NewIppAttributeWithOneValue(IPP_TAG_ENUM)
	p.SetName("printer-type")
	p.SetValue([]byte{0})
	g1.AddAttribute(*p)
	q := NewIppAttributeWithOneValue(IPP_TAG_ENUM)
	q.SetName("printer-type-mask")
	q.SetValue([]byte{CUPS_PRINTER_CLASS})
	g1.AddAttribute(*q)
	fmt.Println(g1)
	m.AddAttributeGroup(*g1)
	fmt.Println(m)
	c.DoRequest(m)
}

func (c *CupsServer) DoRequest(m *IppMessage) {
	fi, _ := os.Create("/tmp/pooper")
	defer fi.Close()
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, int8(1))
	binary.Write(b, binary.BigEndian, int8(1))
	binary.Write(b, binary.BigEndian, uint16(CUPS_GET_PRINTERS))
	binary.Write(b, binary.BigEndian, uint32(1))
	binary.Write(b, binary.BigEndian, uint8(IPP_TAG_OPERATION))

	binary.Write(b, binary.BigEndian, uint8(IPP_TAG_CHARSET))
	binary.Write(b, binary.BigEndian, uint16(len("attributes-charset")))
	binary.Write(b, binary.BigEndian, []byte("attributes-charset"))
	binary.Write(b, binary.BigEndian, uint16(len("utf-8")))
	binary.Write(b, binary.BigEndian, []byte("utf-8"))

	binary.Write(b, binary.BigEndian, uint8(IPP_TAG_LANGUAGE))
	binary.Write(b, binary.BigEndian, uint16(len("attributes-natural-language")))
	binary.Write(b, binary.BigEndian, []byte("attributes-natural-language"))
	binary.Write(b, binary.BigEndian, uint16(len("en-us")))
	binary.Write(b, binary.BigEndian, []byte("en-us"))

	binary.Write(b, binary.BigEndian, uint8(IPP_TAG_KEYWORD))
	binary.Write(b, binary.BigEndian, uint16(len("requested-attributes")))
	binary.Write(b, binary.BigEndian, []byte("requested-attributes"))
	binary.Write(b, binary.BigEndian, uint16(len("printer-name")))
	binary.Write(b, binary.BigEndian, []byte("printer-name"))

	binary.Write(b, binary.BigEndian, uint8(IPP_TAG_ENUM))
	binary.Write(b, binary.BigEndian, uint16(len("printer-type")))
	binary.Write(b, binary.BigEndian, []byte("printer-type"))
	binary.Write(b, binary.BigEndian, uint16(4))
	binary.Write(b, binary.BigEndian, int32(0))

	binary.Write(b, binary.BigEndian, uint8(IPP_TAG_ENUM))
	binary.Write(b, binary.BigEndian, uint16(len("printer-type-mask")))
	binary.Write(b, binary.BigEndian, []byte("printer-type-mask"))
	binary.Write(b, binary.BigEndian, uint16(4))
	binary.Write(b, binary.BigEndian, uint32(1))


	binary.Write(b, binary.BigEndian, int8(3))
  //fi.Write(b.Bytes())
	resp, err := http.Post("http://localhost:631", "application/ipp", b)
	if err != nil {
		// handle error
	}
  body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		// handle error
	}
  fmt.Println("working")
  fmt.Println(body)
  fmt.Println("working")
}
