package memberapi

import (
	"net/http"
	"strings"

	"github.com/SeeDAO-OpenSource/sgn/internal/member"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func NewMemberController() MemberController {
	return MemberController{}
}

// @Summary Get all members
// @Description Get all members
// @Tags Member
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /members [get]
// @param page query string false "page"
// @param size query string false "size"
func (c MemberController) GetList(ctx *gin.Context) {
	srv, err := member.NewMemberService()
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
// @Tags Member
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /members/{address} [get]
// @param address path string true "address"
func (c MemberController) GetByAddress(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	srv, err := member.NewMemberService()
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
// @Tags Member
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /members/range/{addresses} [get]
// @param addresses path string true "addresses"
func (c MemberController) GetByAddresses(ctx *gin.Context) {
	param := ctx.Param("addresses")
	if len(param) == 0 {
		mvc.Fail(ctx, http.StatusBadRequest, "addresses is empty")
		return
	}
	srv, err := member.NewMemberService()
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
// @Tags Member
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /members [post]
// @param member body member.Member true "member"
func (c MemberController) Insert(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	model := member.Member{}
	err := ctx.BindJSON(&model)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	srv, err := member.NewMemberService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	err = srv.Insert(&model)
	if err != nil {
		mvc.Error(ctx, err)
	} else {
		mvc.Ok(ctx, true)
	}
}

// @Summary Update member
// @Description Update member
// @Tags Member
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /members/{address} [put]
// @param address path string true "address"
// @param member body member.Member true "member"
func (c MemberController) Update(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	model := member.Member{}
	err := ctx.BindJSON(&model)
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	srv, err := member.NewMemberService()
	if err != nil {
		mvc.Error(ctx, err)
		return
	}
	err = srv.Update(&model)
	if err != nil {
		mvc.Error(ctx, err)
	} else {
		mvc.Ok(ctx, true)
	}
}

// @Summary Delete member
// @Description Delete member
// @Tags Member
// @Accept  json
// @Produce  json
// @Success 200 {object} mvc.DataResult
// @Router /members/{address} [delete]
// @param address path string true "address"
func (c MemberController) Delete(ctx *gin.Context) {
	address := ctx.Param("address")
	if address == "" {
		mvc.Fail(ctx, http.StatusBadRequest, "address is empty")
		return
	}
	srv, err := member.NewMemberService()
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
