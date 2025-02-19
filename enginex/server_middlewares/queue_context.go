package server_middlewares

import (
	"github.com/labstack/echo"
	"github.com/yot-anan-gj/ditp.thaitrade-sale-product/enginex/queue/aws_sqs"
	"github.com/yot-anan-gj/ditp.thaitrade-sale-product/enginex/server_constant"
)

func SqsContextAppender(sqss aws_sqs.AwsSqss) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(sqss) > 0 {
				c.Set(server_constant.SqsContextKey, sqss)
			}

			return next(c)
		}
	}
}
