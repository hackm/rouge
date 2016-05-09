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

func (ps *PaperService) loadLazy(p paper.Paper) {
	const CONTENT string = "content"
	const CONTENT_ID string = "contentId"
	const USER string = "user"
	const USER_ID string = "userId"
	if p != nil {
		_, cc := p[CONTENT]
		_, uc := p[USER]
		if !cc {
			p[CONTENT] = ps.ContentRepository.GetContent(p[CONTENT_ID].(int64))
		}
		if !uc {
			p[USER] = ps.UserRepository.GetUser(p[USER_ID].(int64))
		}
	}
}

func (ps *PaperService) GetPaper(id int64) paper.Paper {
	p := ps.PaperRepository.GetPaper(id)
	ps.loadLazy(p)
	return p
}

func (ps *PaperService) GetUserPapers(id int64) []paper.Paper {
	pl := ps.PaperRepository.GetUserPapers(id);
	for _, p := range pl {
		ps.loadLazy(p)
	}

	return pl
}

func (ps *PaperService) GetTagPapers(tag string) []paper.Paper {
	pl := ps.PaperRepository.GetTagPapers(tag);
	for _, p := range pl {
		ps.loadLazy(p)
	}
	return pl
}

func (ps *PaperService) GetPapers(keyword string) []paper.Paper {
	pl := ps.PaperRepository.GetPapers(keyword);
	for _, p := range pl {
		ps.loadLazy(p)
	}
	return pl
}
