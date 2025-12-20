package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() (string, error) {
	var digits = []rune("0123456789")
	otp := make([]rune, 6)

	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp), nil
}


