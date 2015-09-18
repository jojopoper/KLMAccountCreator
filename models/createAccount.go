package models

import (
	"crypto/rand"
	"fmt"
	"github.com/jojopoper/go-stellar-base"
	"runtime"
	"strings"
	"time"
)

type AccountInfo struct {
	Index      int
	SecretSeed string
	PublicID   string
}

func CreateNewKLMAccount(times, timeout int, filter string, compFlag chan bool, ret []*AccountInfo) {
	tm := time.NewTicker(time.Duration(timeout) * time.Second)
	headStr := "N" + filter
	var currCount int = 0

	defer runtime.GC()

	for {
		select {
		case <-tm.C:
			tm.Stop()
			compFlag <- false
			return
		default:
			{
				b := random(32)
				seed, err := stellarbase.NewRawSeed(b)
				if err == nil {
					pubKey, priKey, err := stellarbase.GenerateKeyFromRawSeed(seed)
					if err == nil {
						addr := pubKey.Address()
						if len(filter) == 0 {
							ret[currCount] = &AccountInfo{
								Index:      currCount + 1,
								SecretSeed: priKey.Seed(),
								PublicID:   addr,
							}
							currCount++
							// fmt.Printf("[%d\t]seed = %s \r\n addr = %s\r\n", currCount, priKey.Seed(), addr)
						} else {
							if strings.HasPrefix(addr, headStr) {
								ret[currCount] = &AccountInfo{
									Index:      currCount + 1,
									SecretSeed: priKey.Seed(),
									PublicID:   addr,
								}
								currCount++
								// fmt.Printf("[%d\t]seed = %s \r\n addr = %s\r\n", currCount, priKey.Seed(), addr)
							}
						}
					}
				}

				if currCount >= times {
					tm.Stop()
					compFlag <- true
					return
				}
			}
		}
	}
}

func random(length int) []byte {
	//rand Read
	k := make([]byte, length)
	if _, err := rand.Read(k); err != nil {
		fmt.Println("rand.Read() error : %v \n", err)
	}
	return k
}
