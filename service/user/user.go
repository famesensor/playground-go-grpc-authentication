package user

type UserService interface {
	// CreateUser(context.Context, *model.CreateUserReq) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}
