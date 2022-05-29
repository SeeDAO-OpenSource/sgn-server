package identity

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/common"
	"github.com/SeeDAO-OpenSource/sgn/pkg/db/mongodb"
)

type IdentityService struct {
	repo MemberRepository
}

func NewIdentityService() (*IdentityService, error) {
	mongoClient, err := mongodb.GetClient(common.MongoOptions)
	if err != nil {
		return nil, err
	}
	repo := NewMongoMemberRepository(mongoClient)
	srv := &IdentityService{
		repo: repo,
	}
	return srv, nil
}

func (s IdentityService) GetList(page int, pageSize int) ([]Member, error) {
	return s.repo.GetList(page, pageSize)
}

func (s IdentityService) GetByAddress(address string) (Member, error) {
	return s.repo.GetByAddress(address)
}

func (s IdentityService) GetByAddresses(addresses []string) ([]Member, error) {
	return s.repo.GetByAddresses(addresses)
}

func (s IdentityService) Insert(member *Member) error {
	return s.repo.Insert(member)
}

func (s IdentityService) InsertManay(members []Member) error {
	return s.repo.InsertManay(members)
}

func (s IdentityService) Update(member *Member) error {
	return s.repo.Update(member)
}

func (s IdentityService) Delete(address string) error {
	return s.repo.Delete(address)
}
