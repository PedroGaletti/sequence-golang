package sequence

import (
	"challenge/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	consecutive_length = 4
)

// ISequenceController: interface of Sequence controller
type ISequenceController interface {
	SequenceRequestValidate(*gin.Context)
	StatsRequestInformation(*gin.Context)
}

// SequenceController: struct of Sequence controller
type SequenceController struct {
	repository ISequenceRepository
}

// NewSequenceController: create a new Sequence controller
func NewSequenceController(repository ISequenceRepository) ISequenceController {
	return &SequenceController{repository}
}

// SequenceRequestValidate: validate the sequence of letters
func (c *SequenceController) SequenceRequestValidate(ctx *gin.Context) {
	letters := &Letters{}

	if err := ctx.BindJSON(&letters); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, SequenceResponseValidate{
			Message: fmt.Sprintf("Something is wrong in your request, err: %s", err.Error()),
			IsValid: false,
		})
		return
	}

	if len(letters.Letters) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, SequenceResponseValidate{
			Message: "You passed an empty array",
			IsValid: false,
		})
		return
	}

	isValid, lMarshal := utils.SequenceProcessValidate(letters.Letters, consecutive_length)

	if err := c.repository.Store(&Sequence{Letters: lMarshal, IsValid: isValid}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, SequenceResponseValidate{
			Message: fmt.Sprintf("Something is wrong in your request, err: %s", err.Error()),
			IsValid: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, SequenceResponseValidate{IsValid: isValid, Message: ""})
}

// StatsRequestInformation: Get the data from database and return the information
func (c *SequenceController) StatsRequestInformation(ctx *gin.Context) {
	// Get all sequences from database
	sequences, err := c.repository.FindAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, StatsInformationResponse{Message: fmt.Sprintf("No records found in this request, err: %s", err.Error())})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, StatsInformationResponse{Message: fmt.Sprintf("Something is wrong in your request, err: %s", err.Error())})
		return
	}

	// Start the vars if default golang value
	var countValid, countInvalid int64

	// Count the valid and invalids sequence inside the database
	for _, sequence := range sequences {
		if sequence.IsValid {
			countValid++
		} else {
			countInvalid++
		}
	}

	calc := float64(countValid) / float64(len(sequences))
	format := fmt.Sprintf("%.3f", calc)
	removeTheRound := format[:len(format)-1]
	ratio, _ := strconv.ParseFloat(removeTheRound, 64)

	// Building the response for the request
	ctx.JSON(http.StatusOK, &StatsInformationResponse{
		CountValid:   countValid,
		CountInvalid: countInvalid,
		Ratio:        ratio,
	})
}
