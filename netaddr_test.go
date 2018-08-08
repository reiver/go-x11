package x11

import (
	"fmt"
	"math/rand"
	"time"

	"testing"
)

func TestDisplayNetAddr(t *testing.T) {

	tests := []struct{
		Value        string
		ExpectedNet  string
		ExpectedAddr string
	}{
		{
			Value: ":0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: ":1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: ":2",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: ":3",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: ":4",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: ":5",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: ":6",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: ":7",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: ":8",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: ":9",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: ":10",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: ":11",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: ":12",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: ":13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: ":98",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: ":99",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: ":100",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: ":101",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: ":102",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: ":103",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: ":998",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: ":999",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: ":1000",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: ":1001",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: ":1002",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: ":1003",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: ":0.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: ":1.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: ":2.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: ":3.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: ":4.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: ":5.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: ":6.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: ":7.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: ":8.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: ":9.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: ":10.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: ":11.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: ":12.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: ":13.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: ":98.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: ":99.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: ":100.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: ":101.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: ":102.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: ":103.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: ":998.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: ":999.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: ":1000.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: ":1001.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: ":1002.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: ":1003.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: ":0.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: ":1.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: ":2.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: ":3.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: ":4.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: ":5.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: ":6.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: ":7.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: ":8.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: ":9.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: ":10.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: ":11.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: ":12.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: ":13.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: ":98.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: ":99.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: ":100.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: ":101.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: ":102.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: ":103.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: ":998.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: ":999.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: ":1000.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: ":1001.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: ":1002.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: ":1003.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: ":0.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: ":1.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: ":2.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: ":3.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: ":4.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: ":5.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: ":6.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: ":7.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: ":8.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: ":9.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: ":10.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: ":11.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: ":12.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: ":13.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: ":98.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: ":99.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: ":100.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: ":101.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: ":102.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: ":103.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: ":998.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: ":999.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: ":1000.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: ":1001.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: ":1002.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: ":1003.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: ":0.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: ":1.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: ":2.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: ":3.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: ":4.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: ":5.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: ":6.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: ":7.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: ":8.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: ":9.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: ":10.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: ":11.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: ":12.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: ":13.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: ":98.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: ":99.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: ":100.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: ":101.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: ":102.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: ":103.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: ":998.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: ":999.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: ":1000.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: ":1001.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: ":1002.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: ":1003.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: "unix:0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: "unix:1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: "unix:2",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: "unix:3",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: "unix:4",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: "unix:5",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: "unix:6",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: "unix:7",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: "unix:8",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: "unix:9",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: "unix:10",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: "unix:11",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: "unix:12",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: "unix:13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: "unix:98",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: "unix:99",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: "unix:100",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: "unix:101",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: "unix:102",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: "unix:103",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: "unix:998",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: "unix:999",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: "unix:1000",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: "unix:1001",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: "unix:1002",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: "unix:1003",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: "unix:0.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: "unix:1.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: "unix:2.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: "unix:3.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: "unix:4.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: "unix:5.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: "unix:6.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: "unix:7.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: "unix:8.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: "unix:9.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: "unix:10.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: "unix:11.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: "unix:12.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: "unix:13.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: "unix:98.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: "unix:99.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: "unix:100.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: "unix:101.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: "unix:102.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: "unix:103.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: "unix:998.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: "unix:999.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: "unix:1000.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: "unix:1001.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: "unix:1002.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: "unix:1003.0",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: "unix:0.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: "unix:1.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: "unix:2.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: "unix:3.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: "unix:4.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: "unix:5.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: "unix:6.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: "unix:7.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: "unix:8.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: "unix:9.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: "unix:10.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: "unix:11.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: "unix:12.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: "unix:13.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: "unix:98.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: "unix:99.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: "unix:100.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: "unix:101.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: "unix:102.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: "unix:103.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: "unix:998.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: "unix:999.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: "unix:1000.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: "unix:1001.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: "unix:1002.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: "unix:1003.1",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: "unix:0.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: "unix:1.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: "unix:2.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: "unix:3.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: "unix:4.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: "unix:5.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: "unix:6.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: "unix:7.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: "unix:8.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: "unix:9.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: "unix:10.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: "unix:11.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: "unix:12.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: "unix:13.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: "unix:98.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: "unix:99.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: "unix:100.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: "unix:101.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: "unix:102.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: "unix:103.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: "unix:998.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: "unix:999.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: "unix:1000.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: "unix:1001.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: "unix:1002.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: "unix:1003.13",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: "unix:0.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X0",
		},
		{
			Value: "unix:1.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1",
		},
		{
			Value: "unix:2.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X2",
		},
		{
			Value: "unix:3.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X3",
		},
		{
			Value: "unix:4.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X4",
		},
		{
			Value: "unix:5.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X5",
		},
		{
			Value: "unix:6.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X6",
		},
		{
			Value: "unix:7.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X7",
		},
		{
			Value: "unix:8.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X8",
		},
		{
			Value: "unix:9.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X9",
		},
		{
			Value: "unix:10.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X10",
		},
		{
			Value: "unix:11.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X11",
		},
		{
			Value: "unix:12.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X12",
		},
		{
			Value: "unix:13.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X13",
		},

		{
			Value: "unix:98.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X98",
		},
		{
			Value: "unix:99.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X99",
		},
		{
			Value: "unix:100.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X100",
		},
		{
			Value: "unix:101.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X101",
		},
		{
			Value: "unix:102.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X102",
		},
		{
			Value: "unix:103.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X103",
		},

		{
			Value: "unix:998.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X998",
		},
		{
			Value: "unix:999.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X999",
		},
		{
			Value: "unix:1000.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1000",
		},
		{
			Value: "unix:1001.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1001",
		},
		{
			Value: "unix:1002.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1002",
		},
		{
			Value: "unix:1003.72374",
			ExpectedNet: "unix",
			ExpectedAddr: "/tmp/.X11-unix/X1003",
		},









		{
			Value: "127.0.0.1:0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6000",
		},
		{
			Value: "127.0.0.1:1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6001",
		},
		{
			Value: "127.0.0.1:2",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6002",
		},
		{
			Value: "127.0.0.1:3",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6003",
		},
		{
			Value: "127.0.0.1:4",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6004",
		},
		{
			Value: "127.0.0.1:5",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6005",
		},
		{
			Value: "127.0.0.1:6",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6006",
		},
		{
			Value: "127.0.0.1:7",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6007",
		},
		{
			Value: "127.0.0.1:8",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6008",
		},
		{
			Value: "127.0.0.1:9",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6009",
		},
		{
			Value: "127.0.0.1:10",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6010",
		},
		{
			Value: "127.0.0.1:11",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6011",
		},
		{
			Value: "127.0.0.1:12",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6012",
		},
		{
			Value: "127.0.0.1:13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6013",
		},

		{
			Value: "127.0.0.1:98",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6098",
		},
		{
			Value: "127.0.0.1:99",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6099",
		},
		{
			Value: "127.0.0.1:100",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6100",
		},
		{
			Value: "127.0.0.1:101",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6101",
		},
		{
			Value: "127.0.0.1:102",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6102",
		},
		{
			Value: "127.0.0.1:103",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6103",
		},

		{
			Value: "127.0.0.1:998",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6998",
		},
		{
			Value: "127.0.0.1:999",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6999",
		},
		{
			Value: "127.0.0.1:1000",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7000",
		},
		{
			Value: "127.0.0.1:1001",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7001",
		},
		{
			Value: "127.0.0.1:1002",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7002",
		},
		{
			Value: "127.0.0.1:1003",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7003",
		},









		{
			Value: "127.0.0.1:0.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6000",
		},
		{
			Value: "127.0.0.1:1.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6001",
		},
		{
			Value: "127.0.0.1:2.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6002",
		},
		{
			Value: "127.0.0.1:3.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6003",
		},
		{
			Value: "127.0.0.1:4.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6004",
		},
		{
			Value: "127.0.0.1:5.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6005",
		},
		{
			Value: "127.0.0.1:6.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6006",
		},
		{
			Value: "127.0.0.1:7.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6007",
		},
		{
			Value: "127.0.0.1:8.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6008",
		},
		{
			Value: "127.0.0.1:9.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6009",
		},
		{
			Value: "127.0.0.1:10.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6010",
		},
		{
			Value: "127.0.0.1:11.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6011",
		},
		{
			Value: "127.0.0.1:12.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6012",
		},
		{
			Value: "127.0.0.1:13.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6013",
		},

		{
			Value: "127.0.0.1:98.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6098",
		},
		{
			Value: "127.0.0.1:99.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6099",
		},
		{
			Value: "127.0.0.1:100.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6100",
		},
		{
			Value: "127.0.0.1:101.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6101",
		},
		{
			Value: "127.0.0.1:102.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6102",
		},
		{
			Value: "127.0.0.1:103.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6103",
		},

		{
			Value: "127.0.0.1:998.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6998",
		},
		{
			Value: "127.0.0.1:999.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6999",
		},
		{
			Value: "127.0.0.1:1000.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7000",
		},
		{
			Value: "127.0.0.1:1001.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7001",
		},
		{
			Value: "127.0.0.1:1002.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7002",
		},
		{
			Value: "127.0.0.1:1003.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7003",
		},









		{
			Value: "127.0.0.1:0.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6000",
		},
		{
			Value: "127.0.0.1:1.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6001",
		},
		{
			Value: "127.0.0.1:2.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6002",
		},
		{
			Value: "127.0.0.1:3.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6003",
		},
		{
			Value: "127.0.0.1:4.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6004",
		},
		{
			Value: "127.0.0.1:5.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6005",
		},
		{
			Value: "127.0.0.1:6.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6006",
		},
		{
			Value: "127.0.0.1:7.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6007",
		},
		{
			Value: "127.0.0.1:8.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6008",
		},
		{
			Value: "127.0.0.1:9.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6009",
		},
		{
			Value: "127.0.0.1:10.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6010",
		},
		{
			Value: "127.0.0.1:11.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6011",
		},
		{
			Value: "127.0.0.1:12.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6012",
		},
		{
			Value: "127.0.0.1:13.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6013",
		},

		{
			Value: "127.0.0.1:98.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6098",
		},
		{
			Value: "127.0.0.1:99.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6099",
		},
		{
			Value: "127.0.0.1:100.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6100",
		},
		{
			Value: "127.0.0.1:101.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6101",
		},
		{
			Value: "127.0.0.1:102.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6102",
		},
		{
			Value: "127.0.0.1:103.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6103",
		},

		{
			Value: "127.0.0.1:998.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6998",
		},
		{
			Value: "127.0.0.1:999.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6999",
		},
		{
			Value: "127.0.0.1:1000.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7000",
		},
		{
			Value: "127.0.0.1:1001.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7001",
		},
		{
			Value: "127.0.0.1:1002.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7002",
		},
		{
			Value: "127.0.0.1:1003.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7003",
		},









		{
			Value: "127.0.0.1:0.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6000",
		},
		{
			Value: "127.0.0.1:1.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6001",
		},
		{
			Value: "127.0.0.1:2.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6002",
		},
		{
			Value: "127.0.0.1:3.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6003",
		},
		{
			Value: "127.0.0.1:4.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6004",
		},
		{
			Value: "127.0.0.1:5.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6005",
		},
		{
			Value: "127.0.0.1:6.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6006",
		},
		{
			Value: "127.0.0.1:7.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6007",
		},
		{
			Value: "127.0.0.1:8.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6008",
		},
		{
			Value: "127.0.0.1:9.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6009",
		},
		{
			Value: "127.0.0.1:10.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6010",
		},
		{
			Value: "127.0.0.1:11.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6011",
		},
		{
			Value: "127.0.0.1:12.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6012",
		},
		{
			Value: "127.0.0.1:13.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6013",
		},

		{
			Value: "127.0.0.1:98.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6098",
		},
		{
			Value: "127.0.0.1:99.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6099",
		},
		{
			Value: "127.0.0.1:100.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6100",
		},
		{
			Value: "127.0.0.1:101.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6101",
		},
		{
			Value: "127.0.0.1:102.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6102",
		},
		{
			Value: "127.0.0.1:103.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6103",
		},

		{
			Value: "127.0.0.1:998.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6998",
		},
		{
			Value: "127.0.0.1:999.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6999",
		},
		{
			Value: "127.0.0.1:1000.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7000",
		},
		{
			Value: "127.0.0.1:1001.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7001",
		},
		{
			Value: "127.0.0.1:1002.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7002",
		},
		{
			Value: "127.0.0.1:1003.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7003",
		},









		{
			Value: "127.0.0.1:0.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6000",
		},
		{
			Value: "127.0.0.1:1.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6001",
		},
		{
			Value: "127.0.0.1:2.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6002",
		},
		{
			Value: "127.0.0.1:3.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6003",
		},
		{
			Value: "127.0.0.1:4.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6004",
		},
		{
			Value: "127.0.0.1:5.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6005",
		},
		{
			Value: "127.0.0.1:6.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6006",
		},
		{
			Value: "127.0.0.1:7.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6007",
		},
		{
			Value: "127.0.0.1:8.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6008",
		},
		{
			Value: "127.0.0.1:9.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6009",
		},
		{
			Value: "127.0.0.1:10.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6010",
		},
		{
			Value: "127.0.0.1:11.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6011",
		},
		{
			Value: "127.0.0.1:12.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6012",
		},
		{
			Value: "127.0.0.1:13.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6013",
		},

		{
			Value: "127.0.0.1:98.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6098",
		},
		{
			Value: "127.0.0.1:99.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6099",
		},
		{
			Value: "127.0.0.1:100.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6100",
		},
		{
			Value: "127.0.0.1:101.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6101",
		},
		{
			Value: "127.0.0.1:102.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6102",
		},
		{
			Value: "127.0.0.1:103.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6103",
		},

		{
			Value: "127.0.0.1:998.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6998",
		},
		{
			Value: "127.0.0.1:999.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:6999",
		},
		{
			Value: "127.0.0.1:1000.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7000",
		},
		{
			Value: "127.0.0.1:1001.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7001",
		},
		{
			Value: "127.0.0.1:1002.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7002",
		},
		{
			Value: "127.0.0.1:1003.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "127.0.0.1:7003",
		},









		{
			Value: "example.com:0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6000",
		},
		{
			Value: "example.com:1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6001",
		},
		{
			Value: "example.com:2",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6002",
		},
		{
			Value: "example.com:3",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6003",
		},
		{
			Value: "example.com:4",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6004",
		},
		{
			Value: "example.com:5",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6005",
		},
		{
			Value: "example.com:6",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6006",
		},
		{
			Value: "example.com:7",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6007",
		},
		{
			Value: "example.com:8",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6008",
		},
		{
			Value: "example.com:9",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6009",
		},
		{
			Value: "example.com:10",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6010",
		},
		{
			Value: "example.com:11",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6011",
		},
		{
			Value: "example.com:12",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6012",
		},
		{
			Value: "example.com:13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6013",
		},

		{
			Value: "example.com:98",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6098",
		},
		{
			Value: "example.com:99",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6099",
		},
		{
			Value: "example.com:100",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6100",
		},
		{
			Value: "example.com:101",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6101",
		},
		{
			Value: "example.com:102",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6102",
		},
		{
			Value: "example.com:103",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6103",
		},

		{
			Value: "example.com:998",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6998",
		},
		{
			Value: "example.com:999",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6999",
		},
		{
			Value: "example.com:1000",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7000",
		},
		{
			Value: "example.com:1001",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7001",
		},
		{
			Value: "example.com:1002",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7002",
		},
		{
			Value: "example.com:1003",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7003",
		},









		{
			Value: "example.com:0.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6000",
		},
		{
			Value: "example.com:1.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6001",
		},
		{
			Value: "example.com:2.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6002",
		},
		{
			Value: "example.com:3.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6003",
		},
		{
			Value: "example.com:4.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6004",
		},
		{
			Value: "example.com:5.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6005",
		},
		{
			Value: "example.com:6.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6006",
		},
		{
			Value: "example.com:7.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6007",
		},
		{
			Value: "example.com:8.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6008",
		},
		{
			Value: "example.com:9.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6009",
		},
		{
			Value: "example.com:10.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6010",
		},
		{
			Value: "example.com:11.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6011",
		},
		{
			Value: "example.com:12.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6012",
		},
		{
			Value: "example.com:13.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6013",
		},

		{
			Value: "example.com:98.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6098",
		},
		{
			Value: "example.com:99.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6099",
		},
		{
			Value: "example.com:100.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6100",
		},
		{
			Value: "example.com:101.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6101",
		},
		{
			Value: "example.com:102.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6102",
		},
		{
			Value: "example.com:103.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6103",
		},

		{
			Value: "example.com:998.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6998",
		},
		{
			Value: "example.com:999.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6999",
		},
		{
			Value: "example.com:1000.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7000",
		},
		{
			Value: "example.com:1001.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7001",
		},
		{
			Value: "example.com:1002.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7002",
		},
		{
			Value: "example.com:1003.0",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7003",
		},









		{
			Value: "example.com:0.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6000",
		},
		{
			Value: "example.com:1.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6001",
		},
		{
			Value: "example.com:2.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6002",
		},
		{
			Value: "example.com:3.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6003",
		},
		{
			Value: "example.com:4.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6004",
		},
		{
			Value: "example.com:5.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6005",
		},
		{
			Value: "example.com:6.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6006",
		},
		{
			Value: "example.com:7.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6007",
		},
		{
			Value: "example.com:8.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6008",
		},
		{
			Value: "example.com:9.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6009",
		},
		{
			Value: "example.com:10.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6010",
		},
		{
			Value: "example.com:11.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6011",
		},
		{
			Value: "example.com:12.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6012",
		},
		{
			Value: "example.com:13.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6013",
		},

		{
			Value: "example.com:98.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6098",
		},
		{
			Value: "example.com:99.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6099",
		},
		{
			Value: "example.com:100.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6100",
		},
		{
			Value: "example.com:101.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6101",
		},
		{
			Value: "example.com:102.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6102",
		},
		{
			Value: "example.com:103.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6103",
		},

		{
			Value: "example.com:998.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6998",
		},
		{
			Value: "example.com:999.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6999",
		},
		{
			Value: "example.com:1000.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7000",
		},
		{
			Value: "example.com:1001.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7001",
		},
		{
			Value: "example.com:1002.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7002",
		},
		{
			Value: "example.com:1003.1",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7003",
		},









		{
			Value: "example.com:0.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6000",
		},
		{
			Value: "example.com:1.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6001",
		},
		{
			Value: "example.com:2.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6002",
		},
		{
			Value: "example.com:3.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6003",
		},
		{
			Value: "example.com:4.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6004",
		},
		{
			Value: "example.com:5.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6005",
		},
		{
			Value: "example.com:6.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6006",
		},
		{
			Value: "example.com:7.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6007",
		},
		{
			Value: "example.com:8.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6008",
		},
		{
			Value: "example.com:9.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6009",
		},
		{
			Value: "example.com:10.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6010",
		},
		{
			Value: "example.com:11.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6011",
		},
		{
			Value: "example.com:12.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6012",
		},
		{
			Value: "example.com:13.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6013",
		},

		{
			Value: "example.com:98.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6098",
		},
		{
			Value: "example.com:99.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6099",
		},
		{
			Value: "example.com:100.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6100",
		},
		{
			Value: "example.com:101.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6101",
		},
		{
			Value: "example.com:102.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6102",
		},
		{
			Value: "example.com:103.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6103",
		},

		{
			Value: "example.com:998.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6998",
		},
		{
			Value: "example.com:999.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6999",
		},
		{
			Value: "example.com:1000.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7000",
		},
		{
			Value: "example.com:1001.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7001",
		},
		{
			Value: "example.com:1002.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7002",
		},
		{
			Value: "example.com:1003.13",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7003",
		},









		{
			Value: "example.com:0.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6000",
		},
		{
			Value: "example.com:1.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6001",
		},
		{
			Value: "example.com:2.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6002",
		},
		{
			Value: "example.com:3.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6003",
		},
		{
			Value: "example.com:4.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6004",
		},
		{
			Value: "example.com:5.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6005",
		},
		{
			Value: "example.com:6.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6006",
		},
		{
			Value: "example.com:7.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6007",
		},
		{
			Value: "example.com:8.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6008",
		},
		{
			Value: "example.com:9.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6009",
		},
		{
			Value: "example.com:10.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6010",
		},
		{
			Value: "example.com:11.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6011",
		},
		{
			Value: "example.com:12.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6012",
		},
		{
			Value: "example.com:13.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6013",
		},

		{
			Value: "example.com:98.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6098",
		},
		{
			Value: "example.com:99.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6099",
		},
		{
			Value: "example.com:100.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6100",
		},
		{
			Value: "example.com:101.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6101",
		},
		{
			Value: "example.com:102.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6102",
		},
		{
			Value: "example.com:103.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6103",
		},

		{
			Value: "example.com:998.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6998",
		},
		{
			Value: "example.com:999.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:6999",
		},
		{
			Value: "example.com:1000.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7000",
		},
		{
			Value: "example.com:1001.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7001",
		},
		{
			Value: "example.com:1002.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7002",
		},
		{
			Value: "example.com:1003.72374",
			ExpectedNet: "tcp",
			ExpectedAddr: "example.com:7003",
		},
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999999999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d", n),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999999999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf(":%d.%d", n, m),
				ExpectedNet: "unix",
				ExpectedAddr: fmt.Sprintf("/tmp/.X11-unix/X%d", n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999999999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999999999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("127.0.0.1:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("127.0.0.1:%d", 6000+n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(9999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(99999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(9999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(99999999999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(99)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(9999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(99999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(9999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			a := randomness.Int63n(255)
			b := randomness.Int63n(255)
			c := randomness.Int63n(255)
			d := randomness.Int63n(255)

			n := randomness.Int63n(99999999999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("%d.%d.%d.%d:%d.%d", a,b,c,d, n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("%d.%d.%d.%d:%d", a,b,c,d, 6000+n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999999999999)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d", n),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 20

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(9999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}

		for i:=0; i<max; i++ {

			n := randomness.Int63n(99999999999999)
			m := randomness.Int63n(99)

			test := struct{
				Value        string
				ExpectedNet  string
				ExpectedAddr string
			}{
				Value: fmt.Sprintf("example.com:%d.%d", n, m),
				ExpectedNet: "tcp",
				ExpectedAddr: fmt.Sprintf("example.com:%d", 6000+n),
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		actualNet, actualAddr, err := displaynetaddr(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.ExpectedNet, actualNet; expected != actual {
			t.Errorf("For test #%d, expected network to be %q, but actually got %q. (Value for address was %q.)", testNumber, expected, actual, actualAddr)
			continue
		}

		if expected, actual := test.ExpectedAddr, actualAddr; expected != actual {
			t.Errorf("For test #%d, expected address to be %q, but actually got %q. (Value for network was %q.)", testNumber, expected, actual, actualNet)
			continue
		}
	}
}

func TestDisplayNetAddrError(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "",
		},



		{
			Value: " ",
		},
		{
			Value: "  ",
		},
		{
			Value: "   ",
		},

		{
			Value: "          ",
		},



		{
			Value: "\n",
		},



		{
			Value: "host:display.screen",
		},
		{
			Value: ":display.screen",
		},
		{
			Value: ":display",
		},



		{
			Value: "DISPLAY",
		},



		{
			Value: "127.0.0.1:123.234.345",
		},
		{
			Value: "127.0.0.1:123.234.345.456",
		},



		{
			Value: ":123.234.345",
		},
		{
			Value: ":123.234.345.456",
		},



		{
			Value: ":127.0.0.1",
		},
	}

	for testNumber, test := range tests {

		_, _, err := displaynetaddr(test.Value)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %q; for %q.", testNumber, err, err, test.Value)
			continue
		}
	}
}
