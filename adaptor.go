package gobot

// DigitalPinOptioner is the interface to provide the possibility to change pin behavior for the next usage
type DigitalPinOptioner interface {
	// SetLabel change the pins label
	SetLabel(string) (changed bool)
	// SetDirectionOutput sets the pins direction to output with the given initial value
	SetDirectionOutput(initialState int) (changed bool)
	// SetDirectionInput sets the pins direction to input
	SetDirectionInput() (changed bool)
}

// DigitalPinOptionApplier is the interface to apply options to change pin behavior immediately
type DigitalPinOptionApplier interface {
	// ApplyOptions apply all given options to the pin immediately
	ApplyOptions(...func(DigitalPinOptioner) bool) error
}

// DigitalPinner is the interface for system gpio interactions
type DigitalPinner interface {
	// Export exports the pin for use by the adaptor
	Export() error
	// Unexport releases the pin from the adaptor, so it is free for the operating system
	Unexport() error
	// Read reads the current value of the pin
	Read() (int, error)
	// Write writes to the pin
	Write(int) error
	// DigitalPinOptionApplier is the interface to change pin behavior immediately
	DigitalPinOptionApplier
}

// DigitalPinValuer is the interface to get pin behavior for the next usage. The interface is and should be rarely used.
type DigitalPinValuer interface {
	// DirectionBehavior gets the direction behavior when the pin is used the next time.
	// This means its possibly not in this direction type at the moment.
	DirectionBehavior() string
}

// DigitalPinnerProvider is the interface that an Adaptor should implement to allow clients to obtain
// access to any DigitalPin's available on that board. If the pin is initially acquired, it is an input.
// Pin direction and other options can be changed afterwards by pin.ApplyOptions() at any time.
type DigitalPinnerProvider interface {
	DigitalPin(id string) (DigitalPinner, error)
}

// PWMPinner is the interface for system PWM interactions
type PWMPinner interface {
	// Export exports the PWM pin for use by the operating system
	Export() error
	// Unexport releases the PWM pin from the operating system
	Unexport() error
	// Enabled returns the enabled state of the PWM pin
	Enabled() (bool, error)
	// SetEnabled enables/disables the PWM pin
	SetEnabled(bool) error
	// Polarity returns true if the polarity of the PWM pin is normal, otherwise false
	Polarity() (bool, error)
	// SetPolarity sets the polarity of the PWM pin to normal if called with true and to inverted if called with false
	SetPolarity(normal bool) error
	// Period returns the current PWM period in nanoseconds for pin
	Period() (uint32, error)
	// SetPeriod sets the current PWM period in nanoseconds for pin
	SetPeriod(uint32) error
	// DutyCycle returns the duty cycle in nanoseconds for the PWM pin
	DutyCycle() (uint32, error)
	// SetDutyCycle writes the duty cycle in nanoseconds to the PWM pin
	SetDutyCycle(uint32) error
}

// PWMPinnerProvider is the interface that an Adaptor should implement to allow
// clients to obtain access to any PWMPin's available on that board.
type PWMPinnerProvider interface {
	PWMPin(id string) (PWMPinner, error)
}

// Adaptor is the interface that describes an adaptor in gobot
type Adaptor interface {
	// Name returns the label for the Adaptor
	Name() string
	// SetName sets the label for the Adaptor
	SetName(string)
	// Connect initiates the Adaptor
	Connect() error
	// Finalize terminates the Adaptor
	Finalize() error
}

// Porter is the interface that describes an adaptor's port
type Porter interface {
	Port() string
}
