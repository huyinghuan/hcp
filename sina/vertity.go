package sina

import (
  "encoding/json"
  "fmt"
  "log"

  sdk "github.com/SinaCloudStorage/SinaCloudStorage-SDK-Go"
)
type bucket struct {
  ConsumedBytes int64  `json:"ConsumeBytes"`
  CreationDate string  `json:"CreationDate"`
  Name string `json:"Name"`
}
type bucketList struct {
  Owner map[string]string `json:"Owner"`
  Buckets []bucket `json:"Buckets"`
}

func Verify(ak string, sk string, bucket string)error{
  scs :=  sdk.NewSCS(ak, sk, "http://sinacloud.net")
  data, err:=scs.Bucket("").ListBucket()
  if err!=nil{
    log.Panicln(err)
  }
  list := bucketList{}
  json.Unmarshal(data, &list)
  exist := false
  for _, item := range list.Buckets{
    if item.Name == bucket{
      exist = true
      break
    }
  }
  if !exist{
   return fmt.Errorf("找不到Buckets:%s, 请在新浪云控制台创建", bucket)
  }
  return nil
}
