package helper

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func LogInfo(format string, args ...interface{}) {
	color.Set(color.FgHiBlue)
	log.Printf(format+"\n", args...)
	color.Unset()
}

func LogNotice(format string, args ...interface{}) {
	color.Set(color.FgCyan)
	log.Printf(format+"\n", args...)
	color.Unset()
}

func LogSuccess(format string, args ...interface{}) {
	color.Set(color.FgGreen)
	log.Printf(format+"\n", args...)
	color.Unset()
}

func LogError(format string, args ...interface{}) {
	color.Set(color.FgRed)
	log.Println("[ERROR] " + fmt.Sprintf(format, args...))
	color.Unset()
}
