package telegram

import (
	"log"
	"net/url"
	"strings"
	"tg-bot/lib/e"
	"tg-bot/storage"
)

const (
	RndCmd  = "/rnd"
	HelpCmd = "/help"
	StarCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		// TODO: AddPage()
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StarCmd:
	default:
	}
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}

	if isExists {
		return p.tg.SendMessage(chatID, "")
	}
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
