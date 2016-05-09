package services

import (
	"../repositories/user"
	"../repositories/paper"
	"../repositories/content"
)

type PaperService struct {
	UserRepository    user.IUserRepository
	PaperRepository   paper.IPaperRepository
	ContentRepository content.IContentRepository
}

func NewPaperService() *PaperService {
	return &PaperService{
		UserRepository: user.NewUserRepository(),
		PaperRepository: paper.NewPaperRepository(),
		ContentRepository: content.NewContentRepository(),
	}
}

// func (p Paper) loadLazy() {
//   _, cc := p["content"]
// 	_, uc := p["user"]
// }

func (ps *PaperService) GetPaper(id int64) paper.Paper {
	p := ps.PaperRepository.GetPaper(id)
	_, cc := p["content"]
	_, uc := p["user"]
	if !cc {
		p["content"] = ps.ContentRepository.GetContent(p["contentId"].(int64));
	}
	if !uc {
		p["user"] = ps.UserRepository.GetUser(p["userId"].(int64));
	}
	return p
}

func (ps *PaperService) GetUserPapers(id int64) []paper.Paper {
	pl := ps.PaperRepository.GetUserPapers(id);
	for _, p := range pl {
		_, cc := p["content"]
		_, uc := p["user"]
		if !cc {
      p["content"] = ps.ContentRepository.GetContent(p["contentId"].(int64));
    }
    if !uc {
      p["user"] = ps.UserRepository.GetUser(p["userId"].(int64));
    }
	}

	return pl
}

func (ps *PaperService) GetTagPapers(tag string) []paper.Paper {
	pl := ps.PaperRepository.GetTagPapers(tag);
	for _, p := range pl {
		_, cc := p["content"]
		_, uc := p["user"]
		if !cc {
      p["content"] = ps.ContentRepository.GetContent(p["contentId"].(int64));
    }
    if !uc {
      p["user"] = ps.UserRepository.GetUser(p["userId"].(int64));
    }
	}
	return pl
}

func (ps *PaperService) GetPapers(keyword string) []paper.Paper {
	pl := ps.PaperRepository.GetPapers(keyword);
	for _, p := range pl {
		_, cc := p["content"]
		_, uc := p["user"]
		if !cc {
      p["content"] = ps.ContentRepository.GetContent(p["contentId"].(int64));
    }
    if !uc {
      p["user"] = ps.UserRepository.GetUser(p["userId"].(int64));
    }
	}
	return pl
}
