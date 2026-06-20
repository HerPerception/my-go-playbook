package main
import "testing"
func TestInitializeArray(t *testing.T) {
got := InitializeArray()
want := [4]int{0, 10, 20, 30}
if got != want {
t.Errorf("got %v, want %v", got, want)
}
// Verify the structural constraint of fixed length
var _ [4]int = got
}