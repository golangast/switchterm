package switchutility

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/golangast/gentil/utility/text"
	"github.com/golangast/switchterm/loggers"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
)

func UP(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	if atline >= 1 {
		atline--
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil
}

func Down(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	linecount := len(list)
	if atline <= linecount-2 {
		atline++
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil

}
func Right(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	linecount := len(list)
	rows := (len(list) + cols - 1) / cols
	if atline <= linecount-rows {
		atline = atline + rows
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil
}

func Left(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	rows := (len(list) + cols - 1) / cols
	if atline >= rows {
		atline = atline - rows
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)

	return atline, false, nil
}
func ClearDirections() {
	fmt.Print("\033[H\033[2J")
	colortermer.ColorizeCol("purple", "purple", "(q-quit) - (c-multiselection) - (r-remove) - (enter-select/execute) - (u-update tag) - down/up/left/right")
	fmt.Println("\n")

}

func Directions() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
}

func PrintColumns(cols, atline int, list, chosen []string, background, foreground string) {
	rows := (len(list) + cols - 1) / cols

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(list) {
				break // This means the last column is not "full"
			}

			if i == atline {

				colortermer.ColorizeCol(background, foreground, list[atline])

			} else {
				if slices.Contains(chosen, list[i]) {
					colortermer.ColorizeCol("purple", foreground, list[i])

				} else {
					fmt.Printf("%-11s%s", list[i], " ")
				}
			}
		}
		fmt.Println() //yes this needs to be here for padding
	}
}

func PrintColumnsWChosen(cols, atline int, list []string, background, foreground string) {
	rows := (len(list) + cols - 1) / cols

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(list) {
				break // This means the last column is not "full"
			}

			if i == atline {
				colortermer.ColorizeCol(background, foreground, list[atline])

			} else {
				fmt.Printf("%-11s%s", list[i], " ")

			}

		}
		fmt.Println() //yes this needs to be here for padding

	}
}

func RemoveItemWChosen(remove bool, list, chosen []string) bool {
	// if remove is true then remove the chosen
	if remove {

		//remove chosen from list
		for _, item := range chosen {
			index := slices.Index(list, item)
			if index > -1 {
				tags.DeleteTag(item)
				fmt.Println("removed: ", item)

			}

			RemoveCMD(item)
		}
	}

	return false

}
func RemoveCMD(cmd string) {

	if err := UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `"github.com/golangast/switchterm/cmd/`+cmd+`"`, "", ""); err != nil {
		Checklogger(err, "trying to update text in cmdrunner.go")
	}

	if err := UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"()", "", ""); err != nil {
		Checklogger(err, "trying to remove call")
	}

	if err := os.RemoveAll("./cmd/" + cmd); err != nil {
		Checklogger(err, "trying to remove folder")
	}

}
func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func BashFileOut(file string, commands []string) (string, string, error) {
	logger := loggers.CreateLogger()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmds := strings.Join(commands, " ")
	f, err := os.Open(file)
	if err != nil {
		logger.Error(
			"opening bash file ",
			slog.String("error: "+file, err.Error()),
		)
	}
	defer f.Close()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/dev/stdin", cmds)
	} else {
		cmd = exec.Command("bash", "/dev/stdin", cmds)
	}
	cmd.Stdin = f
	cmd.Stdout = os.Stdout

	if cmd.Err != nil {
		logger.Error(
			"trying to run bash command",
			slog.String("error: ", err.Error()),
		)
	}

	err = cmd.Run()
	if err != nil {
		logger.Error(
			"running bash command",
			slog.String("error: ", err.Error()),
		)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return stdout.String(), stderr.String(), nil
}

func RunApps(chosen []string) {
	logger := loggers.CreateLogger()

	t, err := tags.GetNoteByChosen(chosen)
	if err != nil {
		logger.Error(
			"getting all tags by note",
			slog.String("error: ", err.Error()),
		)
	}
	var argsset []string
	for _, v := range t {
		colortermer.ColorizeOutPut("dpurple", "purple", "What are the Args for "+v.CMD+"? (please use spaces) ||| Note: "+v.Note+"\n")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		args := scannerdesc.Text()
		strArrayOne := strings.Split(args, " ")
		argsset = append(argsset, strArrayOne...)

		out, outerr, err := BashFileOut("."+v.Bashdir+"/"+v.Bashfile+".bash", argsset)
		if err != nil {
			logger.Error(
				"trying to execute bash command",
				slog.String("error: ", err.Error()),
			)
		}

		if out != "" {
			fmt.Println("\n")
			colortermer.ColorizeOutPut("dpurple", "bpurple", "output: "+out)
			fmt.Println("\n")
		}
		if outerr != "" {
			colortermer.ColorizeOutPut("dpurple", "bpurple", "output: "+outerr)
		}
	}

}

func Execute(file, types string, commands []string) (*exec.Cmd, error) {

	f, e := os.Open(file)
	if e != nil {
		log.Fatal(e.Error())
	}
	defer f.Close()
	cmd := exec.Command(types, commands...)
	cmd.Stdin = f
	cmd.Stdout = os.Stdout

	if cmd.Err != nil {
		return cmd, cmd.Err
	}

	if e := cmd.Run(); e != nil {
		log.Fatal(e.Error())
	}

	return cmd, nil
}
func Delete[T comparable](collection []T, el T) []T {
	idx := Find(collection, el)
	if idx > -1 {
		return slices.Delete(collection, idx, idx+1)
	}
	return collection
}

func Find[T comparable](collection []T, el T) int {
	for i := range collection {
		if collection[i] == el {
			return i
		}
	}
	return -1
}
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GetPropDatatype(prop string) []string {
	var property []string
	var types []string
	var field []string
	var strright string
	s := strings.Split(prop, " ")

	for _, ss := range s {
		sss := strings.Replace(ss, "\"", "", -1)
		property = append(property, TrimDot(sss))
		strright = strings.Replace(TrimDotright(sss), ".", "", 1)
		types = append(types, strright)
	}

	for a, str1_word := range property {
		for b, str2_word := range types {
			if a == b {
				field = append(field, str1_word+" "+str2_word)
			}
		}
	}
	return field
}

func GetField(fields string) ([]string, []string) {
	var field []string
	var property []string
	s := strings.Split(fields, " ")
	for i, ss := range s {
		if i%2 == 0 {
			//get even
			field = append(field, TrimDotright(ss))
		} else {
			property = append(property, TrimDotright(ss))

		}

	}
	return field, property
}
func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
func WritetemplateStruct(temp string, f *os.File, d Data) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}

func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}

type Data struct {
	Name      string
	MapData   map[string]string
	Fields    string
	MapFields map[string]string
	Github    string
	Domain    string
}

func InputScan(s string) string {
	fmt.Println(s)
	scannerdesc := bufio.NewScanner(os.Stdin)
	tr := scannerdesc.Scan()
	if tr {
		dir := scannerdesc.Text()
		stripdir := strings.TrimSpace(dir)
		return stripdir
	} else {
		return ""
	}

}

func InputScanDirections(directions string) string {
	fmt.Println(directions)

	scannerdesc := bufio.NewScanner(os.Stdin)
	tr := scannerdesc.Scan()
	if tr {
		dir := scannerdesc.Text()
		stripdir := strings.TrimSpace(dir)
		return stripdir
	} else {
		return ""
	}

}

func ShellBash(s, errmess string) error {

	out, errout, err := Shellout(s)
	if err != nil {
		return err
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println(errout)
	}

	return nil

}

func UpdateText(file, check, comment, replace string) error {
	if text.FindTextNReturn(file, check) != comment {
		err := text.UpdateText(file, comment, replace+"\n"+comment)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveText(file, check, comment, replace string) error {
	if text.FindTextNReturn(file, check) != comment {
		err := text.UpdateText(file, comment, replace+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCode(file, check, comment, replace string) error {
	if text.FindTextNReturn(file, check) != comment {
		err := text.UpdateText(file, check, replace+"\n"+comment)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateLine(file, check, replace string) error {

	err := text.UpdateText(file, check, replace)
	if err != nil {
		return err
	}

	return nil
}
func ReplaceLine(file, check, replace string) error {
	input, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, check) {
			lines[i] = replace
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(file, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
func Checklogger(err error, s string) {
	logger := loggers.CreateLogger()
	if err != nil {
		logger.Error(
			s,
			slog.String("error: ", err.Error()),
		)
	}
}
func Makefolder(p string) error {
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
		return err
	}
	return nil
}
