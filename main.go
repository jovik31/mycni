package main

import (
	"fmt"
	"regexp"
	//"strings"
)

func main() {


	arg := "118237498247;K8S_POD_NAME=cni0-tets9ecsdcsdc-defcsdvc;K8S_POD_NAMESPACE=netsn12314132"
	var re = regexp.MustCompile(`(-?;K8S_POD_NAME=)(.+);`)
	fn:=re.FindStringSubmatch(arg)
	//first fn element is K8s_POD_NAME=, the second is the pod name
	fmt.Print(fn[2])

	//access kubernetes api to get the pod annotations to know where to schedule
	//IF annotation not present schedule to default tenant
}
