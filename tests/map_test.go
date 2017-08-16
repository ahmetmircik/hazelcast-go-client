package tests

import (
	"github.com/hazelcast/go-client"
	. "github.com/hazelcast/go-client/rc"
	"log"
	"strconv"
	"testing"
)

const DEFAULT_XML_CONFIG string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><hazelcast xsi:schemaLocation=\"http://www.hazelcast.com/schema/config hazelcast-config-3.9.xsd\" xmlns=\"http://www.hazelcast.com/schema/config\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"></hazelcast>"

func TestMain(m *testing.M) {
	remoteController, err := NewRemoteControllerClient("localhost:9701")
	if remoteController == nil || err != nil {
		log.Fatal("create remote controller failed:", err)
	}
	cluster, err := remoteController.CreateCluster("3.9", DEFAULT_XML_CONFIG)
	remoteController.StartMember(cluster.ID)
	m.Run()
	remoteController.ShutdownCluster(cluster.ID)
}
func TestMapProxy_SinglePutGet(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	testKey := "testingKey"
	testValue := "testingValue"
	mp.Put(testKey, testValue)
	res, err := mp.Get(testKey)
	if err != nil {
		t.Error(err)
	} else {
		if res != testValue {
			t.Errorf("get returned a wrong value")
		}
	}
}
func TestMapProxy_Remove(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	testKey := "testingKey"
	testValue := "testingValue"
	mp.Put(testKey, testValue)
	removed, err := mp.Remove(testKey)
	if err != nil {
		t.Error(err)
	} else {
		if removed != testValue {
			t.Errorf("remove returned a wrong value")
		}
	}
	size, err := mp.Size()
	if err != nil {
		t.Error(err)
	} else {
		if size != 0 {
			t.Errorf("Map size should be 0.")
		}
	}
	found, err := mp.ContainsKey(testKey)
	if err != nil {
		t.Error(err)
	} else {
		if found {
			t.Errorf("containsKey returned a wrong result")
		}
	}

	//TODO::Check if map contains "testingkey"
}
func TestMapProxy_ContainsKey(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	testKey := "testingKey1"
	testValue := "testingValue"
	mp.Put(testKey, testValue)
	found, err := mp.ContainsKey(testKey)
	if err != nil {
		t.Error(err)
	} else {
		if !found {
			t.Errorf("containsKey returned a wrong result")
		}
	}
	found, err = mp.ContainsKey("testingKey2")
	if err != nil {
		t.Error(err)
	} else {
		if found {
			t.Errorf("containsKey returned a wrong result")
		}
	}
}
func TestMapProxy_ContainsValue(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	testKey := "testingKey1"
	testValue := "testingValue"
	mp.Put(testKey, testValue)
	found, err := mp.ContainsValue(testValue)
	if err != nil {
		t.Error(err)
	} else {
		if !found {
			t.Errorf("containsValue returned a wrong result")
		}
	}
	found, err = mp.ContainsValue("testingValue2")
	if err != nil {
		t.Error(err)
	} else {
		if found {
			t.Errorf("containsValue returned a wrong result")
		}
	}
}
func TestMapProxy_Clear(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	testKey := "testingKey1"
	testValue := "testingValue"
	mp.Put(testKey, testValue)
	err := mp.Clear()
	if err != nil {
		t.Error(err)
	} else {
		size, err := mp.Size()
		if err != nil {
			t.Error(err)
		} else {
			if size != 0 {
				t.Errorf("Map clear failed.")
			}
		}
	}
}
func TestMapProxy_Delete(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	for i := 0; i < 10; i++ {
		mp.Put("testingKey"+strconv.Itoa(i), "testingValue"+strconv.Itoa(i))
	}
	mp.Delete("testingKey1")
	size, err := mp.Size()
	if err != nil {
		t.Error(err)
	} else {
		if size != 9 {
			t.Errorf("Map Delete failed")
		}
	}
}
func TestMapProxy_IsEmpty(t *testing.T) {
	client := hazelcast.NewHazelcastClient()
	mapName := "myMap"
	mp := client.GetMap(&mapName)
	for i := 0; i < 10; i++ {
		mp.Put("testingKey"+strconv.Itoa(i), "testingValue"+strconv.Itoa(i))
	}
	empty, err := mp.IsEmpty()
	if err != nil {
		t.Error(err)
	} else {
		if empty {
			t.Errorf("Map IsEmpty returned a wrong value")
		}
	}
}
