package ceph

import (
	"Goshop/utils/yml_config"
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var (
	cephConn   *s3.S3
	GWEndpoint = yml_config.CreateYamlFactory().GetString("Store.Ceph.CephGWEndpoint")
	SecretKey  = yml_config.CreateYamlFactory().GetString("Store.Ceph.CephSecretKey")
	AccessKey  = yml_config.CreateYamlFactory().GetString("Store.Ceph.CephAccessKey")
)

// GetCephConnection : 获取ceph连接
func GetCephConnection() *s3.S3 {
	if cephConn != nil {
		return cephConn
	}
	// 1. 初始化ceph的一些信息

	auth := aws.Auth{
		AccessKey: AccessKey,
		SecretKey: SecretKey,
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          GWEndpoint,
		S3Endpoint:           GWEndpoint,
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	// 2. 创建S3类型的连接
	return s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的bucket对象
func GetCephBucket(bucket string) *s3.Bucket {
	conn := GetCephConnection()
	return conn.Bucket(bucket)
}

// PutObject : 上传文件到ceph集群
func PutObject(bucket string, path string, data []byte) error {
	return GetCephBucket(bucket).Put(path, data, "octet-stream", s3.PublicRead)
}
