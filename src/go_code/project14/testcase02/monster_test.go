package monster
import(
	"testing"
)
func TestStore(t *testing.T){
	m:=&Monster{
		Name: "红孩儿",
		Age: 10,
		Skill: "吐火",
	}
	res:=m.Store()
	if !res{
		t.Logf("测试错误，希望值为%v 实际值为%v\n",true,res)

	}
	t.Logf("测试成功！！！")
}
func TestReStore(t *testing.T){
	m:=&Monster{}
	res:=m.ReStore()
	if !res{
		t.Logf("测试错误，希望值为%v 实际值为%v\n",true,res)
	}
	t.Logf("测试成功！！！")
}