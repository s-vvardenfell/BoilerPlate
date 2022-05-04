package somepackage

import (
	"boilerplate/config"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func Test_Work(t *testing.T) {

	log := config.Logrus{
		LogLvl: 0,
		ToFile: false,
		ToJson: false,
		LogDir: "",
	}

	cnfg := &config.Config{
		Field1: "test field 1",
		Field2: "test field 2",
		Logrus: log,
	}

	debug := false
	logrus.SetLevel(logrus.Level(cnfg.Logrus.LogLvl))

	t.Log("\tTesting 'Work'")
	{
		ss := New(cnfg, debug)
		res, err := ss.Work()
		require.NoError(t, err)
		require.EqualValues(t, res, cnfg.Field1+" "+cnfg.Field2)
	}
}
