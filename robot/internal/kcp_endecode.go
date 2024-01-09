package internal

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/gucooing/hkrpg-go/pkg/endec"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	pb "google.golang.org/protobuf/proto"
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

type KcpMsg struct {
	CmdId     uint16
	HeadData  []byte
	ProtoData []byte
}

type ProtoMessage struct {
	cmdId   uint16
	message pb.Message
}

func DecodeBinToPayload(data []byte, kcpMsgList *[]*KcpMsg, xorKey []byte) {
	if xorKey != nil {
		// xor解密
		endec.Xor(data, xorKey)
	}
	DecodeLoop(data, kcpMsgList)
	return
}

func DecodeLoop(data []byte, kcpMsgList *[]*KcpMsg) {
	// 长度太短
	if len(data) < 16 {
		logger.Error("packet len less than 16 byte")
		return
	}
	// 头部幻数错误
	if data[0] != 0x9D || data[1] != 0x74 || data[2] != 0xC7 || data[3] != 0x14 {
		logger.Error("packet head magic 0x9d74c714 error")
		logger.Info("", hex.EncodeToString(data))
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
		logger.Error("packet len too long")
		return
	}
	if len(data) < packetLen {
		logger.Error("packet len not enough")
		return
	}
	// 尾部幻数错误
	if data[int(headLen)+int(protoLen)+12] != 0xD7 || data[int(headLen)+int(protoLen)+13] != 0xA1 || data[int(headLen)+int(protoLen)+14] != 0x52 || data[int(headLen)+int(protoLen)+15] != 0xC8 {
		logger.Error("packet tail magic 0xd7a152c8 error")
		return
	}
	// 头部数据
	headData := data[12 : 12+int(headLen)]
	// proto数据
	protoData := data[12+int(headLen) : 12+int(headLen)+int(protoLen)]
	// 返回数据
	kcpMsg := new(KcpMsg)
	kcpMsg.CmdId = cmdId
	kcpMsg.HeadData = headData
	kcpMsg.ProtoData = protoData
	*kcpMsgList = append(*kcpMsgList, kcpMsg)
	// 有不止一个包 递归解析
	if len(data) > packetLen {
		DecodeLoop(data[packetLen:], kcpMsgList)
	}
}

func EncodePayloadToBin(kcpMsg *KcpMsg, xorKey []byte) (bin []byte) {
	if kcpMsg.HeadData == nil {
		kcpMsg.HeadData = make([]byte, 0)
	}
	if kcpMsg.ProtoData == nil {
		kcpMsg.ProtoData = make([]byte, 0)
	}
	// 检查长度
	packetLen := len(kcpMsg.HeadData) + len(kcpMsg.ProtoData) + 16
	if packetLen > PacketMaxLen {
		logger.Error("packet len too long")
		return make([]byte, 0)
	}
	bin = make([]byte, packetLen)
	// 头部幻数
	bin[0] = 0x9D
	bin[1] = 0x74
	bin[2] = 0xC7
	bin[3] = 0x14

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
	bin[len(bin)-4] = 0xD7
	bin[len(bin)-3] = 0xA1
	bin[len(bin)-2] = 0x52
	bin[len(bin)-1] = 0xC8
	if xorKey != nil {
		// xor解密
		endec.Xor(bin, xorKey)
	}
	return bin
}
