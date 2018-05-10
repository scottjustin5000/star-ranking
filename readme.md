# Star Ranking

> Bayesian approximation for determining an average star rating

For example if an item has 50 5 stars, 90 4 stars, 60 3 stars, 15 2 stars, and 23 1 stars what is its actual, overall, average rating?
The approach implemented here is taken from the [following description](http://www.evanmiller.org/ranking-items-with-star-ratings.html)

### Usage

```go
  stars := []StarItem{StarItem{Count: 50, Type: 5}}
  stars = append(stars, StarItem{Count: 90, Type: 4})
  stars = append(stars, StarItem{Count: 60, Type: 3})
  stars = append(stars, StarItem{Count: 15, Type: 2})
  stars = append(stars, StarItem{Count: 23, Type: 1})

  avg := AvgStars(stars)
  //avg = 3.3899443291304214

```

