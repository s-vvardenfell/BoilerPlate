package somepackage

import (
	"boilerplate/config"
	"time"

	"github.com/sirupsen/logrus"
)

type SomeStruct struct {
	Field1  string
	Field2  string
	isDebug bool
}

func New(cnfg *config.Config, debug bool) *SomeStruct {
	return &SomeStruct{
		Field1:  cnfg.Field1,
		Field2:  cnfg.Field2,
		isDebug: debug,
	}
}

func (ss *SomeStruct) Work() (string, error) {
	for i := 0; i < 2; i++ {
		logrus.Infof("%#v works\n", ss)
		time.Sleep(5 * time.Second)
	}
	return ss.Field1 + " " + ss.Field2, nil
}
