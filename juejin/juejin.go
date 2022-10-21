package juejin

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	BASE_URL        = "https://api.juejin.cn"
	CHECKIN_API     = BASE_URL + "/growth_api/v1/check_in"
	LOTTERY_API     = BASE_URL + "/growth_api/v1/lottery/draw"
	GLOBAL_BIG_API  = BASE_URL + "/growth_api/v1/lottery_history/global_big"
	DIP_LUCKY_API   = BASE_URL + "/growth_api/v1/lottery_lucky/dip_lucky"
	NOT_COLLECT_API = BASE_URL + "/user_api/v1/bugfix/not_collect"
	COLLECT_API     = BASE_URL + "/user_api/v1/bugfix/collect"
)

type JueJin struct {
	Result string        `json:"result"`
	Client *resty.Client `json:"client"`
}

func New() *JueJin {
	j := &JueJin{
		Result: "",
		Client: resty.New(),
	}

	return j
}

func (j *JueJin) SetCookie(cookie string) *JueJin {
	j.Client.SetHeader("cookie", cookie)

	return j
}

func (j *JueJin) AddResult(s string) *JueJin {
	j.Result += s + "\n\n"

	return j
}

func (j *JueJin) GetResult() string {
	return j.Result
}

func (j *JueJin) CheckIn() *JueJin {
	resp, err := j.Client.R().Post(CHECKIN_API)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 签到失败\n❓ 失败原因: %s", err))
	}

	var result Response
	result.Data = new(CheckIn)

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 签到失败\n❓ 失败原因: %s", err))
	}

	data, ok := result.Data.(*CheckIn)
	if !ok && result.ErrNo != 0 {
		return j.AddResult(fmt.Sprintf("😔 签到失败\n❓ 失败原因: %s", result.ErrMsg))
	}

	return j.AddResult(fmt.Sprintf("😊 签到成功🎉🎉🎉\n💎 获得矿石: %d\n💎 全部矿石: %d", data.IncrPoint, data.SumPoint))
}

func (j *JueJin) Lottery() *JueJin {
	resp, err := j.Client.R().Post(LOTTERY_API)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 抽奖失败\n❓ 失败原因: %s", err))
	}

	var result Response
	result.Data = new(LotteryDraw)

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 抽奖失败\n❓ 失败原因: %s", err))
	}

	data, ok := result.Data.(*LotteryDraw)
	if !ok && result.ErrNo != 0 {
		j.Result += fmt.Sprintf("😔 抽奖失败\n❓ 失败原因: %s", result.ErrMsg)
		return j
	}

	return j.AddResult(fmt.Sprintf("😊 抽奖成功🎉🎉🎉\n🎁 成功获得: %s", data.LotteryName))
}

func (j *JueJin) GetLuckyUsers() ([]LuckyUser, error) {
	resp, err := j.Client.R().Post(GLOBAL_BIG_API)
	if err != nil {
		return nil, err
	}

	var result Response
	result.Data = new(LotteryHistory)

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	data, ok := result.Data.(*LotteryHistory)
	if !ok && result.ErrNo != 0 {
		return nil, errors.New(result.ErrMsg)
	}

	return data.LuckyUser, nil
}

func (j *JueJin) DipLucky() *JueJin {
	luckyUsers, err := j.GetLuckyUsers()
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 沾沾失败\n❓ 失败原因: %s", err))
	}

	resp, err := j.Client.R().SetBody(map[string]interface{}{
		"lottery_history_id": luckyUsers[0].HistoryID,
	}).Post(DIP_LUCKY_API)

	var result Response
	result.Data = new(DipLucky)

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 沾沾失败\n❓ 失败原因: %s", err))
	}

	data, _ := result.Data.(*DipLucky)
	if result.ErrNo != 0 {
		return j.AddResult(fmt.Sprintf("😔 沾沾失败\n❓ 失败原因: %s", result.ErrMsg))
	}

	return j.AddResult(fmt.Sprintf("😊 沾沾成功🎉🎉🎉\n🍀 沾到幸运: %d\n🍀 当前幸运: %d", data.DipValue, data.TotalValue))
}

func (j *JueJin) GetBugs() (*[]Bug, error) {
	resp, err := j.Client.R().Post(NOT_COLLECT_API)
	if err != nil {
		return nil, err
	}

	var result Response
	result.Data = new([]Bug)

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	data, ok := result.Data.(*[]Bug)
	if !ok && result.ErrNo != 0 {
		return nil, errors.New(result.ErrMsg)
	}

	return data, nil
}

func (j *JueJin) CollectBug() *JueJin {
	bugList, err := j.GetBugs()
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 BugFix失败\n❓ 失败原因: %s", err))
	}

	if len(*bugList) == 0 {
		return j.AddResult(fmt.Sprintf("😔 BugFix失败\n❓ 失败原因: 没有可fix的bug!"))
	}

	for _, v := range *bugList {
		j.Client.R().SetBody(map[string]interface{}{
			"bug_time": v.BugTime,
			"bug_type": v.BugType,
		}).Post(COLLECT_API)

	}

	return j.AddResult(fmt.Sprintf("😊 BugFix完成🎉🎉🎉"))
}
