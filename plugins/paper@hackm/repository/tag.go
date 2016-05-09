package repository

import (
  "strings"
  "../entity"
)

type ITagRepository interface {
  Get(name string) *entity.Tag
  GetList(startsWith string, take int) []*entity.Tag
  Add(t *entity.Tag) *entity.Tag
}

type MemTagRepository struct {
  store []entity.Tag
}

func NewMemTagRepository() *MemTagRepository {
  return &MemTagRepository{}
}

func (r *MemTagRepository) Get(name string) *entity.Tag {
  for _, t := range r.store {
    if t.Name == name {
      return &t
    }
  }
  return nil
}
func (r *MemTagRepository) GetList(startsWith string, take int) []*entity.Tag {
  var l []*entity.Tag
  i := 0
  for _, t := range r.store {
    if strings.HasPrefix(t.Name, startsWith) {
      l = append(l, &t)
      i += i + 1
      if i == take {
        break
      }
    }
  }
  return l
}
var tIdx int64 = 0
func (r *MemTagRepository) Add(t *entity.Tag) *entity.Tag {
  tIdx = tIdx + 1
  r.store = append(r.store, *t)
  return t
}
