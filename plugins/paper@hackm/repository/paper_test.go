package repository

import (
  "testing"
  "../entity"
)

func addPaper(r IPaperRepository) *entity.Paper {
  return r.Add(&entity.Paper{})
}
func resetPaperIndex(){
  pIdx = int64(0)
}

func TestGetPaper_Should_Be_Nil(t *testing.T) {
  r := NewMemPaperRepository()
  p := r.Get(int64(1000))
  if p != nil {
    t.Errorf("A paper should be nil but %v", p)
  }
}
func TestAddPaper_Should_Be_Created(t *testing.T) {
  r := NewMemPaperRepository()
  resetPaperIndex()
  p := addPaper(r)
  if p == nil {
    t.Errorf("A paper should be not nil")
  } else if p.Id != 1 {
    t.Errorf("A paper should be Id == 1 but %v", p)    
  }
}
func TestGetPaper_Should_Be_Id1(t *testing.T) {
  r := NewMemPaperRepository()
  resetPaperIndex()  
  p := addPaper(r)
  if p == nil {
    t.Errorf("A paper should be not nil")
  } else if p.Id != 1 {
    t.Errorf("A paper should be Id == 1 but %v", p)
  }
}
func TestGetPapers_Should_Be_Length2(t *testing.T) {
  r := NewMemPaperRepository()
  resetPaperIndex()  
  _ = addPaper(r)
  _ = addPaper(r)
  l := r.GetList(0, 10)
  if len(l) != 2 {
    t.Errorf("Length of paper list should be 2 but %d", len(l))
  }
}
func TestRemovePaper_Should_Be_Deleted(t *testing.T) {
  r := NewMemPaperRepository()
  resetPaperIndex()  
  _ = addPaper(r)
  p := r.Remove(int64(1))
  l := r.GetList(0, 10)
  if p == nil {
    t.Errorf("A paper should be not nil")    
  } else if p.Id != 1 {
    t.Errorf("A paper should be Id == 1 but %v", p)
  } else if len(l) != 0 {
    t.Errorf("Length of paper list should be 0 but %d", len(l))    
  }
}
