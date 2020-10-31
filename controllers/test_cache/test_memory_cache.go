package test_cache

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type TestMemoryCacheController struct {
	beego.Controller
}

func (t *TestMemoryCacheController) Get() {
	adapter, _ := cache.NewCache("memory", `{"interval":60}`)
	// 對數據進行操作
	adapter.Put("name", "梁靜茹演唱會--202020當我們談論愛情", 10)
	adapter.Put("time", "2020/12/26 & 2020/12/27", 10)
	adapter.Put("number", 22, 60)
	// 獲取1個key的值
	adapter.Get("name")

	// 獲取多個key的值
	keys := []string{"name", "time"}
	adapter.GetMulti(keys)

	// 判斷是否存在，返回bool類型
	is_exist := adapter.IsExist("name")
	fmt.Println(is_exist)

	// 刪除
	adapter.Delete("name")

	adapter.ClearAll()

	// 遞增 遞減
	adapter.Incr("age")
	adapter.Decr("age")

	t.TplName = "demo.html"
}
