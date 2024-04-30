package main

import (
	"biliVideoDownload/pkg/api"
)

const dest = "C:\\Users\\mikumifa\\GolandProjects\\biliVideoDownload\\test.flv"

func main() {
	//api.Download("BV1oJ4m1n7TM", dest, api.Single)
	api.Download("BV1LF4m1T7d4", dest, api.Part, api.MakeQn(api.Audio132K, api.Qn360PFirst), nil)

}
