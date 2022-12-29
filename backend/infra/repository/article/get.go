package article

import (
	"backend/domain/model"
)

func (r *articleRepository) ListArticles() (articles []*model.Article, err error) {
	err = r.DB.Order("created_at desc").Preload("Tags").Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepository) GetArticle(id string) (article *model.Article, err error) {
	err = r.DB.Where("id = ?", id).Preload("Tags").Find(&article).Error

	if err != nil {
		return nil, err
	}

	return article, nil
}
