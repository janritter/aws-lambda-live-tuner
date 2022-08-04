package helper

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func LogInfo(format string, args ...interface{}) {
	color.Set(color.FgHiBlue)
	log.Println(fmt.Sprintf(format, args...))
	color.Unset()
}

func LogNotice(format string, args ...interface{}) {
	color.Set(color.FgCyan)
	log.Println(fmt.Sprintf(format, args...))
	color.Unset()
}

func LogSuccess(format string, args ...interface{}) {
	color.Set(color.FgGreen)
	log.Println(fmt.Sprintf(format, args...))
	color.Unset()
}

func LogError(format string, args ...interface{}) {
	color.Set(color.FgRed)
	log.Println(fmt.Sprintf(format, args...))
	color.Unset()
}
