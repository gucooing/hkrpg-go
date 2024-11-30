package alg

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

func GetEc2b() *random.Ec2b {
	open, err := os.Open("data/Ec2b.bin")
	defer open.Close()
	if err != nil {
		ec2p := random.NewEc2b()
		ioutil.WriteFile("data/Ec2b.bin", ec2p.Bytes(), 0644)
		logger.Info(text.GetText(10))
		return ec2p
	} else {
		ec2p, err := io.ReadAll(open)
		if err != nil {
			logger.Error(text.GetText(11), err)
			return nil
		}
		defer open.Close()
		ec2b, err := random.LoadEc2bKey(ec2p)
		if err != nil {
			logger.Error(text.GetText(12), err)
			return nil
		}
		logger.Info(text.GetText(13))
		return ec2b
	}
}
