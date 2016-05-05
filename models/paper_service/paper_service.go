package paper_service

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

func (ps *PaperService) GetPaper(id int64) paper.Paper {
	p := ps.PaperRepository.GetPaper(id)
	_, cc := p["content"]
	_, uc := p["user"]
	if !cc {
		p["content"] = ps.ContentRepository.GetContent(p["contentId"]);
	}
	if !uc {
		p["user"] = ps.UserRepository.GetUser(p["userId"]);
	}
	return p
}

func (ps *PaperService) GetUserPapers(id int64) paper.Paper {
	pl := ps.PaperRepository.GetUserPapers(id);
	for p := range pl {
		_, cc := p["content"]
		_, uc := p["user"]
		if !cc {
			p["content"] = ps.ContentRepository.GetContent(p["contentId"]);
		}
		if !uc {
			p["user"] = ps.UserRepository.GetUser(p["userId"]);
		}
	}

	return pl
}

func (ps *PaperService) GetTagPapers(tag string) paper.Paper {
	pl := ps.PaperRepository.GetTagPapers(tag);
	for p := range pl {
		_, cc := p["content"]
		_, uc := p["user"]
		if !cc {
			p["content"] = ps.ContentRepository.GetContent(p["contentId"]);
		}
		if !uc {
			p["user"] = ps.UserRepository.GetUser(p["userId"]);
		}
	}
	return pl
}

func (ps *PaperService) GetPapers(keyword string) paper.Paper {
	pl := ps.PaperRepository.GetPapers(keyword);
	for p := range pl {
		_, cc := p["content"]
		_, uc := p["user"]
		if !cc {
			p["content"] = ps.ContentRepository.GetContent(p["contentId"]);
		}
		if !uc {
			p["user"] = ps.UserRepository.GetUser(p["userId"]);
		}
	}
	return pl
}
