package uuidutil

import (
	"crypto/md5"

	"github.com/google/uuid"
)

func GenerateOfflineUUID(username string) uuid.UUID {
	md5uuid := md5.Sum([]byte("OfflinePlayer:" + username))
	md5uuid[6] = (md5uuid[6] & 0x0f) | uint8((3&0xf)<<4)
	md5uuid[8] = (md5uuid[8] & 0x3f) | 0x80
	return md5uuid
}
