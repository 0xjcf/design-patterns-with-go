package builder

import "testing"

func TestBuildWithoutDirector(t *testing.T) {
	customSportsCarBuilder := CustomSportsCarBuilder{}
	customSportsCarBuilder.SetColor("Green")
	customSportsCarBuilder.SetSeats(2)
	customSportsCarBuilder.SetDoors(2)
	customSportsCarBuilder.SetTransmition("Manual")
	customSportsCarBuilder.SetWheels(4)
	sportsCar := customSportsCarBuilder.GetVehicle()

	if sportsCar.Color != "Green" {
		t.Errorf("Expected color to be 'Green' but got %s\n instead", sportsCar.Color)
	}
	if sportsCar.Seats != 2 {
		t.Errorf("Expected seats to be 2 but got %d\n instead", sportsCar.Seats)
	}
	if sportsCar.Doors != 2 {
		t.Errorf("Expected doors to be 2 but got %d\n instead", sportsCar.Doors)
	}
	if sportsCar.Transmission != "Manual" {
		t.Errorf("Expected transmission to be 'Manual' but got %s\n instead", sportsCar.Transmission)
	}
	if sportsCar.Wheels != 4 {
		t.Errorf("Expected wheel count to be 4 but got %d\n instead", sportsCar.Wheels)
	}
}

func TestBuildWithDirector(t *testing.T) {
	manufacturer := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturer.SetBuilder(carBuilder)
	manufacturer.BuildVehicle()
	car := carBuilder.GetVehicle()

	if car.Color != "White" {
		t.Errorf("Expected color to be 'White' but got %s\n instead", car.Color)
	}
	if car.Seats != 5 {
		t.Errorf("Expected seats to be 2 but got %d\n instead", car.Seats)
	}
	if car.Doors != 4 {
		t.Errorf("Expected seats to be 4 but got %d\n instead", car.Doors)
	}
	if car.Transmission != "Automatic" {
		t.Errorf("Expected transmission to be 'Manual' but got %s\n instead", car.Transmission)
	}
	if car.Wheels != 4 {
		t.Errorf("Expected wheel count to be 4 but got %d\n instead", car.Wheels)
	}
}
