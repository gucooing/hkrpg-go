package text

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	"github.com/hjson/hjson-go/v4"
)

func getTextMapName() map[spb.LanguageType]string {
	textMapName := map[spb.LanguageType]string{
		spb.LanguageType_LANGUAGE_SC: "TextMap/TextMapCN.json",
		spb.LanguageType_LANGUAGE_TC: "TextMap/TextMapCHT.json",
		spb.LanguageType_LANGUAGE_EN: "TextMap/TextMapEN.json",
	}
	return textMapName
}

type TextMap struct {
	language   spb.LanguageType //  语言
	dataPrefix string
	textMap    map[spb.LanguageType]map[int32]string
}

var TEXT *TextMap

func NewTextMap(language string, dataPrefix string) {
	if dataPrefix == "" {
		dataPrefix = "./data/"
	}

	l := GetLanguageTypeByS(language)
	if l == spb.LanguageType_LANGUAGE_NONE {
		l = spb.LanguageType_LANGUAGE_SC
	}

	TEXT = &TextMap{
		language:   l,
		dataPrefix: dataPrefix,
		textMap:    make(map[spb.LanguageType]map[int32]string),
	}
	TEXT.loadTextMap()
	if TEXT.textMap[TEXT.language] == nil {
		if TEXT.textMap[spb.LanguageType_LANGUAGE_SC] != nil {
			TEXT.language = spb.LanguageType_LANGUAGE_SC
		}
	}
}

func GetLanguageTypeByS(s string) spb.LanguageType {
	switch s {
	case "cn":
		return spb.LanguageType_LANGUAGE_SC // 简体中文
	case "cht":
		return spb.LanguageType_LANGUAGE_TC // 繁体中文
	case "en":
		return spb.LanguageType_LANGUAGE_EN // 英文
	default:
		return spb.LanguageType_LANGUAGE_NONE
	}
}

func getText() *TextMap {
	if TEXT == nil {
		NewTextMap("", "")
	}
	return TEXT
}

func GetText(v int32) string {
	text := getText()
	if text.textMap == nil {
		text.loadTextMap()
	}
	k, ok := text.textMap[text.language][v]
	if !ok {
		return fmt.Sprintf("未找到语言:%v", v)
	}
	return k
}

func GetTextByL(l spb.LanguageType, v int32) string {
	text := getText()
	if text.textMap == nil {
		text.loadTextMap()
	}
	k, ok := text.textMap[l][v]
	if !ok {
		return fmt.Sprintf("未找到语言:%v", v)
	}
	return k
}

func (t *TextMap) loadTextMap() {
	if t.textMap == nil {
		t.textMap = make(map[spb.LanguageType]map[int32]string)
	}
	for l, file := range getTextMapName() {
		if t.textMap[l] == nil {
			t.textMap[l] = make(map[int32]string)
		}
		tm := make(map[int32]string)
		data, err := os.ReadFile(t.dataPrefix + file)
		if err != nil {
			logger.Error("load %s error:%s", file, err.Error())
			continue
		}
		err = hjson.Unmarshal(data, &tm)
		if err != nil {
			logger.Error("load %sn error:%s", file, err.Error())
			continue
		}
		logger.Info("load %s success", file)
		t.textMap[l] = tm
	}
}
