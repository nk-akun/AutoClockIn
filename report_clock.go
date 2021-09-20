package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/iris-contrib/schema"
)

type Report struct {
	Uid                   string `json:"uid" schema:"uid"`                          // 用户id
	Date                  string `json:"data" schema:"created"`                     // 上报日期
	VaccineRelatedExplain string `json:"vaccine_related_explain" schema:"ymjzxgqk"` // 疫苗接种相关情况
	VaccineInScool        string `json:"vaccine_in_school" schema:"xwxgymjzqk"`     // 校内疫苗相关接种情况  0未接种 1第一针 2第二针 3两针
	BodyHeat              string `json:"body_heat" schema:"tw"`                     // 体温，1:35度以下 2:35-36.5 3:36.6-36.9 ......
	Address               string `json:"address" schema:"address"`                  // 地址
	Area                  string `json:"area" schema:"area"`                        // 地区
	Province              string `json:"province" schema:"province"`                // 省份
	City                  string `json:"city" schema:"city"`                        // 城市
	GeoApiInfo            string `json:"geo_api_info" schema:"geo_api_info"`        // 地理位置信息
	CreateTime            string `json:"create_time" schema:"created"`              // 创建时间
	Sfzx                  string `json:"sfzx" schema:"sfzx"`                        // 不知何意
	SchoolReason          string `json:"school_reason" schema:"fxyy"`               // 返校原因
	HasVaccine            string `json:"has_vaccine" schema:"sfjzxgym"`             // 是否已经接种疫苗
	Sfjzdezxgym           string `json:"sfjzdezxgym" schema:"sfjzdezxgym"`          // 不知何意
	Id                    string `json:"id" schema:"id"`                            // 不知何意 13578779
	EaiSess               string `json:"eai-sess" schema:"-"`                       // 身份标识符
}

var (
	formEncoder = schema.NewEncoder()
)

func DoReport() error {
	log.Printf("开始打卡~")

	reportFile, err := os.Open("/root/go/src/github.com/nk-akun/AutoClockIn/report_conf.json")
	if err != nil {
		log.Printf("open file error=%+v\n", err)
		return err
	}
	defer reportFile.Close()

	// 读取配置文件
	decoder := json.NewDecoder(reportFile)
	stus := []*Report{}
	err = decoder.Decode(&stus)
	if err != nil {
		log.Printf("parse report conf error=%+v\n", err)
		return err
	}

	for _, stu := range stus {
		client := &http.Client{}
		stu.FixParams()

		// 读取form参数
		form := url.Values{}
		err := formEncoder.Encode(stu, form)
		if err != nil {
			log.Printf("form encode error,student=%+v,err=%+v\n", stu, err)
			continue
		}

		// 发送请求
		req, _ := http.NewRequest(http.MethodPost, "https://app.bupt.edu.cn/ncov/wap/default/save", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", fmt.Sprintf("eai-sess=%s", stu.EaiSess))

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("postForm error,student=%+v,err=%+v\n", stu, err)
			continue
		}

		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("打卡日志:student=%+v\nresp=%+v\nresponse_body=%s\n", stu, resp, string(respBody))
	}
	return nil
}

func (stu *Report) FixParams() {
	now := time.Now()
	nowDate := now.Format("20060102")
	nowStamp := now.Unix()
	stu.Date = nowDate
	stu.CreateTime = strconv.FormatInt(nowStamp, 10)
}
