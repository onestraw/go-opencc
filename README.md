# go-opencc

Go wrapper for OpenCC

## Example

```Go
package main

import (
	"fmt"

	"github.com/onestraw/go-opencc"
)

func main() {
	config := "/usr/local/share/opencc/t2s.json"
	input := "中國臺灣"

	c, _ := opencc.New(config)
	defer c.Free()
	output, _ := c.Convert(input)
	fmt.Println(output)
	// Output: 中国台湾
}
```

## Reference

- [OpenCC](https://github.com/byvoid/opencc)
- [OpenCC C API](http://byvoid.github.io/OpenCC/1.0.4/group__opencc__c__api.html)
