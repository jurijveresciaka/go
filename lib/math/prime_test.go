package math

import (
  "testing"
)

func TestIsPrime(t *testing.T) {
  type testDataStruct struct {
    number uint64
    isPrime bool
  }

  var testDataList = []testDataStruct{
    {0, false}, {1, false}, {2, true}, {3, true},
    {4, false}, {5, true}, {6, false}, {7, true},
    {8, false}, {9, false},
    {1000000, false},
  }

  for _, testData := range testDataList {
    isPrime := IsPrime(testData.number)

    if isPrime != testData.isPrime {
      t.Errorf("IsPrime(%d) == %q, want %d", testData.number, isPrime, testData.isPrime)
    }
  }
}
