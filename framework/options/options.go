package options

import context "golang.org/x/net/context"

type Options struct {
	ctx        context.Context
	devices    []device
	batchSize  uint
	symbol     []byte
	weights    []byte
	inputNodes []inputNode
	outputNode string
}

type Option func(*Options)

func Context(c context.Context) Option {
	return func(o *Options) {
		o.ctx = c
	}
}

func BatchSize(n uint) Option {
	return func(o *Options) {
		o.batchSize = n
	}
}

func Device(deviceType DeviceType, id int) Option {
	return func(o *Options) {
		o.devices = append(o.devices, device{deviceType: deviceType, id: id})
	}
}

func Symbol(sym []byte) Option {
	return func(o *Options) {
		o.symbol = sym
	}
}

func Weights(w []byte) Option {
	return func(o *Options) {
		o.weights = w
	}
}

func InputNode(key string, shape []uint) Option {
	return func(o *Options) {
		o.inputNodes = append(
			o.inputNodes,
			inputNode{
				key:   key,
				shape: shape,
			},
		)
	}
}

func OutputNode(output string) Option {
	return func(o *Options) {
		o.outputNode = output
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		ctx:       context.Background(),
		batchSize: Config.BatchSize,
		devices: []device{
			device{deviceType: CPU_DEVICE, id: 0},
		},
	}

	for _, o := range opts {
		o(options)
	}

	for ii, inputNode := range options.inputNodes {
		if len(options.inputNodes[ii].shape) == 3 {
			options.inputNodes[ii].shape = append([]uint{options.batchSize}, inputNode.shape...)
		} else {
			options.inputNodes[ii].shape[0] = options.batchSize
		}
	}

	return options
}
