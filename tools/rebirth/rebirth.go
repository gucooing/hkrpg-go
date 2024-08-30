package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

var initialKey = make(map[uint32][]byte)
var sessionKey []byte
var kcpMap map[string]*kcp.KCP

func main() {
	fileName := "./test.pcapng"
	captureHandler, err := pcap.OpenOffline(fileName)
	if err != nil {
		log.Println("Could not open pacp file", err)
		return
	}
	readKeys()
	startSniffer(captureHandler)
}

func startSniffer(captureHandler *pcap.Handle) {
	defer captureHandler.Close()

	err := captureHandler.SetBPFFilter("udp portrange 23301-23302")
	if err != nil {
		log.Println("Could not set the filter of capture")
		return
	}

	packetSource := gopacket.NewPacketSource(captureHandler, captureHandler.LinkType())
	packetSource.NoCopy = true

	kcpMap = make(map[string]*kcp.KCP)

	for packet := range packetSource.Packets() {

		// capTime := packet.Metadata().Timestamp
		data := packet.ApplicationLayer().Payload()
		udp := packet.TransportLayer().(*layers.UDP)
		fromServer := udp.SrcPort == 23301 || udp.SrcPort == 23302

		if len(data) <= 20 {
			continue
		}

		handleKcp(data, fromServer)
	}
}

func handleKcp(data []byte, fromServer bool) {
	data = reformData(data)

	conv := binary.LittleEndian.Uint64(data)
	key := strconv.Itoa(int(conv))
	if fromServer {
		key += "svr"
	} else {
		key += "cli"
	}

	if _, ok := kcpMap[key]; !ok {
		kcpInstance := kcp.NewKCP(conv, func(buf []byte, size int) {})
		kcpInstance.WndSize(256, 256)
		kcpInstance.SetMtu(1200)
		kcpMap[key] = kcpInstance
	}
	kcpInstance := kcpMap[key]
	_ = kcpInstance.Input(data, true, true)

	size := kcpInstance.PeekSize()
	for size > 0 {
		kcpBytes := make([]byte, size)
		kcpInstance.Recv(kcpBytes)
		handleProtoPacket(kcpBytes, conv)
		size = kcpInstance.PeekSize()
	}
	kcpInstance.Update()
}

func reformData(data []byte) []byte {
	i := 0
	tokenSizeTotal := 0
	var messages [][]byte
	for i < len(data) {
		convId := data[i : i+8]
		remainingHeader := data[i+8 : i+28]
		contentLen := int(binary.LittleEndian.Uint32(data[i+24 : i+28]))
		content := data[i+28 : (i + 28 + contentLen)]

		formattedMessage := make([]byte, 28+contentLen)
		copy(formattedMessage, convId)
		copy(formattedMessage[8:], remainingHeader)
		copy(formattedMessage[28:], content)
		i += 28 + contentLen
		tokenSizeTotal += 4
		messages = append(messages, formattedMessage)
	}

	return bytes.Join(messages, []byte{})
}

func handleProtoPacket(data []byte, conv uint64) {
	kcpMsgList := make([]*alg.PackMsg, 0)
	key := binary.BigEndian.Uint32(data[:8])
	key = key ^ 0x9D74C714 // Magic Start for SR
	var xorPad []byte

	if sessionKey != nil {
		xorPad = sessionKey
	} else {
		if len(initialKey[key]) == 0 {
			log.Println("Could not found initial key to decrypt", key)
			return
			// closeHandle()
		}
		xorPad = initialKey[key]
	}
	if xorPad == nil {
		log.Println("Could not found key to decrypt", key)
		return
	}
	alg.DecodeBinToPayload(data, &kcpMsgList, xorPad)

	for _, msg := range kcpMsgList {
		if msg.CmdId == cmd.PlayerGetTokenScRsp {
			handlePlayerGetTokenScRspPacket(msg.ProtoData)
		}
		playerRegisterMessage(msg, conv)
	}
}

func handlePlayerGetTokenScRspPacket(playerMsg []byte) {
	rsp := new(proto.PlayerGetTokenScRsp)
	pb.Unmarshal(playerMsg, rsp)
	if rsp.SecretKeySeed != 0 {
		sessionKey = random.CreateXorPad(rsp.SecretKeySeed, false)
	}
}

func playerRegisterMessage(msg *alg.PackMsg, conv uint64) {
	protoObj := alg.DecodePayloadToProto(msg)
	playerBinRegisterMessage(protoObj, msg.CmdId, conv)
	// data := protojson.Format(protoObj)
	// log.Printf("NAME: %s KcpMsg: \n%s\n\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(msg.CmdId), data)
}

func readKeys() {
	var initialKeyJson map[uint32]string
	file, err := ioutil.ReadFile("./data/Keys.json")
	if err != nil {
		log.Fatal("Could not load initial key @ ./data/Keys.json #1", err)
	}
	err = json.Unmarshal(file, &initialKeyJson)
	if err != nil {
		log.Fatal("Could not load initial key @ ./data/Keys.json #2", err)
	}

	for k, v := range initialKeyJson {
		decode, _ := base64.RawStdEncoding.DecodeString(v)
		initialKey[k] = decode
	}
}
