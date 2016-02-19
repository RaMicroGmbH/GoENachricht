package utils

import (
	"github.com/metakeule/fmtdate"
	"strings"
	"time"
)

var d = GetDate()

func GetDay() string {
	return strings.Split(d, ".")[0]
}

func GetMonth() string {
	return strings.Split(d, ".")[1]
}

func GetYear() string {
	var year = strings.Split(d, ".")
	return strings.SplitAfterN(year[2], "", 3)[2]
}

func GetDate() string {
	date := fmtdate.Format("DD.MM.YYYY", time.Now())
	return date
}
