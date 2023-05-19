package huggingface

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	api "code.gitea.io/gitea/modules/structs"
	"net/http"
)

func HandleWhoAmI(ctx *context.APIContext) {
	token := ctx.Req.Header.Get("Bearer")

	//todo how to getWhoAmIResponseByToken, mapping into this response
	log.Debug(token)

	whoAmIResponse := api.WhoAmIResponse{
		Id:            "63d46a7c1da8c28b7b8abb0a",
		Type:          "user",
		Name:          "qunfei",
		FullName:      "Qunfei Wu",
		Email:         "wu.qunfei@gmail.com",
		CanPay:        false,
		Plan:          "NO_PLAN",
		AvatarUrl:     "https://aeiljuispo.cloudimg.io/v7/https://s3.amazonaws.com/moonup/production/uploads/1674866533852-63d46a7c1da8c28b7b8abb0a.png?w=200&h=200&f=face",
		PeriodEnd:     0,
		Auth:          api.AuthResponse{},
		EmailVerified: true,
		IsPro:         false,
		Orgs:          nil,
		Scope:         nil,
	}
	ctx.JSON(http.StatusOK, whoAmIResponse)
}
