package main

import (
	"log"

	"github.com/nivin-studio/juejin-auto/dingtalk"
	"github.com/nivin-studio/juejin-auto/juejin"
	"github.com/nivin-studio/juejin-auto/utils"
)

func main() {
	msg := juejin.New().
		SetCookie(utils.Env("JUEJIN_COOKIE", ``)).
		CheckIn().
		LotteryDraw().
		LotteryDip().
		CollectBugs().
		GetResult()

	err := dingtalk.NewRobot().
		SetWebhook(utils.Env("DINGTALK_WEBHOOK", ``)).
		SetSecret(utils.Env("DINGTALK_SECRET", ``)).
		SendMessage(dingtalk.NewTextMessage(msg))
	if err != nil {
		log.Println(err)
	}

	log.Println("Hello, 世界!")
}
