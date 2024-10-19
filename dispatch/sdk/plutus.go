package sdk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/constant"
)

func (s *Server) listPayPlat(c *gin.Context) {
	payPlats := make([]*constant.PayPlat, 0)
	payPlats = append(payPlats, &constant.PayPlat{
		IsRecommend: false,
		MarketWord:  "",
		MerchantInfo: struct {
			MerchantId        string `json:"merchant_id"`
			MiniProgramId     string `json:"mini_program_id"`
			OfficialAccountId string `json:"official_account_id"`
		}{
			MerchantId:        "cjsuajvnn00qi4dnl02g",
			MiniProgramId:     "",
			OfficialAccountId: "wx4b7e11b9e935040b",
		},
		PayPlat:                  "wechatpay",
		PayType:                  "wechatpay",
		PayTypeIconUrl:           "https://sdk-webstatic.mihoyo.com/sdk-payment-upload/2023/01/11/872a682b635c801b568214479464092b_4558682075688947537.png",
		PayTypeName:              "andy pg",
		MarketId:                 "",
		UseMarket:                false,
		UseOfficialQrcode:        false,
		RealCashierId:            "4b1ca92d-acc0-4a93-afb2-b3e3b5c0ac65",
		PromotionText:            "",
		Desc:                     "",
		Fold:                     false,
		DistinctQrCodePayInfo:    nil,
		PayAbility:               "jsapi_pay",
		IsPromotionTextClickable: false,
		ChannelPromotion:         nil,
		OpenId:                   "",
	})
	resp := constant.ListPayPlatResp{
		Retcode: 0,
		Message: "OK",
		Data: &constant.ListPayPlatData{
			IsNewUser:              false,
			PayPlats:               payPlats,
			SelectedPayType:        "andy pg",
			FoldOther:              false,
			FoldOtherFinal:         false,
			ExpandMixedQrcode:      true,
			ExpandMixedQrcodeFinal: true,
			UseOfficialQrcode:      false,
			Reports:                nil,
		},
	}
	c.JSON(http.StatusOK, resp)
}
