package sys_core

import (
	"context"
)

type BaseProcess interface {
	Input()
	Process()

	Invoke(*context.Context, ...OperationBase) bool
}

type OperationBase func(*context.Context)

type ServiceBase struct {
	Request map[string]interface{}
	Body    map[string]interface{}

	Response map[string]interface{} // code, result, error
}

func (p *ServiceBase) Input() {
	// return err or failed check result
	p.Request = make(map[string]interface{})
	p.Response = map[string]interface{}{
		"error": "err",
	}
}

func (p *ServiceBase) Process() {
	// return err or exception

}

func (p *ServiceBase) Invoke(c *context.Context, funcs ...OperationBase) bool {
	for _, f := range funcs {
		f(c)
		if p.Response["error"] != nil {
			return false
		}
	}
	return true
}
