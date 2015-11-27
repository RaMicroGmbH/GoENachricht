go-bindata -pkg="assets" views\. views\base\.
move /Y bindata.go system\assets\bindata.go
go-bindata-assetfs -pkg="assetfs" static\css\. static\fonts\. static\js\.
move /Y bindata_assetfs.go system\assetfs\bindata_assetfs.go
go build 

