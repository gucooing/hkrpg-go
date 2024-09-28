package alg

import (
	"encoding/binary"
	"errors"
	"log"

	"github.com/gucooing/hkrpg-go/pkg/endec"
)

// sr游戏协议编解码

/*
							《崩坏：星穹铁道》TCP协议(带*为xor加密数据)
0			1			2					4					 6  					16(字节)
+---------------------------------------------------------------------------------------+
| 				0x9d74c714					|		cmdId		|		headLen 		|
+---------------------------------------------------------------------------------------+
|				payloadLen  				|					head*					|
+---------------------------------------------------------------------------------------+
|								payload*						|		0xd7a152c8		|
+---------------------------------------------------------------------------------------+
*/

func TcpDecodeBinToPayload(bin []byte, kcpMsgList *[]*PackMsg, xorKey []byte) error {
	var err error
	TcpDecodeLoop(bin, kcpMsgList, xorKey, err)
	return err
}

func TcpDecodeLoop(bin []byte, kcpMsgList *[]*PackMsg, xorKey []byte, err error) {
	// 头部幻数错误
	if binary.BigEndian.Uint32(bin[:4]) != 0x9d74c714 {
		err = errors.New("packet head magic 0x9d74c714 error")
		return
	}
	// 协议号
	cmdId := binary.BigEndian.Uint16(bin[4:6])
	// 头部长度
	headLen := binary.BigEndian.Uint16(bin[6:8])
	// proto长度
	protoLen := binary.BigEndian.Uint32(bin[8:12])
	// 检查长度
	packetLen := int(headLen) + int(protoLen) + 16
	if packetLen > PacketMaxLen {
		err = errors.New("packet len too long")
		return
	}
	if len(bin) < packetLen {
		return
	}
	// 尾部幻数错误
	if binary.BigEndian.Uint32(bin[packetLen-4:packetLen]) != 0xd7a152c8 {
		err = errors.New("packet tail magic 0xd7a152c8 error")
		return
	}
	data := bin[12 : 12+int(headLen)+int(protoLen)]
	if xorKey != nil {
		endec.Xor(data, xorKey)
	}
	// 头部数据
	headData := data[int(headLen):]
	// proto数据
	protoData := data[int(headLen) : int(headLen)+int(protoLen)]
	// 返回数据
	kcpMsg := new(PackMsg)
	kcpMsg.CmdId = cmdId
	kcpMsg.HeadData = make([]byte, headLen)
	kcpMsg.ProtoData = make([]byte, protoLen)
	kcpMsg.Length = packetLen
	copy(kcpMsg.HeadData, headData)
	copy(kcpMsg.ProtoData, protoData)
	*kcpMsgList = append(*kcpMsgList, kcpMsg)

	// 有不止一个包 递归解析
	if len(bin) > packetLen+16 {
		TcpDecodeLoop(bin[packetLen:], kcpMsgList, xorKey, err)
	}
}

func TcpEncodePayloadToBin(kcpMsg *PackMsg, xorKey []byte) (bin []byte) {
	if kcpMsg.HeadData == nil {
		kcpMsg.HeadData = make([]byte, 0)
	}
	if kcpMsg.ProtoData == nil {
		kcpMsg.ProtoData = make([]byte, 0)
	}
	// 头部长度
	headLen := len(kcpMsg.HeadData)
	// proto长度
	protoLen := len(kcpMsg.ProtoData)
	// 检查长度
	packetLen := headLen + protoLen + 16
	if packetLen > PacketMaxLen {
		log.Println("packet len too long")
		return make([]byte, 0)
	}
	bin = make([]byte, packetLen)
	// 头部幻数
	binary.BigEndian.PutUint32(bin[:4], 0x9d74c714)

	// 协议号
	binary.BigEndian.PutUint16(bin[4:6], kcpMsg.CmdId)
	// 头部长度
	binary.BigEndian.PutUint16(bin[6:8], uint16(len(kcpMsg.HeadData)))
	// proto长度
	binary.BigEndian.PutUint32(bin[8:12], uint32(len(kcpMsg.ProtoData)))
	data := append(kcpMsg.HeadData, kcpMsg.ProtoData...)
	// xor加密
	if xorKey != nil {
		endec.Xor(data, xorKey)
	}
	// 头部数据
	copy(bin[12:], data[:headLen])
	// proto数据
	copy(bin[12+headLen:], data[headLen:protoLen])
	// 尾部幻数
	binary.BigEndian.PutUint32(bin[len(bin)-4:], 0xd7a152c8)

	return bin
}
