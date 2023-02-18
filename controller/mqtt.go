package controller

import (
	"ChallengeCup/common"
	"ChallengeCup/model"
	"ChallengeCup/service"

	"github.com/kataras/iris/v12"
)

func MqttSubscribe(ctx iris.Context) {
}

func MqttPublish(ctx iris.Context) {
}

// func GetData(ctx iris.Context) {
// 	data := service.Service.MqttService.GetDataNow()
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetDataByTime(ctx iris.Context) {
// 	start := ctx.FormValue("start")
// 	end := ctx.FormValue("end")
// 	data := service.Service.MqttService.GetDataByTime(start, end)
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetTemperature(ctx iris.Context) {
// 	data := service.Service.MqttService.GetTemperatureNow()
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetTemperatureByTime(ctx iris.Context) {
// 	start := ctx.FormValue("start")
// 	end := ctx.FormValue("end")
// 	data := service.Service.MqttService.GetDataByTime(start, end)
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetHeartRate(ctx iris.Context) {
// 	data := service.Service.MqttService.GetTemperatureNow()
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetHeartRateByTime(ctx iris.Context) {
// 	start := ctx.FormValue("start")
// 	end := ctx.FormValue("end")
// 	data := service.Service.MqttService.GetDataByTime(start, end)
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetBloodOxygen(ctx iris.Context) {
// 	data := service.Service.MqttService.GetTemperatureNow()
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

// func GetBloodOxygenByTime(ctx iris.Context) {
// 	start := ctx.FormValue("start")
// 	end := ctx.FormValue("end")
// 	data := service.Service.MqttService.GetDataByTime(start, end)
// 	ctx.JSON(model.Result{
// 		Code:    common.SUCCESS,
// 		Message: common.Message(common.SUCCESS),
// 		Data:    data,
// 	})
// }

func GetMqttData(ctx iris.Context) {
	dataType := ctx.FormValue("type")
	streamData := ctx.FormValue("stream")
	start := ctx.FormValue("start")
	end := ctx.FormValue("end")
	// start是起始时间戳, end是结束时间戳
	// stream是是否实时获取数据, true是实时获取, false是获取历史数据
	var data interface{}
	switch streamData {
	case "true":
		switch dataType {
		case "temperature":
			data = service.Service.MqttService.GetTemperatureNow()
		case "heart_rate":
			data = service.Service.MqttService.GetHeartRateNow()
		case "blood_oxygen":
			data = service.Service.MqttService.GetBloodOxygenNow()
		case "all":
			data = service.Service.MqttService.GetDataNow()
		default:
			data = nil
			ctx.JSON(model.Result{
				Code:    common.CLIENT_ERROR,
				Message: common.Message(common.CLIENT_ERROR),
				Data:    data,
			})
			return
		}
	case "false":
		switch dataType {
		case "temperature":
			data = service.Service.MqttService.GetTemperatureByTime(start, end)
		case "heart_rate":
			data = service.Service.MqttService.GetHeartRateByTime(start, end)
		case "blood_oxygen":
			data = service.Service.MqttService.GetBloodOxygenByTime(start, end)
		case "all":
			data = service.Service.MqttService.GetDataByTime(start, end)
		default:
			data = nil
			ctx.JSON(model.Result{
				Code:    common.CLIENT_ERROR,
				Message: common.Message(common.CLIENT_ERROR),
				Data:    data,
			})
			return
		}

	}
	ctx.JSON(model.Result{
		Code:    common.SUCCESS,
		Message: common.Message(common.SUCCESS),
		Data:    data,
	})
}
