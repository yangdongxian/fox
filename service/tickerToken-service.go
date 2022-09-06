package service

import (
	"context"
	"encoding/json"
	"fmt"
	"fox/common"
	"fox/dto"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type TickerToken interface {
	GetRedisAccessToken() (token string)

	ExecuteGenerateToken()
}

type tickerTokenService struct {
	rdb          *redis.Client
	tokenService WxTokenService
}

func NewTickerTokenService(r *redis.Client, token WxTokenService) TickerToken {
	wx := tickerTokenService{rdb: r, tokenService: token}
	wx.tickerAccessToken()

	return &tickerTokenService{
		rdb:          r,
		tokenService: token,
	}
}

func (ticker *tickerTokenService) generateToken() []byte {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env files")
	}
	Appid := os.Getenv("APPID")
	Secret := os.Getenv("SECRET")
	GrantType := os.Getenv("GrantType")

	u, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		log.Fatal(err)
	}
	params := &url.Values{}
	params.Set("appid", Appid)
	params.Set("secret", Secret)
	params.Set("grant_type", GrantType)

	u.RawQuery = params.Encode()

	resp, resErr := http.Get(u.String())
	if err != nil {
		log.Fatal(resErr)
	}
	respByte, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Printf("request token is failed,error:%v resByte:%v", readErr, string(respByte))
	}
	now := time.Now()
	currentTime := now.Format("2006/01/02 15:04:05.000")
	fmt.Printf("%v  generateToken  response:%v \n", currentTime, string(respByte))
	return respByte
}
func (ticker *tickerTokenService) saveAccessToken(res []byte) {
	ctx := context.Background()
	// response is contain errcode and errmsg key
	// {"errcode": 40125, "errmsg": ""}
	var wxTokenDTO dto.WxTokenCreateDTO
	resString := string(res) //[]byte转string
	resMap := make(map[string]interface{})
	unErr := json.Unmarshal(res, &resMap)
	if unErr != nil {
		log.Fatal("json.Unmarshal is error:", unErr)
	}
	if strings.Contains(resString, "errcode") || strings.Contains(resString, "errmsg") { //异常返回：
		log.Fatal("request wxtoken is error:", resString)
		return
	} else { //正常返回
		//access_token = resMap["access_token"].(string)
		//expiresIn := resMap["expires_in"].(float64)
		//expiresInInt64 := int64(expiresIn * math.Pow10(0))
		//
		wxTokenDTO.AccessToken = resMap["access_token"].(string)
		expiresFloat64 := resMap["expires_in"].(float64)
		expiresInInt64 := int64(expiresFloat64 * math.Pow10(0))
		wxTokenDTO.ExpiresIn = expiresInInt64
	}
	startTime := time.Now().Unix()

	result, err := ticker.rdb.HSet(ctx, "wxAccessToken", "access_token", wxTokenDTO.AccessToken, "expires_in", wxTokenDTO.ExpiresIn, "start_time", startTime).Result()
	if err != nil {
		fmt.Printf("saveAccessToken rdb.set err:%v result:%v", err, result)
	}
	wxTokenDTO.ID = uuid.NewV4().String()
	wxTokenDTO.CreateAt = time.Now()
	wxTokenDTO.UpdateAt = time.Now()
	wxTokenDTO.IsDelete = false
	wxTokenDTO.DeleteAt = time.Now()
	ticker.tokenService.Insert(wxTokenDTO)
}
func (ticker *tickerTokenService) GetRedisAccessToken() (token string) {
	ctx := context.Background()
	accessToken, err := ticker.rdb.HGet(ctx, common.WxAccessToken, "access_token").Result()
	if err != nil {
		log.Fatal(err)
	}
	return accessToken
}

//func (ticker *tickerTokenService) GetDBAccessToken() (token string) {
//	token = ticker.tokenService.FindByAccessToken().AccessToken
//	return token
//}

func (t *tickerTokenService) ExecuteGenerateToken() {
	res := t.generateToken()
	t.saveAccessToken(res)
}
func (t *tickerTokenService) tickerAccessToken() {
	//ticker := time.Tick(120 * time.Second)
	//go func() {
	//	for range ticker {
	//		res := wx.generateToken()
	//		wx.saveAccessToken(res)
	//	}
	//}()

	ticker := time.NewTicker(common.GapTime * time.Second)
	t.ExecuteGenerateToken()
	go func() {
		for range ticker.C {
			t.ExecuteGenerateToken()
		}
	}()

}
