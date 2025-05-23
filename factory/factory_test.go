package factory

import "testing"

func TestCat_Sound(t *testing.T) {
	animal := AnimalFactory("cat")

	if animal.Sound() != "Meow!" {
		t.Error("Cat sound error")
	}
}

func TestDog_Sound(t *testing.T) {
	animal := AnimalFactory("dog")

	if animal.Sound() != "Wang!" {
		t.Error("Dog sound error")
	}
}
