package mockinator

import (
	. "github.com/franela/goblin"
	"github.com/pkg/errors"
	"testing"
)

const expectedIntReturn = 1

var expectedError = errors.New("teste")

type mockExemple struct {
	Mock Mockinator
}

func (d *mockExemple) functionExample() (int, error) {
	a, b := d.Mock.Execute(d.functionExample)
	return a.(int), b
}

func TestMock(t *testing.T) {

	mock := Mockinator{}
	mock.MustInit()
	d := mockExemple{Mock: mock}

	g := Goblin(t)

	g.Describe("should mockinator work as expected", func() {

		g.It("should mock return as expected", func() {

			d.Mock.SetError(d.functionExample, expectedError)
			d.Mock.SetReturn(d.functionExample, expectedIntReturn)

			number, err := d.functionExample()

			g.Assert(number).Equal(expectedIntReturn)
			g.Assert(err).Equal(expectedError)
		})
	})

}
