package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"

	"github.com/skip2/go-qrcode"
)

func main() {
	qrcodeKey := "d2fdb28c0a60d50a582c4fc9517c0457"
	uid := uint32(20458)
	aid := uint64(4122935097)
	key := genQrcodeKey(uid, aid, "xxxxx")
	fmt.Println(key)
	fmt.Println("is equal:", key == qrcodeKey)
	operatorId := ""     // 如果为空，不添加到参数中
	styleData := ""      // 如果为空，不添加到参数中
	bindSourceType := "" // 如果为空，不添加到参数中
	h5Params := url.Values{}
	h5Params.Set("qrcode_key", qrcodeKey)
	h5Params.Set("advertiser_id", strconv.FormatUint(uint64(uid), 10))
	h5Params.Set("adgroup_id", strconv.FormatUint(aid, 10))
	if operatorId != "" {
		h5Params.Set("operator_id", operatorId)
	}
	if styleData != "" {
		h5Params.Set("style_data", styleData)
	}
	if bindSourceType != "" {
		h5Params.Set("bind_source_type", bindSourceType)
	}
	h5Url := "https://ad.qq.com/atlas/h5_super_preview?" + h5Params.Encode()
	ssoParams := url.Values{}
	ssoParams.Set("service_tag", "61")
	ssoParams.Set("sso_redirect_uri", h5Url)
	ssoUrl := "https://sso.e.qq.com/login?" + ssoParams.Encode()
	fmt.Println("h5Url: ", h5Url)
	fmt.Println("ssoUrl: ", ssoUrl)
	// generateQR("gdt.png", url)
}

// generateQR 将对应的内容生成的文件里, 也可以将其放到字节流里
func generateQR(fileName string, url string) {
	qrCode, _ := qrcode.New(url, qrcode.Medium)
	err := qrCode.WriteFile(256, fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("QR code generated and saved as %v", fileName))
}

func genQrcodeKey(uid uint32, aid uint64, qrCodeSecret string) string {
	qrcodeKey := strconv.FormatUint(uint64(uid), 10) + "-" + strconv.FormatUint(aid, 10) + "-" + qrCodeSecret
	return MD5([]byte(qrcodeKey))
}

func MD5(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}
