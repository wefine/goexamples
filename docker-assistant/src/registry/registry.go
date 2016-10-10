package main

import "net/http"
import "io/ioutil"
import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
)

type Catalog struct {
	Repositories []string `json:"repositories"`
}

type ImageInfo struct {
	Name string `json:"name"`
	Tags []string `json:"tags"`
}

func (i ImageInfo) String() string {
	slice := []string{}
	for _, tag := range i.Tags {
		slice = append(slice, i.Name + ":" + tag)
	}

	return strings.Join(slice, "\n")
}

var registryUrl string = "-1"
var pattern string = "-1"

func main() {
	switch len(os.Args) {
	case 2:
		registryUrl = os.Args[1]
	case 3:
		registryUrl = os.Args[1]
		pattern = os.Args[2]
	default:
		fmt.Println("Need parameters!")
		os.Exit(1)
	}

	v2_catalog := registryUrl + "/v2/_catalog";
	content := remoteResponse(v2_catalog);

	res := Catalog{}
	json.Unmarshal([]byte(content), &res)

	imageMap := map[string]ImageInfo{}
	for _, v := range res.Repositories {
		imageMap[v] = imageInfo(v)
	}

	slice := []string{}
	for _, imageInfo := range imageMap {
		v := imageInfo.String()
		if pattern == "-1" || pattern != "-1" && strings.Contains(v, pattern) {
			slice = append(slice, v)
		}
	}

	for _, v := range slice {
		fmt.Println(v)
	}
}

func remoteResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body);
}

func imageInfo(name string) ImageInfo {
	tags_list := registryUrl + "/v2/" + name + "/tags/list"
	content := remoteResponse(tags_list)

	info := ImageInfo{}
	json.Unmarshal([]byte(content), &info)

	return info
}