package service

import (
	"cschain-bond/logger"
	"cschain-bond/oss"
	"cschain-bond/utils"
	"encoding/hex"
	"fmt"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/capsule"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/keys"
	"gitlab.bianjie.ai/csrb-bond/umbral-go/pre"
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
	// TODO remove the filename to the correct position
	filename = "../scripts/0817testUpload.text"

	plainText := utils.GetOriginTest()
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

	err = oss.UploadFile(filename, "0817testUpload.text")
	if err != nil {
		return err
	}
	return nil
}

func ReEncryptFile() error {
	if err := oss.DownloadFile("0817testUpload.text", "./0817testUpload.text"); err != nil {
		return err
	}

	bz, err := ioutil.ReadFile("./0817testUpload.text")
	if err != nil {
		return err
	}

	data := strings.Split(string(bz), "#")
	if len(data) != 2 {
		logger.Error("read the data of size must equals 2")
		return err
	}

	capsuleHex := data[1]
	capsuleData := capsule.NewCapsule()
	if err = capsuleData.FromHex(capsuleHex); err != nil {
		return err
	}

	if v, err := pre.GenProxyReEncryptionKey(delegatorPrivKey, receiverPubKey); err != nil {
		return err
	} else {
		var cfragHexes []string
		if res, err := pre.ReEncrypt(v, capsuleData); err != nil {
			return err
		} else {
			for _, v := range res {
				cfragHexes = append(cfragHexes, hex.EncodeToString(v))
			}
			// upload the reEncrypt file
			if err = ioutil.WriteFile("./0817reEncrypt.text", []byte(string(data[0])+"#"+capsuleHex+"#"+cfragHexes[0]), os.ModePerm); err != nil {
				return err
			}
			if err = oss.UploadFile("./0817reEncrypt.text", "0817reEncrypt.text"); err != nil {
				return err
			}
		}
	}
	return nil
}

func PreDecryptFile() error {
	if err := oss.DownloadFile("0817reEncrypt.text", "./0817reEncrypt.text"); err != nil {
		return err
	}

	bz, err := ioutil.ReadFile("./0817reEncrypt.text")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	data := strings.Split(string(bz), "#")
	if len(data) != 3 {
		logger.Error("read the data of size must equals 3")
		return err
	}
	capsuleHex := data[1]
	capsuleData := capsule.NewCapsule()
	if err := capsuleData.FromHex(capsuleHex); err != nil {
		return err
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
		return err
	}
	cfragBytes = append(cfragBytes, b)
	capsuleData.AttachCFragBytes(cfragBytes)

	_, err = pre.Decrypt(receiverPrivKey, []byte(data[0]), capsuleData)
	if err != nil {
		return err
	}
	return nil
}

func TempTest() {
	// 1. get ciphertext and capsule encrypted by delegator public key
	plaintext := []byte("hello world!")
	ciphertext, capsuleData, err := pre.Encrypt(delegatorPubKey, plaintext)
	if err != nil {
		panic(err)
	}

	// 2. generate rk
	if v, err := pre.GenProxyReEncryptionKey(delegatorPrivKey, receiverPubKey); err != nil {
		panic(err)
	} else {
		// 3. re-encrypt
		if cfragBytes, err := pre.ReEncrypt(v, capsuleData); err != nil {
			panic(err)
		} else {
			if len(cfragBytes) > 0 {
				capsuleData.AttachCFragBytes(cfragBytes)
			}
			// 4. decrypt re-encryption ciphertext
			if v, err := pre.Decrypt(receiverPrivKey, ciphertext, capsuleData); err != nil {
				panic(err)
			} else {
				fmt.Println(v)
			}
		}
	}
}
