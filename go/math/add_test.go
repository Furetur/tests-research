package math

import "testing"
import "fmt"
import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

func TestAdd(t *testing.T) {
	if Add(1, 5) != 6 {
		t.Fail()
	}
}

func TestAddMany(t *testing.T) {
	numbers := [6]int {1, 2, 3, 4, 5, 6}
	for _, n := range numbers {
		t.Run(fmt.Sprintf("%d+%d", n, n) , func(t *testing.T) {
			if Add(n, n) == 2*n {
				t.Fail()
			}
		})
	}
}


func TestGinkgo(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Books Suite")
}

var _ = Describe("Add Test Suite", func() {
	var i int = -1
	BeforeEach(func() {
		i = 0
		fmt.Println(i)
	})
	AfterEach(func() {
		i = 1
	})
	It("should work correctly", func() {
		Expect(Add(1, 10)).To(Equal(12))
	})
})