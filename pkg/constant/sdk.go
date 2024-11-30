package constant

type UrlList struct {
	Version        string
	MdkResVersion  string
	IfixVersion    string
	IfixUrl        string
	LuaUrl         string
	ExResourceUrl  string
	AssetBundleUrl string
	Time           string
}

type Login struct {
	Data    *LoginData `json:"data"`
	Message string     `json:"message"`
	Retcode any        `json:"retcode"`
}
type LoginData struct {
	Account             *LoginAccount `json:"account"`
	DeviceGrantRequired bool          `json:"device_grant_required"`
	SafeMoblieRequired  bool          `json:"safe_moblie_required"`
	RealpersonRequired  bool          `json:"realperson_required"`
	ReactivateRequired  bool          `json:"reactivate_required"`
	RealnameOperation   string        `json:"realname_operation"`
}
type LoginAccount struct {
	UID               string `json:"uid"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	IsEmailVerify     string `json:"is_email_verify"`
	Realname          string `json:"realname"`
	IdentityCard      string `json:"identity_card"`
	Token             string `json:"token"`
	SafeMobile        string `json:"safe_mobile"`
	FacebookName      string `json:"facebook_name"`
	GoogleName        string `json:"google_name"`
	TwitterName       string `json:"twitter_name"`
	GameCenterName    string `json:"game_center_name"`
	AppleName         string `json:"apple_name"`
	SonyName          string `json:"sony_name"`
	TapName           string `json:"tap_name"`
	Country           string `json:"country"`
	ReactivateTicket  string `json:"reactivate_ticket"`
	AreaCode          string `json:"area_code"`
	DeviceGrantTicket string `json:"device_grant_ticket"`
	SteamName         string `json:"steam_name"`
	UnmaskedEmail     string `json:"unmasked_email"`
	UnmaskedEmailType int    `json:"unmasked_email_type"`
}

type GranterApiGetConfig struct {
	Data    *GranterApiGetConfigData `json:"data"`
	Message string                   `json:"message"`
	Retcode any                      `json:"retcode"`
}
type GranterApiGetConfigData struct {
	Protocol                bool           `json:"protocol"`
	QrEnabled               bool           `json:"qr_enabled"`
	LogLevel                string         `json:"log_level"`
	AnnounceURL             string         `json:"announce_url"`
	PushAliasType           int            `json:"push_alias_type"`
	DisableYsdkGuard        bool           `json:"disable_ysdk_guard"`
	EnableAnnouncePicPopup  bool           `json:"enable_announce_pic_popup"`
	AppName                 string         `json:"app_name"`
	QrEnabledApps           *QrEnabledApps `json:"qr_enabled_apps"`
	QrAppIcons              *QrAppIcons    `json:"qr_app_icons"`
	QrCloudDisplayName      string         `json:"qr_cloud_display_name"`
	EnableUserCenter        bool           `json:"enable_user_center"`
	FunctionalSwitchConfigs struct{}       `json:"functional_switch_configs"`
}
type QrEnabledApps struct {
	Bbs   bool `json:"bbs"`
	Cloud bool `json:"cloud"`
}
type QrAppIcons struct {
	App   string `json:"app"`
	Bbs   string `json:"bbs"`
	Cloud string `json:"cloud"`
}

type Check struct {
	Data    *CheckData `json:"data"`
	Message string     `json:"message"`
	Retcode any        `json:"retcode"`
}
type CheckData struct {
	Id      string   `json:"id"`
	Action  string   `json:"action"`
	Geetest *Geetest `json:"geetest"`
}
type Geetest struct {
	Challenge  string `json:"challenge"`
	Gt         string `json:"gt"`
	NewCaptcha string `json:"new_captcha"`
	Success    string `json:"success"`
}
type RiskyApiCheck struct {
	ActionType string `json:"action_type"`
	APIName    string `json:"api_name"`
	Username   string `json:"username"`
}

type LoginAccountRequestJson struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsCrypto bool   `json:"is_crypto"`
}

type ComboTokenReq struct {
	AppID     any    `json:"app_id"`
	ChannelID any    `json:"channel_id"`
	Data      string `json:"data"`
	Device    string `json:"device"`
	Sign      string `json:"sign"`
}
type ComboTokenReqLoginTokenData struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
	Guest bool   `json:"guest"`
}

type ComboTokenRsp struct {
	Retcode int                     `json:"retcode"`
	Message string                  `json:"message"`
	Data    *ComboTokenRspLoginData `json:"data"`
}
type ComboTokenRspLoginData struct {
	ComboID       string      `json:"combo_id"`
	OpenID        string      `json:"open_id"`
	ComboToken    string      `json:"combo_token"`
	Data          string      `json:"data"`
	Heartbeat     bool        `json:"heartbeat"`
	AccountType   int         `json:"account_type"`
	FatigueRemind interface{} `json:"fatigue_remind"`
}

type LoginTokenRequest struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

type CreateOrderReq struct {
	Who struct {
		Account string `json:"account"`
		Token   string `json:"token"`
	} `json:"who"`
	Order struct {
		ChannelId  string `json:"channel_id"`
		Account    string `json:"account"`
		PayPlat    string `json:"pay_plat"`
		Country    string `json:"country"`
		Currency   string `json:"currency"`
		Amount     int    `json:"amount"`
		Game       string `json:"game"`
		Region     string `json:"region"`
		Uid        string `json:"uid"`
		GoodsId    string `json:"goods_id"`
		GoodsTitle string `json:"goods_title"`
		GoodsNum   string `json:"goods_num"`
		ClientType int    `json:"client_type"`
		Device     string `json:"device"`
		PriceTier  string `json:"price_tier"`
	} `json:"order"`
	Sign             string `json:"sign"`
	DoNotNoticeAgain bool   `json:"do_not_notice_again"`
}

type CreateOrderResp struct {
	Retcode int             `json:"retcode"`
	Message string          `json:"message"`
	Data    CreateOrderData `json:"data"`
}

type CreateOrderData struct {
	GoodsId              string `json:"goods_id"`
	OrderNo              string `json:"order_no"`
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	RedirectUrl          string `json:"redirect_url"`
	ForeignSerial        string `json:"foreign_serial"`
	EncodeOrder          string `json:"encode_order"`
	Account              string `json:"account"`
	CreateTime           string `json:"create_time"`
	ExtInfo              string `json:"ext_info"`
	Balance              string `json:"balance"`
	Method               string `json:"method"`
	Action               string `json:"action"`
	SessionToken         string `json:"session_token"`
	DisplayDontShowAgain bool   `json:"display_dont_show_again"`
	NoticeAmount         int    `json:"notice_amount"`
	Cluster              string `json:"cluster"`
}

type ListPayPlatResp struct {
	Retcode int              `json:"retcode"`
	Message string           `json:"message"`
	Data    *ListPayPlatData `json:"data"`
}

type ListPayPlatData struct {
	IsNewUser              bool        `json:"is_new_user"`
	PayPlats               []*PayPlat  `json:"pay_plats"`
	SelectedPayType        string      `json:"selected_pay_type"`
	FoldOther              bool        `json:"fold_other"`
	FoldOtherFinal         bool        `json:"fold_other_final"`
	ExpandMixedQrcode      bool        `json:"expand_mixed_qrcode"`
	ExpandMixedQrcodeFinal bool        `json:"expand_mixed_qrcode_final"`
	UseOfficialQrcode      bool        `json:"use_official_qrcode"`
	Reports                interface{} `json:"reports"`
}

type PayPlat struct {
	IsRecommend  bool   `json:"is_recommend"`
	MarketWord   string `json:"market_word"`
	MerchantInfo struct {
		MerchantId        string `json:"merchant_id"`
		MiniProgramId     string `json:"mini_program_id"`
		OfficialAccountId string `json:"official_account_id"`
	} `json:"merchant_info"`
	PayPlat                  string      `json:"pay_plat"`
	PayType                  string      `json:"pay_type"`
	PayTypeIconUrl           string      `json:"pay_type_icon_url"`
	PayTypeName              string      `json:"pay_type_name"`
	MarketId                 string      `json:"market_id"`
	UseMarket                bool        `json:"use_market"`
	UseOfficialQrcode        bool        `json:"use_official_qrcode"`
	RealCashierId            string      `json:"real_cashier_id"`
	PromotionText            string      `json:"promotion_text"`
	Desc                     string      `json:"desc"`
	Fold                     bool        `json:"fold"`
	DistinctQrCodePayInfo    interface{} `json:"distinct_qr_code_pay_info"`
	PayAbility               string      `json:"pay_ability"`
	IsPromotionTextClickable bool        `json:"is_promotion_text_clickable"`
	ChannelPromotion         interface{} `json:"channel_promotion"`
	OpenId                   string      `json:"open_id"`
}

type CurrencyAndCountryByIpResp struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Currency         string `json:"currency"`
		Country          string `json:"country"`
		PriceTierVersion string `json:"price_tier_version"`
	} `json:"data"`
}

type GateGetPlayerComboToken struct {
	Retcode    int    `json:"retcode"`
	AccountId  string `json:"account_id"`
	ComboToken string `json:"combo_token"`
}
