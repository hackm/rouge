package content

type Content map[string]interface{}

type IContentRepository interface {
  GetContent(id int64) Content
}

type ContentRepository struct {}

func NewContentRepository() *ContentRepository {
  return &ContentRepository{}
}

func (r *ContentRepository) GetContent(id int64) Content {
  c := make(Content)
  c["id"] = id
  c["title"] = "Sample"
  c["body"] = `
    <p>this is body</p>
    <ul>
      <li>Hello</li>
      <li>World</li>
    </ul>`
  return c
}
