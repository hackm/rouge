package paper

import (
	"testing"
)

func TestGetPaper(t *testing.T) {
  r := NewPaperRepository()
  p := r.GetPaper(1000)
  if p == nil {
    t.Errorf("nil from GetPaper")
  }
}

func TestGetPapers(t *testing.T) {
  r := NewPaperRepository()
  pl := r.GetPapers("")
  if len(pl) == 0 {
    t.Errorf("length 0 from GetPapers ")
  }
}

func TestCreatePapers(t *testing.T) {
  r := NewPaperRepository()
  r.CreatePaper(Paper{
    "userId": 1000,
    "contentId": 3000,
  })
  pl := r.GetPapers("")
  if len(pl) != 2 {
    t.Errorf("lack of length: %v", pl)
  }
}
func TestUpdatePapers(t *testing.T) {
  r := NewPaperRepository()
  p := r.GetPaper(1000)
  p["contentId"] = 10000
  r.UpdatePaper(p)
  p = r.GetPaper(1000)
  if p["contentId"] != 10000 {
    t.Errorf("not equal: %v", p)
  }
}
func TestDeletePapers(t *testing.T) {
  r := NewPaperRepository()
  r.DeletePaper(1000)
  pl := r.GetPapers("")
  if len(pl) != 0 {
    t.Errorf("%v", pl)
  }
}
