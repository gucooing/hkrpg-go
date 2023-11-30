package SDK

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
	Protocol               bool          `json:"protocol"`
	QrEnabled              bool          `json:"qr_enabled"`
	LogLevel               string        `json:"log_level"`
	AnnounceURL            string        `json:"announce_url"`
	PushAliasType          int           `json:"push_alias_type"`
	DisableYsdkGuard       bool          `json:"disable_ysdk_guard"`
	EnableAnnouncePicPopup bool          `json:"enable_announce_pic_popup"`
	AppName                string        `json:"app_name"`
	QrEnabledApps          QrEnabledApps `json:"qr_enabled_apps"`
	QrAppIcons             QrAppIcons    `json:"qr_app_icons"`
	QrCloudDisplayName     string        `json:"qr_cloud_display_name"`
	EnableUserCenter       bool          `json:"enable_user_center"`
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
