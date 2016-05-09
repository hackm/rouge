package service

import (
  "../entity"
  repo "../repository"
)

const SERVICE string = "Service:paper@hackm"
func Key() string {
  return SERVICE
}

type Service struct {
  contents repo.IContentRepository
  papers repo.IPaperRepository
  tags repo.ITagRepository
}

func NewService() *Service {
  return  &Service{
    contents: repo.NewMemContentRepository(),
    papers: repo.NewMemPaperRepository(),
  }
}

func (s *Service) CreatePaper(p *entity.Paper) *entity.Paper {
  c := s.contents.Add(p.Content)
  p.ContentId = c.Id
  return s.papers.Add(p)
}
func (s *Service) UpdatePaper(p *entity.Paper) error {
  c := s.contents.Add(p.Content)
  p.ContentId = c.Id
  s.papers.Update(p)
  return nil
}
func (s *Service) DeletePaper(id int64) *entity.Paper {
  return s.papers.Remove(id)
}
func (s *Service) GetPaper(id int64) *entity.Paper {
  return s.papers.Get(id)
}
func (s *Service) GetPapers(skip int, take int) []*entity.Paper {
  return s.papers.GetList(skip, take)
}
func (s *Service) GetPaperContents(paperId int64) []*entity.Content {
  return s.contents.GetList(paperId)
}
