package store

type Car struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Price int    `json:"price"`
}

type Cars []Car

func NewThreeCar() *Cars {
	cars := &Cars{
		Car{},
	}
	return cars
}
