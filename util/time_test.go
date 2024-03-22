package util

import "testing"

func TestStringTotime(t *testing.T) {

	convertedTime := StringToTime("2019-02-12T18:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que o ano seja 2019")
	}
}
