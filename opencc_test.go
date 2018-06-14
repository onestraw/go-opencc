package opencc_test

import (
	"fmt"

	"github.com/onestraw/go-opencc"
)

func ExampleS2T() {
	config := "/usr/local/share/opencc/s2t.json"
	input := "中华文化"
	c, err := opencc.New(config)
	if err != nil {
		fmt.Printf("New err=%v", err)
		return
	}
	defer c.Free()
	output, err := c.Convert(input)
	if err != nil {
		fmt.Printf("Convert err=%v", err)
		return
	}
	fmt.Println(output)
	// Output: 中華文化
}

func ExampleT2S() {
	config := "/usr/local/share/opencc/t2s.json"
	input := "中國臺灣"
	c, _ := opencc.New(config)
	defer c.Free()
	output, _ := c.Convert(input)
	fmt.Println(output)
	// Output: 中国台湾
}
