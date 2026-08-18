package orchestration

import (
	"context"
	"sync"
)

// Operation 表示一个可被编排的操作。
type Operation func(context.Context) error

// Sequential 串行执行多个 Operation。
// → 一个接一个
// → 出错立即停止后续
// → 返回错误
func Sequential(
	ctx context.Context,
	ops ...Operation,
) error {

	for _, op := range ops {
		if err := op(ctx); err != nil {
			return err
		}
	}

	return nil
}

// ParallelFailFast 并行执行多个 Operation。
// → 全部并行启动
// → 任意失败
// → cancel siblings
// → 等待全部结束
// → 返回错误
func ParallelFailFast(
	ctx context.Context,
	ops ...Operation,
) error {

	if len(ops) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup

	errCh := make(chan error, len(ops))

	wg.Add(len(ops))

	for _, op := range ops {
		go func(op Operation) {
			defer wg.Done()

			if err := op(ctx); err != nil {
				select {
				case errCh <- err:
					cancel()
				default:
				}
			}

		}(op)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		return err
	}

	return nil
}

//把下面 6 个行为钉死：
//1. Sequential success
//2. Sequential error → stop
//3. FailFast error → cancel siblings
//4. WaitAll error → siblings continue
//5. Parent context cancellation → all stop
//6. Empty operations
