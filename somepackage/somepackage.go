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

func (ss *SomeStruct) Work() {
	for {
		logrus.Info("%#v works\n", ss)
		time.Sleep(5 * time.Second)
	}
}
