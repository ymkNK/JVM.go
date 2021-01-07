package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath:= filepath.Join(jreDir,"lib", "*")
	c.bootClasspath=newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir,"lib","ext", "*")
	c.extClasspath=newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption ==""{
		cpOption="."
	}
	c.userClasspath=newEntry(cpOption)
}

//
/*
1 优先使用用户输入的-Xjre选项作为jre目录
2 如果没有输入该 选项，则在当前目录下寻找jre目录
3. 如果找不到，尝试使用 JAVA_HOME环境变量
 */
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _,err :=os.Stat(path); err !=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}