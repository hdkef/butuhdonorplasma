package models

type Captcha struct {
	ImgURL string `json:"imgurl"`
	ID     int64  `json:"id"`
}

type CaptchaCheck struct {
	ID   int64  `json:"id"`
	Answ string `json:"answ"`
}
