package optimize

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"

	compression "github.com/nurlantulemisov/imagecompression"
)

func Optimizer(imgin string) {
	file, err := os.Open(imgin)

	if err != nil {
		log.Fatalf(err.Error())
	}

	img, err := png.Decode(file)

	if err != nil {
		log.Fatalf(err.Error())
	}
	fi, err := os.Stat(imgin)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// get the size

	fmt.Println("before: ", fi.Size())

	compressing, _ := compression.New(50)
	compressingImage := compressing.Compress(img)

	f, err := os.Create(imgin)
	if err != nil {
		log.Fatalf("error creating file: %s", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(f)

	err = png.Encode(f, compressingImage)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fif, err := os.Stat(imgin)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("after: ", fif.Size())
}

func Minifycss(infile, outfile string) {

	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	b, err := os.ReadFile(infile)
	check(err)

	mb, err := m.Bytes("text/css", b)
	check(err)

	err = os.WriteFile(outfile, mb, 0644)
	check(err)

}

func Minifyjs(infile, outfile string) {
	m := minify.New()

	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	b, err := os.ReadFile(infile)
	check(err)

	mb, err := m.Bytes("application/javascript", b)
	check(err)

	err = os.WriteFile(outfile, mb, 0644)
	check(err)

}

func Concat(files []string, fileout string) {

	var buf bytes.Buffer
	for _, file := range files {
		b, err := os.ReadFile(file)
		check(err)

		buf.Write(b)
	}

	err := os.WriteFile(fileout, buf.Bytes(), 0644)
	check(err)

}

func Getfiles(in, ext string) (error, []string) {
	var filelist []string
	files, err := os.ReadDir(in)
	check(err)
	for _, file := range files {
		if strings.Contains(file.Name(), ext) {
			filelist = append(filelist, in+"/"+file.Name())
		}
	}
	return nil, filelist
}
func GetImageFiles(in string) (error, []string) {
	var filelist []string
	files, err := os.ReadDir(in)
	check(err)
	for _, file := range files {
		if strings.Contains(file.Name(), ".png") || strings.Contains(file.Name(), ".jpg") || strings.Contains(file.Name(), ".webp") {
			filelist = append(filelist, in+"/"+file.Name())
		}
	}
	return nil, filelist
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
