package idv1

import (
	"net/http"
	"strings"

	"github.com/SeeDAO-OpenSource/sgn/internal/identity"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/gin-gonic/gin"
)

type IdentityController struct {
}

func newIdentityController() IdentityController {
	return IdentityController{}
}

func (c IdentityController) GetList(ctx *gin.Context) {
	srv, err := identity.NewIdentityService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	pageIndex, pageSize := mvc.PageQuery(ctx)
	data, err := srv.GetList(pageIndex, pageSize)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	mvc.Ok(ctx, data)
}

func (c IdentityController) GetByAddress(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	srv, err := identity.NewIdentityService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	member, err := srv.GetByAddress(address)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	mvc.Ok(ctx, member)
}

func (c IdentityController) GetByAddresses(ctx *gin.Context) {
	param := ctx.Param("addresses")
	if len(param) == 0 {
		mvc.Fail(ctx, http.StatusBadRequest, "addresses is empty")
		return
	}
	srv, err := identity.NewIdentityService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	addresses := strings.Split(param, ",")
	members, err := srv.GetByAddresses(addresses)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	mvc.Ok(ctx, members)
}

func (c IdentityController) Insert(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	member := identity.Member{}
	err := ctx.BindJSON(&member)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	srv, err := identity.NewIdentityService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	err = srv.Insert(&member)
	if err != nil {
		mvc.Error(ctx, err)
	} else {
		mvc.Ok(ctx, true)
	}
}

func (c IdentityController) Update(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	member := identity.Member{}
	err := ctx.BindJSON(&member)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	srv, err := identity.NewIdentityService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	err = srv.Update(&member)
	if err != nil {
		mvc.Error(ctx, err)
	} else {
		mvc.Ok(ctx, true)
	}
}

func (c IdentityController) Delete(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	srv, err := identity.NewIdentityService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	if err := srv.Delete(address); err != nil {
		mvc.Error(ctx, err)
	} else {
		mvc.Ok(ctx, true)
	}
}
