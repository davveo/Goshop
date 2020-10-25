package oss

import (
	"Goshop/utils/yml_config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var (
	ossCli          *oss.Client
	endPoint        = yml_config.CreateYamlFactory().GetString("Store.Oss.OSSEndpoint")
	accessKeyId     = yml_config.CreateYamlFactory().GetString("Store.Oss.OSSAccesskeyID")
	accessKeySecret = yml_config.CreateYamlFactory().GetString("Store.Oss.OSSAccessKeySecret")
	bucket          = yml_config.CreateYamlFactory().GetString("Store.Oss.OSSBucket")
)

func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}

	ossCli, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return ossCli
}

// Bucket : 获取bucket存储空间
func Bucket() *oss.Bucket {
	cli := Client()
	if cli != nil {
		bucket, err := cli.Bucket(bucket)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		return bucket
	}
	return nil
}

// DownloadURL : 临时授权下载url
func DownloadURL(objName string) string {
	signedURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedURL
}

// BuildLifecycleRule : 针对指定bucket设置生命周期规则
func BuildLifecycleRule(bucketName string) {
	// 表示前缀为test的对象(文件)距最后修改时间30天后过期。
	ruleTest1 := oss.BuildLifecycleRuleByDays("rule1", "test/", true, 30)
	rules := []oss.LifecycleRule{ruleTest1}

	_ = Client().SetBucketLifecycle(bucketName, rules)
}

// GenFileMeta :  构造文件元信息
func GenFileMeta(metas map[string]string) []oss.Option {
	var options []oss.Option
	for k, v := range metas {
		options = append(options, oss.Meta(k, v))
	}
	return options
}
