package piyographql_test

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/Khan/genqlient/graphql"
	"github.com/NishimuraTakuya-nt/go-rest-chi/tests/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/adapter/secondary/piyographql"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/common/apperror"
	"github.com/NishimuraTakuya-nt/go-rest-chi/internal/domain/interface/external/piyographqls"
)

var _ = Describe("BaseClient", func() {
	var (
		baseClient piyographqls.BaseClient
		ctx        context.Context
		testHelper *helper.TestHelper
	)

	BeforeEach(func() {
		ctx = context.Background()
		testHelper = helper.NewTestHelper()
		baseClient = piyographql.NewBaseClient(testHelper.Logger, testHelper.Config)
	})

	Describe("Execute", func() {
		Context("正常系", func() {
			It("クエリが成功した場合、結果を返すこと", func() {
				expectedResult := map[string]interface{}{"data": "test"}
				query := func() (interface{}, error) {
					return expectedResult, nil
				}

				result, err := baseClient.Execute(ctx, "TestOperation", query)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(expectedResult))
			})
		})

		Context("異常系", func() {
			Context("GraphQLエラーの場合", func() {
				It("バリデーションエラーを適切に処理すること", func() {
					gqlErrors := gqlerror.List{
						{
							Message: "validation failed",
							Path:    ast.Path{ast.PathName("fieldName")},
							Extensions: map[string]interface{}{
								"code": "GRAPHQL_VALIDATION_FAILED",
							},
						},
					}
					query := func() (interface{}, error) {
						return nil, gqlErrors
					}

					result, err := baseClient.Execute(ctx, "TestOperation", query)

					Expect(result).To(BeNil())
					var validationErr *apperror.ValidationErrors
					Expect(errors.As(err, &validationErr)).To(BeTrue())
				})

				It("認証エラーを適切に処理すること", func() {
					gqlErrors := gqlerror.List{
						{
							Message: "unauthorized",
							Extensions: map[string]interface{}{
								"code": "UNAUTHENTICATED",
							},
						},
					}
					query := func() (interface{}, error) {
						return nil, gqlErrors
					}

					result, err := baseClient.Execute(ctx, "TestOperation", query)

					Expect(result).To(BeNil())
					var unauthorizedErr *apperror.AppError
					Expect(errors.As(err, &unauthorizedErr)).To(BeTrue())
				})

				It("Not Foundエラーを適切に処理すること", func() {
					gqlErrors := gqlerror.List{
						{
							Message: "resource not found",
							Extensions: map[string]interface{}{
								"code": "NOT_FOUND",
							},
						},
					}
					query := func() (interface{}, error) {
						return nil, gqlErrors
					}

					result, err := baseClient.Execute(ctx, "TestOperation", query)

					Expect(result).To(BeNil())
					var notFoundErr *apperror.AppError
					Expect(errors.As(err, &notFoundErr)).To(BeTrue())
				})

				It("レート制限エラーを適切に処理すること", func() {
					gqlErrors := gqlerror.List{
						{
							Message: "rate limit exceeded",
							Extensions: map[string]interface{}{
								"code": "RATE_LIMITED",
							},
						},
					}
					query := func() (interface{}, error) {
						return nil, gqlErrors
					}

					result, err := baseClient.Execute(ctx, "TestOperation", query)

					Expect(result).To(BeNil())
					var rateLimitErr *apperror.AppError
					Expect(errors.As(err, &rateLimitErr)).To(BeTrue())
				})
			})

			Context("ネットワークエラーの場合", func() {
				It("タイムアウトエラーを適切に処理すること", func() {
					netErr := &net.OpError{
						Op:  "dial",
						Net: "tcp",
						Err: fmt.Errorf("timeout"),
					}
					query := func() (interface{}, error) {
						return nil, netErr
					}

					result, err := baseClient.Execute(ctx, "TestOperation", query)

					Expect(result).To(BeNil())
					var externalErr *apperror.AppError
					Expect(errors.As(err, &externalErr)).To(BeTrue())
				})

				It("コンテキストのタイムアウトを適切に処理すること", func() {
					query := func() (interface{}, error) {
						return nil, context.DeadlineExceeded
					}

					result, err := baseClient.Execute(ctx, "TestOperation", query)

					Expect(result).To(BeNil())
					var timeoutErr *apperror.AppError
					Expect(errors.As(err, &timeoutErr)).To(BeTrue())
				})
			})
		})
	})

	Describe("GqlClient", func() {
		It("設定されたGraphQLクライアントを返すこと", func() {
			client := baseClient.GqlClient()
			Expect(client).NotTo(BeNil())
			_, ok := client.(graphql.Client)
			Expect(ok).To(BeTrue())
		})
	})
})
