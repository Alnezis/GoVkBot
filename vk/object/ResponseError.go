package object

type ResponseError struct {
	ErrorCode  int            `json:"error_code"`
	Error      string         `json:"error_msg"`
	Params     []RequestParam `json:"request_params"`
	CaptchaSid string         `json:"captcha_sid,omitempty"`
	CaptchaImg string         `json:"captcha_img,omitempty"`
}
