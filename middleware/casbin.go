package middleware

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/response"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if claims, ok := c.Get("claims"); !ok {
			response.Fail(c, response.TokenFailed, nil)
			c.Abort()
			return
		} else {
			userinfo := claims.(*model.Claims)
			obj := c.Request.URL.Path
			act := c.Request.Method
			sub := userinfo.RoleID
			e := SyncedEnforcer()
			// 判断策略中是否存在
			if ok, err := e.Enforce(sub, obj, act); ok {
				c.Next()
			} else {
				response.Fail(c, response.CasbinFailed, err)
				c.Abort()
				return
			}
		}
	}
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func SyncedEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDBUseTableName(global.Db, "", "tb_casbin_rule")
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.Cfg.Casbin.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
