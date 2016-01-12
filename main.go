package main

import (
	"flag"
	"fmt"
	"GoENachricht/controllers"
	"log"
	"net/smtp"
	"os"
	"sync"
)

//global vars
var (
	Version       string = "1.2.0"
	RADatPath     string
	RAKDNr        string
	protPortFlag  string
	uiPortFlag    string
	OversizeFlag  bool
	sentAttCount  int
	totalAttCount int
	MailBoxLimit  int64 = 10 * 1000 * 1024
	auth          smtp.Auth
)

func init() {
	flag.StringVar(&RADatPath, "datpath", "/ra", "RA-Micro Datenpfad")
	//flag.StringVar(&protPortFlag, "prot", "24881", "Protokollanzeige auf Port")
	flag.StringVar(&uiPortFlag, "ui", "50001", "Anzeige des UI auf Port")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	controllers.HandleFiles()
	go startUIService(uiPortFlag)
	wg.Wait()
}

func panicExit(msg string, logKey string, exitCode int) {
	showMsg(msg, true, logKey)
	os.Exit(exitCode)
}

func showMsg(msg string, logThis bool, logKey string) {
	if logThis {
		log.Println(msg)
	} else {
		fmt.Println(msg)
	}
}
