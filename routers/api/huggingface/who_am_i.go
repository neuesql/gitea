package huggingface

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	api "code.gitea.io/gitea/modules/structs"
	"net/http"
)

func WhoAMIHandler(ctx *context.APIContext) {
	token := ctx.Req.Header.Get("Bearer")

	//todo how to getWhoAmIResponseByToken, mapping into this response
	log.Debug(token)
	user := ctx.Doer

	whoAmIResponse := api.WhoAmIResponse{
		Id:            string(user.ID),
		Type:          "user",
		Name:          user.Name,
		FullName:      user.FullName,
		Email:         user.Email,
		CanPay:        false,
		Plan:          "NO_PLAN",
		AvatarUrl:     "https://aeiljuispo.cloudimg.io/v7/https://s3.amazonaws.com/moonup/production/uploads/1674866533852-63d46a7c1da8c28b7b8abb0a.png?w=200&h=200&f=face",
		PeriodEnd:     0,
		Auth:          api.AuthResponse{},
		EmailVerified: user.IsActive,
		IsPro:         false,
		Orgs:          nil,
	}
	ctx.JSON(http.StatusOK, whoAmIResponse)
}
