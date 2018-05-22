package action

import (
	"github.com/urfave/cli"
	"fmt"
	"strings"
	"strconv"
	"hcp/utils"
	"hcp/cloud"
)
var(
	Cloud = []string{"[1]阿里云", "[2]腾讯云", "[3]新浪云"}
	CloudMap =[]string{"aliyun", "tencent", "sina"}

	DescribeArray = []string{"id", "secretKey", "bucket"}
	Describe = map[string]map[string]string{
		"aliyun": map[string]string{
			"id": "请输入SecretId:  ",
			"secretKey": "请输入SecretKey:  ",
			"bucket": "请输入Bucket名称(xxxx-125xxx):  ",
		},
		"tencent": map[string]string{
			"id": "请输入SecretId:  ",
			"secretKey": "请输入SecretKey:  ",
			"bucket": "请输入Bucket名称(xxxx-125xxx):  ",
		},
		"sina": map[string]string{
			"id": "请输入Access Key:  ",
			"secretKey": "请输入Secret Key:  ",
			"bucket": "请输入Bucket名称:  ",
		},
	}
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


func getStringInput(desc string) string{
	target := ""
	for true{
		fmt.Print(desc)
		if _, err :=fmt.Scanln(&target); err!=nil{
			continue
		}else{
			break
		}
	}
	return strings.TrimSpace(target)
}

func InitAction(c *cli.Context) (err error){
	cloudName:=getCloud()
	descMap := Describe[cloudName]
	inputMap := map[string]string{}
	for _, key := range DescribeArray{
		inputMap[key] = getStringInput(descMap[key])
	}
	service:= cloud.GetService(cloudName, inputMap)
	service.Verity("")
	utils.WriteConfig(cloudName, inputMap["id"], inputMap["secretKey"], inputMap["bucket"])
	fmt.Println("初始化完成。。")
	return nil
}
