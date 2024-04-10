package sdk

import (
	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Aliyun struct {
	AccessKeyId     string
	AccessKeySecret string
	client          *bssopenapi.Client
}

//var client *bssopenapi.Client

var aliyunClient = &Aliyun{}
var AliyunSDK SDK = aliyunClient

func NewClient(id, secret string) (err error) {
	config := &openapi.Config{
		AccessKeyId:     &id,
		AccessKeySecret: &secret,
	}
	config.Endpoint = tea.String("business.aliyuncs.com")
	aliyunClient.client, err = bssopenapi.NewClient(config)
	if err != nil {
		logrus.Fatal("can not create aliyun client:", err)
	}

	return nil
}

func (a *Aliyun) GetBalance() (float64, error) {
	result, err := aliyunClient.client.QueryAccountBalance()
	if err != nil {
		logrus.Fatal("can not get account balance:", err)
	}
	res, err := strconv.ParseFloat(*result.Body.Data.AvailableAmount, 32)
	if err != nil {
		logrus.Fatal("can not parse account balance:", err)
	}
	return res, nil
}

func (a *Aliyun) GetTodayBill() (float64, error) {

	// 获得今日账单
	month := time.Now().Format("2006-01")
	date := time.Now().Format(time.DateOnly)
	req := &bssopenapi.QueryInstanceBillRequest{
		BillingCycle: tea.String(month),
		BillingDate:  tea.String(date),
		Granularity:  tea.String("DAILY"),
	}
	result, err := aliyunClient.client.QueryInstanceBill(req)
	if err != nil {
		logrus.Fatal("can not get instance bill:", err)
	}
	items := result.Body.Data.Items.Item
	var res float32
	for _, item := range items {
		if *item.ProductCode == "ecs" {
			res += *item.CashAmount
		}
	}
	return float64(res), nil
}
