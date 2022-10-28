package juejin

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	BASE_URL            = "https://api.juejin.cn"
	CHECK_IN_API        = BASE_URL + "/growth_api/v1/check_in"
	LOTTERY_DRAW_API    = BASE_URL + "/growth_api/v1/lottery/draw"
	LOTTERY_USERS_API   = BASE_URL + "/growth_api/v1/lottery_history/global_big"
	LOTTERY_DIP_API     = BASE_URL + "/growth_api/v1/lottery_lucky/dip_lucky"
	NOT_COLLECT_BUG_API = BASE_URL + "/user_api/v1/bugfix/not_collect"
	COLLECT_API         = BASE_URL + "/user_api/v1/bugfix/collect"
)

type JueJin struct {
	Result string        `json:"result"`
	Client *resty.Client `json:"client"`
}

// 创建掘金实例
func New() *JueJin {
	j := &JueJin{
		Result: "",
		Client: resty.New(),
	}

	return j
}

// 设置cookie
func (j *JueJin) SetCookie(cookie string) *JueJin {
	j.Client.SetHeader("cookie", cookie)

	return j
}

// 添加结果
func (j *JueJin) AddResult(s string) *JueJin {
	j.Result += s + "\n\n"

	return j
}

// 获取结果
func (j *JueJin) GetResult() string {
	return j.Result
}

// 签到
func (j *JueJin) CheckIn() *JueJin {
	resp, err := j.Client.R().Post(CHECK_IN_API)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 签到失败\n❓ 失败原因: %s", err))
	}

	log.Println("签到请求结果:", string(resp.Body()))

	var result CheckInResp
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 签到失败\n❓ 失败原因: %s", err))
	}

	if result.ErrNo != 0 {
		return j.AddResult(fmt.Sprintf("😔 签到失败\n❓ 失败原因: %s", result.ErrMsg))
	}

	return j.AddResult(fmt.Sprintf("😊 签到成功\n💎 获得矿石: %d\n💎 全部矿石: %d", result.Data.IncrPoint, result.Data.SumPoint))
}

// 抽奖
func (j *JueJin) LotteryDraw() *JueJin {
	resp, err := j.Client.R().Post(LOTTERY_DRAW_API)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 抽奖失败\n❓ 失败原因: %s", err))
	}

	log.Println("抽奖请求结果:", string(resp.Body()))

	var result LotteryDrawResp
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 抽奖失败\n❓ 失败原因: %s", err))
	}

	if result.ErrNo != 0 {
		j.Result += fmt.Sprintf("😔 抽奖失败\n❓ 失败原因: %s", result.ErrMsg)
		return j
	}

	return j.AddResult(fmt.Sprintf("😊 抽奖成功\n🎁 成功获得: %s", result.Data.LotteryName))
}

// 中奖用户
func (j *JueJin) LotteryUsers() ([]LotteryUser, error) {
	resp, err := j.Client.R().Post(LOTTERY_USERS_API)
	if err != nil {
		return nil, err
	}

	log.Println("中奖用户请求结果:", string(resp.Body()))

	var result LuckyUsersResp
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	if result.ErrNo != 0 {
		return nil, errors.New(result.ErrMsg)
	}

	return result.Data.LotteryUsers, nil
}

// 沾喜气
func (j *JueJin) LotteryDip() *JueJin {
	luckyUsers, err := j.LotteryUsers()
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 沾沾失败\n❓ 失败原因: %s", err))
	}

	resp, err := j.Client.R().SetBody(map[string]interface{}{
		"lottery_history_id": luckyUsers[0].HistoryID,
	}).Post(LOTTERY_DIP_API)

	log.Println("沾喜气请求结果:", string(resp.Body()))

	var result LotteryDipResp
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return j.AddResult(fmt.Sprintf("😔 沾沾失败\n❓ 失败原因: %s", err))
	}

	if result.ErrNo != 0 {
		return j.AddResult(fmt.Sprintf("😔 沾沾失败\n❓ 失败原因: %s", result.ErrMsg))
	}

	return j.AddResult(fmt.Sprintf("😊 沾沾成功\n🍀 沾到幸运: %d\n🍀 当前幸运: %d", result.Data.DipValue, result.Data.TotalValue))
}

// 未收集BUG
func (j *JueJin) NotCollectBugs() ([]Bug, error) {
	resp, err := j.Client.R().Post(NOT_COLLECT_BUG_API)
	if err != nil {
		return nil, err
	}

	log.Println("未收集BUG请求结果:", string(resp.Body()))

	var result NotCollectBugsResp
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	if result.ErrNo != 0 {
		return nil, errors.New(result.ErrMsg)
	}

	return result.Data, nil
}

// 收集BUG
func (j *JueJin) CollectBugs() *JueJin {

	var sum int

	for {
		bugs, err := j.NotCollectBugs()
		if err != nil {
			return j.AddResult(fmt.Sprintf("😔 Bug收集失败\n❓ 失败原因: %s", err))
		}

		len := len(bugs)
		if len == 0 {
			if sum == 0 {
				return j.AddResult(fmt.Sprintf("😔 Bug收集失败\n❓ 失败原因: 没有可fix的bug!"))
			}
			break
		}

		for _, b := range bugs {
			j.Client.R().SetBody(map[string]interface{}{
				"bug_time": b.BugTime,
				"bug_type": b.BugType,
			}).Post(COLLECT_API)
		}

		sum += len
	}

	return j.AddResult(fmt.Sprintf("😊 Bug收集完成\n🐛 收集Bug: %d", sum))
}
