package sacloud

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/cloud/model/object_storage"
)

func fillFieldsForBucketControl(controls []object_storage.BucketControlsInner) (bucketControls []object_storage.BucketControlsInner) {
	for _, control := range controls {
		f := false
		newControl := object_storage.BucketControlsInner{
			BucketName: control.BucketName,
			CanRead:    &f,
			CanWrite:   &f,
		}
		if control.CanRead != nil {
			newControl.CanRead = control.CanRead
		}
		if control.CanWrite != nil {
			newControl.CanWrite = control.CanWrite
		}
		bucketControls = append(bucketControls, newControl)
	}
	return bucketControls
}

func (c *CloudAPIClient) CreateObjectStorageSiteAccount() (*object_storage.Account, error) {
	req := c.ObjectStorageAPI.DefaultApi.CreateAccount(context.Background())
	res, rawResp, err := req.Execute()
	if err != nil {
		_ = helper.PrintJson(res)
		return nil, err
	}
	if rawResp.StatusCode-200 > 100 {
		_ = helper.PrintJson(res)
		return nil, fmt.Errorf("%s", rawResp.Status)
	}
	return res, nil
}

func (c *CloudAPIClient) DeleteObjectStorageSiteAccount() error {
	rawResp, err := c.ObjectStorageAPI.DefaultApi.DeleteAccount(context.Background()).Execute()
	if err != nil {
		return err
	}
	if rawResp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", rawResp.Status)
	}
	return nil
}

func (c *CloudAPIClient) ListSites() ([]object_storage.ModelCluster, error) {
	res, rawResp, err := c.ObjectStorageAPI.DefaultApi.GetClusters(context.Background()).Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode == 200 || rawResp.StatusCode == 400 {
		return res.Data, nil
	}
	_ = helper.PrintJson(res)
	return nil, fmt.Errorf("%s", rawResp.Status)
}

func (c *CloudAPIClient) ListAccountKeys() ([]object_storage.AccountKeysDataInner, error) {
	res, rawResp, err := c.ObjectStorageAPI.DefaultApi.GetAccountKeys(context.Background()).Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode == 200 || rawResp.StatusCode == 400 {
		return res.Data, nil
	}
	_ = helper.PrintJson(res)
	return nil, fmt.Errorf("%s", rawResp.Status)
}

func (c *CloudAPIClient) CreateAccountKey() (*object_storage.AccountKey, error) {
	req := c.ObjectStorageAPI.DefaultApi.CreateAccountKey(context.Background())
	res, rawResp, err := req.Execute()
	if err != nil {
		_ = helper.PrintJson(res)
		return nil, err
	}
	if rawResp.StatusCode-200 > 100 {
		_ = helper.PrintJson(res)
		return nil, fmt.Errorf("%s", rawResp.Status)
	}
	return res, nil
}

func (c *CloudAPIClient) DeleteAccountKey(id string) error {
	req := c.ObjectStorageAPI.DefaultApi.DeleteAccountKey(context.Background(), id)
	res, err := req.Execute()
	if err != nil {
		_ = helper.PrintJson(res)
		return err
	}
	return nil
}

func (c *CloudAPIClient) CreateBucket(name string) (*object_storage.ModelBucket, error) {
	req := c.ObjectStorageAPI.DefaultApi.CreateBucket(context.Background(), name)
	req = req.HandlerPutBucketReqBody(object_storage.HandlerPutBucketReqBody{ClusterId: "isk01"})
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode-200 > 100 {
		return nil, fmt.Errorf("%s", rawResp.Status)
	}
	return res.Data, nil
}

func (c *CloudAPIClient) DeleteBucket(name string) error {
	req := c.ObjectStorageAPI.DefaultApi.DeleteBucket(context.Background(), name)
	res, err := req.Execute()
	if err != nil {
		_ = helper.PrintJson(res)
		return err
	}
	return nil
}

func (c *CloudAPIClient) CreatePermission(name string, controls []object_storage.BucketControlsInner) (*object_storage.Permission, error) {
	req := c.ObjectStorageAPI.DefaultApi.CreatePermission(context.Background())
	permCfg := object_storage.PermissionBucketControlsBody{
		DisplayName:    &name,
		BucketControls: fillFieldsForBucketControl(controls),
	}
	req = req.PermissionBucketControlsBody(permCfg)
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode-200 > 100 {
		_ = helper.PrintJson(res)
		return nil, fmt.Errorf("%s", rawResp.Status)
	}
	return res, nil
}

func (c *CloudAPIClient) GetPermissions(id string) (*object_storage.Permission, error) {
	req := c.ObjectStorageAPI.DefaultApi.GetPermission(context.Background(), id)
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode == 200 || rawResp.StatusCode == 400 {
		return res, nil
	}
	_ = helper.PrintJson(res)
	return nil, fmt.Errorf("%s", rawResp.Status)
}

func (c *CloudAPIClient) ListPermissions() ([]object_storage.PermissionsDataInner, error) {
	req := c.ObjectStorageAPI.DefaultApi.GetPermissions(context.Background())
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode == 200 || rawResp.StatusCode == 400 {
		return res.Data, nil
	}
	_ = helper.PrintJson(res)
	return nil, fmt.Errorf("%s", rawResp.Status)
}

func (c *CloudAPIClient) UpdatePermission(id string, name string, controls []object_storage.BucketControlsInner) (*object_storage.Permission, error) {
	req := c.ObjectStorageAPI.DefaultApi.UpdatePermission(context.Background(), id)
	req = req.PermissionBucketControlsBody(object_storage.PermissionBucketControlsBody{
		DisplayName:    &name,
		BucketControls: controls,
	})
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode-200 > 100 {
		_ = helper.PrintJson(res)
		return nil, fmt.Errorf("%s", rawResp.Status)
	}
	return res, nil
}

func (c *CloudAPIClient) DeletePermission(id string) error {
	req := c.ObjectStorageAPI.DefaultApi.DeletePermission(context.Background(), id)
	_, err := req.Execute()
	if err != nil {
		return err
	}
	return nil
}

func (c *CloudAPIClient) CreatePermissionKey(permissionId string) (*object_storage.PermissionKey, error) {
	req := c.ObjectStorageAPI.DefaultApi.CreatePermissionKey(context.Background(), permissionId)
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode != http.StatusCreated {
		_ = helper.PrintJson(res)
		return nil, fmt.Errorf("%s", rawResp.Status)
	}
	return res, nil
}

func (c *CloudAPIClient) ListBuckets() (buckets []types.Bucket, err error) {
	listBucketsOutput, err := c.S3Client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list buckets, %v", err)
	}
	return listBucketsOutput.Buckets, nil
}

func (c *CloudAPIClient) ListPermissionKeys(permissionId string) ([]object_storage.PermissionKeysDataInner, error) {
	req := c.ObjectStorageAPI.DefaultApi.GetPermissionKeys(context.Background(), permissionId)
	res, rawResp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if rawResp.StatusCode == 200 || rawResp.StatusCode == 400 {
		return res.Data, nil
	}
	_ = helper.PrintJson(res)
	return nil, fmt.Errorf("%s", rawResp.Status)
}

func (c *CloudAPIClient) DeletePermissionKey(permissionId string, kerId string) error {
	req := c.ObjectStorageAPI.DefaultApi.DeletePermissionKey(context.Background(), permissionId, kerId)
	res, err := req.Execute()
	if err != nil {
		return err
	}
	if res.StatusCode-200 > 100 {
		_ = helper.PrintJson(res)
		return fmt.Errorf("%s", res.Status)
	}
	return nil
}

// DecomposeS3Uri URIを元にバケット名およびキーの基点パスを返却します。キー基点パスとは
// (1) `/`以外で始まる2文字以上のパス
// (2) `/`のみ
// のいずれかの文字列を指し、2つ以上の`/`文字が連続しないことが保証されます。
func DecomposeS3Uri(s3ObjectUri string) (bucket string, keyBase string, err error) {
	if strings.Index(s3ObjectUri, "s3://") != 0 || len(s3ObjectUri) < 5 {
		return "", "", fmt.Errorf("invalid s3 uri prefix")
	}
	s3ObjectUri = s3ObjectUri[5:]
	s3ObjectUri = strings.TrimLeft(s3ObjectUri, "/")
	s3ObjectUri = strings.ReplaceAll(s3ObjectUri, "//", "/")
	s3ObjectUri = strings.TrimRight(s3ObjectUri, "/")
	s := strings.Split(s3ObjectUri, "/")
	if len(s) < 1 {
		return "", "", fmt.Errorf("invalid resource description")
	} else if len(s) > 1 {
		var keySwap []string
		for _, p := range s[1:] {
			if p != "" {
				keySwap = append(keySwap, p)
			}
		}
		keyBase = strings.Join(keySwap, "/")
	} else {
		keyBase = ""
	}
	return s[0], keyBase, nil
}

func (c *CloudAPIClient) PutObject(source string, dest string) error {
	var (
		f                 io.Reader
		err               error
		isSmallerThan5GiB = true
	)
	if source == "-" {
		f = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(source)
		if err != nil {
			return err
		}
		fInfo, err := file.Stat()
		if err != nil {
			return err
		}
		if fInfo.Size() >= 5*1000^3 {
			isSmallerThan5GiB = false
		}
		f = file
	}
	bucket, key, err := DecomposeS3Uri(dest)
	if err != nil {
		return err
	}
	pathParts := strings.Split(key, "/")
	var destFileName string
	if key == "" || dest != strings.TrimRight(dest, "/") {
		//バケットルートもしくはディレクトリの場合、ソースパスのファイル名をキー末尾に付加
		if source == "-" {
			return fmt.Errorf("invalid stdin redirect: target file name should be specified")
		}
		s := strings.Split(source, "/")
		if len(s) == 0 {
			return fmt.Errorf("invalid source file name: %s", source)
		}
		destFileName = s[len(s)-1]
		if key == "" {
			key = "/"
		}
		key = fmt.Sprintf("%s%s", key, destFileName)
	} else {
		destFileName = pathParts[len(pathParts)-1]
	}
	if c.Debug {
		log.Println("desf: ", destFileName)
	}
	if c.Debug {
		log.Println("key: ", key)
		log.Println("bucket: ", bucket)
	}
	if isSmallerThan5GiB {
		return c.putObjectInternal(f, bucket, key)
	} else {
		return c.uploadObjectInternal(f, bucket, key)
	}
}

func (c *CloudAPIClient) putObjectInternal(f io.Reader, bucket string, key string) error {
	if c.Debug {
		log.Println("bucket", bucket)
		log.Println("key", key)
	}
	_, err := c.S3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   f,
	})
	return err
}

func (c *CloudAPIClient) uploadObjectInternal(f io.Reader, bucket string, key string) error {
	uploader := manager.NewUploader(c.S3Client)
	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   f,
	})
	return err
}

func (c *CloudAPIClient) GetObject(source string) (io.Reader, error) {
	bucket, key, err := DecomposeS3Uri(source)
	if err != nil {
		return nil, err
	}
	res, err := c.S3Client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

func (c *CloudAPIClient) RemoveObject(uri string) error {
	bucket, key, err := DecomposeS3Uri(uri)
	if err != nil {
		return err
	}
	_, err = c.S3Client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	return err
}

func (c *CloudAPIClient) RemoveObjectRecursively(directryUri string) error {
	bucket, _, err := DecomposeS3Uri(directryUri)
	if err != nil {
		return err
	}
	list, err := c.ListObjects(directryUri)
	if err != nil {
		return err
	}
	var objects []types.ObjectIdentifier
	for _, objKey := range list {
		objects = append(objects, types.ObjectIdentifier{Key: &objKey})
		log.Println("obj to be deleted:", objKey)
	}
	_, err = c.S3Client.DeleteObjects(context.Background(), &s3.DeleteObjectsInput{
		Bucket: &bucket,
		Delete: &types.Delete{
			Objects: objects,
		},
	})
	return err
}

func (c *CloudAPIClient) ListObjects(key string) ([]string, error) {
	bucket, key, err := DecomposeS3Uri(key)
	if err != nil {
		return nil, err
	}
	if len(key) != 0 && key[len(key)-1] == '/' {
		key += "/"
	}
	rawResp, err := c.S3Client.ListObjects(context.Background(), &s3.ListObjectsInput{
		Bucket: &bucket,
		Prefix: &key,
	})
	if c.Debug {
		helper.PrintJson(rawResp) // nolint
	}
	if err != nil {
		return nil, err
	}
	var res []string
	for _, r := range rawResp.Contents {
		res = append(res, *r.Key)
	}
	return res, nil
}

func (c *CloudAPIClient) CheckObjectExistence(key string) error {
	bucket, key, err := DecomposeS3Uri(key)
	if err != nil {
		return err
	}
	res, err := c.S3Client.HeadObject(context.Background(), &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return err
	}
	if c.Debug {
		helper.PrintJson(res) // nolint
	}
	return nil
}
