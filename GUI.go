package main

import (
	"./controllers"
	"./routes"
	"./system/assetfs"
	"net/http"
)

/*
	TODO FOR USE ASSETS

	Install go-bindata and go-bindata-assets

	go get -u github.com/jteeuwen/go-bindata/...

	Download the package for go-bindata-assetfs from:
	https://github.com/baxterbln/go-bindata-assetfs-subpgk

	unpack the go-bindata-assetfs file
	go into the folder go-bindata-assetfs/go-bindata-assetfs

	build the go-bindata-assetfs with
	go build -o go-bindata-assetfs main.go

	check if the path to the go-bindata and go-bindata-assetfs in %PATH% (on linux on $PATH)


	execute:

    On Linux:
	go-bindata -pkg="assets" views/* apps/example/views/*
	go-bindata-assetfs -pkg="assetfs" static/*

    On Windows:
    go-bindata -pkg="assets" views\. views\base\. apps\example\views\.
    go-bindata-assetfs -pkg="assetfs" static\css\. static\fonts\. static\js\.

	move the file new file bindata_assetfs.go to system/assetfs/bindata_assetfs.go

	execute:

    On Linux
	go-bindata -pkg="assets" views/* apps/example/views/*

    On Windows:
    go-bindata -pkg="assets" views\. views\base\.  apps\example\views\.

	move the file new file bindata.go to system/assets/bindata.go

	run the app:
	go run example.go

	build for windows:
	env GOOS=windows GOARCH=386 go build -o example.exe example.go
	env GOOS=windows GOARCH=amd64 go build -o example.exe example.go

	to build without a terminal window:
	env GOOS=windows GOARCH=386 go build -ldflags "-H windowsgui"  -o example.exe example.go
	env GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui"  -o example.exe example.go
*/

func startUIService(Port string) {

	controllers.InitContent("RA-Micro Go App Volltext Suche", Version, protPortFlag)

	mux := http.NewServeMux()

	// Mapping to the static folders, is importend for mapping to assets
	mux.Handle("/js/", http.FileServer(assetfs.AssetFS()))
	mux.Handle("/css/", http.FileServer(assetfs.AssetFS()))
	mux.Handle("/fonts/", http.FileServer(assetfs.AssetFS()))

	// Add routes
	routes.Include(mux)

	http.ListenAndServe(":"+Port, mux)

}
