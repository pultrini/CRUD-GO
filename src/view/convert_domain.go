package view

import (
	"github.com/pultrini/CRUD-GO/src/controller/model/response"
	"github.com/pultrini/CRUD-GO/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse{
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}