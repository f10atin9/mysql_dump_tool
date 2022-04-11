package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/yunify/qingstor-sdk-go/config"
	qs "github.com/yunify/qingstor-sdk-go/service"
	"io"
	"io/ioutil"
	"os"
	"time"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var setupLog = ctrl.Log.WithName("setup")

func subMain() error {
	ctrl.SetLogger(zap.New(zap.UseDevMode(QSconfig.development)))

	config, err := config.New(QSconfig.AccessKey, QSconfig.SecretKey)
	if err != nil {
		setupLog.Error(err, "create a config failed")
		return err
	}
	qsService, _ := qs.Init(config)
	bucketService, _ := qsService.Bucket(QSconfig.BucketName, QSconfig.Zone)

	if err != nil {
		setupLog.Error(err, "Create config failed")
		return err
	}

	dateStr := time.Now().Format("2006010215")
	dirPath := QSconfig.LocalPath + "/" + dateStr
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		setupLog.Error(err, "read sql dir failed")
		return err
	}
	for _, fi := range dir {
		filePath := dirPath + "/" + fi.Name()
		err := uploadSQL(filePath, dateStr, bucketService)
		if err != nil {
			setupLog.Error(err, "upload sql failed")
			return err
		}
	}
	return nil
}

func uploadSQL(filePath, dateStr string, bucketService *qs.Bucket) error {
	file, err := os.Open(filePath)
	if err != nil {
		setupLog.Error(err, "os open file failed")
	}
	defer func() {
		_ = file.Close()
	}()

	hash := md5.New()
	io.Copy(hash, file)
	hashInBytes := hash.Sum(nil)[:16]
	md5String := hex.EncodeToString(hashInBytes)
	toPtr := func(s string) *string { return &s }
	input := &qs.PutObjectInput{
		ContentMD5:      toPtr(md5String),
		Body:            file,
		XQSStorageClass: toPtr("STANDARD"),
	}
	objectKey := QSconfig.UploadPath + "/" + dateStr + "/" + file.Name()
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Put object to bucket(name: %s) failed with given error: %s\n", QSconfig.BucketName, err)
	} else {
		fmt.Printf("%s has been uploaded to bucket. Status code: %d \n", file.Name(), *output.StatusCode)
		setupLog.Info("%s has been uploaded to bucket. Status code: %d \n", file.Name(), *output.StatusCode)
	}
	return err
}
