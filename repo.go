package main

type UserRepository interface {
	Get(id string) (User, error)
}

