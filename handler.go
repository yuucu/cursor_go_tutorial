package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloRequest はhelloエンドポイントのリクエスト構造体
type HelloRequest struct {
	Message string `json:"message"`
}

// HelloResponse はhelloエンドポイントのレスポンス構造体
type HelloResponse struct {
	Message string `json:"message"`
}

// HelloHandler はhelloメッセージをそのまま返すハンドラー
func HelloHandler(c echo.Context) error {
	var req HelloRequest

	// リクエストボディをバインド
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// レスポンスを作成
	res := HelloResponse{
		Message: req.Message,
	}

	return c.JSON(http.StatusOK, res)
}
