package verify

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"path"

	"ChallengeCup/config"
	log "ChallengeCup/utils/logger"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/kataras/iris/v12"
	"gopkg.in/gomail.v2"
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

func MailSendCode(ctx iris.Context, email string, code string, timer string) error {
	conf := config.LoadConfig().Mail
	templateFile := "../../resources/template/SendMailCode.liquid"

	binds := map[string]interface{}{
		"code":  code,
		"timer": timer,
	}

	template, _ := template.New(path.Base(templateFile)).Funcs(template.FuncMap{
		"RouteName2URL": ctx.GetCurrentRoute().Path,
	}).ParseFiles(templateFile)

	var tpl bytes.Buffer
	if err := template.Execute(&tpl, binds); err != nil {
		log.Error(err)
		return err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", conf.From)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "验证码")
	message.SetBody("text/html", tpl.String())
	if err := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password).DialAndSend(message); err != nil {
		log.Error(err)
		return err
	}

	log.Infof("send code to %s success", email)

	return nil
}
