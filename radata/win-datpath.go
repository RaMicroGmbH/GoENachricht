// +build windows

package radata

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"../utils"
	"golang.org/x/sys/windows/registry"
)

const (
	raDir_GO = "go"
	raDir_AR = "AR"
	raDir_EN = "ENachrichtDemo"
)

var (
	enachrichtsPath string
)

// get registry keys

func RamicroBenutzernummer() string {
	return readReg("Benutzernummer")
}

func RamicroDatPath() string {
	return readReg("DatPath")
}

func readReg(key string) string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\RA-MICRO\RA-MICRO\Global`, registry.QUERY_VALUE)
	defer k.Close()
	if err != nil {
		//log.Println("error during OpenKey, returning default")
		return ""
	}
	s, _, err := k.GetStringValue(key)

	if err != nil {
		log.Println("Fehler beim Lesen des Registry-Keys " + key)
		return ""
	}
	return strings.TrimSpace(s)
}

func GetENachrichtPath() string {

	if enachrichtsPath == "" {
		//host, _ := os.Hostname()
		enachrichtsPath = filepath.Join(RamicroDatPath(), raDir_GO, raDir_EN, raDir_AR, utils.GetYear(), utils.GetMonth(), utils.GetDay())
		_, err := os.Stat(enachrichtsPath)
		if err != nil {
			os.MkdirAll(enachrichtsPath, 0777)
		}
		enachrichtsPath = filepath.Join(enachrichtsPath)
	}
	return enachrichtsPath
}

func DatPathIsThere(dp string) bool {
	if dp != "" {
		fi, err := os.Stat(dp)
		if err != nil {
			return false
		}
		return fi.IsDir()
	}
	return false
}
