package paper

type Paper map[string]interface{}

type IPaperRepository interface {
	GetPaper(id int64) Paper
	GetUserPapers(userId int64) []Paper
	GetTagPapers(tag string) []Paper
	GetPapers(keyword string) []Paper

	CreatePaper(p Paper) (Paper, error)
	UpdatePaper(p Paper) error
	DeletePaper(id int64) error
}

type PaperRepository struct{
	store []Paper	
}

func NewPaperRepository() *PaperRepository {
	return &PaperRepository{
		store: []Paper{
			Paper{
				"id": int64(1000),
				"userId": int64(1000),
				"contentId": int64(2000),
			},
		},
	}
}

func (r *PaperRepository) GetPaper(id int64) Paper {
	var sp Paper
	for _, v := range r.store {
		if v["id"] == id {
			sp = v
			break
		}
	}
	return sp
}

func (r *PaperRepository) GetUserPapers(userId int64) []Paper {
	return r.store
}

func (r *PaperRepository) GetTagPapers(tag string) []Paper {
	return r.store
}

func (r *PaperRepository) GetPapers(keyword string) []Paper {
	return r.store
}

var index = int64(1000)
func (r *PaperRepository) CreatePaper(p Paper) (Paper, error) {
	index = index + 1
	p["id"] = index
	r.store = append(r.store, p)
	return p, nil
}

func (r *PaperRepository) UpdatePaper(p Paper) error {
	var sp Paper
	for _, v := range r.store {
		if v["id"] == p["id"] {
			sp = v
			break
		}
	}
	if sp == nil{
		// return error
	}
	sp["id"] = p["id"]
	sp["contentId"] = p["contentId"]
	return nil
}

func (r *PaperRepository) DeletePaper(id int64) error {
	idx := int64(-1)
	for i, v := range r.store {
		if v["id"] == id {
			idx = int64(i)
			break
		}
	}
	if idx == -1 {
		// return error
	}
	r.store = append(r.store[:idx], r.store[idx+1:]...)
	return nil
}
