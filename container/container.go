// Package container provides a container for find data content by address
// and store data content by address.
package container

type Container interface {
	// Get returns the data content by address.
	// If the address is not found, returns nil.
	Get(address Address) ([]byte, error)

	// Put stores the data content by address.
	// If the address is already exists, the old data content will be overwritten.
	Put(address Address, data []byte) error

	// Delete deletes the data content by address.
	// If the address is not found, returns false.
	Delete(address Address) error

	// Exists checks if an address exists in the container.
	Exists(address Address) bool

	// Keys returns all addresses in the container.
	Keys() []Address

	// Size returns the number of data contents in the container.
	Size() int
}

// Address is a unique identifier for data content.
type Address interface {
	// String returns the string representation of the address.
	String() string

	// Hash returns a unique hash value for the address.
	Hash() string

	// Equals compares this address with another for equality.
	Equals(other Address) bool
}
