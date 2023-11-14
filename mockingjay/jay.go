package mockingjay

type RealmService interface {
	RandNumb() int
}

type Realm struct {
	service RealmService
}

func NewNumber(service RealmService) *Realm {
	return &Realm{service: service}
}

func (r *Realm) AddNumb() int {
	firstInt := r.service.RandNumb()
	secondInt := 9
	return firstInt + secondInt
}
