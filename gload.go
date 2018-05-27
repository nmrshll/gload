package gload

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dwarvesf/glod"
	"github.com/dwarvesf/glod/chiasenhac"
	"github.com/dwarvesf/glod/facebook"
	"github.com/dwarvesf/glod/nhaccuatui"
	"github.com/dwarvesf/glod/soundcloud"
	"github.com/dwarvesf/glod/vimeo"
	"github.com/dwarvesf/glod/youtube"
	"github.com/dwarvesf/glod/zing"
)

const (
	initNhacCuaTui string = "nhaccuatui"
	initZingMp3    string = "zing"
	initYoutube    string = "youtube"
	initSoundCloud string = "soundcloud"
	initChiaSeNhac string = "chiasenhac"
	initFacebook   string = "facebook"
	initVimeo      string = "vimeo"
)

func Download(link string) {
	var glod glod.Source
	if strings.Contains(link, initNhacCuaTui) {
		glod = &nhaccuatui.NhacCuaTui{}
	} else if strings.Contains(link, initZingMp3) {
		glod = &zing.Zing{}
	} else if strings.Contains(link, initYoutube) {
		glod = &youtube.Youtube{}
	} else if strings.Contains(link, initSoundCloud) {
		glod = &soundcloud.SoundCloud{}
	} else if strings.Contains(link, initChiaSeNhac) {
		glod = &chiasenhac.ChiaSeNhac{}
	} else if strings.Contains(link, initFacebook) {
		glod = &facebook.Facebook{}
	} else if strings.Contains(link, initVimeo) {
		glod = &vimeo.Vimeo{}
	}

	directLinks, err := glod.GetDirectLink(link)
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(directLink, initYoutube) || strings.Contains(link, initZingMp3) || strings.Contains(link, initVimeo) {
		splitUrl := strings.Split(_temp, "~")
		temp = splitUrl[0]
	}

	resp, _ := http.Get(temp)
	defer resp.Body.Close()

	var nameSanitized string

	if strings.Contains(link, initNhacCuaTui) {
		splitName := strings.Split(temp, "/")
		nameSanitized = strings.Trim(splitName[len(splitName)-1], " ")

	} else if strings.Contains(link, initZingMp3) {
		splitName := strings.Split(_temp, "~")
		nameSanitized = strings.Trim(splitName[1]+".mp3", " ")

	} else if strings.Contains(link, initYoutube) {
		splitName := strings.Split(_temp, "~")
		nameSanitized = strings.Trim(splitName[1], " ")

	} else if strings.Contains(link, initSoundCloud) {
		splitName := strings.Split(temp, "/")
		nameSanitized = strings.Trim(splitName[4]+".mp3", " ")

	} else if strings.Contains(link, initChiaSeNhac) {
		splitName := strings.Split(temp, "~")
		splitNameSplash := strings.Split(splitName[0], "/")
		var nameBeforeSanitize = splitNameSplash[len(splitNameSplash)-1]
		nameSanitized = strings.Replace(nameBeforeSanitize, "%20", " ", -1)
		nameSanitized = strings.Trim(nameSanitized, " ")

	} else if strings.Contains(link, initFacebook) {
		splitName := strings.Split(link, "/")
		nameSanitized = strings.Trim(splitName[len(splitName)-2]+".mp4", " ")

	} else if strings.Contains(link, initVimeo) {
		splitName := strings.Split(_temp, "~")
		nameSanitized = strings.Trim(splitName[1]+".mp4", " ")
	}

	name = append(name, nameSanitized)
}
