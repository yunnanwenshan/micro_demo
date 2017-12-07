package run

type FetchOptions struct {
	Update bool
}

type FetchOption func(o *FetchOptions)

// Update tells Fetch to update the source
func Update(b bool) FetchOption {
	return func(o *FetchOptions) {
		o.Update = b
	}
}
