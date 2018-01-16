package base

import (
	"strings"
	"time"
	"encoding/base64"
)

type Info struct {
	Code int32
	Msg string
}

func Success() Info  {
	return Info{Code:22000, Msg:"success"}
}


func LoginParamError() Info {
	return Info{Code:23000, Msg:"params error"}
}

func LoginTokenError() Info  {
	return Info{Code:23001, Msg:"token error"}
}

func RegisterParamError() Info {
	return Info{Code:23002, Msg:"params error"}
}

func GenToken(args ...string) string  {
	res := strings.Join(args, "")
	res = res + time.Now().String()
	s := []byte(res)
	return base64.StdEncoding.EncodeToString(s)
}