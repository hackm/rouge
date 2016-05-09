package repository

import (
  "time"
  "../entity"
)

type IPaperRepository interface {
  Get(id int64) *entity.Paper
  GetList(skip int, take int) []*entity.Paper
  Add(p *entity.Paper) *entity.Paper
  Update(p *entity.Paper) error
  Remove(id int64) *entity.Paper
}

type MemPaperRepository struct {
  store []entity.Paper
}

func NewMemPaperRepository() *MemPaperRepository {
  return &MemPaperRepository{}
}

func (r *MemPaperRepository) Get(id int64) *entity.Paper {
  for _, p := range r.store {
    if p.Id == id {
      return &p
    }
  }
  return nil
}
func (r *MemPaperRepository) GetList(skip int, take int) []*entity.Paper {
  var l []*entity.Paper
  i := 0
  j := 0
  for _, p := range r.store {
    j = j + 1
    if j >= skip {
      l = append(l, &p)
      i += i + 1
      if i == take {
        break
      }
    }
  }
  return l
}
var pIdx int64 = 0
func (r *MemPaperRepository) Add(p *entity.Paper) *entity.Paper {
  pIdx = pIdx + 1
  p.Id = pIdx
  p.CreatedAt = time.Now()
  p.UpdatedAt = time.Now()
  r.store = append(r.store, *p)
  return p
}
func (r *MemPaperRepository) Update(p *entity.Paper) error {
  sp := r.Get(p.Id)
  sp.ContentId = p.ContentId
  sp.UpdatedAt = time.Now()
  return nil
}
func (r *MemPaperRepository) Remove(id int64) *entity.Paper {
  for i, p := range r.store {
    if p.Id == id {
      r.store = append(r.store[:i],r.store[i+1:]...)
      return &p
    }
  }
  return nil
}