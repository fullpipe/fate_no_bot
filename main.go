package main

import (
	"fmt"
	"log"
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

		totalRoll := rollText(m.Text)
		if totalRoll > 0 {
			b.Reply(m, fmt.Sprintf("You rolled: %d", totalRoll))
		}
	})

	b.Start()
}
