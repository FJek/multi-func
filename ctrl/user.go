package ctrl

import (
	"fmt"
	"fzw/proto/util"
	"image/color"
	"net/http"

	"github.com/skip2/go-qrcode"
)

// 生成用户二维码
func GenerQrcode() {
	url := ""
	err := qrcode.WriteFile(url, qrcode.Medium, 256, `qr1.png`)
	if err != nil {
		fmt.Println("二维码生成失败！！")
		return
	}
}

//生成彩色二维码
func GenerQrcodeColor() {
	// 参数：URL、容错级别、二维码宽高、背景色、前景色、保存路径
	url := ""
	err := qrcode.WriteColorFile(url, qrcode.Medium, 256, color.Black, color.White, `qr2.png`)
	if err != nil {
		fmt.Println(err)
	}
}

// 获取用户ip地址
func GetUserIp(w http.ResponseWriter, r *http.Request) {
	ip := util.GetIp(w, r)
	fmt.Printf("ip: %s\n", ip)
}
