package hashutils

import (
	"crypto/md5"
	"encoding/hex"
)

func IsTargetHash(sourceStr, targetHash string) bool {
	hash := md5.Sum([]byte(sourceStr))
	hexEnc := hex.EncodeToString(hash[:])
	return hexEnc == targetHash
}
