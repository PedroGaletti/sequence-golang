package sequence

import (
	"github.com/jinzhu/gorm"
)

const (
	tbSequences = "sequences"
)

// ISequenceRepository: interface of Sequence repository
type ISequenceRepository interface {
	Store(*Sequence) error
	FindAll() ([]Sequence, error)
}

// SequenceRepository: struct of Sequence repository
type SequenceRepository struct {
	db *gorm.DB
}

// NewSequenceRepository: create a new Sequence repository
func NewSequenceRepository(db *gorm.DB) ISequenceRepository {
	return &SequenceRepository{db}
}

// Save: save the sequence
func (r *SequenceRepository) Store(sequence *Sequence) error {
	return r.db.Table(tbSequences).Create(&sequence).Error
}

// FindAll: get all the sequences inside the database
func (r *SequenceRepository) FindAll() ([]Sequence, error) {
	sequences := []Sequence{}

	err := r.db.Table(tbSequences).Model(&Sequence{}).Find(&sequences).Error

	return sequences, err
}
