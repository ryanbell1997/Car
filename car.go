package main

import (
	"errors"
	"fmt"
	"time"
)

type Car interface {
	SpeedUp() error
	SlowDown() error
}

type Mercedes struct {
	MaxSpeed     int
	CurrentSpeed int
}

func (m *Mercedes) SpeedUp() error {
	if m.CurrentSpeed >= m.MaxSpeed {
		fmt.Printf("Car has reached or exceeded max speed of: %d, Current speed is: %d\n", m.MaxSpeed, m.CurrentSpeed)
		return errors.New("At max speed")
	}

	m.CurrentSpeed += 5
	fmt.Printf("Current speed is: %d\n", m.CurrentSpeed)
	return m.SpeedUp()
}

func (m *Mercedes) SlowDown() error {
	if m.CurrentSpeed <= 0 {
		fmt.Printf("Car has stopped\n")
		//Ensure that the current speed has gone under 0.
		m.CurrentSpeed = 0
		return errors.New("Have slowed down to a stop")
	}

	fmt.Printf("Car slowing down. Current speed: %d\n", m.CurrentSpeed)
	m.CurrentSpeed -= 10
	return m.SlowDown()
}

func driveCar(car Car) {
	err := car.SpeedUp()
	if err != nil {
		car.SlowDown()
	}
}

func main() {
	mercedes := Mercedes{MaxSpeed: 100, CurrentSpeed: 0}

	go driveCar(&mercedes)
	time.Sleep(10 * time.Second)
}
