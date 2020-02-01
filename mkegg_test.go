package main_test

import (
	"github.com/johan-bolmsjo/mkegg/internal/eggs"
	"io/ioutil"
	"testing"
)

func Test_ReadData(t *testing.T) {
	dataReader := eggs.TXT_HelloWorld()

	data, err := ioutil.ReadAll(dataReader)
	if err != nil {
		t.Fatalf("failed to read data from egg: %s", err)
	}

	testDataFilename := "data/hello_world.txt"
	expectedData, err := ioutil.ReadFile(testDataFilename)
	if err != nil {
		t.Fatalf("failed to read test data from %q: %s", testDataFilename, err)
	}
	if string(data) != expectedData {
		t.Fatalf("unexpected egg data %q; want %q", string(data), expectedData)
	}
}
