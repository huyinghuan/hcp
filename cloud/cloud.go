package cloud

import (
	"time"
	"log"
	"hcp/dir"
)

type Cloud interface {
	Init()
	//基本校验
	Verity(bucketName string)(error)
	// Put 文件上传
	Put(fileName string, remoteFilePath string, params *PutParams) (string, error)
}

type PutParams struct{
	BucketName string
	Public bool
	Expire bool
	ExpireDate time.Time
}


func GetService(cloudName string, config map[string]string) Cloud{
	id := config["id"]
	secretKey:=config["secretKey"]
	bucket:=config["bucket"]
	switch cloudName {
	case "sina": return getSinaService(id, secretKey, bucket)
	}
	log.Panicf("找不到云服务商: %s", cloudName)
	return nil
}


func UploadFiles(souceFilesList []string, targetFilePath string){
	files:=dir.GetTargetFiles(souceFilesList)

}