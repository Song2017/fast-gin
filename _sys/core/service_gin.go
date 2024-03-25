package sys_core

import (
	t "server/_sys/type"

	"github.com/gin-gonic/gin"
)

type GinProcess interface {
	Input()
	Process()

	Invoke(*gin.Context, ...OperationBase) bool
}
type Operation func(*gin.Context)

type ServiceGin struct {
	ServiceBase
}

func (p *ServiceGin) Input() {
	p.Request = make(map[string]interface{})
	p.Response = map[string]interface{}{
		"code": 400,
	}
}

func (p *ServiceGin) Error(c *gin.Context, result t.M) {
	code := result["code"].(int)
	c.AbortWithStatusJSON(code, result["error"])
}

// useless, it calls methods of ServiceGin
func (p *ServiceGin) Invoke(c *gin.Context, funcs ...Operation) bool {
	for _, f := range funcs {
		f(c)
		if p.Response["error"] != nil {
			p.Error(c, p.Response)
			return false
		}
	}
	return true
}
