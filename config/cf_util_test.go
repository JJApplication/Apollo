/*
Project: dirichlet cf_util_test.go
Created: 2021/11/22 by Landers
*/

package config

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInitGlobalConfig(t *testing.T) {
	err := InitGlobalConfig()
	t.Logf("%+v", DirichletConf)
	t.Log(err)
}

func TestSaveGlobalConfig(t *testing.T) {
	DirichletConf = DConfig{
		ServiceRoot: ".",
		APPRoot:     ".",
		APPManager:  ".",
		APPCacheDir: ".",
		APPLogDir:   ".",
		APPTmpDir:   ".",
		APPBackUp:   ".",
		Log:         DLog{},
		DB:          DDb{},
		Server:      Server{},
	}

	err := SaveGlobalConfig()
	if err != nil {
		t.Log(err)
	}
}

func TestSync(t *testing.T) {
	e := InitGlobalConfig()
	if e != nil {
		t.Error(e)
	}
	fmt.Println(DirichletConf)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
			DirichletConf.Sync()
		}(i)
	}

	for i := 0; i < 10; i++ {
		data := <-ch
		fmt.Println(data)
	}
}

func TestUpdate(t *testing.T) {
	e := InitGlobalConfig()
	if e != nil {
		t.Error(e)
	}
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			str := fmt.Sprintf("%d", rand.Int())
			DirichletConf.Update(&DConfig{APPRoot: str, Log: DLog{EnableLog: "ok"}})
			ch <- i
		}(i)
	}

	for i := 0; i < 5; i++ {
		data := <-ch
		fmt.Printf("changed: %+v %d\n", &DirichletConf, data)
	}
}
