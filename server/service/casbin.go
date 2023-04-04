package service

import (
	"strconv"
	"sync"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

type CasbinService struct{}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (c *CasbinService) Update(id uint, endpoints []request.Endpoint) error {
	roleId := strconv.Itoa(int(id))
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
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
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
