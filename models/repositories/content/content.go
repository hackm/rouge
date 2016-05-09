package content

type Content map[string]interface{}

type IContentRepository interface {
  GetContent(id int64) Content
  GetPaperContents(paperId int64) []Content
  CreateContent(c Content) (Content, error)
}

type ContentRepository struct {
  store []Content
}

func NewContentRepository() *ContentRepository {
  return &ContentRepository{
    store: []Content {
      Content {
        "id": int64(2000),
        "title": "Sample",
        "body":`
          <p>this is body</p>
          <ul>
            <li>Hello</li>
            <li>World</li>
          </ul>`,
      },
    },
  }
}

func (r *ContentRepository) GetContent(id int64) Content {
  var c Content
  for _, v := range r.store {
		if v["id"] == id {
			c = v
			break
		}
	}
  return c
}

func (r *ContentRepository) GetPaperContents(paperId int64) []Content {
  cl := []Content{}
  for _, v := range r.store {
		if v["paperId"] == paperId {
			cl = append(cl, v)
		}
	}
  return cl
}

var index = int64(2000)
func (r *ContentRepository) CreateContent(c Content) (Content, error) {
  index = index + 1
  c["id"] = index
  r.store = append(r.store, c)
  return c, nil
}
