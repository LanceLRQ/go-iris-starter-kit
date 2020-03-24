package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/kataras/iris"
	"github.com/lancelrq/go-iris-starter-kit/src/structs"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 进行签名运算
func GenerateSignatures(params *structs.QuerySignatureParams, token string) {
	s256 := sha256.New()
	msg := []string { token, params.Nonce, params.TimeStamp }
	sort.Strings(msg)
	msgJoined := []byte(strings.Join(msg, ""))
	s256.Write(msgJoined)
	hex := s256.Sum(nil)
	params.Signature = fmt.Sprintf("%x", hex)
	return
}

// 检查签名
func CheckSingatures(signature string, params *structs.QuerySignatureParams, token string) bool {
	GenerateSignatures(params, token)
	return strings.Trim(signature, " ") != "" && signature == params.Signature
}

// 从Iris的上下文里获取并校验签名信息
func CheckSignaturesFromIrisContext(ctx iris.Context, token string) bool {
	nonce := ctx.URLParamDefault("nonce", "")
	timeStamp := ctx.URLParamDefault("timestamp", "")
	sign := ctx.URLParamDefault("signature", "")

	timeStampNumber, err := strconv.ParseInt(timeStamp, 10, 64)
	nowTime := time.Now().Unix()
	if err != nil || math.Abs(float64(timeStampNumber - nowTime)) > 600 {
		return false
	}

	success := CheckSingatures(sign, &structs.QuerySignatureParams{
		Nonce: nonce,
		TimeStamp: timeStamp,
	}, token)

	return success
}

// 创建签名参数
func CreateSignatures() (params structs.QuerySignatureParams, token string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	nonceBytes := make([]byte, 16)
	for i := 0 ; i < 16 ; i++ {
		nonceBytes[i] = byte(r.Intn(26) + 97)
	}
	nonce := string(nonceBytes)
	timeStamp := fmt.Sprintf("%d", time.Now().Unix())
	params = structs.QuerySignatureParams{
		Nonce: nonce,
		TimeStamp: timeStamp,
	}
	GenerateSignatures(&params, token)
	return
}