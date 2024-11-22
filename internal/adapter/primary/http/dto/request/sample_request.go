package request

import (
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
)

// SampleRequest
// @Description Sample information
type SampleRequest struct {
	// refs: https://github.com/swaggo/swag#example-value-of-struct
	ID                      string        `json:"id" validate:"sampleId"`
	StringVal               string        `json:"string_val" validate:"required,min=0,max=50"`
	IntVal                  int           `json:"int_val" validate:"required,gte=1"`
	ArrayVal                []string      `json:"array_val"`
	Email                   string        `json:"email" validate:"omitempty,email" example:"test@example.com"`
	SampleDetailRequired    *SampleDetail `json:"sample_detail_required" validate:"required"`
	SampleDetailNotRequired *SampleDetail `json:"sample_detail_not_required" validate:"omitempty"`
}

// SampleDetail
// @Description Sample detail information
type SampleDetail struct {
	ID    int    `json:"id" validate:"required,gte=1"`
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Price int    `json:"price" validate:"omitempty,gte=1"`
}

func (r SampleRequest) ToSampleModel() *model.Sample {
	return &model.Sample{
		ID:        r.ID,
		StringVal: r.StringVal,
		IntVal:    r.IntVal,
		ArrayVal:  r.ArrayVal,
		Email:     r.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
