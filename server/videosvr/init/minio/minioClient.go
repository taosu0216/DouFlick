package minio

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"strings"
	"sync"
	"videosvr/config"
	"videosvr/init/snowflake"
	"videosvr/log"
)

type Minio struct {
	MinioClient  *minio.Client
	EndPoint     string
	Port         string
	VideoBuckets string
	PicBuckets   string
}

var (
	Client    Minio
	OnceMinio sync.Once
)

func GetMinio() Minio {
	OnceMinio.Do(func() {
		initMinio()
	})
	return Client
}

func initMinio() {
	info := config.GetGlobalConfig().MinioConfig
	endpoint := info.Host
	port := info.Port
	endpoint = endpoint + ":" + port
	accessKeyID := info.AccessKeyId
	secretKey := info.SecretAccessKey
	videoBucket := info.VideoBuckets
	pictureBucket := info.PicBuckets

	//初始化minio的Client
	minioClient, err := minio.New(endpoint, accessKeyID, secretKey, false)
	if err != nil {
		panic(err)
	}
	//创建存储桶
	createBucket(minioClient, videoBucket)
	createBucket(minioClient, pictureBucket)
	Client = Minio{
		MinioClient:  minioClient,
		EndPoint:     endpoint,
		Port:         port,
		VideoBuckets: videoBucket,
		PicBuckets:   pictureBucket,
	}
}

func createBucket(m *minio.Client, bucketName string) {
	exists, err := m.BucketExists(bucketName)
	if err != nil {
		log.Error("check bucket err:%s", err.Error())
	}
	if !exists {
		err = m.MakeBucket(bucketName, "us-east-1")
		if err != nil {
			panic(err)
		}
	}

	//设置桶策略
	policy := `{"Version": "2012-10-17",
				"Statement": 
					[{
						"Action":["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucketName + `/*"],
						"Sid": ""
					}]
				}`
	err = m.SetBucketPolicy(bucketName, policy)
	if err != nil {
		log.Error("set bucket policy err:%s", err.Error())
	}
	log.Infof("set bucket %s success", bucketName)
}

func (m *Minio) UploadFile(fileType, file, userID string) (string, error) {
	var filename strings.Builder
	var contentType, Suffix, bucket string
	if fileType == "video" {
		fmt.Println("upload video")
		contentType = "video/mp4"
		Suffix = ".mp4"
		bucket = m.VideoBuckets
	} else {
		fmt.Println("upload pic")
		contentType = "image/jpeg"
		Suffix = ".jpg"
		bucket = m.PicBuckets
	}
	filename.WriteString(userID)
	filename.WriteString("_")
	snowFlakeID := snowflake.GenerateID()
	filename.WriteString(snowFlakeID)
	filename.WriteString(Suffix)
	fmt.Println(bucket, filename.String(), file)
	n, err := m.MinioClient.FPutObject(bucket, filename.String(), file, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error("upload file err:%s", err.Error())
		return "", err
	}
	log.Infof("upload %d bytes file success, filename:%s", n, filename.String())
	url := "http://" + m.EndPoint + "/" + bucket + "/" + filename.String()
	return url, nil
}
