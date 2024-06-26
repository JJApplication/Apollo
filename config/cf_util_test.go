/*
Project: Apollo cf_util_test.go
Created: 2021/11/22 by Landers
*/

package config

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestInitGlobalConfig(t *testing.T) {
	err := InitGlobalConfig()
	t.Logf("%+v", ApolloConf)
	t.Log(err)
}

func TestSaveGlobalConfig(t *testing.T) {
	ApolloConf = DConfig{
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

func TestGenerateConfig(t *testing.T) {
	cf := DConfig{}
	data, _ := json.MarshalIndent(cf, "", "  ")
	err := os.WriteFile("test_gen.pig", data, 0644)
	if err != nil {
		t.Log(err)
	}
}

func TestSync(t *testing.T) {
	e := InitGlobalConfig()
	if e != nil {
		t.Error(e)
	}
	fmt.Println(ApolloConf)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
			ApolloConf.Sync()
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
			ApolloConf.Update(&DConfig{APPRoot: str, Log: DLog{EnableLog: true}})
			ch <- i
		}(i)
	}

	for i := 0; i < 5; i++ {
		data := <-ch
		fmt.Printf("changed: %+v %d\n", &ApolloConf, data)
	}
}
