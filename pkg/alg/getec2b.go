package alg

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/random"
)

func GetEc2b() *random.Ec2b {
	open, err := os.Open("data/Ec2b.bin")
	defer open.Close()
	if err != nil {
		ec2p := random.NewEc2b().Bytes()
		ioutil.WriteFile("data/Ec2b.bin", ec2p, 0644)
		log.Println("ec2b不存在,生成ec2b文件中")
		ec2b, err := random.LoadEc2bKey(ec2p)
		if err != nil {
			log.Printf("parse region ec2b error: %v\n", err)
			return nil
		}
		return ec2b
	} else {
		ec2p, err := io.ReadAll(open)
		if err != nil {
			log.Println("read Ec2b error")
			return nil
		}
		defer open.Close()
		ec2b, err := random.LoadEc2bKey(ec2p)
		if err != nil {
			log.Printf("parse region ec2b error: %v\n", err)
			return nil
		}
		return ec2b
	}
}
