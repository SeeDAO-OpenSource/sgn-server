package memebersv1

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

// @Summary Get all members
// @Description Get all members
// @Tags Identity
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /identity/v1 [get]
// @param page query string false "page"
// @param size query string false "size"
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

// @Summary Get member by address
// @Description Get member by address
// @Tags Identity
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /identity/v1/{address} [get]
// @param address path string true "address"
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

// @Summary Get members by addresses
// @Description Get members by addresses
// @Tags Identity
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /identity/v1/range/{addresses} [get]
// @param addresses path string true "addresses"
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

// @Summary Insert member
// @Description Insert member
// @Tags Identity
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /identity/v1 [post]
// @param member body identity.Member true "member"
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

// @Summary Update member
// @Description Update member
// @Tags Identity
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /identity/v1/{address} [put]
// @param address path string true "address"
// @param member body identity.Member true "member"
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

// @Summary Delete member
// @Description Delete member
// @Tags Identity
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /identity/v1/{address} [delete]
// @param address path string true "address"
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
