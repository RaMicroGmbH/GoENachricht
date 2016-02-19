package radata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	_ "golang.org/x/sys/windows/registry" //get package with "go get godoc.org/golang.org/x/sys"
)

const (
	BenutzerInifile          = "\\benutzer\\BENUTZER.INI"
	EMailInifile             = "\\benutzer\\32EMAIL.INI"
	ENachrichtBetreffMskfile = "\\mas\\ENachrichtBetreff.msk"
)

var (
	RegexpFindUser = regexp.MustCompile(`\[[0-9]+\]`)
	UserSettings   = map[string]string{}
	CurrentUserId  = RamicroBenutzernummer()
	AllUsers       []string
	AllSubjects    SubjectSettings
	JsonInitData   []byte

	RADatPath = RamicroDatPath()

	benutzerInifiledata []string
	emailInifiledata    []string
	betreffmskfiledata  []string

	benutzerinireadstastus string
	emailInireadstatus     string
	betreffmskreadstatus   string
)

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
	Default string
	ZWahl   string
	Samples []string
}

func HandleFiles() {

	benutzerInifiledata, benutzerinireadstastus := readFile(RADatPath + BenutzerInifile)
	if benutzerinireadstastus == nil {
		getUserNamesFromBenutzerIni(benutzerInifiledata)
	}

	emailInifiledata, emailInireadstatus := readFile(RADatPath + EMailInifile)
	if emailInireadstatus == nil {
		getUserSettingsFromEMailIni(emailInifiledata)
	}

	betreffmskfiledata, betreffmskreadstatus := readFile(RADatPath + ENachrichtBetreffMskfile)
	if betreffmskreadstatus == nil {
		getSubjectsettingsFromENachrichtBetreff(betreffmskfiledata)
	}

	JsonInitData, err := json.Marshal(InitData{UserSettings, AllSubjects, AllUsers})
	Checkerr(err)

	fmt.Println("INI_DATA BYTE_ARRAY:", JsonInitData)
	//ioutil.WriteFile("MyFile1.txt", JsonInitData, 0x777)

}

func getSubjectsettingsFromENachrichtBetreff(data []string) {

	if data != nil {
		AllSubjects.Samples = data[1:]
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
	Function to read the Ini-files and  split the data to lines.
*/
func readFile(fileName string) ([]string, error) {
	var err error
	if fileExists(fileName) {
		content, err := ioutil.ReadFile(fileName)
		Checkerr(err)
		data := strings.Split(toUtf8(content), "\r\n")
		return data, err

	} else {
		return nil, err
	}
}

//func writeFiles(fileName string) []string {
//	content, err := ioutil.WriteFile(fileName, 0x777)
//	Checkerr(err)
//}

/*
	Function to save a new subject after adding
*/
func SaveSubject(data []string) {

}

/*--------------------------------------------------------------------------
					UTILITY FUNCTIONS
--------------------------------------------------------------------------*/

/*
  Handling special characters
  Note : rune is an alias for int32
  Character-set : iso8859-1
*/
func toUtf8(data []byte) string {
	buf := make([]rune, len(data))
	for i, b := range data {
		buf[i] = rune(b)
	}
	return string(buf)
}

func fileExists(fnam string) bool {
	if _, err := os.Stat(fnam); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Checkerr(err error) {
	if err != nil {
		//panic(err)
		return
	}
}
