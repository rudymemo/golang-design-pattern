package observer

import "fmt"

type Observer interface {
	Update(message string)
}

type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Notify(message string)
}

type NewsAgency struct {
	observers []Observer
}

func NewNewsAgency() *NewsAgency {
	return &NewsAgency{
		observers: make([]Observer, 0),
	}
}

func (na *NewsAgency) Subscribe(observer Observer) {
	na.observers = append(na.observers, observer)
}

func (na *NewsAgency) Unsubscribe(observer Observer) {
	for i, obs := range na.observers {
		if obs == observer {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			break
		}
	}
}

func (na *NewsAgency) Notify(message string) {
	for _, observer := range na.observers {
		observer.Update(message)
	}
}

type NewsChannel struct {
	name string
	news string
}

func NewNewsChannel(name string) *NewsChannel {
	return &NewsChannel{name: name}
}

func (nc *NewsChannel) Update(message string) {
	nc.news = message
	fmt.Printf("%s received news: %s\n", nc.name, message)
}

func (nc *NewsChannel) GetNews() string {
	return nc.news
}