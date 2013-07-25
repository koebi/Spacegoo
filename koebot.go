package main

import (
	. "github.com/Merovius/spacegoo"
	"github.com/Merovius/spacegoo/boilerplate"
)

type KoeBot struct {
	myHomePlanet    Planet
	enemyHomePlanet Planet
	myPlanetList    Planets
}

func (bot *KoeBot) Move(state GameState) Move {
	//conquering inital planets

	//set home planets and create planet-list (by distance)
	if state.Round == 0 {
		for _, myp := range state.MyPlanets() {
			bot.myHomePlanet = myp
		}
		for _, theirp := range state.TheirPlanets() {
			bot.enemyHomePlanet = theirp
		}
		bot.myPlanetList = state.NeutralPlanets().SortByDist(bot.myHomePlanet.X, bot.myHomePlanet.Y)
		var PlanetList Planets
		for i, p := range bot.myPlanetList {
			if p.Dist(bot.enemyHomePlanet.X, bot.enemyHomePlanet.Y) > p.Dist(bot.myHomePlanet.X, bot.myHomePlanet.Y) {
				PlanetList = append(PlanetList, bot.myPlanetList[i])
			}
		}
		bot.myPlanetList = PlanetList
		return Nop{}
	}

	// calculate rounds to conquer nearest planets
	rounds := (len(state.NeutralPlanets()) - 1) / 2

	// conquer initial planets
	var conquer Planet
	var oneShip Ships
	if len(bot.myPlanetList) < rounds {
		rounds = len(bot.myPlanetList)
	}
	if state.Round <= rounds {
		index := state.Round - 1
		conquer = bot.myPlanetList[index]
		oneShip = Ships{5, 5, 5}
		shipFleet := conquer.Ships.Add(oneShip)
		return Send{bot.myHomePlanet, conquer, shipFleet}
	}

	// using robot strategy

	if state.Round > rounds {
		for _, myp := range state.MyPlanets() {
			min := 9999999
			meiner := myp
			for _, myp2 := range state.MyPlanets() {
				if myp2.Ships.Sum() > myp.Ships.Sum() {
					meiner = myp2
				}
			}
			var deiner Planet
			for _, yourp := range state.NotMyPlanets() {
				if myp.Dist(yourp.X, yourp.Y) < min {
					min = myp.Dist(yourp.X, yourp.Y)
					deiner = yourp
				}
			}
			return Send{meiner, deiner, meiner.Ships}
		}
	}
	return Nop{}

}

func main() {
	boilerplate.Run(&KoeBot{})
}
