package constant

type contextKey string

const (
	UserCtxKey contextKey = "user"
)

var (
	SkipPath = map[string]struct{}{
		"/proto.AuthService/SignIn": {},
		"/proto.AuthService/SignUp": {},
	}
)
