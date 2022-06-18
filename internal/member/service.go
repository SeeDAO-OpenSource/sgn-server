package member

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberService struct {
	repo MemberRepository
}

func NewMemberService(mongoClient *mongo.Client) (*MemberService, error) {
	repo := NewMongoMemberRepository(mongoClient)
	srv := &MemberService{
		repo: repo,
	}
	return srv, nil
}

func (s MemberService) GetList(page int, pageSize int) ([]Member, error) {
	return s.repo.GetList(page, pageSize)
}

func (s MemberService) GetByAddress(address string) (Member, error) {
	return s.repo.GetByAddress(address)
}

func (s MemberService) GetByAddresses(addresses []string) ([]Member, error) {
	return s.repo.GetByAddresses(addresses)
}

func (s MemberService) Insert(member *Member) error {
	return s.repo.Insert(member)
}

func (s MemberService) InsertManay(members []Member) error {
	return s.repo.InsertManay(members)
}

func (s MemberService) Update(member *Member) error {
	return s.repo.Update(member)
}

func (s MemberService) Delete(address string) error {
	return s.repo.Delete(address)
}
