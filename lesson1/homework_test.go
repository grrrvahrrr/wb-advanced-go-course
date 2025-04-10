package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	//для меня это новое поэтому подробно, биты на практике двигать не приходилось =)

	return (((number & 0x000000FF) << 24) | //берем последний байт числа и сдвигаем вправо на 24 бита до позиции первого байта
		((number & 0x0000FF00) << 8) | // берем предпоследний байт и сдвигаем на 8 бит на позицию второго байта нового числа
		((number & 0x00FF0000) >> 8) | // берем третий байт и сдвигаем враво на 8 бит на позицию предпоследнего байта
		((number & 0xFF000000) >> 24)) // берем четверты байт и сдвигаем на позицию последнего байта в числе
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
