// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MsgQueryContent struct {
	_tab flatbuffers.Table
}

func GetRootAsMsgQueryContent(buf []byte, offset flatbuffers.UOffsetT) *MsgQueryContent {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MsgQueryContent{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MsgQueryContent) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MsgQueryContent) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MsgQueryContent) ID(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MsgQueryContent) IDLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MsgQueryContent) IDBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MsgQueryContent) MutateID(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *MsgQueryContent) WantRetrieve() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MsgQueryContent) MutateWantRetrieve(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func MsgQueryContentStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MsgQueryContentAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0)
}
func MsgQueryContentStartIDVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MsgQueryContentAddWantRetrieve(builder *flatbuffers.Builder, WantRetrieve bool) {
	builder.PrependBoolSlot(1, WantRetrieve, false)
}
func MsgQueryContentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}