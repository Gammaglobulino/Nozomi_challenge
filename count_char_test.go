package Nozomi_challenge

import (
	"../Nozomi_challenge"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppearingMoreThanTwiceChars(t *testing.T) {
	testString := "MynameisEarlbabe"

	resultString, err := Nozomi_challenge.AppearingMoreThanTwiceChars(testString, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.EqualValues(t, "meab", resultString)

}

func TestInputEmptyString(t *testing.T) {
	testString := ""
	resultString, err := Nozomi_challenge.AppearingMoreThanTwiceChars(testString, true)
	assert.NotNil(t, err, err.Error())
	assert.EqualValues(t, "", resultString)

}

func TestSpaceIncludedAsChar(t *testing.T) {
	testString := "My name is Earl babe"

	resultString, err := Nozomi_challenge.AppearingMoreThanTwiceChars(testString, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.EqualValues(t, "m eab", resultString)
}

func TestSpaceNotIncludedAsChar(t *testing.T) {
	testString := "My name is Earl babe"

	resultString, err := Nozomi_challenge.AppearingMoreThanTwiceChars(testString, false)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.EqualValues(t, "meab", resultString)
}
