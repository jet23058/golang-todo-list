package responses

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, res *Success) {
	c.JSON(res.StatusCode, gin.H{
		"data":    res.Data,
		"status":  res.Status,
		"message": res.Message,
	})
}
