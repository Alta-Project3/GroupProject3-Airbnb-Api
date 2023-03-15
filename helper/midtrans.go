package helper

import (
	"github.com/Findryankp/snapMidtransGo"
)

var ServerKey = ""

func PostMidtrans(data map[string]any) (string, error) {
	var postData = snapMidtransGo.DataPostMidtrans{
		OrderId:   data["order_id"].(string),
		Nominal:   data["nominal"].(int),
		FirstName: data["firstname"].(string),
		LastName:  data["lastname"].(string),
		Email:     data["email"].(string),
		Phone:     data["phone"].(string),
		ServerKey: ServerKey,
	}

	test, err := snapMidtransGo.SanboxRequestSnapMidtrans(postData)
	return test, err
}
