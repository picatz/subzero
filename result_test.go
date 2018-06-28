package subzero

import "testing"
import "fmt"
import "errors"
import "reflect"

func TestResult(t *testing.T) {
	var units = []struct {
		exp Result
		got string
	}{
		{Result{Type: "example", Success: "info.bing.com"}, "info.bing.com"},
		{Result{Type: "example", Failure: errors.New("failed")}, "failed"},
	}
	for _, u := range units {
		if u.exp.Failure != nil {
			if !reflect.DeepEqual(u.exp.Failure.Error(), u.got) {
				t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
			}
		} else {
			if !reflect.DeepEqual(u.exp.Success, u.got) {
				t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
			}
		}
	}
}

func TestResultIsSuccess(t *testing.T) {
	var units = []struct {
		exp Result
		got bool
	}{
		{Result{Type: "example", Success: "info.bing.com"}, true},
		{Result{Type: "example", Failure: errors.New("failed")}, false},
	}
	for _, u := range units {
		if u.exp.IsSuccess() != u.got {
			t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
		}
	}
}

func TestResultIsFailure(t *testing.T) {
	var units = []struct {
		exp Result
		got bool
	}{
		{Result{Type: "example", Success: "info.bing.com"}, false},
		{Result{Type: "example", Failure: errors.New("failed")}, true},
	}
	for _, u := range units {
		if u.exp.IsFailure() != u.got {
			t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
		}
	}
}

func TestResultHasType(t *testing.T) {
	var units = []struct {
		exp Result
		got bool
	}{
		{Result{Type: "example", Success: "info.bing.com"}, true},
		{Result{Type: "example", Failure: errors.New("failed")}, true},
		{Result{}, false},
	}
	for _, u := range units {
		if u.exp.HasType() != u.got {
			t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
		}
	}
}

func ExampleResult() {
	result := Result{Type: "example", Success: "info.bing.com"}
	if result.Failure != nil {
		fmt.Println(result.Type, ":", result.Failure)
	} else {
		fmt.Println(result.Type, ":", result.Success)
	}
	// Output: example : info.bing.com
}

func ExampleResultIsSuccess() {
	result := Result{Success: "wiggle.github.com"}
	if result.IsSuccess() {
		fmt.Println(result.Success)
	}
	// Output: wiggle.github.com
}

func ExampleResultIsFailure() {
	result := Result{Failure: errors.New("failed to party")}
	if result.IsFailure() {
		fmt.Println(result.Failure.Error())
	}
	// Output: failed to party
}

func ExampleResultHasType() {
	result := Result{Type: "example"}
	fmt.Println(result.HasType())
	// Output: true
}
