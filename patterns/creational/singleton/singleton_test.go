package singleton_test

import (
	"testing"

	"github.com/manishrw/go-design-patterns/patterns/creational/singleton"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSingleton(t *testing.T) {
	RegisterFailHandler(Fail)
	// RunSpecs(t, "Singleton Pattern Suite")
}

var _ = Describe("Singleton", func() {
	Describe("GetInstance", func() {
		Context("when called multiple times", func() {
			It("should return the same instance", func() {
				// Get first instance
				instance1 := singleton.GetInstance()

				// Get second instance
				instance2 := singleton.GetInstance()

				// Verify both instances are the same object
				Expect(instance1).To(BeIdenticalTo(instance2))
			})
		})

		Context("when storing and retrieving values", func() {
			It("should maintain the same state across different references", func() {
				// Store a value using first reference
				instance1 := singleton.GetInstance()
				testValue := "test value"
				instance1.SetValue(testValue)

				// Get a second reference and check the value
				instance2 := singleton.GetInstance()
				retrievedValue := instance2.GetValue()

				// Value should be preserved
				Expect(retrievedValue).To(Equal(testValue))
			})

			It("should allow updating values", func() {
				instance := singleton.GetInstance()

				// Set initial value
				initialValue := 42
				instance.SetValue(initialValue)
				Expect(instance.GetValue()).To(Equal(initialValue))

				// Update value
				updatedValue := "updated"
				instance.SetValue(updatedValue)
				Expect(instance.GetValue()).To(Equal(updatedValue))
			})
		})
	})

	Context("when used in concurrent scenarios", func() {
		It("should maintain singleton property", func() {
			// Create channels to coordinate goroutines
			instanceChan := make(chan *singleton.Singleton, 10)
			done := make(chan bool)

			// Launch multiple goroutines to get the instance
			for i := 0; i < 10; i++ {
				go func() {
					instanceChan <- singleton.GetInstance()
					done <- true
				}()
			}

			// Wait for all goroutines to complete
			for i := 0; i < 10; i++ {
				<-done
			}
			close(instanceChan)

			// Get the first instance for comparison
			var firstInstance *singleton.Singleton
			for instance := range instanceChan {
				if firstInstance == nil {
					firstInstance = instance
				} else {
					// All instances should be identical
					Expect(instance).To(BeIdenticalTo(firstInstance))
				}
			}
		})
	})
})
