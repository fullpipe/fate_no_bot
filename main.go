package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		diceMenu = &tb.ReplyMarkup{
			OneTimeKeyboard: true,
			Selective:       true,
		}

		d4   = diceMenu.Text("🎲 d4")
		d6   = diceMenu.Text("🎲 d6")
		d8   = diceMenu.Text("🎲 d8")
		d10  = diceMenu.Text("🎲 d10")
		d12  = diceMenu.Text("🎲 d12")
		d20  = diceMenu.Text("🎲 d20")
		d100 = diceMenu.Text("🎲 d100")

		dices = []tb.Btn{d4, d6, d8, d10, d12, d20, d100}
	)

	diceMenu.Reply(
		diceMenu.Row(d4, d6, d8),
		diceMenu.Row(d10, d12, d100),
		diceMenu.Row(d20),
	)

	b.Handle("/roll", func(m *tb.Message) {
		b.Reply(m, "Choose the dice", diceMenu)
	})

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Chat, "I would help you with your choice. Send me options to choose from. For example:\n"+
			"comma separated\n"+
			"@fate_no_bot tea, coffee, water\n"+
			"or just space separated\n"+
			"@fate_no_bot tea coffee water\n"+
			"you could mention people\n"+
			"@fate_no_bot @one, @two, @someone")
	})

	for _, d := range dices {
		b.Handle(&d, func(m *tb.Message) {
			b.Reply(m, fmt.Sprintf("You rolled: %d", rollText(m.Text)))
		})
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.Sender.IsBot {
			return
		}

		if !m.Private() && !strings.HasPrefix(m.Text, "@fate_no_bot") {
			return
		}

		message := strings.TrimPrefix(m.Text, "@fate_no_bot")
		message = strings.TrimSpace(message)

		totalRoll := rollText(message)
		if totalRoll > 0 {
			b.Reply(m, fmt.Sprintf("You rolled: %d", totalRoll))
			return
		}

		//make choice without dice
		choice, err := choose(message)
		if err != nil {
			b.Reply(m, err)
			return
		}

		b.Reply(m, fmt.Sprintf("Your choice: %s", choice))
	})

	b.Start()
}

func choose(text string) (string, error) {
	//try to split by comma
	choices := strings.Split(text, ",")
	if len(choices) < 2 {
		choices = strings.Fields(text)
	}

	if len(choices) < 2 {
		return "", errors.New("Nothing to choose")
	}

	choice := choices[rand.Intn(len(choices))]

	return strings.TrimSpace(choice), nil
}
