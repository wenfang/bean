package retry

// Retry Strategy
type Retry struct {
	strategies []Strategy
	canceled   bool
}

func New(strategies []Strategy) *Retry {
	return &Retry{strategies: strategies}
}

func (r *Retry) Next() bool {
	for _, s := range r.strategies {
		if !s.Next() {
			return false
		}
	}
	return true
}

func (r *Retry) Reset() {
	for _, s := range r.strategies {
		s.Reset()
	}
}

func (r *Retry) Do(action func() error) error {
	var err error
	r.Reset()
	for r.Next() && !r.canceled {
		if err = action(); err == nil {
			return nil
		}
	}
	return err
}

func (r *Retry) Cancel() {
	r.canceled = true
}
