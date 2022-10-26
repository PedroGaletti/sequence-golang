package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// If you pass an empty the controller will return a error.

var (
	valid_array           = []string{"DUHBHB", "DDBUHD", "UBDUHU", "BHBDHH", "DDDDUB", "UDBDUH"}
	invalid_array         = []string{"CCHBHB", "DUDUHD", "UBUUHU", "BHBDDH", "DDBDUB", "UDBDUH"}
	valid_oblique_array   = []string{"D", "D", "D", "D", "U", "H", "D", "B", "B", "D", "U", "U", "B", "U", "H", "B", "U", "H", "D", "D", "H", "U", "H", "H", "B", "D", "B", "B", "H", "U", "D", "D", "H", "D", "U", "B"}
	invalid_oblique_array = []string{"C", "D", "D", "D", "U", "H", "D", "B", "B", "D", "U", "U", "B", "U", "H", "B", "U", "H", "D", "D", "H", "U", "H", "H", "B", "D", "B", "B", "H", "U", "D", "D", "H", "D", "U", "B"}
	valid_marshal_array   = `["DUHBHB","DDBUHD","UBDUHU","BHBDHH","DDDDUB","UDBDUH"]`
	invalid_marshal_array = `["CCHBHB","DUDUHD","UBUUHU","BHBDDH","DDBDUB","UDBDUH"]`
	valid_length          = 4
	invalid_length        = 5
)

func TestHorizontal(t *testing.T) {
	t.Run("Horizontal - Test a valid case", func(t *testing.T) {
		assert.Equal(t, int64(1), Horizontal(valid_array, int64(valid_length)))
	})

	t.Run("Horizontal - Test an invalid array case", func(t *testing.T) {
		assert.Equal(t, int64(0), Horizontal(invalid_array, int64(valid_length)))
	})

	t.Run("Horizontal - Test a valid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), Horizontal(valid_array, int64(invalid_length)))
	})

	t.Run("Horizontal - Test an invalid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), Horizontal(invalid_array, int64(invalid_length)))
	})
}

func TestVertical(t *testing.T) {
	t.Run("Vertical - Test a valid case", func(t *testing.T) {
		assert.Equal(t, int64(1), Vertical(valid_array, int64(valid_length)))
	})

	t.Run("Vertical - Test an invalid array case", func(t *testing.T) {
		assert.Equal(t, int64(0), Vertical(invalid_array, int64(valid_length)))
	})

	t.Run("Vertical - Test a valid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), Vertical(valid_array, int64(invalid_length)))
	})

	t.Run("Vertical - Test an invalid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), Vertical(invalid_array, int64(invalid_length)))
	})
}

func TestBottomUpOblique(t *testing.T) {
	t.Run("BottomUpOblique - Test a valid case", func(t *testing.T) {
		assert.Equal(t, int64(1), BottomUpOblique(valid_array, int64(valid_length)))
	})

	t.Run("BottomUpOblique - Test an invalid array case", func(t *testing.T) {
		assert.Equal(t, int64(0), BottomUpOblique(invalid_array, int64(valid_length)))
	})

	t.Run("BottomUpOblique - Test a valid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), BottomUpOblique(valid_array, int64(invalid_length)))
	})

	t.Run("BottomUpOblique - Test an invalid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), BottomUpOblique(invalid_array, int64(invalid_length)))
	})
}

func TestUpBottomOblique(t *testing.T) {
	t.Run("UpBottomOblique - Test a valid case", func(t *testing.T) {
		assert.Equal(t, int64(1), UpBottomOblique(valid_array, int64(valid_length)))
	})

	t.Run("UpBottomOblique - Test an invalid array case", func(t *testing.T) {
		assert.Equal(t, int64(0), UpBottomOblique(invalid_array, int64(valid_length)))
	})

	t.Run("UpBottomOblique - Test a valid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), UpBottomOblique(valid_array, int64(invalid_length)))
	})

	t.Run("UpBottomOblique - Test an invalid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), UpBottomOblique(invalid_array, int64(invalid_length)))
	})
}

func TestCount(t *testing.T) {
	t.Run("Count - Test a valid oblique array case", func(t *testing.T) {
		assert.Equal(t, int64(1), Count(valid_oblique_array, int64(valid_length)))
	})

	t.Run("Count - Test an invalid oblique array case", func(t *testing.T) {
		assert.Equal(t, int64(0), Count(invalid_oblique_array, int64(valid_length)))
	})

	t.Run("Count - Test a valid oblique array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), Count(valid_oblique_array, int64(invalid_length)))
	})

	t.Run("Count - Test an invalid array with invalid length case", func(t *testing.T) {
		assert.Equal(t, int64(0), Count(invalid_oblique_array, int64(invalid_length)))
	})
}

func TestValidateLetters(t *testing.T) {
	t.Run("ValidateLetters - Test a valid array case", func(t *testing.T) {
		assert.Equal(t, true, ValidateLetters(valid_array))
	})

	t.Run("ValidateLetters - Test an invalid array case", func(t *testing.T) {
		assert.Equal(t, false, ValidateLetters(invalid_array))
	})
}

func TestSequenceProcessValidate(t *testing.T) {
	t.Run("SequenceProcessValidate - Test a valid array and valid length case", func(t *testing.T) {
		isValid, lMarshal := SequenceProcessValidate(valid_array, int64(valid_length))

		assert.Equal(t, true, isValid)
		assert.Equal(t, valid_marshal_array, lMarshal)
	})

	t.Run("SequenceProcessValidate - Test a valid array and invalid length case", func(t *testing.T) {
		isValid, lMarshal := SequenceProcessValidate(valid_array, int64(invalid_length))

		assert.Equal(t, false, isValid)
		assert.Equal(t, valid_marshal_array, lMarshal)
	})

	t.Run("SequenceProcessValidate - Test a invalid array and invalid length case", func(t *testing.T) {
		isValid, lMarshal := SequenceProcessValidate(invalid_array, int64(invalid_length))

		assert.Equal(t, false, isValid)
		assert.Equal(t, invalid_marshal_array, lMarshal)
	})

	t.Run("SequenceProcessValidate - Test a invalid array and valid length case", func(t *testing.T) {
		isValid, lMarshal := SequenceProcessValidate(invalid_array, int64(valid_length))

		assert.Equal(t, false, isValid)
		assert.Equal(t, invalid_marshal_array, lMarshal)
	})
}
