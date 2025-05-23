package factory

type Animal interface {
	Sound() string
}

type Dog struct{}

type Cat struct{}

func (dog *Dog) Sound() string {
	return "Wang!"
}

func (cat *Cat) Sound() string {
	return "Meow!"
}

func AnimalFactory(animalType string) Animal {
	switch animalType {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}
