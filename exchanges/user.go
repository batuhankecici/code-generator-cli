package exchanges

import "github.com/biges/hybrone-sentinel-backend/models"

type UserService interface {
	UserCreate(UserCreateRequest, UserCreateResponse) error
	UserUpdate(UserUpdateRequest, UserUpdateResponse) error
	UserDelete(UserDeleteRequest, UserDeleteResponse) error
	UserGet(UserGetRequest, UserGetResponse) error
	UserList(UserListRequest, UserListResponse) error
}

// UserCreateRequest holds data for creating new User.
type UserCreateRequest struct {
	User *models.User `json:"user" validate:"required"`
}

// UserCreateResponse holds newly created User instance data.
type UserCreateResponse struct {
	BaseResponse
	User *models.User `json:"user"`
}

// UserUpdateRequest holds data for updating a User.
type UserUpdateRequest struct {
	User *models.User `json:"user" validate:"required"`
}

// UserUpdateResponse holds newly updated User instance data.
type UserUpdateResponse struct {
	BaseResponse
	User *models.User `json:"user"`
}

// UserDeleteRequest holds data for deleting new User.
type UserDeleteRequest struct {
	ID uint `json:"id"`
}

// UserDeleteResponse holds newly deleted User query error.
type UserDeleteResponse struct {
	BaseResponse
	ID uint `json:"id"`
}

// UserGetRequest holds data for get User.
type UserGetRequest struct {
	User *models.User `json:"user" validate:"required"`
}

// UserGetResponse holds User instance data.
type UserGetResponse struct {
	BaseResponse
	User *models.User `json:"user"`
}

// UserListRequest holds filters for User list.
type UserListRequest struct {
	// filters
}

// UserListResponse holds Users data.
type UserListResponse struct {
	BaseResponse
	Users *[]models.User `json:"users"`
}