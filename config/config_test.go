package config

import (
	"io"
	"testing"
)

// TestEval determines whether Eval produces a valid Config object from ad-hoc
// example data.
func TestEval(t *testing.T) {
	var testData string
	var err error
	
	// Define an IO pipe to send/receive data locally
	reader, writer := io.Pipe()
	defer reader.Close()
	defer writer.Close()
	
	// First test: Determine whether Eval() produces results from test data.
	
	testData = `
		{
			"debug": true,
			"domains": [{
				"name": "woodnsword.io"
			}],
			"registrar": {
				"address": "woodnsword.net",
				"port": "7776"
			}
		}
	`
	go writer.Write([]byte(testData))
	
	_, err = Eval(reader)
	if err != nil {
		t.Error(err)
	}
	
	// Second test: Determine whether Eval() panics when it receives invalid
	// test data.
	
	testData = `blah blah`
	go writer.Write([]byte(testData))
	
	_, err = Eval(reader)
	if err == nil {
		t.Error("Eval() did not produce an error when expected to.")
	}
}
