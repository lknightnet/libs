package genToken

type ServiceToken struct {
	API string
}

func NewServiceToken(api string) *ServiceToken {
	return &ServiceToken{
		API: api,
	}
}
