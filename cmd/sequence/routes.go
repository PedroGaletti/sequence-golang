package sequence

import (
	"github.com/gin-gonic/gin"
)

// Router : starting Sequence handler
func Router(r gin.IRoutes, sequenceController ISequenceController) {
	r.POST("/", sequenceController.SequenceRequestValidate)
	r.GET("/stats", sequenceController.StatsRequestInformation)
}
