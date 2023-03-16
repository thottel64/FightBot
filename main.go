package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Global variables listed together
var (
	Token     string
	initiator Fighter
	responder Fighter
	fightInit bool
)

type Fighter struct {
	ID   string
	HP   int
	turn bool
}

func main() {
	Token = os.Getenv("DISCORD_TOKEN")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session, \n", err)
		return
	}
	ticker := time.NewTicker(24 * time.Hour)
	for now := range ticker.C {
		if now.Weekday() == time.Tuesday && now.Hour() == 8 {
			_, err = dg.ChannelMessageSend("541777196960972823", "https://cdn.discordapp.com/attachments/541777196960972823/1080147477962969169/trim.D7D0EB9D-23C2-4D9E-BE42-1F248EE65D14.mp4")
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}

	// Register the Fightbot func as a callback for messages that meet the required parameters for events.
	dg.AddHandler(FightBot)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening. Luckily, the DiscordGo package makes this easy by calling the .Open method
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,\n", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func FightBot(s *discordgo.Session, m *discordgo.MessageCreate) {
	var err error
	rand.NewSource(time.Now().UnixNano())
	if m.Author.ID == s.State.User.ID {
		return
	}
	var turncounter int

	if strings.ToLower(m.Content) == "fightbot help" {
		_, err = s.ChannelMessageSend(m.ChannelID, "Fightbot is here to help. Initiate a fight by typing `fight` followed by a space and a ping to the user you wish to fight. Ex. `fight @User1` . Then to attack, simply type in `punch`. Once a player's HP reaches 0, they lose. ")
		if err != nil {
			log.Println("Could not send message \n", err)
		}
	}
	if strings.ToLower(m.Content) == "chicken salad" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541777196960972823/1083122814116188280/Snapinsta.app_334279527_508487954817631_2059612099667519088_n-1.mp4")
	}
	if strings.ToLower(m.Content) == "ppme" {
		var pp []byte
		pp = append(pp, 56)
		for i := 1; i <= rand.Intn(20); i++ {
			pp = append(pp, 61)
		}
		pp = append(pp, 68)
		_, err = s.ChannelMessageSend(m.ChannelID, string(pp))
		if err != nil {
			log.Println(err)
			return
		}
	}
	if strings.ToLower(m.Content) == "yeah" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://media.tenor.com/CVZlYWQibqoAAAAC/eli-drake-la-knight.gif")
		if err != nil {
			log.Println(err)
		}
	}
	// if the user types in fight followed by another user's mention, the bot initiates a fight between the two users
	if len(m.Content) >= 8 && (strings.ToLower(m.Content[0:8]) == "fight <@") && string(m.Content[len(m.Content)-1]) == ">" && fightInit == false {
		if m.Content == strings.ToLower("fight <@>") {
			return
		}
		responder.ID = m.Content[6:len(m.Content)]
		initiator.ID = m.Author.Mention()
		responder.HP, initiator.HP = 100, 100
		_, err = s.ChannelMessageSend(m.ChannelID, initiator.ID+" is requesting to fight "+responder.ID)
		if err != nil {
			log.Println("Could not send message \n", err)
		}
		turncounter = rand.Intn(1)
		fightInit = true
		if turncounter == 0 {
			initiator.turn = true
			_, err = s.ChannelMessageSend(m.ChannelID, initiator.ID+" it's your turn.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
		if turncounter == 1 {
			responder.turn = true
			_, err = s.ChannelMessageSend(m.ChannelID, responder.ID+" it's your turn.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
		if responder.ID == initiator.ID {
			_, err = s.ChannelMessageSend(m.ChannelID, "Congrats you win... and lose. I'm not sure what you were trying to accomplish here.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
			fightInit = false
		}
	}

	if strings.ToLower(m.Content) == "punch" && m.Author.Mention() == initiator.ID && initiator.turn == true && fightInit == true {
		Dmg := rand.Intn(60)
		responder.HP = responder.HP - Dmg
		if responder.HP < 0 {
			responder.HP = 0
		}
		dmgstring := strconv.Itoa(Dmg)
		hpstring := strconv.Itoa(responder.HP)
		initiator.turn = false
		responder.turn = true
		if Dmg <= 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, "you missed your punch. "+responder.ID+" takes no damage.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
		if Dmg > 0 {

			_, err = s.ChannelMessageSend(m.ChannelID, "you punched "+responder.ID+" for "+dmgstring+" damage. Leaving them with "+hpstring+" HP left.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
		if responder.HP <= 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, initiator.ID+" is your winner of the round!")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
			responder.turn = false
			initiator.turn = false
			fightInit = false
		}
		if responder.HP > 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, responder.ID+" it's your turn.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}

	}

	if strings.ToLower(m.Content) == "punch" && m.Author.Mention() == responder.ID && responder.turn == true && fightInit == true {
		Dmg := rand.Intn(60)
		initiator.HP = initiator.HP - Dmg
		if initiator.HP < 0 {
			initiator.HP = 0
		}
		dmgstring := strconv.Itoa(Dmg)
		hpstring := strconv.Itoa(initiator.HP)
		initiator.turn = true
		responder.turn = false
		if Dmg <= 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, "you missed your punch. "+initiator.ID+" takes no damage.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
		if Dmg > 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, "you punched "+initiator.ID+" for "+dmgstring+" damage. Leaving them with "+hpstring+" HP left.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
		if initiator.HP <= 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, responder.ID+" is your winner of the round!")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
			responder.turn = false
			initiator.turn = false
			fightInit = false
		}
		if initiator.HP > 0 {
			_, err = s.ChannelMessageSend(m.ChannelID, initiator.ID+" it's your turn.")
			if err != nil {
				log.Println("Could not send message \n", err)
			}
		}
	}
	if strings.ToLower(m.Content) == "surrender" && (m.Author.Mention() == initiator.ID || m.Author.Mention() == responder.ID) && fightInit == true {
		responder.turn = false
		initiator.turn = false
		fightInit = false
		_, err = s.ChannelMessageSend(m.ChannelID, "No contest. The fight is over.")
		if err != nil {
			log.Println("Could not send message \n", err)
		}
	}
}
