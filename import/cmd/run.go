package cmd

import (
	"errors"
	"fmt"
	"github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/service"
	qs "github.com/yunify/qingstor-sdk-go/service"
	"io"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var setupLog = ctrl.Log.WithName("setup")

func subMain() error {
	ctrl.SetLogger(zap.New(zap.UseDevMode(QSconfig.development)))

	if QSconfig.AccessKey == "" && QSconfig.SecretKey == "" {
		err := errors.New("AccessKey and SecretKey cannot be empty string")
		setupLog.Error(err, "AccessKey and SecretKey cannot be empty string")
		return err
	}
	downloadConfig, err := config.New(QSconfig.AccessKey, QSconfig.SecretKey)
	if err != nil {
		setupLog.Error(err, "create a config failed")
		return err
	}
	qsService, _ := qs.Init(downloadConfig)
	bucket, _ := qsService.Bucket(QSconfig.BucketName, QSconfig.Zone)

	bOutput, err := bucket.ListObjects(&qs.ListObjectsInput{Prefix: &QSconfig.BucketPath})

	for _, obj := range bOutput.Keys {
		err := downloadSQL(bucket, obj)
		if err != nil {
			setupLog.Error(err, "download filed failed")
		}
	}
	return nil
}

func downloadSQL(bucket *qs.Bucket, obj *qs.KeyType) error {

	getOutput, err := bucket.GetObject(
		*obj.Key,
		&service.GetObjectInput{},
	)

	if err != nil {
		setupLog.Error(err, fmt.Sprintf("get object %s failed", *obj.Key))
		return err
	} else {
		defer getOutput.Close() // 一定记得关闭GetObjectOutput, 否则容易造成链接泄漏
		f, err := os.OpenFile(QSconfig.LocalPath+"/"+*obj.Key, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			setupLog.Error(err, "Open file failed")
			return err
		}
		defer f.Close()
		// Download with 32k temporary buffer

		_, err = io.CopyBuffer(f, getOutput.Body, make([]byte, *obj.Size))
		if err != nil {
			return err
		}
	}
	return nil
}
