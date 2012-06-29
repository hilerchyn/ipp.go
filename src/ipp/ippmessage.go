package ipp

import (
	"bytes"
	"encoding/binary"
	"log"
)

func populateByteArray(v int64, i interface{}, s int) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, v)
	if err != nil {
		log.Fatal("binary.Write failed:", err)
	}
	q := buf.Bytes()[buf.Len()-s:]
	for n := 0; n < s; n++ {
		i.([s]byte)[n] = q[n]
	}
}

type IppAdditionalValue struct {
	valueTag    byte
	nameLength  [2]byte
	valueLength [2]byte
	value       string
}

func NewIppAdditionalValue() *IppAdditionalValue {
	var v IppAdditionalValue
	populateByteArray(int64(0x0000), v.nameLength, 2)
	return &v
}

type IppAttributeWithOneValue struct {
	valueTag    byte
	nameLength  [2]byte
	name        string
	valueLength [2]byte
	value       string
}

type IppAttribute struct {
	attributeWithOneValue IppAttributeWithOneValue
	additionalValues      []IppAdditionalValue
}

type IppAttributeGroup struct {
	beginAttributeGroupTag byte
	attribute              []IppAttribute
}

type IppMessage struct {
	VersionNumber         [2]byte
	OperationIdStatusCode [2]byte
	RequestId             [4]byte
	AttributeGroup        []IppAttributeGroup
	EndAttributeTag       byte
	Data                  []byte
}

func NewIppMessage() *IppMessage {
	var m IppMessage
	m.EndAttributeTag = 0x03
	m.setVersion()
	return &m
}

func (m *IppMessage) setVersion() {
	m.VersionNumber[0] = IPP_MAJOR_VERSION
	m.VersionNumber[1] = IPP_MINOR_VERSION
}

func (m *IppMessage) SetOperationIdStatusCode(d int16) {
	//buf := new(bytes.Buffer)
	pi := int64(d)
	//err := binary.Write(buf, binary.BigEndian, pi)
	//if err != nil {
	//	log.Fatal("binary.Write failed:", err)
	//}
	//q := buf.Bytes()[buf.Len()-2:]
	//m.OperationIdStatusCode[0] = q[0]
	//m.OperationIdStatusCode[1] = q[1]
	populateByteArray(pi, m.OperationIdStatusCode, 2)
}

func (m *IppMessage) SetRequestId(d int32) {
	buf := new(bytes.Buffer)
	pi := int64(d)
	err := binary.Write(buf, binary.BigEndian, pi)
	if err != nil {
		log.Fatal("binary.Write failed:", err)
	}
	q := buf.Bytes()[buf.Len()-4:]
	m.RequestId[0] = q[0]
	m.RequestId[1] = q[1]
	m.RequestId[2] = q[2]
	m.RequestId[3] = q[3]
}
