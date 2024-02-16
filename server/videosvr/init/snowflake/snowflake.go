package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
	"sync"
	"time"
	"usersvr/log"
	"videosvr/config"
)

var (
	node *snowflake.Node
	Once sync.Once
)

func initSlowFlake(startTime string, machineID int) {
	st, err := time.Parse("2006-01-02 00:00:00", startTime)
	if err != nil {
		log.Fatalf("parse time err:%s", err.Error())
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(int64(machineID))
	if err != nil {
		log.Fatalf("init snowflake err:%s", err.Error())
	}
	return
}
func GenerateID() string {
	Once.Do(func() {
		initSlowFlake(time.Now().Format("2006-01-02 00:00:00"), config.GetGlobalConfig().SvrConfig.MachineID)
	})
	return strconv.FormatInt(node.Generate().Int64(), 10)
}
