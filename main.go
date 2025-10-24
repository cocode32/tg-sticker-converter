package main

import "fmt"
import "io/ioutil"
import "path/filepath"
import "os"

import "github.com/cocode32/tg-sticker-converter/libtgsconverter"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: tg-sticker-converter path/to/tgs format (apng, gif, png, webp)")
		os.Exit(1)
	}
	extension := os.Args[2]
	if !libtgsconverter.SupportsExtension(extension) {
		fmt.Println("Unsupported extension: " + extension)
		os.Exit(2)
	}
	filename := os.Args[1]
	opt := libtgsconverter.NewConverterOptions()
	opt.SetExtension(extension)
	ret, err := libtgsconverter.ImportFromFile(filename, opt)
	if err != nil {
		fmt.Println("Error in tg-sticker-converter.ImportFromFile:" + err.Error())
		os.Exit(3)
	}
	tgs := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(tgs)]
	ioutil.WriteFile(name+"."+extension, ret, 0666)
}
