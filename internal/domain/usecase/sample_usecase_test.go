package usecase

import (
	"context"
	"testing"

	"github.com/NishimuraTakuya-nt/go-rest-chi/config"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/mock/mockpiyographql"
	"github.com/golang/mock/gomock"
)

func Test_sampleUsecase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockpiyographql.NewMockSampleClient(ctrl)
	target := NewSampleUsecase(logger.NewLogger(&config.AppConfig{}), mockClient) // fixme test cfg

	t.Run("get", func(t *testing.T) {
		ID := "123"
		// モックの振る舞いを設定
		mockClient.EXPECT().Sample(context.Background(), ID).
			Return(&model.Sample{ID: "123", StringVal: "Test Sample"}, nil)

		// テストケースを実行
		sample, err := target.Get(context.Background(), ID)

		assert.NoError(t, err)
		assert.Equal(t, "123", sample.ID)
	})

}
