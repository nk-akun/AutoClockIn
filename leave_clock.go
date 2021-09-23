package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Doleave() {
	log.Printf("开始申请~")

	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	goutTimeStr := now.Add(-8 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	backTime, _ := time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02")+" 15:57:41", loc)
	backTimeStr := backTime.Format("2006-01-02T15:04:05.000Z")
	unknownTime, _ := time.ParseInLocation("2006-01-02 15:04:05", now.AddDate(0, 0, -1).Format("2006-01-02")+" 16:00:00", loc)
	unknownTimeStr := unknownTime.Format("2006-01-02T15:04:05.000Z")

	data := fmt.Sprintf("{\"app_id\":\"578\",\"form_data\":{\"1716\":{\"User_5\":\"马荣坤\",\"User_7\":\"2020141056\",\"User_9\":\"计算机学院（国家示范性软件学院）\",\"User_11\":\"13210123307\",\"SelectV2_58\":[{\"name\":\"西土城校区\",\"value\":\"2\",\"default\":0,\"imgdata\":\"\"}],\"UserSearch_60\":{\"uid\":72870,\"name\":\"付泓霖\",\"number\":\"2010813766\"},\"Calendar_62\":%s,\"Calendar_50\":%s,\"Calendar_47\":%s,\"Input_28\":\"银行\",\"MultiInput_30\":\"取钱\",\"Radio_52\":{\"value\":\"1\",\"name\":\"本人已阅读并承诺\"},\"Validate_63\":\"\",\"Alert_65\":\"\",\"Validate_66\":\"\",\"Alert_67\":\"\"}}}", unknownTimeStr, goutTimeStr, backTimeStr)
	starterDepartId := "181789"

	form := url.Values{}
	form["data"] = []string{data}
	form["starter_depart_id"] = []string{starterDepartId}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "https://service.bupt.edu.cn/site/apps/launch", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "_ga=GA1.3.1487030133.1601201572;PHPSESSID=ST-3093201-FhytAodxpKecDeFmxbma-IfHi-cas-1631548189222;vjuid=196680;vjvd=7027e8c1faeec374cfa5332a1abd51c4;vt=147268262")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("Referer", "https://service.bupt.edu.cn/v2/matter/m_start?id=578")
	req.Header.Set("Origin", "https://service.bupt.edu.cn")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("postForm error,err=%+v\n", err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Printf("打卡日志:resp=%+v\nresponse_body=%s\n", resp, string(respBody))
}
