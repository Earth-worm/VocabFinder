package input

import (
	"github.com/Earth-worm/VocabFinder/domain/line"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type Global struct {
	LambdaContext *lambdacontext.LambdaContext
	LineConfig    *line.Config
}