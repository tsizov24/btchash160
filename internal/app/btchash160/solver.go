package btchash160

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"time"

	"btchash160/pkg/logger"

	"github.com/tsizov24/btclib"
)

var (
	hash160     []byte
	min, max    *big.Int
	lim         int
	numOfBlocks *big.Int
)

func init() {
	var err error
	hash160, err = hex.DecodeString(conf.Hash160)
	logger.Log(err, logger.PanicLevel)

	min = getInt(conf.Min)
	max = getInt(conf.Max)

	lim = 1 << 32

	numOfBlocks = big.NewInt(1)
	numOfBlocks.Add(numOfBlocks, max)
	numOfBlocks.Sub(numOfBlocks, min)
	numOfBlocks.Div(numOfBlocks, big.NewInt(int64(lim)))
}

func getInt(s string) *big.Int {
	res, ok := new(big.Int).SetString(s, 16)
	if !ok {
		logger.Log(fmt.Errorf("wrong int %s", s), logger.PanicLevel)
	}
	return res
}

func Start() {
	for i := 1; i < runtime.NumCPU(); i++ {
		go count()
	}
	count()
}

func count() {
	var k *big.Int
	for {
		t := time.Now()
		k = getRand()
		if solve(k) {
			break
		}
		fmt.Println(time.Now(), time.Since(t))
	}
	fmt.Println(k.Text(16))
}

func getRand() *big.Int {
	res, err := rand.Int(rand.Reader, numOfBlocks)
	logger.Log(err, logger.PanicLevel)
	res.Mul(res, big.NewInt(int64(lim)))
	res.Add(res, min)
	return res
}

func solve(k *big.Int) (ok bool) {
	gx, gy := btclib.GetBasePoint()
	x, y := btclib.PrivToPub(k)
	for i := 0; i < lim; i++ {
		if reflect.DeepEqual(hash160, btclib.PubToHash160(x, y, true)) {
			ok = true
			k.Add(k, big.NewInt(int64(i)))
			break
		}
		x, y = btclib.AddPoints(x, y, gx, gy)
	}
	return
}
