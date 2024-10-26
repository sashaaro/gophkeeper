package auth_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sashaaro/gophkeeper/internal/auth"
	"github.com/sashaaro/gophkeeper/internal/entity"
)

func TestRegisterer_Register(t *testing.T) {
	type hasherResult struct {
		res string
		err error
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
			name:              "fail password hash",
			expectedHasherRes: hasherResult{err: errors.New("random error")},
			args:              args{login: "test", password: "pwd"},
			wantErr:           true,
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

			hasherMock := NewMockPasswordHasher(ctrl)
			hasherMock.EXPECT().Hash(gomock.Any()).Return(tt.expectedHasherRes.res, tt.expectedHasherRes.err)

			repoMock := NewMockUserCreator(ctrl)
			repoMock.EXPECT().CreateUser(gomock.Any()).AnyTimes().Return(tt.expectedRepoRes)

			r := auth.NewRegisterer(hasherMock, repoMock)

			got, err := r.Register(tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
