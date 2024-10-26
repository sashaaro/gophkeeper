package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/service/mocks"
	"github.com/stretchr/testify/require"
)

func TestUserService_Login(t *testing.T) {
	type fields struct {
		hasher PasswordHasher
		repo   UserRepository
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				hasher: tt.fields.hasher,
				repo:   tt.fields.repo,
			}
			got, err := s.Login(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.NotEqual(t, uuid.Nil, got.ID)
			got.ID = uuid.Nil
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type hasherResult struct {
		res string
	}
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name              string
		args              args
		expectedHasherRes hasherResult
		expectedRepoRes   error
		want              *entity.User
		wantErr           bool
	}{
		{
			name:              "success",
			expectedHasherRes: hasherResult{res: "hashed res"},
			args:              args{login: "test", password: "pwd"},
			want:              &entity.User{Login: "test", Password: "hashed res"},
		},
		{
			name:            "fail repo",
			args:            args{login: "test", password: "pwd"},
			expectedRepoRes: errors.New("random error"),
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			hasherMock := mocks.NewMockPasswordHasher(ctrl)
			hasherMock.EXPECT().Hash(gomock.Any()).Return(tt.expectedHasherRes.res, nil)

			repoMock := mocks.NewMockUserRepository(ctrl)
			repoMock.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().Return(tt.expectedRepoRes)

			r := NewUserService(hasherMock, repoMock)

			got, err := r.Create(context.Background(), tt.args.login, tt.args.password)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotEqual(t, uuid.Nil, got.ID)
			got.ID = uuid.Nil
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
