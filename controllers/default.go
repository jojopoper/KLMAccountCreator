package controllers

import (
	"jojopoper/KLMAccountCreator/models"
	// "fmt"
	"github.com/astaxie/beego"
)

var createTimes int = 10
var letters string = ""

type MainController struct {
	beego.Controller

	l_IsChinese bool
	l_IsEnglish bool
}

func (this *MainController) Get() {
	ck, err := this.Ctx.Request.Cookie("Language")
	if err != nil {
		this.Data["L_IsChinese"] = true
		this.Data["L_IsEnglish"] = false
	} else {
		blang := ck.Value == "CH"
		this.Data["L_IsChinese"] = blang
		this.Data["L_IsEnglish"] = !blang
	}

	// counts := this.Input().Get("Acc_Count")
	// letters := this.Input().Get("ACC_Letters")

	// fmt.Println(counts)
	// fmt.Println(letters)

	// createTimes, err := strconv.Atoi(counts)
	// fmt.Println(createTimes)
	// if err != nil {
	// 	this.Data["AccountTabelVisible"] = "hidden"
	// 	return
	// }

	var flag chan bool

	flag = make(chan bool)

	var ret bool
	result := make([]*models.AccountInfo, createTimes)
	go models.CreateNewKLMAccount(createTimes, 30, letters, flag, result)

	ret = <-flag
	if !ret {
		this.Data["NumberResult"] = "Timeout!"
	} else {
		this.Data["NumberResult"] = len(result)
	}

	this.Data["AccountTabelVisible"] = "visible"
	this.Data["result"] = result

	this.TplNames = "new_account.html"
}

func (this *MainController) Post() {
	lang := this.Input().Get("languageBtn")
	// this.Data["L_IsChinese"] = true
	// this.Data["L_IsEnglish"] = false
	bLang := lang == "EN"
	this.Data["L_IsEnglish"] = bLang
	this.Data["L_IsChinese"] = !bLang

	this.Ctx.SetCookie("Language", lang, 0)

	// createTimes := 100
	// letters := ""

	var flag chan bool

	flag = make(chan bool)

	var ret bool
	result := make([]*models.AccountInfo, createTimes)
	go models.CreateNewKLMAccount(createTimes, 30, letters, flag, result)

	ret = <-flag
	if !ret {
		this.Data["NumberResult"] = "Timeout!"
	} else {
		this.Data["NumberResult"] = len(result)
	}

	this.Data["AccountTabelVisible"] = "visible"
	this.Data["result"] = result

	this.TplNames = "new_account.html"
}
