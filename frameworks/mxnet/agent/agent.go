package agent

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"google.golang.org/grpc"

	"github.com/pkg/errors"
	"github.com/rai-project/dlframework/frameworks/mxnet"
	rgrpc "github.com/rai-project/grpc"
	context "golang.org/x/net/context"

	dl "github.com/rai-project/dlframework"

	"github.com/rai-project/uuid"

	common "github.com/rai-project/dlframework/frameworks/common/agent"
	predict "github.com/rai-project/dlframework/frameworks/mxnet/predict"
)

type registryServer struct {
	common.Registry
}

type predictorServer struct {
	common.Predictor
}

func (p *predictorServer) Predict(ctx context.Context, req *dl.PredictRequest) (*dl.PredictResponse, error) {
	_, model, err := p.FindFrameworkModel(ctx, req)
	if err != nil {
		return nil, err
	}

	predictor, err := predict.New(*model)
	if err != nil {
		return nil, err
	}
	defer predictor.Close()
	if err := predictor.Download(); err != nil {
		return nil, err
	}

	reader, err := p.InputReaderCloser(ctx, req)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read input as image")
	}
	data, err := predictor.Preprocess(img)
	if err != nil {
		return nil, err
	}

	probs, err := predictor.Predict(data)
	if err != nil {
		return nil, err
	}

	probs.Sort()

	if req.GetLimit() != 0 {
		trunc := probs.Take(int(req.GetLimit()))
		probs = &trunc
	}

	return &dl.PredictResponse{
		Id:       uuid.NewV4(),
		Features: *probs,
		Error:    nil,
	}, nil
}

func RegisterRegistryServer() *grpc.Server {
	var grpcServer *grpc.Server
	grpcServer = rgrpc.NewServer(dl.RegistryServiceDescription)
	svr := &registryServer{
		Registry: common.Registry{
			Base: common.Base{
				Framework: mxnet.FrameworkManifest,
			},
		},
	}
	svr.PublishInRegistery()
	dl.RegisterRegistryServer(grpcServer, svr)
	return grpcServer
}

func RegisterPredictorServer() *grpc.Server {
	var grpcServer *grpc.Server
	grpcServer = rgrpc.NewServer(dl.PredictorServiceDescription)
	svr := &predictorServer{
		Predictor: common.Predictor{
			Base: common.Base{
				Framework: mxnet.FrameworkManifest,
			},
		},
	}
	svr.PublishInRegistery()
	dl.RegisterPredictorServer(grpcServer, svr)
	return grpcServer
}
