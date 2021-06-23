package controller

import (
	"butuhdonorplasma/models"
	"errors"
	"math"
	"math/rand"
)

var captchaAnswer map[int64]string = map[int64]string{
	1:  "21421",
	2:  "32314",
	3:  "46653",
	4:  "29314",
	5:  "62021",
	6:  "43574",
	7:  "01295",
	8:  "34598",
	9:  "21786",
	10: "89723",
}
var captchaImgURL map[int64]string = map[int64]string{
	1:  "captcha/1.png",
	2:  "captcha/2.png",
	3:  "captcha/3.png",
	4:  "captcha/4.png",
	5:  "captcha/5.png",
	6:  "captcha/6.png",
	7:  "captcha/7.png",
	8:  "captcha/8.png",
	9:  "captcha/9.png",
	10: "captcha/10.png",
}

func GetCaptcha() models.Captcha {
	randomID := int64(math.Round(1 + rand.Float64()*(10-1)))

	imgurl := captchaImgURL[randomID]

	return models.Captcha{
		ImgURL: imgurl,
		ID:     randomID,
	}
}

func CheckCaptcha(captchaID int64, captcha string) error {
	answ := captchaAnswer[captchaID]
	if answ == captcha {
		return nil
	}
	return errors.New("captcha not match")
}
