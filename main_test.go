package main

import (
	"bytes"
	"testing"
)

const testFullResult = `
├───project
│ ├───file.txt (19b)
│ └───gopher.png (70372b)
├───static
│ ├───a_lorem
│ │ ├───ipsum
│ │ │ └───gopher.png (70372b)
│ │ ├───dolor.txt (empty)
│ │ └───gopher.png (70372b)
│ ├───css
│ │ └───body.css (28b)
│ ├───html
│ │ └───index.html (57b)
│ ├───js
│ │ └───site.js (10b)
│ ├───z_lorem
│ │ ├───ipsum
│ │ │ └───gopher.png (70372b)
│ │ ├───dolor.txt (empty)
│ │ └───gopher.png (70372b)
│ └───empty.txt (empty)
├───zline
│ ├───lorem
│ │ ├───ipsum
│ │ │ └───gopher.png (70372b)
│ │ ├───dolor.txt (empty)
│ │ └───gopher.png (70372b)
│ └───empty.txt (empty)
└───zzfile.txt (empty)
`

func TestTreeFull(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", nil)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testFullResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testFullResult)
	}
}