package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "invafresh.com/identity-provider/api/docs"
	"invafresh.com/identity-provider/config"
	"invafresh.com/identity-provider/token"

	swaggerfiles "github.com/swaggo/files"
)

type Api struct {
	Config *config.SConfig
}

func NewApi(config *config.SConfig) Api {
	return Api{
		Config: config,
	}
}

func (a *Api) Start() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/token")
		{
			eg.POST("", a.createToken)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(fmt.Sprintf("%s:%s", a.Config.Api.Host, a.Config.Api.Port))
}

// Create coupon
//
//	@Tags		token
//	@Summary	create token
//	@Accept		json
//	@Produce	json
//	@Param		params			body		TokenRequest				false	"Query Params"
//	@Success	200				{object}	TokenResponse
//	@Failure	401				{object}	string
//	@Router		/token [POST]
func (a *Api) createToken(c *gin.Context) {
	request := token.TokenRequest{}

	if err := c.BindJSON(&request); err != nil {
		return
	}

	con := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		a.Config.MsSql.Server, a.Config.MsSql.User, a.Config.MsSql.Password, a.Config.MsSql.Port, a.Config.MsSql.Database)
	fmt.Println("con->", con)

	validater := token.NewOwnerPasswordValidator(con)

	user, ok := validater.Validate(request)

	if ok {
		jwtToken, _ := token.GenerateJWT(user.User_Id)

		response := token.TokenResponse{
			Token:      jwtToken,
			ExpiryTime: 123,
		}

		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, "invalid")
	}
}
