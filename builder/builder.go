package builder

type House struct {
	windows int
	doors   int
	rooms   int
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (builder *HouseBuilder) SetWindows(windows int) *HouseBuilder {
	builder.house.windows = windows
	return builder
}

func (builder *HouseBuilder) SetDoors(doors int) *HouseBuilder {
	builder.house.doors = doors
	return builder
}

func (builder *HouseBuilder) SetRooms(rooms int) *HouseBuilder {
	builder.house.rooms = rooms
	return builder
}

func (builder *HouseBuilder) Build() House {
	return builder.house
}
