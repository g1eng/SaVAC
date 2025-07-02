package sacloud

import (
	"crypto/md5" //nolint
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateOnetimePath(secret string, path string, expiredAt time.Time) string {
	hasher := md5.New() //nolint

	hexExpirationTime := fmt.Sprintf("%x", expiredAt.Unix())
	seed := "/" + path + "/" + secret + "/" + hexExpirationTime + "/"
	fmt.Println(seed)
	hasher.Write([]byte(seed))

	hash := hex.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("%s?webaccel_secure_hash=%s&webaccel_secure_time=%s", path, hash, hexExpirationTime)
}
