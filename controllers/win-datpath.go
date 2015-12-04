// +build windows

package controllers

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"strings"
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
