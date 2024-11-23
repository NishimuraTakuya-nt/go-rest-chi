package usecase_test

import (
	"context"
	"errors"

	"github.com/NishimuraTakuya-nt/go-rest-chi/tests/helper"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/usecase"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/mock/mockpiyographql"
)

var _ = Describe("SampleUsecase", func() {
	var (
		// テストで使用する変数を定義
		ctrl          *gomock.Controller
		mockClient    *mockpiyographql.MockSampleClient
		sampleUsecase usecase.SampleUsecase
		ctx           context.Context
		testHelper    *helper.TestHelper
	)

	// 各テストの前に実行される
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = mockpiyographql.NewMockSampleClient(ctrl)
		testHelper = helper.NewTestHelper()
		sampleUsecase = usecase.NewSampleUsecase(
			testHelper.Logger,
			mockClient,
		)
		ctx = context.Background()
	})

	// 各テストの後に実行される
	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Get", func() {
		Context("正常系", func() {
			It("指定したIDのSampleを取得できること", func() {
				// テストデータ準備
				expectedSample := &model.Sample{
					ID:        "test-id-123",
					StringVal: "Test Sample",
				}

				// モックの設定
				mockClient.EXPECT().
					Sample(ctx, expectedSample.ID).
					Return(expectedSample, nil)

				// テスト実行
				result, err := sampleUsecase.Get(ctx, expectedSample.ID)

				// 検証
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(expectedSample))
			})
		})

		Context("異常系", func() {
			Context("IDが空文字の場合", func() {
				It("エラーを返すこと", func() {
					// テスト実行
					result, err := sampleUsecase.Get(ctx, "")

					// 検証
					Expect(err).To(HaveOccurred())
					Expect(result).To(BeNil())
				})
			})

			Context("クライアントがエラーを返す場合", func() {
				It("エラーを返すこと", func() {
					id := "test-id-123"
					expectedErr := errors.New("client error")

					// モックの設定
					mockClient.EXPECT().
						Sample(ctx, id).
						Return(nil, expectedErr)

					// テスト実行
					result, err := sampleUsecase.Get(ctx, id)

					// 検証
					Expect(err).To(Equal(expectedErr))
					Expect(result).To(BeNil())
				})
			})
		})
	})
})
