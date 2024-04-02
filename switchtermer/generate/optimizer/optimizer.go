package optimizer

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/golangast/gentil/utility/text"
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"

	compression "github.com/nurlantulemisov/imagecompression"
)

func Optimizes() {

	d, err := domain.GetStringDomains()
	switchutility.Checklogger(err, "getting all domains for grid")

	chosendomain := switchselector.MenuInstuctions(d, 1, "purple", "purple", "Which website are you going to run the database server for?")

	//get paths of asset folders from config file
	cssin := chosendomain + GetAssetDir(chosendomain+"/assets/config/assetdirectory.yaml", "cssin")
	jsin := chosendomain + GetAssetDir(chosendomain+"/assets/config/assetdirectory.yaml", "jsin")
	cssout := chosendomain + GetAssetDir(chosendomain+"/assets/config/assetdirectory.yaml", "cssout")
	jsout := chosendomain + GetAssetDir(chosendomain+"/assets/config/assetdirectory.yaml", "jsout")
	imgin := chosendomain + GetAssetDir(chosendomain+"/assets/config/assetdirectory.yaml", "imgin")
	// get all assets file
	err, files := Getfiles(cssin, ".css")
	check(err)
	err, jsfiles := Getfiles(jsin, ".js")
	check(err)
	err, imgfiles := GetImageFiles(imgin)
	check(err)

	//concatenate all assets
	Concat(files, cssout)
	Concat(jsfiles, jsout)

	//minify all assets
	Minifycss(cssout, cssout)
	Minifyjs(jsout, jsout)

	//optimize images
	if len(imgfiles) > 1 {
		for _, imgins := range imgfiles {
			go func() {
				Optimizer(imgins)
			}()
		}
	} else {
		for _, imgins := range imgfiles {
			Optimizer(imgins)
		}
	}

	fmt.Println("assets optimized...")
}

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
func GetAssetDir(file, line string) string {
	return strings.Trim(TrimDotright(text.FindLineNReturn(file, line)), ".")
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}
