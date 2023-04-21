package service

import (
	"strconv"
	"sync"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

type CasbinService struct{}

var CasbinServiceInstance = new(CasbinService)
var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (c *CasbinService) List(id int) (list []response.CasbinPolicy) {
	e := GetEnforcer()
	roleId := strconv.Itoa(id)
	t := e.GetFilteredPolicy(0, roleId)
	for _, v := range t {
		list = append(list, response.CasbinPolicy{
			Path:   v[1],
			Method: v[2],
		})
	}
	return list
}

func (c *CasbinService) Update(id uint, endpoints []request.CasbinPolicy) error {
	roleId := strconv.Itoa(int(id))
	CasbinServiceInstance.ClearCasbin(0, roleId) // clean up casbin for the role first
	rules := [][]string{}
	for _, v := range endpoints {
		rules = append(rules, []string{roleId, v.Path, v.Method})
	}
	e := GetEnforcer()
	if success, err := e.AddPolicies(rules); !success {
		return err
	}
	return nil
}

func (c *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := GetEnforcer()
	if hasPolicy := e.GetFilteredPolicy(v, p...); hasPolicy != nil {
		success, _ := e.RemoveFilteredPolicy(v, p...)
		return success
	}
	return true
}

func GetEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		text := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.LOGGER.Error("fail to get casbin enforcer", zap.Error(err))
			return
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a)
	})
	syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
