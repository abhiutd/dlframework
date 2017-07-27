package predict

import (
	"bufio"
	"image"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/anthonynsimon/bild/parallel"
	"github.com/anthonynsimon/bild/transform"
	"github.com/pkg/errors"
	"github.com/rai-project/dlframework"
	common "github.com/rai-project/dlframework/frameworks/common/predict"
	"github.com/rai-project/downloadmanager"
	gomxnet "github.com/songtianyi/go-mxnet-predictor/mxnet"
)

type ImagePredictor struct {
	common.ImagePredictor
	workDir   string
	features  []string
	predictor *gomxnet.Predictor
}

func New(model dlframework.ModelManifest) (common.Predictor, error) {
	modelInputs := model.GetInputs()
	if len(modelInputs) != 1 {
		return nil, errors.New("number of inputs not supported")
	}
	firstInputType := modelInputs[0].GetType()
	if strings.ToLower(firstInputType) != "image" {
		return nil, errors.New("input type not supported")
	}
	return newImagePredictor(model)
}

func newImagePredictor(model dlframework.ModelManifest) (*ImagePredictor, error) {
	framework, err := model.ResolveFramework()
	if err != nil {
		return nil, err
	}

	workDir, err := model.WorkDir()
	if err != nil {
		return nil, err
	}

	ip := &ImagePredictor{
		ImagePredictor: common.ImagePredictor{
			Base: common.Base{
				Framework: framework,
				Model:     model,
			},
		},
		workDir: workDir,
	}

	return ip, nil
}

func (p *ImagePredictor) GetWeightsUrl() string {
	model := p.Model
	if model.GetModel().GetIsArchive() {
		return model.GetModel().GetBaseUrl()
	}
	return path.Join(model.GetModel().GetBaseUrl(), model.GetModel().GetWeightsPath())
}

func (p *ImagePredictor) GetGraphUrl() string {
	model := p.Model
	if model.GetModel().GetIsArchive() {
		return model.GetModel().GetBaseUrl()
	}
	return path.Join(model.GetModel().GetBaseUrl(), model.GetModel().GetGraphPath())
}

func (p *ImagePredictor) GetFeaturesUrl() string {
	model := p.Model
	params := model.GetOutput().GetParameters()
	pfeats, ok := params["features_url"]
	if !ok {
		return ""
	}
	return pfeats.Value
}

func (p *ImagePredictor) GetGraphPath() string {
	model := p.Model
	graphPath := model.GetModel().GetGraphPath()
	return filepath.Join(p.workDir, graphPath)
}

func (p *ImagePredictor) GetWeightsPath() string {
	model := p.Model
	graphPath := model.GetModel().GetWeightsPath()
	return filepath.Join(p.workDir, graphPath)
}

func (p *ImagePredictor) GetFeaturesPath() string {
	model := p.Model
	return filepath.Join(p.workDir, model.GetName()+".features")
}

func (p *ImagePredictor) Preprocess(input interface{}) (interface{}, error) {
	img, ok := input.(image.Image)
	if !ok {
		return nil, errors.New("expecting an image input")
	}

	imageDims, err := p.GetImageDimensions()
	if err != nil {
		return nil, err
	}
	img = transform.Resize(img, int(imageDims[2]), int(imageDims[3]), transform.Linear)

	meanImage, err := p.GetMeanImage()
	if err != nil || meanImage == nil {
		meanImage = []float32{0, 0, 0}
	}

	b := img.Bounds()
	height := b.Max.Y - b.Min.Y // image height
	width := b.Max.X - b.Min.X  // image width

	res := make([]float32, 3*height*width)
	parallel.Line(height, func(start, end int) {
		w := width
		h := height
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				r, g, b, _ := img.At(x+b.Min.X, y+b.Min.Y).RGBA()
				res[y*w+x] = float32(r>>8) - meanImage[0]
				res[w*h+y*w+x] = float32(g>>8) - meanImage[1]
				res[2*w*h+y*w+x] = float32(b>>8) - meanImage[2]
			}
		}
	})
	return res, nil
}

func (p *ImagePredictor) Download() error {
	if _, err := downloadmanager.Download(p.GetGraphUrl(), p.GetGraphPath()); err != nil {
		return err
	}
	if _, err := downloadmanager.Download(p.GetWeightsUrl(), p.GetWeightsPath()); err != nil {
		return err
	}
	if _, err := downloadmanager.Download(p.GetFeaturesUrl(), p.GetFeaturesPath()); err != nil {
		return err
	}
	return nil
}

func (p *ImagePredictor) getPredictor() error {
	if p.predictor != nil {
		return nil
	}
	symbol, err := ioutil.ReadFile(p.GetGraphPath())
	if err != nil {
		return errors.Wrapf(err, "cannot read %s", p.GetGraphPath())
	}
	params, err := ioutil.ReadFile(p.GetWeightsPath())
	if err != nil {
		return errors.Wrapf(err, "cannot read %s", p.GetWeightsPath())
	}

	var features []string
	f, err := os.Open(p.GetFeaturesPath())
	if err != nil {
		return errors.Wrapf(err, "cannot read %s", p.GetFeaturesPath())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		features = append(features, line)
	}

	p.features = features

	inputDims, err := p.GetImageDimensions()
	if err != nil {
		return err
	}
	modelInputShape := make([]uint32, len(inputDims))
	for ii, v := range inputDims {
		modelInputShape[ii] = uint32(v)
	}

	pred, err := gomxnet.CreatePredictor(symbol,
		params,
		gomxnet.Device{gomxnet.CPU_DEVICE, 0},
		[]gomxnet.InputNode{{Key: "data", Shape: modelInputShape}},
	)
	if err != nil {
		return err
	}
	p.predictor = pred

	return nil
}

func (p *ImagePredictor) Predict(input interface{}) (*dlframework.PredictionFeatures, error) {

	if err := p.getPredictor(); err != nil {
		return nil, err
	}

	data, ok := input.([]float32)
	if !ok {
		return nil, errors.New("expecting []float32 input in predict function")
	}

	if err := p.predictor.SetInput("data", data); err != nil {
		return nil, err
	}

	if err := p.predictor.Forward(); err != nil {
		return nil, err
	}

	probabilities, err := p.predictor.GetOutput(0)
	if err != nil {
		return nil, err
	}

	rprobs := make([]*dlframework.PredictionFeature, len(probabilities))
	for ii, prob := range probabilities {
		rprobs[ii] = &dlframework.PredictionFeature{
			Index:       int64(ii),
			Probability: prob,
		}
	}

	res := dlframework.PredictionFeatures(rprobs)
	return &res, nil
}

func (p *ImagePredictor) Close() error {
	if p.predictor != nil {
		p.predictor.Free()
	}
	return nil
}
