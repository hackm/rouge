package paper

type Paper map[string]interface{}

type IPaperRepository interface {
  GetPaper(id int64) Paper
  GetUserPapers(userId int64) []Paper
  GetTagPapaers(tag string) []Paper
  GetPapaers(keyword string) []Paper
  
  CreatePaper(p Paper) error
  UpdatePaper(p Paper) error
  DeletePaper(id int64) error
}

type PaperRepository struct {}

func NewPaperRepository *PaperRepository{
  return &PaperRepository{}
}

func (r *PaperRepository) GetPaper(id int64) Paper {
  p := make(Paper)
  p["id"] = id
  p["userId"] = 1000
  p["contentId"] = 2000
  
  return p
}
func (r *PaperRepository) GetUserPapers(userId int64) []Paper {
  p1 := make(Paper)
  p1["id"] = 1000
  p1["userId"] = 1000
  p1["contentId"] = 2000
  
  p2 := make(Paper)
  p2["id"] = 1001
  p2["userId"] = 1000
  p2["contentId"] = 2001
  
  return []Paper{p1, p2}
}
func (r *PaperRepository) GetTagPapaers(tag string) []Paper {
  p1 := make(Paper)
  p1["id"] = 1000
  p1["userId"] = 1000
  p1["contentId"] = 2000
  
  p2 := make(Paper)
  p2["id"] = 1001
  p2["userId"] = 1000
  p2["contentId"] = 2001
  
  return []Paper{p1, p2}
}
func (r *PaperRepository) GetPapaers(keyword string) []Paper {
  p1 := make(Paper)
  p1["id"] = 1000
  p1["userId"] = 1000
  p1["contentId"] = 2000
  
  p2 := make(Paper)
  p2["id"] = 1001
  p2["userId"] = 1000
  p2["contentId"] = 2001
  
  return []Paper{p1, p2}
}

func (r *PaperRepository) CreatePaper(p Paper) error {
}
func (r *PaperRepository) UpdatePaper(p Paper) error {
}
func (r *PaperRepository) DeletePaper(p Paper) error {
}