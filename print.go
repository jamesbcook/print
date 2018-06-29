package print

import (
	"fmt"
	"io"
	"os"
)

const (
	green  = "\x1b[32m"
	red    = "\x1b[31m"
	reset  = "\033[0m"
	yellow = "\x1b[33m"
	blue   = "\x1b[34m"
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

//Debugf prints a debug statement in prinf format
func Debugf(debugging bool, w *io.Writer) func(format string, v ...interface{}) {
	return func(format string, v ...interface{}) {
		if debugging {
			output := fmt.Sprintf("[Debug] ")
			output += format
			fmt.Fprintf((*w), output, v...)
		}
	}
}

//Debugln prints a debug statement in println format
func Debugln(debugging bool, w *io.Writer) func(v ...interface{}) {
	return func(v ...interface{}) {
		if debugging {
			output := fmt.Sprintf("[Debug] ")
			v = append([]interface{}{output}, v...)
			fmt.Fprintln((*w), v...)
		}
	}
}

//Error for writing an error to a io writer
func Error(w *io.Writer) func(v error) {
	return func(v error) {
		output := fmt.Sprintf("[Error] ")
		fmt.Fprintln((*w), output, v)
	}
}
