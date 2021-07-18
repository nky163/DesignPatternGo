package main

// One Time Password(OTP)
// SMS OTP, EMAIL OTP があるが、OTPの処理自体は同じ
// ・Cenerate a random n digit number
// ・Save this number in the cache for later verification
// ・Perpare the content
// ・Send the notification
// ・Publish the metrics

// iOtp is interface
// OTPタイプそれぞれで実装するメソッドセット
type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishNetrics()
}

// テンプレートクラス
type otp struct {
	iOtp iOtp // インターフェースを持つ
}

// テンプレートメソッド
// インターフェースを使って実装する
// インターフェースの実装は継承した構造体で行う
func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishNetrics()
	return nil
}
