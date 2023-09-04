package usecase_test

import (
	"fmt"
	mock_repository "go-rest-api/mock/repository"
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/usecase"
	"go-rest-api/validator"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func Test_userUsecase_SignUp(t *testing.T) {
	type fields struct {
		ur func(*testing.T) repository.IUserRepository
		uv validator.IUserValidator
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UserResponse
		wantErr bool
	}{
		// 正常終了
		{
			name: "正常終了",
			fields: fields{
				ur: func(t *testing.T) repository.IUserRepository {
					t.Helper()                      // おまじない
					ctrl := gomock.NewController(t) // おまじない
					mock := mock_repository.NewMockIUserRepository(ctrl)
					mock.EXPECT().CreateUser(
						gomock.Any(), // モックに期待する引数を指定（スキップ）
					).Return(
						nil, // モックの返り値を指定
					)
					return mock
				},
				uv: validator.NewUserValidator(), // 実体を渡す（テスト対象外）
			},
			args: args{
				user: model.User{
					Email:    "user1@test.com",
					Password: "password",
				},
			},
			want: model.UserResponse{
				ID:    0,
				Email: "user1@test.com",
			},
			wantErr: false,
		},
		// データベース処理失敗
		{
			name: "データベース処理失敗",
			fields: fields{
				ur: func(t *testing.T) repository.IUserRepository {
					t.Helper()                      // おまじない
					ctrl := gomock.NewController(t) // おまじない
					mock := mock_repository.NewMockIUserRepository(ctrl)
					mock.EXPECT().CreateUser(
						gomock.Any(),
					).Return(
						fmt.Errorf("failed to create user"),
					)
					return mock
				},
				uv: validator.NewUserValidator(), // 実体を渡す（テスト対象外）
			},
			args: args{
				user: model.User{
					Email:    "user1@test.com",
					Password: "password",
				},
			},
			want:    model.UserResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt // おまじない
		t.Run(tt.name, func(t *testing.T) {
			// 準備
			t.Parallel()
			uu := usecase.NewUserUsecase(tt.fields.ur(t), tt.fields.uv)
			// 実行
			got, err := uu.SignUp(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 検証
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("userUsecase.SignUp() mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
