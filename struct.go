package GoStruct

type People struct {
	ID      int64
	Name    string
	Surname string
}

func (p *People) SetName(name string) {
	p.Name = name
}

func (p *People) SetSurname(Surname string) {
	p.Surname = Surname
}

func (p *People) SetID(id int64) {
	p.ID = id
}

func (p *People) GetName() string {
	return p.Name
}

func (p *People) GetSurname() string {
	return p.Surname
}
func (p *People) GetID() int64 {
	return p.ID
}
