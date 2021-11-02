package routers

import (
	"blockchain/apiControl"

	"github.com/astaxie/beego"
)

func init() {
	blockController := &apiControl.Block_control{}
	hashController := &apiControl.Hash_control{}
	chainController := &apiControl.Chain_control{}

	beego.Router("/blockinfo", blockController, "get:Blockget;post:Blockpost")
	beego.Router("/hash", hashController, "get:Hashget")
	beego.Router("/chain", chainController, "get:Blockchaiget;post:Blockchainpost")
	beego.Router("/", hashController, "get:Hashget")
}
