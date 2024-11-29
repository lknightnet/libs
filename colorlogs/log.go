package colorlogs

import (
	"fmt"
	"log"
)

type CustomLog struct {
	Color int
}

func NewCustomLog(color int) *CustomLog {
	return &CustomLog{Color: color}
}

func (cs *CustomLog) Println(v ...interface{}) {
	switch cs.Color {
	case COLOR_DEFAULT:
		log.Println(fmt.Sprintf("%s", v))
	case COLOR_BLUE:
		log.Println(fmt.Sprintf("\033[34m%s\033[0m", v))
	case COLOR_RED:
		log.Println(fmt.Sprintf("\033[31m%s\033[0m", v))
	case COLOR_GREEN:
		log.Println(fmt.Sprintf("\033[32m%s\033[0m", v))
	}
}

func BlueLog(v ...interface{}) {
	log.Println(fmt.Sprintf("\033[34m%s\033[0m", v))
}

func GreenLog(v ...interface{}) {
	log.Println(fmt.Sprintf("\033[32m%s\033[0m", v))
}

func RedLog(v ...interface{}) {
	log.Println(fmt.Sprintf("\033[31m%s\033[0m", v))
}

const (
	COLOR_DEFAULT = iota
	COLOR_RED
	COLOR_GREEN
	COLOR_BLUE
)
