/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type FundStatementsDetailedGetExample struct {
	TAds                          *ads.SDKClient
	AccessToken                   string
	AdvertiserId                  int64
	FundType                      string
	DateRange                     model.DateRangeTransaction
	FundStatementsDetailedGetOpts *api.FundStatementsDetailedGetOpts
}

func (e *FundStatementsDetailedGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AdvertiserId = 789
	e.FundType = "FUND_TYPE_GIFT"
	e.DateRange = model.DateRangeTransaction{
		StartDate: "REPORT START DATE",
		EndDate:   "REPORT END DATE",
	}
	e.FundStatementsDetailedGetOpts = &api.FundStatementsDetailedGetOpts{

		Fields: optional.NewInterface([]string{"time", "external_bill_no", "trade_type", "amount", "description"}),
	}
}

func (e *FundStatementsDetailedGetExample) RunExample() (model.FundStatementsDetailedGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.FundStatementsDetailed().Get(ctx, e.AdvertiserId, e.FundType, e.DateRange, e.FundStatementsDetailedGetOpts)
}

func main() {
	e := &FundStatementsDetailedGetExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
