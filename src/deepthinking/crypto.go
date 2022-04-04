package deepthinking

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

func Encrypt(text *string) {
	md5 := md5.New()
	bi := big.NewInt(0)
	md5.Write([]byte(*text))
	hexstr := hex.EncodeToString(md5.Sum(nil))
	bi.SetString(hexstr, 16)
	*text = bi.String()
}

func Decrypt(text *string) {

}
