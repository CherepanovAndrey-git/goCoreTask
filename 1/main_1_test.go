package main

import (
	"1/pkg"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestStrConv(t *testing.T) {
	intDec := 42
	intOct := 052
	intHex := 0x2A
	floatVar := 3.14
	stringVar := "hello"
	boolVar := true
	complexVar := complex64(complex(1, 2))
	vars := []interface{}{intDec, intOct, intHex, floatVar, stringVar, boolVar, complexVar}
	want := "4242423.14hellotrue(1+2i)"

	got := pkg.StrConv(vars...)
	if got != want {
		t.Errorf("convertToStrings() = %v, want %v", got, want)
	}
}

func TestAddHash(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "Non-empty input", input: "Test"},
		{name: "Empty input", input: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pkg.AddSaltAndHash(tt.input)
			if len(got) != 64 {
				t.Errorf("addSaltAndHash() want: 64, got %d", len(got))
			}
		})
	}
}

func TestHashWithSalt(t *testing.T) {
	input := "qwerty"
	expected := "qwego-2024rty"

	saltedStr := pkg.AddSaltToRunes(input)
	if saltedStr != expected {
		t.Errorf("addSaltToRunes() want %s, got %s", expected, saltedStr)
	}
}

func TestHashWithoutSalt(t *testing.T) {
	input := "test"
	hashWithSalt := pkg.AddSaltAndHash(input)
	hashWithoutSalt := fmt.Sprintf("%x", sha256.Sum256([]byte(input)))

	if hashWithSalt == hashWithoutSalt {
		t.Error("Hash with salt must be different from hash without salt")
	}
}
