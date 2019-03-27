package index

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSV(t *testing.T) {
	hsiCSV, err := os.Open("testdata/hsio190322.csv")
	if err != nil {
		t.Fatal("TestParseCSV: open testdata/hsio190322.csv:", err)
	}
	defer hsiCSV.Close()

	hsiContracts, err := ParseCSV(hsiCSV)
	assert.Len(t, hsiContracts, 1890)

	hhiCSV, err := os.Open("testdata/hhio190322.csv")
	if err != nil {
		t.Fatal("TestParseCSV: open testdata/hhio190322.csv:", err)
	}
	defer hhiCSV.Close()

	hhiContracts, err := ParseCSV(hhiCSV)
	assert.Len(t, hhiContracts, 1784)
}
