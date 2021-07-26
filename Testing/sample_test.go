package Testing

import (
	"fmt"
	"os"
	"testing"
)

// выполняется перед запуском тестов пакета Testing если в други пакетах есть TestMain,
//они запускаться не будут. если нету - запустятся после выполнения тестов пакета TestMain
func TestMain(m *testing.M) {
	fmt.Println("Setup")
	res := m.Run() //
	fmt.Println("Teardown")
	os.Exit(res)
}

// Tests hierarcy
func TestMultiply(t *testing.T) {
	// Setup
	// insert data to db

	// Group of parallel test
	t.Run("GroupA", func(t *testing.T) {
		t.Run("simple", func(t *testing.T) {
			t.Parallel() // Optional - paralel run
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiply(x, y)
			if realResult != result {
				t.Errorf("%d != %d", realResult, result)
			}
			t.Run("1", func(t *testing.T) {
				realResult := Multiply(1, 1)
				if realResult != 1 {
					t.Errorf("Failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284

			realResult := Multiply(x, y)
			if realResult != result {
				t.Errorf("%d != %d", realResult, result)
			}
		})

		t.Run("negative", func(t *testing.T) {
			var x, y, result = -2, 4, -8

			realResult := Multiply(x, y)
			if realResult != result {
				t.Errorf("%d != %d", realResult, result)
			}
		})
	})

	//Teardown
	// delete data in db
}

func TestAdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		var x, y, result = 2, 2, 4

		realResult := Add(x, y)
		if realResult != result {
			t.Errorf("%d != %d", realResult, result)
		}
		t.Run("1", func(t *testing.T) {
			realResult := Add(1, 1)
			if realResult != 2 {
				t.Errorf("Failed")
			}
		})
	})

	// Running tests in terminal 	go test -v -run path/name_test //
	// Running all /simple/1 tests 	go test -v -run /simple/1

}
