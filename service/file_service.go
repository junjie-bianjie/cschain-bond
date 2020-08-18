package service

import (
	"context"
	"cschain-bond/logger"
	"cschain-bond/oss"
	"cschain-bond/utils"
	"encoding/hex"
	"fmt"
	preserver "gitlab.bianjie.ai/csrb-bond/pre-server/protos"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/capsule"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/keys"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/kfrag"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/pre"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"strings"
)

const (
	delegatorPrivKeyHex = "3a86bcd5e5b43f112614c3dbe48e56a3a85b39a39ff879cfa3e968bc3c5047d0"
	receiverPrivKeyHex  = "bf749f8b336f697475a86550c36ed1ad8262dd21fcfa7d6838e20b75b6e2611e"
)

var (
	delegatorPrivKey *keys.PrivateKey
	delegatorPubKey  *keys.PublicKey
	receiverPrivKey  *keys.PrivateKey
	receiverPubKey   *keys.PublicKey
)

func init() {
	var err error
	delegatorPrivKey, err = keys.NewPrivateKeyFromHex(delegatorPrivKeyHex)
	if err != nil {
		panic(err)
	}
	receiverPrivKey, err = keys.NewPrivateKeyFromHex(receiverPrivKeyHex)
	if err != nil {
		panic(err)
	}

	delegatorPubKey = delegatorPrivKey.PubKey
	receiverPubKey = receiverPrivKey.PubKey
}

func UploadEncryptFile(filename string) error {
	plainText := utils.GetFile(filename)
	ciphertext, capsuleData, err := pre.Encrypt(delegatorPubKey, plainText)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	data := hex.EncodeToString(ciphertext) + "#" + capsuleData.Hex()
	err = ioutil.WriteFile(filename, []byte(data), os.ModePerm)
	if err != nil {
		logger.Error(fmt.Sprintf("wireFile %s faild: %v", filename, err))
		return err
	}

	err = oss.UploadFile(filename)
	if err != nil {
		return err
	}
	_ = os.Remove(filename)
	return nil
}

func ReEncryptFile(filename, encryptFilename string) error {
	tempFile := utils.RandStringOfLength(8)
	if err := oss.DownloadFile(filename, tempFile); err != nil {
		return err
	}

	bz, err := ioutil.ReadFile(tempFile)
	if err != nil {
		return err
	}
	_ = os.Remove(tempFile)

	data := strings.Split(string(bz), "#")
	if len(data) != 2 {
		logger.Error("read the data of size must equals 2")
		return err
	}

	capsuleHex := data[1]

	capsuleData := capsule.NewCapsule()
	var cfragHexes []string
	if err = capsuleData.FromHex(capsuleHex); err != nil {
		return err
	}

	if kFrags, err := pre.GenProxyReEncryptionKey(delegatorPrivKey, receiverPubKey); err != nil {
		return err
	} else {
		if len(kFrags) == 0 {
			// todo error
			panic(err)
		}
		var rk []string

		for _, v := range kFrags {
			kFragHex := v.Hex()
			newKFrag := kfrag.NewKFrag()
			if err := newKFrag.FromHex(kFragHex); err != nil {
				// todo error
				panic(err)
			} else {
				// todo coreect
			}

			rk = append(rk, kFragHex)
		}

		reqBody := &preserver.ReEncryptReq{
			ReEncryptionKey: [][]byte{[]byte(rk[0])},
			Capsule:         []byte(capsuleHex),
		}

		// 连接grpc服务器
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		// 初始化客户端
		c := preserver.NewPreServerClient(conn)
		if res, err := c.ReEncrypt(context.Background(), reqBody); err != nil {
			return err
		} else {
			for _, v := range res.CFrags {
				cfragHexes = append(cfragHexes, hex.EncodeToString(v))
			}
			// upload the reEncrypt file
			if err = ioutil.WriteFile(encryptFilename, []byte(string(data[0])+"#"+capsuleHex+"#"+cfragHexes[0]), os.ModePerm); err != nil {
				return err
			}
			if err = oss.UploadFile(encryptFilename); err != nil {
				return err
			}
			_ = os.Remove(encryptFilename)
		}
	}
	return nil
}

func PreDecryptFile(filename string) (string, error) {
	tempFile := utils.RandStringOfLength(8)
	if err := oss.DownloadFile(filename, tempFile); err != nil {
		return "", err
	}

	bz, err := ioutil.ReadFile(tempFile)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	_ = os.Remove(tempFile)

	data := strings.Split(string(bz), "#")
	if len(data) != 3 {
		logger.Error("read the data of size must equals 3")
		return "", err
	}
	capsuleHex := data[1]
	capsuleData := capsule.NewCapsule()
	if err := capsuleData.FromHex(capsuleHex); err != nil {
		return "", err
	}
	//var cfragBytes [][]byte
	//for _, v := range cfragHexes {
	//	if b, err := hex.DecodeString(v); err != nil {
	//		continue
	//	} else {
	//		cfragBytes = append(cfragBytes, b)
	//	}
	//}
	var cfragBytes [][]byte
	b, err := hex.DecodeString(data[2])
	if err != nil {
		return "", err
	}

	cfragBytes = append(cfragBytes, b)
	capsuleData.AttachCFragBytes(cfragBytes)
	plainText, err := hex.DecodeString(data[0])
	if err != nil {
		return "", err
	}

	res, err := pre.Decrypt(receiverPrivKey, plainText, capsuleData)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// todo encapsulation grpc client
func GrpClient() {

}
