package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

const RegexPhrase string = "[0-9]{4}[A-Z]{2}"

func ParseZipCodes(zipCode string) (int, error) {
	r, err := regexp.Compile(RegexPhrase)
	if err != nil {
		return 0, err
	}

	if r.Match([]byte(zipCode)) {
		pc4, err := stripPCA4Code(zipCode)
		if err != nil {
			return 0, nil
		}
		return pc4, nil
	} else {
		return 0, fmt.Errorf("could not parse zip code %s -- must be in form %s", zipCode, RegexPhrase)
	}
}

func stripPCA4Code(zipCode string) (int, error) {
	pc4 := zipCode[0:4]
	pc4Int, err := strconv.Atoi(pc4)
	if err != nil {
		return 0, err
	}
	return pc4Int, nil
}
