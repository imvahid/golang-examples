package ncode

import "errors"

var cityCode map[string]string = map[string]string{
	"011": "تهران جنوب",
	"020": "تهران شرق",
}

func getCityName(code string) (string, error) {
	val, ok := cityCode[code]
	if ok {
		return val, nil
	}
	return "", errors.New("city doesn't exist")
}
