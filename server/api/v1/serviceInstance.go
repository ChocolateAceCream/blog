/*
* @fileName serviceInstance.go
* @author Di Sheng
* @date 2022/10/18 08:50:27
* @description Warp out each service instance, so that service can be shared between each api
 */

package apiV1

import (
	"github.com/ChocolateAceCream/blog/initializer"
	"github.com/ChocolateAceCream/blog/service"
)

var userService = new(service.UserService)
var casbinService = new(service.CasbinService)
var roleService = new(service.RoleService)
var menuService = new(service.MenuService)
var initService = new(initializer.InitService)
var endpointService = new(service.EndpointService)
