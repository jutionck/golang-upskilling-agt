package manager

import "github.com/jutionck/golang-upskilling-agt/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

// UserUseCase implements UseCaseManager.
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
