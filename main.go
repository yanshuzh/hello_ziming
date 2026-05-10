package main

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	h.GET("/uppercase", func(ctx context.Context, c *app.RequestContext) {
		input := c.Query("text")
		if input == "" {
			c.JSON(consts.StatusBadRequest, utils.H{
				"error": "text parameter is required",
			})
			return
		}

		result := strings.ToUpper(input)
		c.JSON(consts.StatusOK, utils.H{
			"original": input,
			"result":   result,
		})
	})

	h.POST("/uppercase", func(ctx context.Context, c *app.RequestContext) {
		type Request struct {
			Text string `json:"text"`
		}

		var req Request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(consts.StatusBadRequest, utils.H{
				"error": "invalid JSON format",
			})
			return
		}

		if req.Text == "" {
			c.JSON(consts.StatusBadRequest, utils.H{
				"error": "text field is required",
			})
			return
		}

		result := strings.ToUpper(req.Text)
		c.JSON(consts.StatusOK, utils.H{
			"original": req.Text,
			"result":   result,
		})
	})

	h.Spin()
}
