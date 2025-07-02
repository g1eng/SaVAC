package cloud_actions

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/cloud/model/object_storage"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/urfave/cli/v3"
)

func (g *CloudActionGenerator) GenerateObjectStorageCreateSiteAction(_ context.Context, _ *cli.Command) error {
	sites, err := g.ApiClient.CreateObjectStorageSiteAccount()
	if err != nil {
		return err
	}
	return helper.PrintJson(sites)
}

func (g *CloudActionGenerator) GenerateObjectStorageListSiteAction(_ context.Context, _ *cli.Command) error {
	sites, err := g.ApiClient.ListSites()
	if err != nil {
		return err
	}
	return helper.PrintJson(sites)
}

func (g *CloudActionGenerator) GenerateObjectStorageDeleteSiteAction(_ context.Context, _ *cli.Command) error {
	return g.ApiClient.DeleteObjectStorageSiteAccount()
}

func (g *CloudActionGenerator) GenerateObjectStorageListActionMeta(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() == 0 {
		return g.GenerateObjectStorageListBucketAction(ctx, cmd)
	} else if cmd.NArg() == 1 {
		return g.GenerateObjectStorageListObjectAction(ctx, cmd)
	} else {
		return fmt.Errorf("too many arguments")
	}
}
func (g *CloudActionGenerator) GenerateObjectStorageListObjectAction(_ context.Context, cmd *cli.Command) error {
	buckets, err := g.ApiClient.ListObjects(cmd.Args().First())
	if err != nil {
		return err
	}
	for _, b := range buckets {
		fmt.Println(b)
	}
	return nil
}

func (g *CloudActionGenerator) GenerateObjectStorageListBucketAction(_ context.Context, _ *cli.Command) error {
	buckets, err := g.ApiClient.ListBuckets()
	if err != nil {
		return err
	}
	return helper.PrintJson(buckets)
}

func (g *CloudActionGenerator) GenerateCreateAccountKeyAction(ctx context.Context, cmd *cli.Command) error {
	key, err := g.ApiClient.CreateAccountKey()
	if err != nil {
		return err
	}
	return helper.PrintJson(key)
}

func (g *CloudActionGenerator) GenerateDeleteAccountKeyAction(ctx context.Context, cmd *cli.Command) error {
	keys, err := g.ApiClient.ListAccountKeys()
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return fmt.Errorf("no account key")
	}
	return g.ApiClient.DeleteAccountKey(*keys[0].Id)
}

func (g *CloudActionGenerator) GenerateListAccountKeyAction(_ context.Context, cmd *cli.Command) error {
	keys, err := g.ApiClient.ListAccountKeys()
	if err != nil {
		return err
	}
	return helper.PrintJson(keys)
}

func (g *CloudActionGenerator) GenerateCreateBucketAction(_ context.Context, cmd *cli.Command) error {
	bucket, err := g.ApiClient.CreateBucket(cmd.Args().First())
	if err != nil {
		return err
	}
	return helper.PrintJson(bucket)
}

func (g *CloudActionGenerator) GenerateDeleteBucketAction(_ context.Context, cmd *cli.Command) error {
	return g.ApiClient.DeleteBucket(cmd.Args().First())
}

func (g *CloudActionGenerator) GenerateCreatePermissionAction(ctx context.Context, cmd *cli.Command) error {
	rwBuckets := cmd.StringSlice("rw")
	roBuckets := cmd.StringSlice("ro")
	woBuckets := cmd.StringSlice("wo")
	var bucketNames []string
	if cmd.Bool("force") {
		bucketNames = append(rwBuckets, append(roBuckets, woBuckets...)...)
	} else {
		buckets, err := g.ApiClient.ListBuckets()
		if err != nil {
			return err
		}
		bucketNames = func([]types.Bucket) []string {
			var res []string
			for _, b := range buckets {
				res = append(res, *b.Name)
			}
			return res
		}(buckets)
		time.Sleep(time.Millisecond * 300)
		for _, b := range append(rwBuckets, append(roBuckets, woBuckets...)...) {
			isExist := false
			for _, b2 := range bucketNames {
				if b2 == b {
					isExist = true
					break
				}
			}
			if !isExist {
				return fmt.Errorf("no such bucket: %s", b)
			}
		}
	}

	var controls []object_storage.BucketControlsInner
	t, f := true, false
	for _, b := range bucketNames {
		c := object_storage.BucketControlsInner{
			BucketName: &b,
			CanRead:    &f,
			CanWrite:   &f,
		}
		for _, cmpBucket := range rwBuckets {
			//fmt.Println("cmpBucket:" + cmpBucket)
			if *c.BucketName == cmpBucket {
				c.CanRead = &t
				c.CanWrite = &t
			}
		}
		for _, cmpBucket := range woBuckets {
			if *c.BucketName == cmpBucket {
				c.CanWrite = &t
			}
		}
		for _, cmpBucket := range roBuckets {
			if *c.BucketName == cmpBucket {
				c.CanRead = &t
			}
		}
		controls = append(controls, c)
	}
	p, err := g.ApiClient.CreatePermission(cmd.Args().First(), controls)
	if err != nil {
		return err
	}
	return helper.PrintJson(p)
}

func (g *CloudActionGenerator) normalizePermissionValue(s string) (string, error) {
	_, err := strconv.Atoi(s)
	if err != nil {
		perms, err := g.ApiClient.ListPermissions()
		if err != nil {
			return "", err
		}
		time.Sleep(time.Millisecond * 300)
		isMatched := false
		for _, p := range perms {
			if s == *p.DisplayName {
				s = fmt.Sprintf("%d", *p.Id)
				isMatched = true
				break
			}
		}
		if !isMatched {
			return "", fmt.Errorf("no permissions found: %s", s)
		}
	}
	return s, nil
}

func (g *CloudActionGenerator) GenerateDeletePermissionAction(ctx context.Context, cmd *cli.Command) error {
	a, err := g.normalizePermissionValue(cmd.Args().First())
	if err != nil {
		return err
	}
	return g.ApiClient.DeletePermission(a)
}

func (g *CloudActionGenerator) GenerateListPermissionAction(ctx context.Context, cmd *cli.Command) error {
	perms, err := g.ApiClient.ListPermissions()
	if err != nil {
		return err
	}
	return helper.PrintJson(perms)
}

func (g *CloudActionGenerator) GenerateUpdatePermissionAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.StringSlice("rw") == nil && cmd.StringSlice("ro") == nil && cmd.StringSlice("wo") == nil {
		return fmt.Errorf("one or more target buckets must be specified")
	}
	pid, err := g.normalizePermissionValue(cmd.Args().First())
	if err != nil {
		return err
	}
	perm, err := g.ApiClient.GetPermissions(pid)
	if err != nil {
		return err
	}
	rwBuckets := cmd.StringSlice("rw")
	roBuckets := cmd.StringSlice("ro")
	woBuckets := cmd.StringSlice("wo")
	for _, b := range append(rwBuckets, append(roBuckets, woBuckets...)...) {
		isExist := false
		for _, c := range perm.Data.BucketControls {
			if *c.BucketName == b {
				isExist = true
				break
			}
		}
		if !isExist {
			return fmt.Errorf("no such bucket: %s", b)
		}
	}
	t, f := true, false
	for i, c := range perm.Data.BucketControls {
		for _, cmpBucket := range rwBuckets {
			if *c.BucketName == cmpBucket {
				perm.Data.BucketControls[i].CanRead = &t
				perm.Data.BucketControls[i].CanWrite = &t
			}
		}
		for _, cmpBucket := range woBuckets {
			if *c.BucketName == cmpBucket {
				perm.Data.BucketControls[i].CanWrite = &t
				perm.Data.BucketControls[i].CanRead = &f
			}
		}
		for _, cmpBucket := range roBuckets {
			if *c.BucketName == cmpBucket {
				perm.Data.BucketControls[i].CanWrite = &f
				perm.Data.BucketControls[i].CanRead = &t
			}
		}
	}
	p, err := g.ApiClient.UpdatePermission(pid, *perm.Data.DisplayName, perm.Data.BucketControls)
	if err != nil {
		return err
	}
	return helper.PrintJson(p)
}

func (g *CloudActionGenerator) GenerateCreatePermissionKeyAction(_ context.Context, cmd *cli.Command) error {
	p, err := g.normalizePermissionValue(cmd.Args().First())
	if err != nil {
		return err
	}
	k, err := g.ApiClient.CreatePermissionKey(p)
	if err != nil {
		return err
	}
	return helper.PrintJson(k)
}

func (g *CloudActionGenerator) GenerateDeletePermissionKeyAction(ctx context.Context, cmd *cli.Command) error {
	p, err := g.normalizePermissionValue(cmd.Args().First())
	if err != nil {
		return err
	}
	keys, err := g.ApiClient.ListPermissionKeys(p)
	if err != nil {
		return err
	} else if len(keys) == 0 {
		return fmt.Errorf("no permission key")
	}
	return g.ApiClient.DeletePermissionKey(p, *keys[0].Id)
}

func (g *CloudActionGenerator) GenerateListPermissionKeyAction(_ context.Context, cmd *cli.Command) error {
	a, err := g.normalizePermissionValue(cmd.Args().First())
	if err != nil {
		return err
	}
	keys, err := g.ApiClient.ListPermissionKeys(a)
	if err != nil {
		return err
	}
	return helper.PrintJson(keys)
}

func (g *CloudActionGenerator) GeneratePutAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 2 {
		return fmt.Errorf("just two arguments should be specified")
	}
	a := cmd.Args().Slice()
	if !cmd.Bool("recursive") {
		return g.ApiClient.PutObject(a[0], a[1])
	} else {
		src := a[0]
		bucket, key, err := sacloud.DecomposeS3Uri(a[1])
		destDir := fmt.Sprintf("s3://%s/%s", bucket, key)
		if err != nil {
			return fmt.Errorf("invlaid S3 URI: %w", err)
		}
		var filePaths []string

		err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			// Handle any walking errors
			if err != nil {
				return err
			}
			if !info.IsDir() {
				filePaths = append(filePaths, path)
			}
			return nil
		})
		if err != nil {
			return err
		}

		if len(filePaths) == 0 {
			return fmt.Errorf("no objects")
		}
		if len(filePaths) == 1 && filePaths[0] == src {
			return g.ApiClient.PutObject(a[0], a[1])
		}

		for _, obj := range filePaths {
			if len(obj) < len(src) {
				return fmt.Errorf("invalid object name")
			}
			println("dest", fmt.Sprintf("%s%s", destDir, obj[len(src):]))
			err := g.ApiClient.PutObject(obj, fmt.Sprintf("%s%s", destDir, obj[len(src):]))
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (g *CloudActionGenerator) GenerateGetAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 2 {
		return fmt.Errorf("just two arguments should be specified")
	}
	writeObject := func(src string, dest string) error {
		r, err := g.ApiClient.GetObject(src)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		b, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		f, err := os.Create(dest)
		if err != nil {
			return err
		}
		_, err = f.Write(b)
		return err
	}
	if !cmd.Bool("recursive") {
		src := cmd.Args().First()
		dest := cmd.Args().Slice()[1]
		return writeObject(src, dest)
	} else {
		src := cmd.Args().First()
		bucket, _, err := sacloud.DecomposeS3Uri(src)
		srcDir, _ := strings.CutPrefix(src, "s3://"+bucket+"/")
		srcDir = strings.TrimRight(srcDir, "/") + "/"
		//log.Println("srcDir", srcDir)
		if err != nil {
			return fmt.Errorf("invlaid S3 URI: %w", err)
		}
		list, err := g.ApiClient.ListObjects(fmt.Sprintf("s3://%s/%s", bucket, srcDir))
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return fmt.Errorf("no objects")
		}
		//--recursive but for single file
		if len(list) == 1 && list[0] == src {
			dest := cmd.Args().Slice()[1]
			return writeObject(src, dest)
		}

		//recursively get objects
		destDir := cmd.Args().Slice()[1]
		for _, obj := range list {
			r, err := g.ApiClient.GetObject(fmt.Sprintf("s3://%s/%s", bucket, obj))
			if err != nil {
				return err
			}
			if _, err := os.ReadDir(destDir); err != nil {
				err = os.MkdirAll(destDir, 0o0700)
				if err != nil {
					return err
				}
			}
			if len(obj) <= len(srcDir) {
				return fmt.Errorf("unknown condition")
			}
			b, err := io.ReadAll(r)
			if err != nil {
				return err
			}
			//log.Println("obj", obj)
			//log.Println("obj[len(srcDir):]", obj[len(srcDir):])
			objLocalPath := fmt.Sprintf("%s/%s", destDir, obj[len(srcDir):])
			objParentDir, _ := path.Split(objLocalPath)
			if _, err := os.ReadDir(objParentDir); err != nil {
				err = os.MkdirAll(objParentDir, 0o0700)
				if err != nil {
					return err
				}
			}
			err = os.WriteFile(objLocalPath, b, 0o0600)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (g *CloudActionGenerator) GenerateListAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return fmt.Errorf("just an argument should be specified")
	}
	r, err := g.ApiClient.ListObjects(cmd.Args().First())
	if err != nil {
		return err
	}
	return helper.PrintJson(r)
}

func (g *CloudActionGenerator) GenerateRmAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return fmt.Errorf("just an argument should be specified")
	}
	if cmd.Bool("recursive") {
		return g.ApiClient.RemoveObjectRecursively(cmd.Args().First())
	}
	return g.ApiClient.RemoveObject(cmd.Args().First())
}

func (g *CloudActionGenerator) GenerateCheckAction(_ context.Context, cmd *cli.Command) error {
	return g.ApiClient.CheckObjectExistence(cmd.Args().First())
}
