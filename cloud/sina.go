package cloud

import (
  "encoding/json"
  "fmt"
  "log"

  sdk "github.com/SinaCloudStorage/SinaCloudStorage-SDK-Go"
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

func GetSinaService(accessKey string, secretKey string, bucket string)*sina{
  s:= &sina{
    AccessKey:accessKey,
    SecretKey:secretKey,
    Bucket: bucket,
  }
  s.Init()
  return s
}

func (cloud *sina) Init()  {
  cloud.Scs =  sdk.NewSCS(cloud.AccessKey, cloud.SecretKey, "http://sinacloud.net")
}

func (cloud *sina) Verify(bucketName string) error{
  data, err:=cloud.Scs.Bucket("").ListBucket()
  if err!=nil{
    log.Panicln(err)
  }
  list := sinaBucketList{}
  json.Unmarshal(data, &list)
  exist := false
  for _, item := range list.Buckets{
    if item.Name == cloud.Bucket{
      exist = true
      break
    }
  }
  if !exist{
   return fmt.Errorf("找不到Buckets:%s, 请在新浪云控制台创建", cloud.Bucket)
  }
  return nil
}

func (cloud *sina) Put(bucketName string) (string, error){
  if bucketName == ""{
    bucketName = cloud.Bucket
  }
  bucket := cloud.Scs.Bucket(bucketName)

  return "", nil
}