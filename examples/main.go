package main

import (
	"fmt"
	"github.com/yukselcodingwithyou/stringbyteconverter"
)

func main() {
	toStringConverter := stringbyteconverter.NewToStringConverter()
	toByteConverter := stringbyteconverter.NewToByteConverter()

	trendyol := toStringConverter.Convert([]byte(`yuksel`))
	fmt.Println(trendyol)

	bytes := toByteConverter.Convert("yuksel")
	fmt.Println(string(bytes))

}
