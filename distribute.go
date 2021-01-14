package main
import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"strings"
	toml "github.com/pelletier/go-toml"
)
// MemberData is struct of pet
type MemberData struct {
	Name string   `toml:"name"`
	Site string   `toml:"site"`
	Twitter string   `toml:"twitter"`
	Github string   `toml:"github"`
	Hackerone string   `toml:"hackerone"`
	Bugcrowd string   `toml:"bugcrowd"`
	Intigriti string   `toml:"intigriti"`
	Yeswehack string   `toml:"yeswehack"`
	Youtube string   `toml:"youtube"`
	Instagram string   `toml:"instagram"`
	Facebook string   `toml:"facebook"`
	Reddit string `toml:"reddit"`
	Stackoverflow string `toml:"stackoverflow"`
}
func main(){
	MakeReadme()
}
// MakeReadme is head + auto write body + foot
func MakeReadme() {
	readme := "| Name | Site | Metadata | Social |\n| ------- | ------- | ------- | ------- |\n"
	_=readme
	files, err := ioutil.ReadDir("./data")
	check(err)
	for _, f := range files {
		if strings.Contains(f.Name(), ".toml") {
			memberData := MemberData{}
			sn := ""
			mt := ""
			_= mt
			_=sn
			fmt.Println("# " + f.Name())
			data, err := ioutil.ReadFile("./data/" + f.Name())
			check(err)
			toml.Unmarshal(data, &memberData)
			site := emptyCheck(memberData.Site,"-")
			twitter := emptyCheck(memberData.Twitter,"")
			github := emptyCheck(memberData.Github,"")
			// Merge Metadata
			if twitter != "" {
				mt = mt + "![](https://img.shields.io/twitter/follow/"+twitter+"?label=followers&logo=twitter&color=white&logoColor=white&style=flat-square) "
			}
			if github != "" {
				mt = mt + "![](https://img.shields.io/github/stars/"+github+"?logo=gitHub&style=flat-square&color=white&logoColor=white) "
				mt = mt + "![](https://img.shields.io/github/followers/"+github+"?logo=gitHub&style=flat-square&color=white&logoColor=white) "
			}
			// Merge Social Network
			sn = sn + makeSocial(memberData.Twitter,"twitter","https://twitter.com/")
			sn = sn + makeSocial(memberData.Github,"github","https://github.com/")
			sn = sn + makeSocial(memberData.Youtube,"youtube","")
			sn = sn + makeSocial(memberData.Instagram,"instagram","https://instagram.com/")
			sn = sn + makeSocial(memberData.Facebook,"facebook","https://facebook.com/")
			sn = sn + makeSocial(memberData.Hackerone,"hackerone","https://hackerone.com/")
			sn = sn + makeSocial(memberData.Bugcrowd,"bugcrowd","https://bugcrowd.com/")
			sn = sn + makeSocial(memberData.Intigriti,"intigriti","https://intigriti.com/")
			sn = sn + makeSocial(memberData.Yeswehack,"yeswehack","https://yeswehack.com/")
			sn = sn + makeSocial(memberData.Reddit,"reddit","https://www.reddit.com/r/")
			sn = sn + makeSocial(memberData.Stackoverflow,"stackoverflow","https://stackoverflow.com/users/")
			readme = readme + memberData.Name + " | " + site + " | " + mt + " | " + sn + " |\n"
		}
	}
	fmt.Println(readme)
	top, err := os.Open("template/head.md")
	check(err)
	headData, _ := ioutil.ReadAll(top)
	foot, err := os.Open("template/foot.md")
	check(err)
	footData, _ := ioutil.ReadAll(foot)
	_ = headData
	_ = footData
	t := time.Now().Format("2006-01-02 15:04:05")
	body := string(headData) + readme + string(footData) + "\nlast changed "+t
	fmt.Println("======================result====================")
	fmt.Println(body)
	fmt.Println("======================command===================")
	fmt.Println("git add README.md ; git commit -m \"update readme using distribute\"; git push")
	file, err := os.OpenFile(
		"README.md",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	check(err)
	defer file.Close()
	_, err = file.Write([]byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func emptyCheck(src,dst string) string {
	if src != "" {
		return strings.Replace(src, "|", "\\|", -1)
	}
	return dst
}
func makeSocial(src, stype, address string) string {
	if src != "" {
		return "[![](/assets/"+stype+".svg)]("+address+src+")"
	}
	return ""
}
