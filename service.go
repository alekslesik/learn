package main

type userAgeService struct {
	userRepo UserRepository
}

func NewUserAgeService(userRepo UserRepository) userAgeService {
	return userAgeService{userRepo}
}

func (r *userAgeService) CanBuyAlcohol(id string) (bool, error) {

	user, err := r.userRepo.Get(id)
	if err != nil {
		return false, err
	}

	if user.age < 21 {
		return false, nil
	}

	return true, nil
}
