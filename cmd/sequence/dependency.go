package sequence

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	Repository ISequenceRepository
	Controller ISequenceController
)

// InjectDependency: create the injections of Sequence
func InjectDependency(router gin.IRoutes, db *gorm.DB) {
	// Repository
	Repository := NewSequenceRepository(db)

	// Controller
	Controller := NewSequenceController(Repository)

	// Routes
	Router(router, Controller)
}
