// Copyright 2017 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package grpc_ctxtags

var (
	defaultOptions = &options{
		requestFieldsFunc: nil,
	}
)

type options struct {
	requestFieldsFunc RequestFieldExtractorFunc
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

// WithFieldExtractor customizes the function for extracting log fields from protobuf messages.
func WithFieldExtractor(f RequestFieldExtractorFunc) Option {
	return func(o *options) {
		o.requestFieldsFunc = f
	}
}
