package whats_util

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gaes"
	whatsin "hotgo/internal/model/input/whats"
	"strings"
)

func AccountDetailToByte(detail *whatsin.WhatsAccountUploadInp, key, vi []byte) ([]byte, error) {
	s := fmt.Sprintf("%s~%s~%s~%s~%s~%s", detail.Account,
		detail.PublicKey,
		detail.PrivateKey,
		detail.PublicMsgKey,
		detail.PrivateMsgKey,
		detail.Identify,
	)
	return gaes.Encrypt([]byte(s), key, vi)
}

func ByteToAccountDetail(content, key, vi []byte) (*whatsin.WhatsAccountUploadInp, error) {
	decrypt, err := gaes.Decrypt(content, key, vi)
	if err != nil {
		return nil, err
	}
	detailSp := strings.Split(string(decrypt), "~")
	if len(detailSp) != 6 {
		return nil, errors.New("the account details are incorrect")
	}
	detail := &whatsin.WhatsAccountUploadInp{
		Account:       detailSp[0],
		PublicKey:     detailSp[1],
		PrivateKey:    detailSp[2],
		PublicMsgKey:  detailSp[3],
		PrivateMsgKey: detailSp[4],
		Identify:      detailSp[5],
	}
	return detail, nil
}
