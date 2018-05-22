package cloud

import (
  "encoding/json"
  "fmt"
  "log"

  sdk "github.com/SinaCloudStorage/SinaCloudStorage-SDK-Go"
  "time"
)

type sinaBucket struct {
  ConsumedBytes int64  `json:"ConsumeBytes"`
  CreationDate string  `json:"CreationDate"`
  Name string `json:"Name"`
}
type sinaBucketList struct {
  Owner map[string]string `json:"Owner"`
  Buckets []sinaBucket `json:"Buckets"`
}


type sina struct{
	AccessKey string
	SecretKey string
	Bucket string
	Scs *sdk.SCS
}

func (cloud *sina) Init(){
	cloud.Scs = sdk.NewSCS(cloud.AccessKey, cloud.SecretKey, "http://sinacloud.net")
}
func (cloud *sina) Verity(bucket string) error{
	bucketName := cloud.Bucket
	if bucket != ""{
		bucketName = bucket
	}
	log.Println("正在校验数据...")
	data, err:=cloud.Scs.Bucket("").ListBucket()
	if err!=nil{
		log.Panicln(err)
	}
	list := sinaBucketList{}
	json.Unmarshal(data, &list)
	exist := false
	for _, item := range list.Buckets{
		if item.Name == bucketName{
			exist = true
			break
		}
	}
	if !exist{
		return fmt.Errorf("找不到Buckets:%s, 请在新浪云控制台创建", bucketName)
	}
	log.Println("校验完成")
	return nil
}
func (cloud *sina) Put(fileName string, remoteFilePath string, params *PutParams) (string, error){
	bucketName := cloud.Bucket

	if params.BucketName != ""{
		bucketName = params.BucketName
	}
	ALC := sdk.Private
	if params.Public {
		ALC = sdk.PublicRead
	}

	bucket := cloud.Scs.Bucket(bucketName)

	if params.Expire {
		if err := bucket.PutExpire(fileName, remoteFilePath, ALC, time.Now().Add(60*time.Second)); err!=nil{
			return "", err
		}
	}else{
		if err := bucket.Put(fileName, remoteFilePath, ALC);err!=nil{
			return "", err
		}
	}
	return fmt.Sprintf("https://%s.sinacloud.net%s", bucketName, remoteFilePath), nil
}


func getSinaService(accessKey string, secretKey string, bucket string) *sina{
	s:= &sina{
		AccessKey:accessKey,
		SecretKey:secretKey,
		Bucket: bucket,
	}
	s.Init()
	return s
}