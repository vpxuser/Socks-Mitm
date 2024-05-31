package main

import (
	"fmt"
	yaklog "github.com/yaklang/yaklang/common/log"
	"socks2https/setting"
	"socks2https/socks"
)

const PROGRAM_NAME = "\n███████╗ ██████╗  ██████╗██╗  ██╗███████╗██████╗ ██╗  ██╗████████╗████████╗██████╗ ███████╗\n██╔════╝██╔═══██╗██╔════╝██║ ██╔╝██╔════╝╚════██╗██║  ██║╚══██╔══╝╚══██╔══╝██╔══██╗██╔════╝\n███████╗██║   ██║██║     █████╔╝ ███████╗ █████╔╝███████║   ██║      ██║   ██████╔╝███████╗\n╚════██║██║   ██║██║     ██╔═██╗ ╚════██║██╔═══╝ ██╔══██║   ██║      ██║   ██╔═══╝ ╚════██║\n███████║╚██████╔╝╚██████╗██║  ██╗███████║███████╗██║  ██║   ██║      ██║   ██║     ███████║\n╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝     ╚══════╝\n                                                                                           \n"

func init() {
	fmt.Println(PROGRAM_NAME)
	yaklog.SetLevel(setting.Config.Log.Level)
}

func main() {
	socks.Run()
}
