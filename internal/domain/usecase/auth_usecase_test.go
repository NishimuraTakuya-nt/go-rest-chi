package usecase_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/model"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/usecase"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/mock/mockservice"
)

var _ = Describe("AuthUsecase", func() {
	var (
		ctrl        *gomock.Controller
		mockToken   *mockservice.MockTokenService
		authUsecase usecase.AuthUsecase
		ctx         context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockToken = mockservice.NewMockTokenService(ctrl)
		authUsecase = usecase.NewAuthUsecase(mockToken)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Authenticate", func() {
		Context("正常系", func() {
			It("有効なトークンから正しいユーザー情報を取得できること", func() {
				// テストデータ
				tokenString := "valid.jwt.token"
				expectedClaims := &model.Claims{
					UserID: "test-user-123",
					Roles:  []string{"admin", "user"},
				}
				expectedUser := &model.User{
					ID:    expectedClaims.UserID,
					Roles: expectedClaims.Roles,
				}

				// モックの設定
				mockToken.EXPECT().
					ValidateToken(ctx, tokenString).
					Return(expectedClaims, nil)

				// テスト実行
				user, err := authUsecase.Authenticate(ctx, tokenString)

				// 検証
				Expect(err).NotTo(HaveOccurred())
				Expect(user).To(Equal(expectedUser))
				Expect(user.ID).To(Equal(expectedClaims.UserID))
				Expect(user.Roles).To(Equal(expectedClaims.Roles))
			})
		})

		Context("異常系", func() {
			Context("トークンが無効な場合", func() {
				It("エラーを返すこと", func() {
					// テストデータ
					invalidToken := "invalid.token"
					expectedErr := errors.New("invalid token")

					// モックの設定
					mockToken.EXPECT().
						ValidateToken(ctx, invalidToken).
						Return(nil, expectedErr)

					// テスト実行
					user, err := authUsecase.Authenticate(ctx, invalidToken)

					// 検証
					Expect(err).To(Equal(expectedErr))
					Expect(user).To(BeNil())
				})
			})

			Context("トークンが空の場合", func() {
				It("エラーを返すこと", func() {
					// テストデータ
					emptyToken := ""
					expectedErr := errors.New("token is empty")

					// モックの設定
					mockToken.EXPECT().
						ValidateToken(ctx, emptyToken).
						Return(nil, expectedErr)

					// テスト実行
					user, err := authUsecase.Authenticate(ctx, emptyToken)

					// 検証
					Expect(err).To(Equal(expectedErr))
					Expect(user).To(BeNil())
				})
			})

			Context("トークンの検証中にエラーが発生した場合", func() {
				It("エラーを返すこと", func() {
					// テストデータ
					token := "valid.looking.token"
					expectedErr := errors.New("validation error")

					// モックの設定
					mockToken.EXPECT().
						ValidateToken(ctx, token).
						Return(nil, expectedErr)

					// テスト実行
					user, err := authUsecase.Authenticate(ctx, token)

					// 検証
					Expect(err).To(Equal(expectedErr))
					Expect(user).To(BeNil())
				})
			})
		})
	})
})
