package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"uvm-backend/model"
	"uvm-backend/utils"
)

//// wx.getUserInfo得到的完整信息
//type WXUser struct {
//	UserInfo      NonSensitiveUser `json:"user_info"`
//	RawData       string     `json:"raw_data"`
//	Signature     string     `json:"signature"`
//	EncryptedData string     `json:"encrypted_data"`
//	IV            string     `json:"iv"`
//}
//// WXUserInfo返回的不包含敏感信息的Non-Sensitive UserInfo
//type NonSensitiveUser struct {
//	OpenID		string		  `json:"open_id, omitempty"`
//	Name        string		  `json:"nick_name"`
//	UnionID		string		  `json:"union_id, omitempty"`
//	AvatarUrl   string 		  `json:"avatar_url"`
//	Gender      int       	  `json:"gender"`
//	Country     string    	  `json:"country"`
//	Province    string    	  `json:"province"`
//	City        string    	  `json:"city"`
//	Language    string        `json:"language"`
//}

// 微信接口服务返回值
// 注意查看api定义的json值
type WXLoginResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

/**
输入小程序发送的code，返回请求微信接口服务进而得到的openID, unionID, session-key和错误情况
目前不需要微信的UserInfo
*/
func WXLogin(code string) (w *WXLoginResp, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.WXLogin: %w", err)
		}
	}()
	appId := "wx7246396bd244fb02"
	appSecret := "ec6c4b6cdbdf81250cf1bc2f9e2e8860"
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 合成url, appId和appSecret可以直接得到
	url = fmt.Sprintf(url, appId, appSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	// 解析body到结构体中
	w = &WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(w); err != nil {
		log.Println(err)
		return nil, err
	}

	// 判断微信接口服务返回的是否为异常情况
	if w.ErrCode != 0 {
		err = errors.New(fmt.Sprintf("ErrCode: %d ErrMsg: %s", w.ErrCode, w.ErrMsg))
		return nil, err
	}
	// 解析
	return w, err
}

/**
用户登录
*/
func UserLogin(openID string) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.UserLogin: %w", err)
		}
	}()
	user := &model.User{
		WXOpenId: openID,
	}
	u, err := user.GetUserByOpenId()
	if err == gorm.ErrRecordNotFound {
		// 没有该用户记录，则需初始化该用户
		user.Name = utils.GetUUID()
		user.BusinessId = 1
		id, err = user.AddUser()
	}
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// 已有该用户记录
	return u.Id, nil
}
