package controllers

import (
	"encoding/json"
	_ "fmt"
	_ "golang.org/x/sys/windows/registry" //get package with "go get godoc.org/golang.org/x/sys"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var RegexpFindUser = regexp.MustCompile(`\[[0-9]+\]`)

var UserSettings = map[string]string{}

var CurrentUserId = RamicroBenutzernummer()

var AllUsers []string

var AllSubjects SubjectSettings

var JsonInitData []byte

var RADatPath = RamicroDatPath()
var BenutzerIni = "\\benutzer\\BENUTZER.INI"
var EMailIni = "\\benutzer\\32EMAIL.INI"
var ENachrichtBetreffMsk = "\\mas\\ENachrichtBetreff.msk"

/*
Data to send to the UI to set all the fields correct.
*/
type InitData struct {
	Settings map[string]string
	Subject  SubjectSettings
	Users    []string
}

/*
Data for the subject settings, including all the examples in a list.
*/
type SubjectSettings struct {
	Default  string
	ZWahl    string
	Examples []string
}

/*

 */
func toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

func HandleFiles() {
	getUserNamesFromBenutzerIni(readFiles(RADatPath + BenutzerIni))
	getUserSettingsFromEMailIni(readFiles(RADatPath + EMailIni))
	getSebjectsettingsFromENachrichtBetreff(readFiles(RADatPath + ENachrichtBetreffMsk))
	JsonInitData, err := json.Marshal(InitData{UserSettings, AllSubjects, AllUsers})
	if err != nil {
		println(err)
		return
	}
	ioutil.WriteFile("MyFile1.txt", JsonInitData, 0x777)
	//fmt.Println(string(JsonInitData))

}

func getSebjectsettingsFromENachrichtBetreff(data []string) {
	AllSubjects.Examples = data[1:]
	settingsArray := strings.Split(data[0], ";")
	for _, element := range settingsArray {
		if strings.HasPrefix(element, "ZWahl") {
			z := strings.Split(element, "=")
			AllSubjects.ZWahl = z[1]
		}
		if strings.HasPrefix(element, "Default") {
			a := strings.Split(element, "=")
			AllSubjects.Default = a[1]
		}
	}
}

func getUserNamesFromBenutzerIni(data []string) {
	log.Println("DATPATH :" + RADatPath)
	var userId string
	for _, element := range data {
		if RegexpFindUser.MatchString(element) == true {

			element = strings.Replace(element, "[", "", 1)
			element = strings.Replace(element, "]", "", 1)
			if len(element) == 1 {
				userId = "00" + element
			} else if len(element) == 2 {
				userId = "0" + element
			} else {
				userId = element
			}
		}
		if strings.HasPrefix(element, "Name") {
			usernameArray := strings.Split(element, "=")
			AllUsers = append(AllUsers, userId+" "+usernameArray[1])
		}
	}
}

func getUserSettingsFromEMailIni(data []string) {
	isCurrentUser := false
	currentUser := "[" + CurrentUserId + "]"
	for _, element := range data {
		if isCurrentUser {
			if RegexpFindUser.MatchString(element) == true {
				return
			}
			if strings.Contains(element, "=") {
				userSetting := strings.Split(element, "=")
				UserSettings[userSetting[0]] = userSetting[1]
			}
		}

		if element == currentUser {
			isCurrentUser = true
		}
	}
}

/*
Function to read the Ini-files and already splitting the data to lines.
*/
func readFiles(fileName string) []string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		println(err)
	}
	data := strings.Split(toUtf8(content), "\r\n")
	return data
}
