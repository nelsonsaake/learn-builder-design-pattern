package main

import "fmt"

// Using the Builder pattern makes sense only when your products
// are quite complex and require extensive configuration.
// The following two products are related, although they don't
// have a common interface
type Car struct {
	// A car can have a GPS, trip computer and some number of seats.
	// Different models of cars (sports car, SUV, cabriolet) might
	// have different features installed or enabled.
}

func newCar() *Car {
	return &Car{}
}

type Manual struct {
	// Each car should have a user manual that corresponds to
	// the car's configuration and describes all its features.
}

func newManual() *Manual {
	return &Manual{}
}

// The builder interface specifies methods for creating the
// different parts of the product objects.
type Builder interface {
	reset()
	setSeats(v ...any)
	setEngine(v ...any)
	setTripComputer(v ...any)
	setGPS(v ...any)
}

// The concrete builder classes follow the builder interface and
// provide specific implementations of the building steps.
// Your program may have several variations of builders, each
// implemented differently
type CarBuilder struct {
	car *Car
}

// reset: clears the object being built
func (c *CarBuilder) reset() {
	c.car = newCar()
}

// setSeats: all production step work with the same product instance.
func (c *CarBuilder) setSeats(v ...any) {
	// Set the number of seats in the car.
}

func (c *CarBuilder) setEngine(v ...any) {
	// Install a given engine.
}

func (c *CarBuilder) setTripComputer(v ...any) {
	// Install a trip computer.
}

func (c *CarBuilder) setGPS(v ...any) {
	// Install a global positioning system.
}

// Concrete builders are supposed to provide their own
// methods for retrieving results. That's because vairous
// types of builders may create entirely different products
// that don't all follow the same interface. Therefore such
// methods can't be declared in the builder interface (at
// least not in a statically-typed programming language).
//
// Usually, after returning the end result to the client, a
// builder instance is expected to be ready to start
// producing another product. That's why it's a usual
// practice to call the reset method at the end of the
// `getProduct` method body. However, this behavior isn't
// mandatory, any you can make your builder wait for an
// explicit reset call from the client code before disposing
// of the previous result.
func (c *CarBuilder) getProduct() *Car {
	var product = c.car
	c.reset()
	return product
}

// constructor
func NewCarBuilder() *CarBuilder {
	carBuilder := &CarBuilder{}
	carBuilder.reset()
	return carBuilder
}

// Unlike other creational patterns, builder lets you construct
// products that don't follow the common interface.
type CarManualBuilder struct {
	manual *Manual
}

func (c *CarManualBuilder) reset() {
	c.manual = newManual()
}

func (c *CarManualBuilder) setSeats(v ...any) {
	// Document care seat features.
}

func (c *CarManualBuilder) setEngine(v ...any) {
	// Add trip computer instructions.
}

func (c *CarManualBuilder) setTripComputer(v ...any) {
	// Add trip computer instructions.
}

func (c *CarManualBuilder) setGPS(v ...any) {
	// Add GPS instructions.
}

func (c *CarManualBuilder) getProduct() *Manual {
	// return the manual and reset the builder
	return c.manual
}

func NewCarManualBuilder() *CarManualBuilder {
	carManualBuilder := &CarManualBuilder{}
	carManualBuilder.reset()
	return carManualBuilder
}

// The director is only responsible for executing the building
// steps in a particular sequence. It's helpful when producing
// products according to a specific order or configuration.
// Strictly speaking, the director class is optional, since the
// client can control builders directly.
type Director struct {
	// The director works with any builder instance that the
	// client code passes to it. This way, the client code may
	// alter the final type of the newly assembled product.
	// The director can construct several product variations
	// using the same building steps.
}

func (d *Director) constructSportsCar(builder Builder) {
	builder.reset()
	builder.setSeats(2)
	builder.setEngine("SportEngine")
	builder.setTripComputer(true)
	builder.setGPS(true)
}

func (d *Director) constructSUV(builder Builder) {
	// ...
}

func NewDirector() *Director {
	return &Director{}
}

// The client code creates a builder Object, passes it to the
// director and then initiates the construction process. The end
// result is retrived from the builder object
type Application struct{}

func (a *Application) makeCar() {

	var director = NewDirector()
	var builder Builder

	builder = NewCarBuilder()
	director.constructSportsCar(builder)
	car := builder.(*CarBuilder).getProduct()

	fmt.Println(car)

	builder = NewCarManualBuilder()
	director.constructSportsCar(builder)
	manual := builder.(*CarManualBuilder).getProduct()

	fmt.Println(manual)

	// The final product is often retrieved from a builder
	// object since the director isn't aware of and not
	// dependent on concrete builders and products.
}
