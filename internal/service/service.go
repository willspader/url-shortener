package service

import (
	"url-shortener/internal/repository"
)

type Service struct {
	repository *repository.Repository
}

func (service Service) LongToShort() string {
	var id int64 = service.repository.NewId()

	var shortUrl string

	for id > 0 {
		var remainder int64 = id % 62

		var ascii byte = fromDecimalToAscii(byte(remainder))

		shortUrl = string(ascii) + shortUrl

		id = id / 62
	}

	return shortUrl
}

/*func (service Service) convertToDecimal(shortUrl string) int64 {
	var sum int64 = 0
	var reversed []byte = reverseString(shortUrl)

	for i := 0; i < len(reversed); i++ {
		var ascii byte = toDecimalFromAscii(reversed[i])

		sum += int64(ascii) * int64((math.Pow(float64(62), float64(i))))
	}

	return sum
}*/

/*func toDecimalFromAscii(char byte) byte {
	if char >= 48 && char <= 57 {
		return char - 48
	}

	if char >= 65 && char <= 90 {
		return char - 55
	}

	return char - 61
}*/

func fromDecimalToAscii(digit byte) byte {
	if digit < 10 {
		return digit + 48
	}

	if digit >= 10 && digit <= 35 {
		return digit + 55
	}

	return digit + 61
}

// reverse a ASCII string
func reverseString(original string) []byte {
	var chars []byte = []byte(original)

	for i, j := 0, len(chars); i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return chars
}

func New(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
