package rank

import (
	"math"
)

type StarItem struct {
	Count int
	Type  int
}

func maxStarType(starItems []StarItem) int {
	max := starItems[0].Type
	for _, item := range starItems {
		if max < item.Type {
			max = item.Type
		}
	}
	return max
}

func rangeOfStars(max int) []int {
	a := make([]int, max)
	for i := range a {
		a[i] = max
		max = max - 1
	}
	return a
}

func sumForCollection(items []StarItem) int {
	sum := 0
	for _, item := range items {
		sum += item.Count
	}
	return sum
}

func squareStars(items []StarItem) []StarItem {
	srd := make([]StarItem, len(items))
	for i, item := range items {
		var s2 = item.Type * item.Type
		srd[i] = StarItem{Count: item.Count, Type: s2}
	}
	return srd
}

func starPower(allStars []int, items []StarItem) float64 {
	sum := sumForCollection(items)
	maxNumStars := maxStarType(items)
	worth := 0.0
	for _, item := range items {
		worth += float64(item.Type*item.Count) / (float64(sum) + float64(maxNumStars))
	}
	return worth
}

/**
* Given a collections of Stars, returns
* For example if an item has 60 5 stars, 80 4 stars, 75 3 stars, 20 2 stars, and 25 1 stars what is its actual, overall, rating?
* This technique implements Bayesian approximation to come up with that score.  Outlined here http://www.evanmiller.org/ranking-items-with-star-ratings.html
**/
func AvgStars(items []StarItem) float64 {
	z := 1.65
	var ms = maxStarType(items)
	var sc = sumForCollection(items)
	var starsSqd = squareStars(items)
	var starRange = rangeOfStars(ms)
	var sp1 = starPower(starRange, items)
	var sp2 = starPower(starRange, starsSqd)
	var x = (sp2 - (sp1 * sp1)) / float64(sc+ms+1)
	var p = sp1 - z*math.Sqrt(x)
	return p

}

/*func main() {
  tmp := []StarItem{StarItem{Count:60, Type:5}}
  tmp = append(tmp, StarItem{Count:80, Type:4})
  tmp = append(tmp, StarItem{Count:75, Type:3})
  tmp = append(tmp, StarItem{Count:20, Type:2})
  tmp = append(tmp, StarItem{Count:25, Type:1})
  z := 1.65
  var ms = maxStarType(tmp)
  var sc = sumForCollection(tmp)
  var starsSqd = squareStars(tmp)
  var starRange = rangeOfStars(ms)
  fmt.Println(ms)
  fmt.Println(sc)
  fmt.Println(starsSqd)
  fmt.Println(starRange)
  var sp1 = starPower(starRange, tmp)
  var sp2 = starPower(starRange, starsSqd)
  fmt.Println(sp1)
  fmt.Println(sp2)
  //return fsns - z*math.sqrt((f(s2, ns)- fsns**2)/(N+K+1))
  //r         - z*sqrt((sp2- fsns**2)/(N+K+1))
  var x = (sp2 - (sp1 * sp1))/float64(sc+ms+1)
  var p = sp1 - z*math.Sqrt(x)
  fmt.Println(p)
}*/
