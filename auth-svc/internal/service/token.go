package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"auth-svc/internal/config"
	"auth-svc/internal/consts"
	"auth-svc/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	Gen() (token string, err error)
	Verify(token string) (err error)
}

type AccessToken struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

func (ac *AccessToken) Gen() (token string, err error) {
	cf := config.GetConfig()
	expTime := gtime.Now().Add(time.Duration(cf.Auth.AccessTokenExpireMinute) * time.Minute)
	ac.Exp = expTime.Unix()

	mapClaims := jwt.MapClaims{}
	dataByte, _ := json.Marshal(ac)
	json.Unmarshal(dataByte, &mapClaims)

	token, err = utility.GenJWT(mapClaims)
	return
}

func (ac *AccessToken) Verify(ctx context.Context, token string) (err error) {
	jwtMap, err := utility.ParseJWT(token)
	if err != nil {
		return err
	}

	dataByte, _ := json.Marshal(jwtMap)
	err = json.Unmarshal(dataByte, &ac)
	if err != nil {
		return err
	}

	key := consts.RedisKeyRefreshToken.Key(ac.Iss)
	err = Redis().CheckKeyExist(ctx, key)
	if err != nil {
		return gerror.NewCode(consts.CodeInvalidToken)
	}
	return
}

type RefreshToken struct {
	Ctx  context.Context
	Uuid string
	Uid  string
	Exp  string
	Nbf  string
}

func (rf *RefreshToken) CheckCtx() (err error) {
	if rf.Ctx == nil {
		return gerror.NewCode(consts.CodeInvalidToken)
	}
	return
}

func (rf *RefreshToken) Gen() (token string, err error) {
	cf := config.GetConfig()
	expTime := gtime.Now().Add(time.Duration(cf.Auth.RefreshTokenExpireMinute) * time.Minute)
	rf.Exp = expTime.String()
	rf.Nbf = fmt.Sprintf("%v", expTime.Unix())
	err = Redis().
		HSetEx(rf.Ctx, consts.RedisKeyRefreshToken.Key(fmt.Sprintf("%v", rf.Uuid)), g.Map{
			"Uid": rf.Uid,
			"Exp": rf.Exp,
			"Nbf": rf.Nbf,
		}, uint(cf.Auth.RefreshTokenExpireMinute)*60)
	if err != nil {
		return
	}
	token = fmt.Sprintf("%v", rf.Uuid)
	return
}

func (rf *RefreshToken) Verify(token string) (err error) {
	// Check key exists
	key := consts.RedisKeyRefreshToken.Key(token)
	err = Redis().CheckKeyExist(rf.Ctx, key)
	if err != nil {
		return gerror.NewCode(consts.CodeInvalidToken)
	}
	// Check expired
	rVal, err := Redis().HGet(rf.Ctx, key)
	if err != nil {
		return
	}
	err = rVal.Scan(rf)
	if err != nil {
		return
	}
	if gtime.Now().Unix() >= utility.String2Int64(rf.Nbf) {
		return gerror.NewCode(consts.CodeTokenExpired)
	}
	return
}
