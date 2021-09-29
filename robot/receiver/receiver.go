package receiver

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/QWERKael/infoflow-robot/robot/sender"
	"github.com/QWERKael/infoflow-robot/util"
	"github.com/QWERKael/utility-go/codec"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Receiver struct {
	Engine *gin.Engine
}

func CheckSignature(rn, timestamp, signature, Token string) bool {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(rn + timestamp + Token))
	md5str := md5Ctx.Sum(nil)
	util.SugarLogger.Infof("验证字符串：%s", hex.EncodeToString(md5str))
	if signature == hex.EncodeToString(md5str) {
		return true
	}
	return false
}

func DecodeMsg(msg []byte) (*MessagePackage, error) {
	AESKey, _ := util.GetAESKey()
	b, err := util.Base64URLDecode(string(msg))
	if err != nil {
		return nil, err
	}
	body := util.AesDecrypt(b, AESKey)
	mp := &MessagePackage{}
	err = codec.DecodeJson(body, mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func ReceivePost(c *gin.Context) {
	echoStr := c.DefaultPostForm("echostr", "")
	if echoStr != "" {
		signature := c.PostForm("signature")
		timestamp := c.PostForm("timestamp")
		rn := c.PostForm("rn")
		util.SugarLogger.Infof("signature:%s\ttimestamp:%s\trn:%s\techoStr:%s", signature, timestamp, rn, echoStr)
		if CheckSignature(rn, timestamp, signature, util.Config.Token) {
			c.String(http.StatusOK, echoStr)
			util.SugarLogger.Infof("验证成功")
		} else {
			c.String(http.StatusOK, "验证失败")
			util.SugarLogger.Infof("验证失败")
		}
	} else {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			util.SugarLogger.Errorf("获取body错误:%s", err.Error())
		}
		var mp *MessagePackage
		mp, err = DecodeMsg(b)
		if err != nil {
			util.SugarLogger.Errorf("解析消息失败：%s", err.Error())
			util.SugarLogger.Errorf("无法解析的消息：%s", b)
		}
		replySender := sender.Sender{
			GroupId:  mp.Message.Header.ToId,
			RobotUrl: util.Config.RobotUrl,
		}
		replySender.SendTextMsg("收到！")
	}
}

func GetReceiver() *Receiver {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		signature := c.Query("signature")
		timestamp := c.Query("timestamp")
		rn := c.Query("rn")
		echoStr := c.Query("echostr")
		fmt.Printf("signature:%s\ttimestamp:%s\trn:%s\techoStr:%s", signature, timestamp, rn, echoStr)
		if CheckSignature(rn, timestamp, signature, "APtwGWnU") {
			c.String(http.StatusOK, echoStr)
			fmt.Printf("验证成功")
		}
		c.String(http.StatusOK, "验证失败")
		fmt.Printf("验证失败")
	})
	r.POST("/", ReceivePost)
	return &Receiver{Engine: r}
}
