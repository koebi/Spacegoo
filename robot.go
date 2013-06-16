package main

import (
	. "github.com/Merovius/spacegoo"
	"github.com/Merovius/spacegoo/boilerplate"
)

type RoBot struct{}

func (bot *RoBot) Move(state GameState) (Move, error) {
	var mod int
    if state.PlayerName(They) == "sendsomewhere" {
        mod = 3
	} else {
        mod = 2
	}
	if state.Round%mod != 0 {
		for _, myp := range state.MyPlanets() {
			min := float64(9999999)
			meiner := myp
			for _, myp2 := range state.MyPlanets() {
				if myp2.Ships.Sum() > myp.Ships.Sum() {
					meiner = myp2
				}
			}
			deiner := Planet{}
			for _, yourp := range state.NeutralPlanets() {
				if myp.Dist(float64(yourp.X), float64(yourp.Y)) < min {
					min = myp.Dist(float64(yourp.X), float64(yourp.Y))
					deiner = yourp
				}
			}
			return state.Send(meiner, deiner, meiner.Ships), nil
		}
		return state.Nop(), nil
	} else {
		for _, myp := range state.MyPlanets() {
			min := float64(9999999)
			meiner := myp
			for _, myp2 := range state.MyPlanets() {
				if myp2.Ships.Sum() > myp.Ships.Sum() {
					meiner = myp2
				}
			}
			deiner := Planet{}
			for _, yourp := range state.TheirPlanets() {
				if myp.Dist(float64(yourp.X), float64(yourp.Y)) < min {
					min = myp.Dist(float64(yourp.X), float64(yourp.Y))
					deiner = yourp
				}
			}
			return state.Send(meiner, deiner, meiner.Ships), nil
		}
		return state.Nop(), nil
	}

}

/*
    for _, myp := range state.MyPlanets() {
		min := float64(99999)
		deiner := Planet{}
		meiner := Planet{}
		for _, myp2 := range state.MyPlanets() {
			for i, _ := range myp2.Ships {
				if myp2.Ships[i] > myp.Ships[v] {
					meiner = myp2
				}
			}
		}
		for _, yourp := range state.NotMyPlanets() {
			if myp.Dist(float64(yourp.X), float64(yourp.Y)) < min {
				min = myp.Dist(float64(yourp.X), float64(yourp.Y))
				deiner = yourp
			}
		}
		return state.Send(meiner, deiner, meiner.Ships), nil
	}
	return state.Nop(), nil
}
*/
func main() {
	boilerplate.Run(&RoBot{})
}
