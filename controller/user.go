package controller

import (
	"github.com/biges/hybrone-sentinel-backend/exchanges"
	"github.com/biges/hybrone-sentinel-backend/models"
)

// UserCreate is a service method of UserService interfaces.
func (s Service) UserCreate(req exchanges.UserCreateRequest) (exchanges.UserCreateResponse, error) {
res := exchanges.UserCreateResponse{}
err := s.UserManager.Create(req.User)
if err != nil {
	res.Message = "Oluşturulamadı."
	return res,err
}
res.User = req.User
return res,err
}

// UserUpdate is a service method of UserService interfaces.
func (s Service) UserUpdate(req exchanges.UserUpdateRequest) (exchanges.UserUpdateResponse, error) {
res := exchanges.UserUpdateResponse{}
err := s.UserManager.Update(req.User)
if err != nil {
	res.Message = "Oluşturulamadı."
	return res,err
}
res.User = req.User
return res,err
}

// UserDelete is a service method of UserService interfaces.
func (s Service) UserDelete(req exchanges.UserDeleteRequest) (exchanges.UserDeleteResponse, error) {
res := exchanges.UserDeleteResponse{}
err := s.UserManager.Delete(req.ID)
if err != nil {
	res.Message = "Silinemedi."
	return res,err
}
res.ID = req.ID
return res,err
}

// UserGet is a service method of UserService interfaces.
func (s Service) UserGet(req exchanges.UserGetRequest) (exchanges.UserGetResponse, error) {
res := exchanges.UserGetResponse{}
err := s.UserManager.Get(req.User)
if err != nil {
	res.Message = "Bulunamadı."
	return res,err
res.User = req.User
return res,err
}

// UserList is a service method of UserService interfaces.
func (s Service) UserList(req exchanges.UserListRequest) (exchanges.UserListResponse, error) {
res := exchanges.UserListResponse{}
users := new([]models.User)
err := s.UserManager.List(users)
if err != nil {
	res.Message = "Bulunamadı."
	return res,err
}
res.Users = users
return res,err
}