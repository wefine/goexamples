package main

import (
	"fmt"
	"net/http"
	"os"
)

func emptyProxy(name string) {
	proxy := os.Getenv(name)

	if len(proxy) > 0 {
		os.Setenv(name, "")
	}
}       //

var name string = "-1"
var tag string = "-1"

func main() {

	switch len(os.Args) {
	case 2:
		name = os.Args[1]
	case 3:
		name = os.Args[1]
		 tag = os.Args[2]
	default:
		fmt.Println("Need parameters!")
		os.Exit(1)
	}

	emptyProxy("HTTP_PROXY")
	emptyProxy("HTTP_PROXY")
	emptyProxy("http_proxy")
	emptyProxy("http_proxy")
	emptyProxy("HTTPS_PROXY")
	emptyProxy("HTTPS_PROXY")
	emptyProxy("https_proxy")
	emptyProxy("https_proxy")

	delete_url := os.Getenv("REGISTRY_URL")

	http_head:="http://"

	//dash_http :=http_head+ delete_url +"/v2/_catalog" //"http://10.114.51.186:5000/v2/_catalog"

	dash_http2 :=http_head+ delete_url+"/v2/"  //"http://10.114.51.186:5000/v2/"

	//dash_lag := "/tags/list" //"/tags/list"

	requestManifest, _ := http.NewRequest("GET",dash_http2 + name +"/manifests/"+ tag, nil)// strings.Split(valtag,",")
	requestManifest.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	client := &http.Client{}
	responseManifest, _ := client.Do(requestManifest)
	var digest string
	digest = responseManifest.Header.Get("Docker-Content-Digest")
	fmt.Println("digest: ", digest)

	//deleteUrl := path.Join("http://10.114.51.186:5000/v2/nginx/manifests/", digest)
	var deleteUrl string
	deleteUrl=dash_http2+ name+"/manifests/"+ digest  // "http://10.114.51.186:5000/v2/nginx/manifests/"+digest
	fmt.Println(deleteUrl)
	requestDelete, _ := http.NewRequest("DELETE", deleteUrl, nil)

	responseDelete, delErr := client.Do(requestDelete)
	if delErr == nil {
		fmt.Println(responseDelete.StatusCode)
	}

	if responseDelete.StatusCode != 202 {
		fmt.Println("error!be careful")
		if (responseDelete.StatusCode == 404){
			fmt.Println("该镜像也被删除过")
		}
		if(responseDelete.StatusCode ==405){
			fmt.Println("该仓库不容许删除镜像，请对仓库进行设置")
		}
	}




}
