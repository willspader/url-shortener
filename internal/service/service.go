package service

import (
	"math"
	"url-shortener/internal/repository"
)

type Service struct {
	repository *repository.Repository
}

func (service Service) LongToShort() string {
	var id int64 = service.repository.NewRecord()
	println(id)
	// TODO -> convert id to base62 representation
	return ""
}

func (service Service) convertToBase(number int64, base int8) {

}

// Convert a [BASE 62] string to its decimal representation [uint64]
func (service Service) convertToDecimal(shortUrl string) int64 {
	var sum int64 = 0
	var reversed []byte = reverseString(shortUrl)

	for i := 0; i < len(reversed); i++ {
		var ascii byte = toDecimalFromAscii(reversed[i])

		sum += int64(ascii) * int64((math.Pow(float64(62), float64(i))))
	}

	return sum
}

// find the integer representation from a byte value
func toDecimalFromAscii(char byte) byte {
	if char >= 48 && char <= 57 {
		return char - 48
	}

	if char >= 65 && char <= 90 {
		return char - 55
	}

	return char - 61
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
