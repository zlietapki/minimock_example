package book_minimock

import "github.com/sirupsen/logrus" // для демонстрации мока внешних модулей

//go:generate minimock -g -i Repo # внутренний интерфейс
//go:generate minimock -g -i github.com/sirupsen/logrus.Formatter # внешний интерфейс

type Repo interface {
	Count(author string) int
}

type Book struct {
	repo      Repo
	formatter logrus.Formatter
}

func New(repo Repo, formatter logrus.Formatter) *Book {
	return &Book{
		repo:      repo,
		formatter: formatter,
	}
}

func (o *Book) CountByAuthor(author string) int {
	o.formatter.Format(nil)
	return o.repo.Count(author)
}
