package trycatch

// TryCache try and cache code block
func TryCatch(tryFunc func(), catchFunc func(interface{})) {
	TryCatchFinally(tryFunc, catchFunc, nil)
}

// TryFinally try, finally code block
func TryFinally(tryFunc, finallyFunc func()) {
	TryCatchFinally(tryFunc, nil, finallyFunc)
}

// TryCacheFinally try, catch,  finally code block
func TryCatchFinally(tryFunc func(), catchFunc func(interface{}), finallyFunc func()) {
	defer func() {
		if err := recover(); err != nil {
			if catchFunc != nil {
				catchFunc(err)
			}
		}
		if finallyFunc != nil {
			finallyFunc()
		}
	}()
	if tryFunc != nil {
		tryFunc()
	}
}
