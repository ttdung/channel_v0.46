package main

import (
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ttdung/channel_v0.46/app"
	"github.com/ttdung/channel_v0.46/cmd/channeld/cmd"
	"os"
)

func main() {
	//arg2 := strings.Split("0hello:1stake", ":")
	//fmt.Printf("len = %v, string:%v \n", len(arg2), arg2[0])

	//var coinA []*sdk.Coin
	//coinA = make([]*sdk.Coin, len(arg2))
	//for i, coin := range arg2 {
	//	decCoin, err := sdk.ParseDecCoin(coin)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	c, _ := sdk.NormalizeDecCoin(decCoin).TruncateDecimal()
	//	coinA[i] = &c
	//	fmt.Println("coinA[i]:", coinA[i])
	//}
	//
	//fmt.Println("coinA:", coinA)

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
