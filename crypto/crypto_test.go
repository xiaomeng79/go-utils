package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	cryptoValues = []struct {
		key string
		val string
	}{
		{"root%!123_md5", "8eed3d99d10f2c06fac5527aa5c448ea"},
		{"root%!123_sha1", "2715741d00e9202b8300954e729b9b351a45ee94"},
		{"root%!123_sha256", "1674b12cf25eee716a8eb278d063133a942307d9ba08079e2911a05a6b77bf07"},
		{"root%!123_sha512", "dda5a15479805ce7e9de21361f6df32c20b115752ad81ed5f8ddd82dce6786d5e4dd3a4f42cb5794fd1f013fa1c9153c7768d15f4f6d9711a17f0755fe17e9bf"},
	}
)

func TestMD5(t *testing.T) {
	for _, v := range cryptoValues {
		res := MD5(v.key)
		if v.key == "root%!123_md5" {
			assert.Equalf(t, v.val, res, "key:%s", v.key)
		} else {
			assert.NotEqualf(t, v.val, res, "key:%s", v.key)
		}
	}
}

func TestSHA1(t *testing.T) {
	for _, v := range cryptoValues {
		res := SHA1(v.key)
		if v.key == "root%!123_sha1" {
			assert.Equalf(t, v.val, res, "key:%s", v.key)
		} else {
			assert.NotEqualf(t, v.val, res, "key:%s", v.key)
		}
	}
}

func TestSHA256(t *testing.T) {
	for _, v := range cryptoValues {
		res := SHA256(v.key)
		if v.key == "root%!123_sha256" {
			assert.Equalf(t, v.val, res, "key:%s", v.key)
		} else {
			assert.NotEqualf(t, v.val, res, "key:%s", v.key)
		}
	}
}

func TestSHA512(t *testing.T) {
	for _, v := range cryptoValues {
		res := SHA512(v.key)
		if v.key == "root%!123_sha512" {
			assert.Equalf(t, v.val, res, "key:%s", v.key)
		} else {
			assert.NotEqualf(t, v.val, res, "key:%s", v.key)
		}
	}
}
