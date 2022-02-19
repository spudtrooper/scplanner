package log

import "github.com/spudtrooper/goutil/colorlog"

const prefix = "[scplanner] "

func Printf(tmpl string, args ...interface{}) {
	colorlog.Printf(prefix+tmpl, args...)
}

func Fatalf(tmpl string, args ...interface{}) {
	colorlog.Fatalf(prefix+tmpl, args...)
}

func Println(s string) {
	colorlog.Println(prefix + s)
}
