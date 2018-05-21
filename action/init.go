package action

import (
	"github.com/urfave/cli"
	"fmt"
	"strings"
	"strconv"
	"log"
)
var(
	Cloud = []string{"[1]阿里云", "[2]腾讯云", "[3]新浪云"}
	CloudMap =[]string{"aliyun", "tencent", "sina"}
)


func getCloud() string{
	flag := true
	target := ""

	for flag{
		fmt.Printf("选择云服务商(默认[1]):\n%s:", strings.Join(Cloud, ","))
		if _, err :=fmt.Scanln(&target); err!=nil{
			return CloudMap[0]
		}
		if server, err:=strconv.Atoi(target); err == nil{
			if server > 0 && server <= len(Cloud){
				return CloudMap[server-1]
			}
		}
	}
	return  CloudMap[0]
}

func InitAction(c *cli.Context) (err error){
	cloud:=getCloud()
	log.Println(cloud)
	return nil
}
