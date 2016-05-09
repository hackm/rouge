package repository

import (
  "testing"
  "../entity"
)

func addCotent(r IContentRepository) *entity.Content{
  return r.Add(&entity.Content {
    PaperId:1,
    Title: "title",
    Body: "body",
  })
}
func resetContentIndex(){
  cIdx = int64(0)
}

func TestGetContent_Should_Be_Nil(t *testing.T) {
  r := NewMemContentRepository()
  c := r.Get(int64(1000))
  if c != nil {
    t.Errorf("A content should be nil but %v", c)
  }
}
func TestAddContent_Should_Be_Created(t *testing.T) {
  r := NewMemContentRepository()
  resetContentIndex()
  c := addCotent(r)
  if c == nil {
    t.Errorf("A content should be not nil")
  } else if c.Id != 1 {
    t.Errorf("A content should be Id == 1 but %v", c)    
  }
}
func TestGetContent_Should_Be_Id1(t *testing.T) {
  r := NewMemContentRepository()
  resetContentIndex()  
  c := addCotent(r)
  if c == nil {
    t.Errorf("A content should be not nil")
  } else if c.Id != 1 {
    t.Errorf("A content should be Id == 1 but %v", c)
  }
}
func TestGetContents_Should_Be_Length2(t *testing.T) {
  r := NewMemContentRepository()
  resetContentIndex()  
  _ = addCotent(r)
  _ = addCotent(r)
  l := r.GetList(int64(1))
  if len(l) != 2 {
    t.Errorf("Length of content list should be 2 but %d", len(l))
  }
}
func TestRemoveContent_Should_Be_Deleted(t *testing.T) {
  r := NewMemContentRepository()
  resetContentIndex()  
  _ = addCotent(r)
  c := r.Remove(int64(1))
  l := r.GetList(int64(1))
  if c == nil {
    t.Errorf("A content should be not nil")    
  } else if c.Id != 1 {
    t.Errorf("A content should be Id == 1 but %v", c)
  } else if len(l) != 0 {
    t.Errorf("Length of content list should be 0 but %d", len(l))    
  }
}
func TestAddContent_when_over10_of_a_paper_Should_Remove_Older(t *testing.T) {
  r := NewMemContentRepository()
  resetContentIndex()  
  _ = addCotent(r);_ = addCotent(r);_ = addCotent(r);_ = addCotent(r);_ = addCotent(r);
  _ = addCotent(r);_ = addCotent(r);_ = addCotent(r);_ = addCotent(r);_ = addCotent(r);
  _ = addCotent(r)
  l := r.GetList(int64(1))
  if len(l) != 10 {
    t.Errorf("Length of content list should be 10 but %d", len(l))    
  }
}