package integers

import (
	"bytes"
	"testing"
)

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}

// 	got := Permeter(rectangle)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f, want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {
	// areaTests := []struct {
	// 	shape Shape
	// 	want  float64
	// }{
	// 	{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
	// 	{Circle{10}, 314.1592653589793},
	// 	{Triangle{12, 6}, 36.0},
	// }

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("%#v: got %g, want %g", shape, got, want)
	// 	}
	// }
	// for _, tt := range areaTests {

	// 	checkArea(t, tt.shape, tt.want)
	// }
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}

// func TestWallet(t *testing.T) {

// 	t.Run("deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))
// 		got := wallet.Balance()

// 		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

// 		want := Bitcoin(10)
// 		if got != want {
// 			t.Errorf("got %s, want %s", want, got)
// 		}
// 	})

// 	t.Run("withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(30)}
// 		wallet.Withdraw(Bitcoin(18))
// 		got := wallet.Balance()

// 		want := Bitcoin(12)
// 		if got != want {
// 			t.Errorf("got %s, want %s", want, got)
// 		}
// 	})

// 	t.Run("withdraw insufficient funds", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(3)}
// 		sum := Bitcoin(90)
// 		err := wallet.Withdraw(sum)

// 		if err == nil {
// 			t.Errorf("wanted error")
// 		}
// 		if wallet.balance != Bitcoin(3) {
// 			t.Errorf("got %s, want %s", wallet.balance, Bitcoin(30))
// 		}

// 	})
// }

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "bla")

	got := buffer.String()
	want := "Hello, bla"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
