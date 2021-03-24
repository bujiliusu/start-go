package remote_package
//https://github.com/easierway/concurrent_map
import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConcurrentMap(t *testing.T)  {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}


//go未解决的依赖问题
//1。同一环境下，不同项目使用同一包的不同版本
//2。无法管理对包的特定版本的依赖
//查找依赖包路径的解决方案如下：
//1。当前包下的rendor目录
//2。向上级目录查找，直到找到src下的vendor目录
//3。在gopath下面查找依赖包
//4。在goroot目录下查找

