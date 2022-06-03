package book_minimock

import (
	"testing"

	"github.com/gojuno/minimock/v3"
)

// Игнорировать параметры, просто возвращать что сказано
func TestBook_CountByAuthor_ignore_params(t *testing.T) {
	// init
	mc := minimock.NewController(t)
	defer mc.Finish() // нужно для проверки "все ли моки были вызваны хоть раз"

	// мокнутые объекты
	bookRepo := NewRepoMock(mc)
	logrusMock := NewFormatterMock(mc)
	b := New(bookRepo, logrusMock)

	// set mocks
	logrusMock.FormatMock.
		Return(nil, nil)
	bookRepo.CountMock.
		Return(222) // безразличны параметры, всегда возвращать 222

	// test
	res := b.CountByAuthor("")
	if res != 222 {
		t.Fail()
	}
}

// Ожидать определенных параметров
func TestBook_CountByAuthor_expect_params(t *testing.T) {
	// init
	mc := minimock.NewController(t)
	defer mc.Finish()

	// мокнутые объекты
	bookRepo := NewRepoMock(mc)
	logrusMock := NewFormatterMock(mc)
	b := New(bookRepo, logrusMock)

	// set mocks
	logrusMock.FormatMock.
		Return(nil, nil)
	bookRepo.CountMock.
		Expect("alpha"). // ожидаемые статичные параметры
		Return(222)

	// test
	res := b.CountByAuthor("alpha")
	if res != 222 {
		t.Fail()
	}
}

// When/Then альтернатива Expect/Return
func TestBook_CountByAuthor_when_then(t *testing.T) {
	// init
	mc := minimock.NewController(t)
	defer mc.Finish()

	// мокнутые объекты
	bookRepo := NewRepoMock(mc)
	logrusMock := NewFormatterMock(mc)
	b := New(bookRepo, logrusMock)

	// set mocks
	logrusMock.FormatMock.
		Return(nil, nil)

	bookRepo.CountMock.
		When("alpha").
		Then(111)

	bookRepo.CountMock.
		When("beta").
		Then(222)

	// tests
	res1 := b.CountByAuthor("alpha")
	if res1 != 111 {
		t.Fail()
	}

	res2 := b.CountByAuthor("beta")
	if res2 != 222 {
		t.Fail()
	}
}

// Вычислять ответ из параметров и (возможно) проверять сами параметры
func TestBook_CountByAuthor_inspect_params(t *testing.T) {
	// init
	mc := minimock.NewController(t)
	defer mc.Finish()

	// мокнутые объекты
	bookRepo := NewRepoMock(mc)
	logrusMock := NewFormatterMock(mc)
	b := New(bookRepo, logrusMock)

	// set mocks
	logrusMock.FormatMock.
		Return(nil, nil)

	bookRepo.CountMock.Set(func(author string) int {
		if author == "alpha" { // вычислять ответ на основе параметров
			return 111
		}
		if author == "beta" {
			return 222
		}

		t.Fail() // ожидается, что сюда выполнение не дойдет
		return 0
	})

	// tests
	res1 := b.CountByAuthor("alpha")
	if res1 != 111 {
		t.Fail()
	}

	res2 := b.CountByAuthor("beta")
	if res2 != 222 {
		t.Fail()
	}
}
