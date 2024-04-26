package router

import (
	"backend/src/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, cc controller.ICuisineController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //corsのミドルウェア
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")}, //デプロイしたときに取得できるドメイン
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, //許可するヘッダーの一覧
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"}, //許可したいメソッド
		AllowCredentials: true,                                     //クッキーの送受信を可能にする
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{ //csrfのミドルウェア
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		//CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode, //postmanで確認のため
	}))
	e.POST("/signup", uc.SignUp) //エンドポイント追加
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	// e.PUT("/update", uc.Update)
	// e.PUT("/update", uc.Update, echojwt.WithConfig(echojwt.Config{
	// 	SigningKey:  []byte(os.Getenv("SECRET")),
	// 	TokenLookup: "cookie:token",
	// }))
	e.GET("/csrf", uc.CsrfToken)

	u := e.Group("/update")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	u.PUT("", uc.Update)

	c := e.Group("/cuisines")
	c.Use(echojwt.WithConfig(echojwt.Config{ //エンドポイントにミドルウェアを追加
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	c.GET("", cc.GetAllCuisines)            //cuisinesのエンドポイントにリクエストがあった場合
	c.GET("/:cuisineId", cc.GetCuisineById) //リクエストパラメーターにcuisineidが入力された場合
	c.POST("", cc.CreateCuisine)
	c.PUT("/:cuisineId", cc.UpdateCuisine) //titleしか更新されない
	c.DELETE("/:cuisineId", cc.DeleteCuisine)

	c.PUT("/image/:cuisineId", cc.UploadImage)
	//c.PUT("/url/:cuisineId", cc.AddURL)
	return e
}
