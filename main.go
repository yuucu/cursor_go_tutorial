package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// リクエスト用の構造体
type HelloRequest struct {
	Text string `json:"text"`
}

// レスポンス用の構造体
type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	// 既存のGETエンドポイント
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 新しいPOSTエンドポイント
	e.POST("/hello", func(c echo.Context) error {
		// リクエストボディをパース
		req := new(HelloRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
		}

		// レスポンスを作成（受け取ったテキストをそのまま返す）
		res := HelloResponse{
			Message: req.Text,
		}

		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
