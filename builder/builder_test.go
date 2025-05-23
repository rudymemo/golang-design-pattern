package builder

import "testing"

func TestBuilder(t *testing.T) {
	const (
		windows = 2
		doors   = 1
		rooms   = 3
	)
	house := NewHouseBuilder().SetWindows(2).SetDoors(1).SetRooms(3).Build()

	if house.windows != windows {
		t.Errorf("Expected %d windows, got %d", windows, house.windows)
	}
	if house.doors != doors {
		t.Errorf("Expected %d doors, got %d", doors, house.doors)
	}
	if house.rooms != rooms {
		t.Errorf("Expected %d rooms, got %d", rooms, house.rooms)
	}
}
