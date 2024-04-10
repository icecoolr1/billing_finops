package sdk

type SDK interface {
	// GetBalance 获得用户余额
	GetBalance() (float64, error)
	// GetTodayBill 获得今日账单
	GetTodayBill() (float64, error)
}
