package mapper

import (
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapters/secondary/piyographql/generated"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/core/domain/models"
)

func ToSampleModel(s *generated.SampleQuerySample) *models.Sample {
	return &models.Sample{
		ID:        s.GetId(),
		StringVal: s.GetStringVal(),
		IntVal:    s.GetIntVal(),
		ArrayVal:  s.GetArrayVal(),
		Email:     s.GetEmail(),
		CreatedAt: s.GetCreatedAt(),
		UpdatedAt: s.GetUpdatedAt(),
	}
}

func ToSampleModelList(list []generated.ListSampleQueryListSample) []*models.Sample {
	var samples []*models.Sample
	for _, s := range list {
		samples = append(samples, &models.Sample{
			ID:        s.GetId(),
			StringVal: s.GetStringVal(),
			IntVal:    s.GetIntVal(),
			ArrayVal:  s.GetArrayVal(),
			Email:     s.GetEmail(),
			CreatedAt: s.GetCreatedAt(),
			UpdatedAt: s.GetUpdatedAt(),
		})
	}
	return samples
}

func ToCreateSampleInput(sample *models.Sample) generated.CreateSampleInput {
	return generated.CreateSampleInput{
		StringVal: sample.StringVal,
		IntVal:    sample.IntVal,
		ArrayVal:  sample.ArrayVal,
		Email:     sample.Email,
	}
}

func ToSampleModelByCreateSample(s *generated.CreateSampleMutationCreateSampleCreateSamplePayloadSample) *models.Sample {
	return &models.Sample{
		ID:        s.GetId(),
		StringVal: s.GetStringVal(),
		IntVal:    s.GetIntVal(),
		ArrayVal:  s.GetArrayVal(),
		Email:     s.GetEmail(),
		CreatedAt: s.GetCreatedAt(),
		UpdatedAt: s.GetUpdatedAt(),
	}
}
