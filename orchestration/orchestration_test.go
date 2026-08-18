package orchestration

import (
	"context"
	"errors"
	"testing"
)

func TestOrchestration_Sequential(t *testing.T) {

	ctx := context.Background()

	var result []string

	err := Sequential(
		ctx,
		func(ctx context.Context) error {
			result = append(result, "A")
			return nil
		},
		func(ctx context.Context) error {
			result = append(result, "B")
			return nil
		},
		func(ctx context.Context) error {
			result = append(result, "C")
			return nil
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestOrchestration_Parallel(t *testing.T) {

	ctx := context.Background()

	var result []string

	err := ParallelFailFast(
		ctx,
		func(ctx context.Context) error {
			result = append(result, "A")
			return nil
		},
		func(ctx context.Context) error {
			result = append(result, "B")
			return nil
		},
		func(ctx context.Context) error {
			result = append(result, "C")
			return nil
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

//PS D:\roc-github\goro-edge> go test ./applications/edge/aggregate/dashboard -v
//=== RUN   TestOrchestration_Sequential
//    orchestration_test.go:34: [A B C]
//--- PASS: TestOrchestration_Sequential (0.00s)
//=== RUN   TestOrchestration_Parallel
//    orchestration_test.go:63: [C A B]
//--- PASS: TestOrchestration_Parallel (0.00s)
//PASS
//ok      github.com/rocwg/goro-edge/applications/edge/aggregate/dashboard        0.015s
//PS D:\roc-github\goro-edge>

func TestOrchestration_SequentialError(t *testing.T) {

	ctx := context.Background()

	var result []string

	errExpected := errors.New("B failed")

	err := Sequential(
		ctx,

		func(ctx context.Context) error {
			result = append(result, "A")
			return nil
		},

		func(ctx context.Context) error {
			result = append(result, "B")
			return errExpected
		},

		func(ctx context.Context) error {
			result = append(result, "C")
			return nil
		},
	)

	if !errors.Is(err, errExpected) {
		t.Fatalf(
			"expected %v, got %v",
			errExpected,
			err,
		)
	}

	t.Log(result)
}

func TestOrchestration_ParallelError(t *testing.T) {

	ctx := context.Background()

	errExpected := errors.New("B failed")

	err := ParallelFailFast(
		ctx,

		func(ctx context.Context) error {
			return nil
		},

		func(ctx context.Context) error {
			return errExpected
		},

		func(ctx context.Context) error {
			return nil
		},
	)

	if !errors.Is(err, errExpected) {
		t.Fatalf(
			"expected %v, got %v",
			errExpected,
			err,
		)
	}
}

func TestOrchestration_Mixed(t *testing.T) {

	ctx := context.Background()

	var result []string

	err := ParallelFailFast(
		ctx,

		func(ctx context.Context) error {
			result = append(result, "A")
			return nil
		},

		func(ctx context.Context) error {
			result = append(result, "B")
			return nil
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	err = Sequential(
		ctx,

		func(ctx context.Context) error {
			result = append(result, "C")
			return nil
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestOrchestration_ParallelCancellation1(t *testing.T) {

	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()

	errExpected := errors.New("B failed")

	err := ParallelFailFast(
		ctx,

		func(ctx context.Context) error {

			<-ctx.Done()

			t.Log("A cancelled")

			return ctx.Err()
		},

		func(ctx context.Context) error {

			cancel()

			return errExpected
		},

		func(ctx context.Context) error {

			<-ctx.Done()

			t.Log("C cancelled")

			return ctx.Err()
		},
	)

	if !errors.Is(err, errExpected) {
		t.Fatalf(
			"expected %v, got %v",
			errExpected,
			err,
		)
	}
}

// 现在补 6 个测试
// 1. Empty
func TestOrchestration_ParallelEmpty(t *testing.T) {

	err := ParallelFailFast(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

// 2. Single
func TestOrchestration_ParallelSingle(t *testing.T) {

	called := false

	err := ParallelFailFast(
		context.Background(),
		func(ctx context.Context) error {
			called = true
			return nil
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("operation was not called")
	}
}

// 3. All success
// TestOrchestration_Parallel
// 4. Error
// TestOrchestration_ParallelError
// 5. Cancellation
func TestOrchestration_ParallelCancellation(t *testing.T) {

	errExpected := errors.New("B failed")

	var cancelledA bool
	var cancelledC bool

	err := ParallelFailFast(
		context.Background(),

		func(ctx context.Context) error {
			<-ctx.Done()
			cancelledA = true
			return ctx.Err()
		},

		func(ctx context.Context) error {
			return errExpected
		},

		func(ctx context.Context) error {
			<-ctx.Done()
			cancelledC = true
			return ctx.Err()
		},
	)

	if !errors.Is(err, errExpected) {
		t.Fatalf(
			"expected %v, got %v",
			errExpected,
			err,
		)
	}

	if !cancelledA {
		t.Fatal("A was not cancelled")
	}

	if !cancelledC {
		t.Fatal("C was not cancelled")
	}
}

// (6)
