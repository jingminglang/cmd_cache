package main
import (
	"fmt"
	"flag"
	"os"
	"os/exec"
	U "github.com/jingminglang/gotools/utils"
)


var cmd_string string
var t_second int

func init() {
	flag.StringVar(&cmd_string,"c","","命令")
	flag.IntVar(&t_second, "t", 0, "缓存时间")
	flag.Parse()
	if(cmd_string == "") {
		flag.Usage()
		os.Exit(2)
	}
}



func main() {
	md5 := U.Md5(cmd_string)
	U.CreateDirIfNotExit("/tmp/cmd_cache/")
	if U.IsFileModifyTimeOlderThanNSecond("/tmp/cmd_cache/"+md5,t_second) {
		out, err := exec.Command("bash", "-c",cmd_string).Output()
		U.Throw(err)
		fmt.Printf("%s\n", out)
		U.WriteStringToFile("/tmp/cmd_cache/"+md5,string(out))
	} else {
		r := U.ReadAllFile("/tmp/cmd_cache/"+md5)
		fmt.Printf("%s\n", r)
	}
}
