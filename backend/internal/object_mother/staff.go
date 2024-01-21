package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
	"time"
)

type StaffMother struct{}

type staffFunc func(p *model.Staff)

func StaffWithID(id int) staffFunc {
	return func(p *model.Staff) {
		p.Id = id
	}
}

func StaffWithName(name string) staffFunc {
	return func(p *model.Staff) {
		p.Name = name
	}
}

func StaffWithBirthday(t time.Time) staffFunc {
	return func(p *model.Staff) {
		p.Birthday = t
	}
}

func StaffWithType(typeStaff string) staffFunc {
	return func(p *model.Staff) {
		p.Type = typeStaff
	}
}

func StaffWithGender(gender string) staffFunc {
	return func(p *model.Staff) {
		p.Gender = gender
	}
}

func StaffWithPhoto(photo []model.Picture) staffFunc {
	return func(p *model.Staff) {
		p.Photo = photo
	}
}

func (s StaffMother) GenerateStaff(opts ...staffFunc) *model.Staff {
	p := &model.Staff{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (s StaffMother) GenerateRandomStaff() *model.Staff {
	gender := []string{"м", "ж"}
	return &model.Staff{
		Id:       rand.Int(),
		Name:     randStringBytes(8),
		Birthday: time.Now(),
		Type:     randStringBytes(8),
		Gender:   gender[rand.Intn(len(gender))],
		Photo:    nil,
	}
}

func (s StaffMother) GenerateRandomStaffSlice(size int) []model.Staff {
	r := make([]model.Staff, 0)
	for i := 0; i < size; i++ {
		r = append(r, *s.GenerateRandomStaff())
	}
	return r
}
