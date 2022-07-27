package response

type CaptchaResp struct {
	CaptchaID     string `json:"captchaID"`
	PicPath       string `json:"picPath"`
	CaptchaLength int `json:"captchaLength"`
}
