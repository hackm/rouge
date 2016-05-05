package paper

type Paper map[string]interface{}

type IPaperRepository interface {
	GetPaper(id int64) Paper
	GetUserPapers(userId int64) []Paper
	GetTagPapers(tag string) []Paper
	GetPapers(keyword string) []Paper

	CreatePaper(p Paper) error
	UpdatePaper(p Paper) error
	DeletePaper(id int64) error
}

type PaperRepository struct{}

func NewPaperRepository() PaperRepository {
	return PaperRepository{}
}

func (r PaperRepository) GetPaper(id int64) Paper {
	p := make(Paper)
	p["id"] = id
	p["userId"] = int64(1000)
	p["contentId"] = int64(2000)

	return p
}
func (r PaperRepository) GetUserPapers(userId int64) []Paper {
	p1 := make(Paper)
	p1["id"] = int64(1000)
	p1["userId"] = int64(1000)
	p1["contentId"] = int64(2000)

	p2 := make(Paper)
	p2["id"] = int64(1001)
	p2["userId"] = int64(1000)
	p2["contentId"] = int64(2001)

	return []Paper{p1, p2}
}
func (r PaperRepository) GetTagPapers(tag string) []Paper {
	p1 := make(Paper)
	p1["id"] = int64(1000)
	p1["userId"] = int64(1000)
	p1["contentId"] = int64(2000)

	p2 := make(Paper)
	p2["id"] = int64(1001)
	p2["userId"] = int64(1000)
	p2["contentId"] = int64(2001)

	return []Paper{p1, p2}
}
func (r PaperRepository) GetPapers(keyword string) []Paper {
	p1 := make(Paper)
	p1["id"] = int64(1000)
	p1["userId"] = int64(1000)
	p1["contentId"] = int64(2000)

	p2 := make(Paper)
	p2["id"] = int64(1001)
	p2["userId"] = int64(1000)
	p2["contentId"] = int64(2001)

	return []Paper{p1, p2}
}

func (r PaperRepository) CreatePaper(p Paper) error {
	return nil
}
func (r PaperRepository) UpdatePaper(p Paper) error {
	return nil
}
func (r PaperRepository) DeletePaper(id int64) error {
	return nil
}
