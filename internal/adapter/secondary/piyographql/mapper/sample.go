package mapper

import (
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/secondary/piyographql/generated"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
)

func ToSampleModel(s *generated.SampleQuerySample) *model.Sample {
	return &model.Sample{
		ID:        s.GetId(),
		StringVal: s.GetStringVal(),
		IntVal:    s.GetIntVal(),
		ArrayVal:  s.GetArrayVal(),
		Email:     s.GetEmail(),
		CreatedAt: s.GetCreatedAt(),
		UpdatedAt: s.GetUpdatedAt(),
	}
}

func ToSampleModelList(list []generated.ListSampleQueryListSample) []*model.Sample {
	var samples []*model.Sample
	for _, s := range list {
		samples = append(samples, &model.Sample{
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

func ToCreateSampleInput(sample *model.Sample) generated.CreateSampleInput {
	return generated.CreateSampleInput{
		StringVal: sample.StringVal,
		IntVal:    sample.IntVal,
		ArrayVal:  sample.ArrayVal,
		Email:     sample.Email,
	}
}

func ToSampleModelByCreateSample(s *generated.CreateSampleMutationCreateSampleCreateSamplePayloadSample) *model.Sample {
	return &model.Sample{
		ID:        s.GetId(),
		StringVal: s.GetStringVal(),
		IntVal:    s.GetIntVal(),
		ArrayVal:  s.GetArrayVal(),
		Email:     s.GetEmail(),
		CreatedAt: s.GetCreatedAt(),
		UpdatedAt: s.GetUpdatedAt(),
	}
}
