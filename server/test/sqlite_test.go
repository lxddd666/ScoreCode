package test

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestSql(t *testing.T) {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Type:    "sqlite",
				Link:    fmt.Sprintf(`sqlite::@file(%s)`, "D:\\main\\grata\\server\\test\\anon.session"),
				Charset: "utf8",
			},
		},
	})
	ctx = gctx.New()
	all, err := g.DB().Ctx(ctx).GetAll(ctx, "select * from sessions")
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range all.List() {
		auth_key := item["auth_key"]
		b, ok := auth_key.([]byte)
		if ok {
			base64Data := base64.StdEncoding.EncodeToString(b)
			fmt.Println(base64Data)
		}
	}

}

type BinaryReader struct {
	stream *bytes.Reader
	_last  interface{}
}

func NewBinaryReader(data []byte) *BinaryReader {
	return &BinaryReader{
		stream: bytes.NewReader(data),
		_last:  nil,
	}
}

func (br *BinaryReader) ReadLong() (uint32, error) {
	var value uint32
	err := binary.Read(br.stream, binary.BigEndian, &value)
	return value, err
}

func (br *BinaryReader) ReadBytes(length int) ([]byte, error) {
	buffer := make([]byte, length)
	_, err := br.stream.Read(buffer)
	return buffer, err
}

func main() {
	data := sha1.Sum([]byte("_key"))
	reader := NewBinaryReader(data[:])

	auxHash, err := reader.ReadLong()
	if err != nil {
		fmt.Println("Failed to read aux_hash:", err)
		return
	}

	_, err = reader.ReadBytes(4)
	if err != nil {
		fmt.Println("Failed to read 4 bytes:", err)
		return
	}

	keyID, err := reader.ReadLong()
	if err != nil {
		fmt.Println("Failed to read key_id:", err)
		return
	}

	fmt.Println("aux_hash:", auxHash)
	fmt.Println("key_id:", keyID)
}
