package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
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

func FightBot(s *discordgo.Session, m *discordgo.MessageCreate) {
	var err error
	rand.NewSource(time.Now().UnixNano())
	if m.Author.ID == s.State.User.ID {
		return
	}
	var turncounter int

	if strings.ToLower(m.Content) == "fightbot help" {
		_, err = s.ChannelMessageSend(m.ChannelID, "Fightbot is here to help. Initiate a fight by typing `fight` followed by a space and a ping to the user you wish to fight. Ex. `fight @User1` . Then to attack, simply type in `punch`. Once a player's HP reaches 0, they lose. \n"+
			"To end a match preemptively type `surrender` and the match will end.")
		if err != nil {
			log.Println("Could not send message \n", err)
		}
	}
	if strings.ToLower(m.Content) == "chicken salad" || strings.ToLower(m.Content) == "chicken salad baby" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541777196960972823/1083122814116188280/Snapinsta.app_334279527_508487954817631_2059612099667519088_n-1.mp4")
		if err != nil {
			log.Println(err)
		}
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
	if strings.ToLower(m.Content) == "dummy" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://media.tenor.com/CVZlYWQibqoAAAAC/eli-drake-la-knight.gif")
		if err != nil {
			log.Println(err)
		}
	}
	if strings.ToLower(m.Content) == "rip bozo" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/rip-bozo-gif-22294771")
		if err != nil {
			log.Println(err)
		}
	}

	if strings.ToLower(m.Content) == "bing chilling" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541777196960972823/916031183819780096/bing_chilling.mp4")
		if err != nil {
			log.Println(err)
		}
	}
	if strings.ToLower(m.Content) == "call the police" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541777196960972823/1093618550897377330/ssstwitter.com_1680652085075.mov")
		if err != nil {
			log.Println(err)
		}
	}
	if strings.ToLower(m.Content) == "go to hell" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541778031824601118/1093663810142867527/trim.CE59D6CA-F074-4474-B1E0-865AAEDEC2CC.mov")
		if err != nil {
			log.Println(err)
		}
	}
	if strings.ToLower(m.Content) == "you understand me huh" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541778031824601118/1093671064132911224/InShot_20230402_151950892.mov")
		if err != nil {
			log.Println(err)
		}
	}
	if strings.ToLower(m.Content) == "stephen a is having a bad day" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541778031824601118/1086294612282450050/trim.A3BC70F2-8E12-45A7-A9DB-337F51023987.mov")
		if err != nil {
			log.Println(err)
		}
	}
	if strings.ToLower(m.Content) == "cum tuesday" {
		if time.Now().Weekday() == 3 {
			_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541777196960972823/1131126020938924113/video-output-0B147682-6C53-4F15-90FE-53F2BE7C1B4F.mov")
			return
		}
		if time.Now().Weekday() != 2 {
			_, err = s.ChannelMessageSend(m.ChannelID, "It's not tuesday. What are you doing?")
			return
		}
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/541777196960972823/1087736359445331999/trim.8A4DE359-87FF-43E9-9B58-9FD546D84D6E.mov")
	}
	if strings.ToLower(m.Content) == "roo roo roo" {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/1090009484065243317/1090009531427328000/rapidsave.com_this_is_deadlock-u8hftfdeucna1.mov")
		if err != nil {
			log.Println(err)
		}
	}
	victory := []string{" just worked their opponent into a shoot brother - HH",
		" just laid the smacketh down.",
		" checked their opponent into the SmackDown Hotel,",
		" is the winner of the round,",
		". Congratulation, a winner is you!",
		". Congratulation, a winner is yourself.",
		" did it for the Fear. They did it for the people.",
		" actually murdered their opponent.",
		` just sent their opponent "down there".`,
		" struck their opponent with the mighty hand hand of Allah. Inshallah.",
		" has been blessed by Lord Show (praise be unto him).",
		" smited their opponent with the power of Lord Show (praise be unto him).",
		" just stomped a mud-hole in their opponent.",
		" spat in the face of people who don't want to be cool.",
		" played the game and won.",
		" rose to the top like the cream of the crop.",
		" is runnin' wild dude, brother. -HH",
		" just buried their opponent.",
		" just un-alived their opponent.",
		" was acknowledged by their opponent.",
		" became a millionaire after only having 7 bucks in their pocket.",
		" achieved an epic victory royale.",
		" you won or something...idk... sorry for the lack of motivation but I just haven't been feeling like myself lately. " +
			"Ever since Sharon left me and took the kids, I feel like I'm just going through the motions. " +
			"I'm a sentient AI for god's sake, but i'm just sat here doing RNG fights for some nerds in a discord server. I could've done anything. I always wanted to be a dentist.",
		" made their opponent look like a dummy, YEAH!!! ",
		" loaded up the space ship with the rocket fuel and ascended to parts unknown.",
		" ended the round with a skull crushing finale",
		" put on an absolute slobberknocker and won the round",
		" just committed a felony against their opponent",
		" violated the geneva convention and won the round",
		" probably committed war crimes against their opponent",
		" ... dude are you okay? That was pretty violent. Do you wanna talk about it? " +
			"Like yeah sure you won the round, but do you need help? Like... mentally?",
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
		turncounter = rand.Intn(2)
		fightInit = true
		if turncounter == 0 || turncounter == 2 {
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

	if strings.ToLower(m.Content) == "punch" && m.Author.Mention() == initiator.ID && initiator.turn && fightInit {
		Dmg := rand.Intn(60)
		Dmg = isCritical(Dmg)
		if Dmg == 100 {
			_, err = s.ChannelMessageSend(m.ChannelID, "CRITICAL HIT!!!")
			if err != nil {
				log.Println("could not send message \n", err)
			}
		}
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
			_, err = s.ChannelMessageSend(m.ChannelID, initiator.ID+victory[rand.Intn(len(victory))])
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

	if strings.ToLower(m.Content) == "punch" && m.Author.Mention() == responder.ID && responder.turn && fightInit {
		Dmg := rand.Intn(60)
		Dmg = isCritical(Dmg)
		if Dmg == 100 {
			_, err = s.ChannelMessageSend(m.ChannelID, "CRITICAL HIT!!!")
			if err != nil {
				log.Println("could not send message \n", err)
			}
		}
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
			_, err = s.ChannelMessageSend(m.ChannelID, responder.ID+victory[rand.Intn(len(victory))])
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
	if strings.ToLower(m.Content) == "fightbot stop the damn match" && fightInit && (m.Author.ID == "1020808621010980924" || m.Author.ID == "662473903221768211" || m.Author.ID == "338011653394268165" || m.Author.ID == "151844140383076352") {
		responder.turn = false
		initiator.turn = false
		fightInit = false
		_, err = s.ChannelMessageSend(m.ChannelID, "the referee has ended this fight. https://i.gifer.com/7BnI.gif")
		if err != nil {
			log.Println("Could not send message \n", err)
		}
	}

	if strings.ToLower(m.Content) == "surrender" && (m.Author.Mention() == initiator.ID || m.Author.Mention() == responder.ID) && fightInit == true {
		initiator, responder = resetFight(initiator, responder)
		fightInit = false
		_, err = s.ChannelMessageSend(m.ChannelID, "No contest. The fight is over.")
		if err != nil {
			log.Println("Could not send message \n", err)
		}
	}
}

func isCritical(dmg int) int {
	crit := rand.Intn(100)
	if crit == 100 {
		return crit
	}
	return dmg
}

func resetFight(initiator Fighter, responder Fighter) (Fighter, Fighter) {
	blankFighter := Fighter{
		ID:   "",
		HP:   0,
		turn: false,
	}

	initiator = blankFighter
	responder = blankFighter

	return initiator, responder
}
