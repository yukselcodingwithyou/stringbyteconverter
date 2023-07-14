package main

import (
	"fmt"
	"github.com/yukselcodingwithyou/stringbyteconverter"
)

func main() {
	converter := stringbyteconverter.New()

	trendyol := converter.ToString([]byte(`trendyol`))
	fmt.Println(trendyol)

	bytes := converter.ToBytes("trendyol")
	fmt.Println(string(bytes))

}
