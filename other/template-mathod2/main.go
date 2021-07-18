package main

import "fmt"

func main() {

	smsOTP := &sms{}

	// 使う側はテンプレート構造体を使う
	// smsOTP はiOtpインターフェースを満たす
	o := otp{iOtp: smsOTP}
	o.genAndSendOTP(4)
	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}
