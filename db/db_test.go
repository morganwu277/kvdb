package db

import (
	"math/rand"
	"testing"
)

func randomChar() rune {
	return rune(int('a') + (rand.Int() % 26))
}
func randomStr(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += string(randomChar())
	}
	return str
}

func TestKVDB_ReadWrite(t *testing.T) {
	db, err := Init("/tmp/lev-test-1")
	defer db.Close()
	if err != nil {
		t.Fatalf("error opening db in /tmp/lev-test-1, error: %v \n", err)
		t.Fail()
	}
	m := map[string]string{}
	strNum := 10
	keyNum := 1000
	for i := 0; i < keyNum; i++ {
		k := randomStr(strNum)
		v := randomStr(strNum)
		t.Logf("write k: %v, v: %v \n", k, v)
		m[k] = v
		db.Write(k, v)
	}

	for k, ev := range m {
		t.Logf("read k: %v \n", k)
		v, err := db.Read(k)
		if err != nil || v != ev {
			t.Fatalf("error reading key: %v, expected value: %v, actual value: %v, error: %v \n",
				k, ev, v, err)
			t.Fail()
		}
	}
}
