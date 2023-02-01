package verify

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"ChallengeCup/config"
	log "ChallengeCup/utils/logger"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func RandomCode() string {
	randomCode, err := rand.Prime(rand.Reader, 32)
	if err != nil {
		fmt.Println(err)
	}

	return randomCode.String()[0:6]
}

func PhoneSendCode(phone string, code string) error {
	conf := config.LoadConfig().SMS
	client, err := sdk.NewClientWithAccessKey(conf.RegionId, conf.AccessKeyID, conf.AppSecret)
	if err != nil {
		log.Errorf("sms connect error: ", err)
		return err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.ApiName = "SendSms"
	request.Version = "2017-05-25"
	request.RegionId = conf.RegionId
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["SignName"] = conf.SignName
	request.QueryParams["TemplateCode"] = conf.TemplateCode
	request.QueryParams["TemplateParam"] = "{\"code\":\"" + code + "\"}"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Errorf("sms send error: ", err)
		return err
	}

	message := struct {
		Code      string `json:"Code"`
		Message   string `json:"Message"`
		BizId     string `json:"BizId"`
		RequestId string `json:"RequestId"`
	}{}

	_ = json.Unmarshal(response.GetHttpContentBytes(), &message)
	if message.Code != "OK" {
		log.Errorf("sms send error: ", message.Message)
		return fmt.Errorf(message.Message)
	}

	log.Infof("sms send success: ", response.GetHttpContentString())

	return nil
}
