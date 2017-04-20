package print

import (
	"fmt"
	"os"
)

const (
	//GREEN Color
	green = "\x1b[32m"
	//RED Color
	red = "\x1b[31m"
	//RESET Color
	reset = "\033[0m"
	//YELLOW Color
	yellow = "\x1b[33m"
	//BLUE Color
	blue = "\x1b[34m"
)

//Goodf printf results
func Goodf(format string, v ...interface{}) {
	output := fmt.Sprintf("%s[+]%s ", green, reset)
	output += format
	fmt.Fprintf(os.Stdout, output, v...)
}

//Goodln println results
func Goodln(v ...interface{}) {
	output := fmt.Sprintf("%s[+]%s", green, reset)
	v = append([]interface{}{output}, v...)
	fmt.Fprintln(os.Stdout, v...)
}

//Badf printf Results
func Badf(format string, v ...interface{}) {
	output := fmt.Sprintf("%s[-]%s ", red, reset)
	output += format
	fmt.Fprintf(os.Stdout, output, v...)
	os.Exit(1)
}

//Badln printf Results
func Badln(v ...interface{}) {
	output := fmt.Sprintf("%s[-]%s", red, reset)
	v = append([]interface{}{output}, v...)
	fmt.Fprintln(os.Stdout, v...)
	os.Exit(1)
}

//Statusf printf results
func Statusf(format string, v ...interface{}) {
	output := fmt.Sprintf("%s[*]%s ", blue, reset)
	output += format
	fmt.Fprintf(os.Stdout, output, v...)
}

//Statusln println results
func Statusln(v ...interface{}) {
	output := fmt.Sprintf("%s[*]%s", blue, reset)
	v = append([]interface{}{output}, v...)
	fmt.Fprintln(os.Stdout, v...)
}

//Warningf printf results
func Warningf(format string, v ...interface{}) {
	output := fmt.Sprintf("%s[!]%s ", yellow, reset)
	output += format
	fmt.Fprintf(os.Stdout, output, v...)
}

//Warningln println results
func Warningln(v ...interface{}) {
	output := fmt.Sprintf("%s[!]%s", yellow, reset)
	v = append([]interface{}{output}, v...)
	fmt.Fprintln(os.Stdout, v...)
}
