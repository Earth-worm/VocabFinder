package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/pkg/errors"
)

func GetParameter(name string) (string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("get ssm parameter error:%s", name))
	}
	svc := ssm.New(sess)
	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("get ssm parameter error:%s", name))
	}
	return *res.Parameter.Value, nil
}