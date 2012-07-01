package ipp

import (
	"encoding/binary"
)

type IppAdditionalValue struct {
	valueTag    byte
	nameLength  int16
	valueLength int16
	value       string
}

func NewIppAdditionalValue(tag byte) *IppAdditionalValue {
	var v IppAdditionalValue
	v.valueTag = tag
	v.nameLength = 0x0000
	return &v
}

func (a *IppAdditionalValue) SetValue(n string) {
	a.value = n
	a.valueLength = int16(len(n))
}

type IppAttributeWithOneValue struct {
	valueTag    byte
	nameLength  int16
	name        string
	valueLength int16
	value       []byte
}

func NewIppAttributeWithOneValue(tag byte) *IppAttributeWithOneValue {
	var o IppAttributeWithOneValue
	o.valueTag = tag
	return &o
}

func (a *IppAttributeWithOneValue) SetName(n string) {
	a.name = n
	a.nameLength = int16(len(n))
}

func (a *IppAttributeWithOneValue) SetValueString(n string) {
	a.value = []byte(n)
	a.valueLength = int16(len(a.value))
}

func (a *IppAttributeWithOneValue) SetValue(n []byte) {
  a.value = n
	a.valueLength = int16(len(a.value))
}

type IppAttribute struct {
	attributeWithOneValue IppAttributeWithOneValue
	additionalValues      []IppAdditionalValue
}

type IppAttributeGroup struct {
	beginAttributeGroupTag byte
	attribute              []IppAttribute
}

func NewIppAttributeGroup(tag byte) *IppAttributeGroup {
  g := IppAttributeGroup{tag, nil}
  return &g
}

func (g *IppAttributeGroup) AddAttribute(oneval IppAttributeWithOneValue) *IppAttribute {
	a := IppAttribute{oneval, nil}
	g.attribute = append(g.attribute, a)
	return &a
}

type IppMessage struct {
	VersionNumber         int16
	OperationIdStatusCode int16
	RequestId             int32
	AttributeGroup        []IppAttributeGroup
	EndAttributeTag       byte
	Data                  []byte
}

func NewIppMessage() *IppMessage {
	var m IppMessage
	m.EndAttributeTag = IPP_TAG_END
	m.setVersion()
	return &m
}

func (m *IppMessage) setVersion() {
	v := []byte{IPP_MAJOR_VERSION, IPP_MINOR_VERSION}
	n, _ := binary.Varint(v)
	m.VersionNumber = int16(n)
}

func (m *IppMessage) SetOperationIdStatusCode(d int16) {
	m.OperationIdStatusCode = d
}

func (m *IppMessage) SetRequestId(d int32) {
	m.RequestId = d
}

func (m *IppMessage) AddAttributeGroup(ag IppAttributeGroup) {
	m.AttributeGroup = []IppAttributeGroup{ag}
}
