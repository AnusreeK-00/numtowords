package numtowords_test

import (
	"testing"

	"github.com/AnusreeK-00/numtowords"
)

func TestInvalidNumber(t *testing.T) {
	_, err := numtowords.Convert(numtowords.MaxNum + 1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}

	_, err = numtowords.Convert(-1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtowords.Convert(0)
	if err != nil {
		t.Log("convert with 0 returned an error")
		t.FailNow()
	}

	if result != "zero" {
		t.Logf("expected 'zero', received %v", result)
		t.FailNow()
	}
}

func TestUnits(t *testing.T) {
	units := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	for index, value := range units {
		result, err := numtowords.Convert(index)
		if err != nil {
			t.Logf("error while converting %v: %v", index, err)
			t.Fail()
		}

		if result != value {
			t.Logf("while converting %v, expected %v, received %v", index, value, result)
			t.Fail()
		}
	}
}

func TestTens(t *testing.T) {
	testcases := map[int]string{
		20: "twenty",
		34: "thirty four",
		48: "forty eight",
		67: "sixty seven",
		99: "ninety nine",
	}

	for num, expected := range testcases {
		result, err := numtowords.Convert(num)
		if err != nil {
			t.Logf("error while converting %v: %v", num, err)
			t.Fail()
		}

		if result != expected {
			t.Logf("while converting %v, expected %v, received %v", num, expected, result)
			t.Fail()
		}
	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		109: "one hundred and nine",
		333: "three hundred and thirty three",
		700: "seven hundred",
		340: "three hundred and forty",
	}

	for num, expected := range testcases {
		result, err := numtowords.Convert(num)
		if err != nil {
			t.Logf("error while converting %v: %v", num, err)
			t.Fail()
		}

		if result != expected {
			t.Logf("while converting %v, expected %v, received %v", num, expected, result)
			t.Fail()
		}
	}
}
