package alg

import (
	"encoding/binary"
	"log"

	"github.com/gucooing/hkrpg-go/pkg/endec"
)

// sr游戏协议编解码

/*
							《崩坏：星穹铁道》KCP协议(带*为xor加密数据)
0			1			2					4											16(字节)
+---------------------------------------------------------------------------------------+
|					sessionId(le)			|					conv(le)				|
+---------------------------------------------------------------------------------------+
|	cmd		|	frg		|		wnd			|					ts						|
+---------------------------------------------------------------------------------------+
|						sn					|					una						|
+---------------------------------------------------------------------------------------+
|						len					|	0x9d74c714*		|		cmdId*			|
+---------------------------------------------------------------------------------------+
|		headLen*		|				payloadLen*				|		head*			|
+---------------------------------------------------------------------------------------+
|								payload*						|		0xd7a152c8*		|
+---------------------------------------------------------------------------------------+
*/

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

type PackMsg struct {
	CmdId     uint16
	HeadData  []byte
	ProtoData []byte
}

func DecodeBinToPayload(data []byte, kcpMsgList *[]*PackMsg, xorKey []byte) {
	// xor解密
	if xorKey != nil {
		endec.Xor(data, xorKey)
	}
	DecodeLoop(data, kcpMsgList)
	return
}

func DecodeLoop(data []byte, kcpMsgList *[]*PackMsg) {
	// 长度太短
	if len(data) < 16 {
		log.Println("packet len less than 16 byte")
		return
	}
	// 头部幻数错误
	if binary.BigEndian.Uint32(data[:4]) != 0x9d74c714 {
		log.Println("packet head magic 0x9d74c714 error")
		return
	}

	// 协议号
	cmdId := binary.BigEndian.Uint16(data[4:6])
	// 头部长度
	headLen := binary.BigEndian.Uint16(data[6:8])
	// proto长度
	protoLen := binary.BigEndian.Uint32(data[8:12])
	// 检查长度
	packetLen := int(headLen) + int(protoLen) + 16
	if packetLen > PacketMaxLen {
		log.Println("packet len too long")
		return
	}
	if len(data) < packetLen {
		log.Println("packet len not enough")
		return
	}
	// 尾部幻数错误
	if binary.BigEndian.Uint32(data[len(data)-4:]) != 0xd7a152c8 {
		log.Println("packet tail magic 0xd7a152c8 error")
		return
	}
	// 头部数据
	headData := data[12 : 12+int(headLen)]
	// proto数据
	protoData := data[12+int(headLen) : 12+int(headLen)+int(protoLen)]
	// 返回数据
	kcpMsg := new(PackMsg)
	kcpMsg.CmdId = cmdId
	kcpMsg.HeadData = make([]byte, headLen)
	kcpMsg.ProtoData = make([]byte, protoLen)
	copy(kcpMsg.HeadData, headData)
	copy(kcpMsg.ProtoData, protoData)
	*kcpMsgList = append(*kcpMsgList, kcpMsg)
	// 有不止一个包 递归解析
	if len(data) > packetLen {
		DecodeLoop(data[packetLen:], kcpMsgList)
	}
}

func EncodePayloadToBin(kcpMsg *PackMsg, xorKey []byte) (bin []byte) {
	if kcpMsg.HeadData == nil {
		kcpMsg.HeadData = make([]byte, 0)
	}
	if kcpMsg.ProtoData == nil {
		kcpMsg.ProtoData = make([]byte, 0)
	}
	// 检查长度
	packetLen := len(kcpMsg.HeadData) + len(kcpMsg.ProtoData) + 16
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
	// 头部数据
	copy(bin[12:], kcpMsg.HeadData)
	// proto数据
	copy(bin[12+len(kcpMsg.HeadData):], kcpMsg.ProtoData)
	// 尾部幻数
	binary.BigEndian.PutUint32(bin[len(bin)-4:], 0xd7a152c8)
	// xor加密
	if xorKey != nil {
		endec.Xor(bin, xorKey)
	}
	return bin
}
