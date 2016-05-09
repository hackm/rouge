package repository

import (
  "testing"
  "../entity"
)

var tags = []string { "go", "goroutine", "rouge" }
var tagsIndex = 0
func addTag(r ITagRepository) *entity.Tag {
  t := r.Add(&entity.Tag {
    Name: tags[tagsIndex],
  })
  tagsIndex = tagsIndex + 1
  return t
}
func resetTagIndex(){
  tIdx = int64(0)
  tagsIndex = 0
}

func TestGetTag_Should_Be_Nil(t *testing.T) {
  r := NewMemTagRepository()
  p := r.Get("go")
  if p != nil {
    t.Errorf("A tag should be nil but %v", p)
  }
}
func TestAddTag_Should_Be_Created(t *testing.T) {
  r := NewMemTagRepository()
  resetTagIndex()
  tag := addTag(r)
  if tag == nil {
    t.Errorf("A tag should be not nil")
  } else if tag.Name != "go" {
    t.Errorf("A tag should be Name == 'go' but %v", tag)    
  }
}
func TestGetTag_Should_Be_Go(t *testing.T) {
  r := NewMemTagRepository()
  resetTagIndex()  
  tag := addTag(r)
  if tag == nil {
    t.Errorf("A tag should be not nil")
  } else if tag.Name != "go" {
    t.Errorf("A tag should be Name == 'go' but %v", tag)
  }
}
func TestGetTags_Should_Be_Length2(t *testing.T) {
  r := NewMemTagRepository()
  resetTagIndex()  
  _ = addTag(r)
  _ = addTag(r)
  _ = addTag(r)
  l := r.GetList("go", 10)
  if len(l) != 2 {
    t.Errorf("Length of paper list should be 2 but %d", len(l))
  }
}