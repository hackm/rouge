package repository

import (
  "time"
  "../entity"
)

type IContentRepository interface {
  Get(id int64) *entity.Content
  GetList(paperId int64) []*entity.Content
  Add(c *entity.Content) *entity.Content
  Remove(id int64) *entity.Content
}

type MemContentRepository struct {
  store []entity.Content
}

func NewMemContentRepository() *MemContentRepository {
  return &MemContentRepository{}
}

func (r *MemContentRepository) Get(id int64) *entity.Content {
  for _, c := range r.store {
    if c.Id == id {
      return &c
    }
  }
  return nil
}

func (r *MemContentRepository) GetList(paperId int64) []*entity.Content {
  var l []*entity.Content
  for _, c := range r.store {
    if c.PaperId == paperId {
      l = append(l, &c)
    }
  }
  return l
}

var cIdx int64 = 0
func (r *MemContentRepository) Add(c *entity.Content) *entity.Content {
  cIdx = cIdx + 1
  c.Id = cIdx
  c.CreatedAt = time.Now()
  c.UpdatedAt = time.Now()
  r.store = append(r.store, *c)
  l := r.GetList(c.PaperId)
  for len(l) > 10 {
    r.Remove(l[0].Id)
    l = l[1:]
  }
  return c
}

func (r *MemContentRepository) Remove(id int64) *entity.Content {
  for i, c := range r.store {
    if c.Id == id {
      r.store = append(r.store[:i],r.store[i+1:]...)
      return &c
    }
  }
  return nil
} 