package ncode

import (
	"errors"
	"fmt"
	"strconv"
)

var placeValue = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

const residualConst = 11
const controlDigitIndex = 9

type NationalID struct {
	stringFormat string
	digitFormat  [10]int
	cityCode     string
	IsValid      bool
}

func NewNationalID(id string) (*NationalID, error) {
	if id == "" {
		return nil, errors.New("national id is empty")
	}

	var err error
	nid := &NationalID{stringFormat: id}
	nid.IsValid, err = nid.validation()
	if err != nil {
		return nil, err
	}
	return nid, nil
}

func (id *NationalID) GetCity() (string, error) {
	return getCityName(id.cityCode)
}

func (id *NationalID) validation() (bool, error) {
	if len(id.stringFormat) == 10 {
		for index, val := range id.stringFormat {
			digit, err := strconv.Atoi(string(val))
			if err != nil {
				return false, err
			}
			id.digitFormat[index] = digit
		}

		ret, err := id.isCheckSumValid()
		if err != nil || ret == false {
			return false, err
		} else {
			id.IsValid = true
			id.cityCode = id.stringFormat[0:3]
			return true, nil
		}
	}
	return false, errors.New("validation error")
}

func (id *NationalID) isCheckSumValid() (bool, error) {
	var (
		sum          int
		resid        int
		controlDigit int
	)
	for index := 0; index < 9; index++ {
		sum = sum + id.digitFormat[index]*placeValue[index]
	}

	resid = sum % residualConst

	controlDigit = id.digitFormat[controlDigitIndex]

	if resid < 2 {
		if resid == controlDigit {
			return true, nil
		}
	} else {
		if controlDigit == (11 - resid) {
			return true, nil
		}
	}

	return false, fmt.Errorf("checksum is invalid, residual: %d, current controlDigit: %d, correct controlDigit: %d", resid, controlDigit, residualConst-resid)
}
