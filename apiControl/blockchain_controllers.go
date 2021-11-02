package apiControl

import (
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

const N = 4

type Chain_control struct {
	beego.Controller
}

func (this *Chain_control) Blockchaiget() {
	this.TplName = "blockchain.html"
}

func (c *Chain_control) Blockchainpost() {
	buf := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &buf)
	if err != nil {
		return
	}
	blockhigh1 := buf["Bl"].(string)
	data := buf["Da"].(string)
	blockhigh, _ := strconv.ParseInt(blockhigh1, 10, 64)
	//chainnum, _ := strconv.ParseInt(chainnum1, 10, 64)

	block := NewBlock_simple(data, blockhigh)
	//hash_data := block.Hash //测试
	pw := NewProofofWork(block)
	nonce, hashdata1 := pw.run()
	hashdata := hex.EncodeToString(hashdata1)
	c.Data["json"] = map[string]interface{}{"hashdata": hashdata, "nonce": nonce}

	c.ServeJSON()
	return
}
