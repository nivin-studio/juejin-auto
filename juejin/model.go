package juejin

type Response struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

type CheckIn struct {
	IncrPoint int `json:"incr_point"`
	SumPoint  int `json:"sum_point"`
}

type LotteryDraw struct {
	Id              int    `json:"id"`
	LotteryId       string `json:"lottery_id"`
	LotteryName     string `json:"lottery_name"`
	LotteryType     int    `json:"lottery_type"`
	LotteryImage    string `json:"lottery_image"`
	LotteryReality  int    `json:"lottery_reality"`
	LotteryDesc     string `json:"lottery_desc"`
	LotteryCost     int    `json:"lottery_cost"`
	DrawLuckyValue  int    `json:"draw_lucky_value"`
	TotalLuckyValue int    `json:"total_lucky_value"`
	HistoryId       string `json:"history_id"`
}

type LotteryHistory struct {
	LuckyUser []LuckyUser `json:"lotteries"`
	Count     int         `json:"count"`
}

type LuckyUser struct {
	UserID            string        `json:"user_id"`
	HistoryID         string        `json:"history_id"`
	UserName          string        `json:"user_name"`
	UserAvatar        string        `json:"user_avatar"`
	LotteryName       string        `json:"lottery_name"`
	LotteryImage      string        `json:"lottery_image"`
	Date              int           `json:"date"`
	DipLuckyUserCount int           `json:"dip_lucky_user_count"`
	DipLuckyUsers     []interface{} `json:"dip_lucky_users"`
}

type DipLucky struct {
	DipAction  int  `json:"dip_action"`
	HasDip     bool `json:"has_dip"`
	TotalValue int  `json:"total_value"`
	DipValue   int  `json:"dip_value"`
}

type Bug struct {
	BugType     int  `json:"bug_type"`
	BugTime     int  `json:"bug_time"`
	BugShowType int  `json:"bug_show_type"`
	IsFirst     bool `json:"is_first"`
}
