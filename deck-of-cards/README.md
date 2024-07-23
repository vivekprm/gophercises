# To make printing struct easier
We have below types defined in this case:

```go
type Suit uint8
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}
```

Now if we want to print cards in readable fashion. We can use below package.
https://pkg.go.dev/golang.org/x/tools/cmd/stringer

Running below command will generate the String implementation for Card type:
```sh 
stringer -type=Card
```

Or we can add a generator as below in cards.go file
```
//go:generate stringer -type=Suit,Rank
```

And then run ```go generate``` command to generate source files.

We can write below function to create a new deck of cards.

```go
func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank:rank})
		}
	}
	return cards
}
```

Whenever we generate deck of card, we will have some set of options. We might want to add joker to the deck, we might want to shuffle our deck, we might want them to be sorted, may be sort them by custom values etc. There might be many sorts of things that you might want to do while creating a new deck of card.

So we are going to look at how we can add these options to our New function. Specifically we want any variable sets of options and have them all run properly. 

One way we could do that is by creating a new struct of Options and pass that to New.

```go
type NewOpts struct {
    ...
    Shuffle bool
    ...
}
```

The problem with this approach is that it's little bit unclear, you almost need pointer to every options, so that you can say if option1 is nil do this, ooption2 is nil do that and so on.

Instead we are going to use **functional options**. Look at this (talk)[https://www.youtube.com/watch?v=6buaPyJ0XeQ]. So in this case we are going to take variadic function options as below:

```go
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank:rank})
		}
	}
	return cards
}
```

Whenever we generate deck of card, we will have some set of options. We might want to add joker to the deck, we might want to shuffle our deck, we might want them to be sorted, may be sort them by custom values etc. There might be many sorts of things that you might want to do while creating a new deck of card.

So we are going to look at how we can add these options to our New function. Specifically we want any variable sets of options and have them all run properly. 

One way we could do that is by creating a new struct of Options and pass that to New.

```go
type NewOpts struct {
    ...
    Shuffle bool
    ...
}
```

The problem with this approach is that it's little bit unclear, you almost need pointer to every options, so that you can say if option1 is nil do this, ooption2 is nil do that and so on.

Instead we are going to use **functional options**. Look at this (talk)[https://www.youtube.com/watch?v=6buaPyJ0XeQ]. So in this case we are going to take variadic function options as below:

```go
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank:rank})
		}
	}
    for _, opt := range opts {
        cards = opt(cards)
    }
	return cards
}
```

# Ranking the cards for sorting purpose
```go
func absRank(c Card) int {
    return int(c.Rank) * int(maxRank) + int(c.Rank)
}
```

# Shuffling
Below is the function to shuffle the deck. 

```go
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(uint64(time.Now().Unix())))

	for i, j := range r.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}
```

But how can we test this function which uses random functions.
One way to do that is, instead of creating random source in the implementation, we pass the source and modify the Shuffle function as below.

```go
func Shuffle(s rand.Source) func(cards []Card) []Card {
    return func(cards []Card) []Card {
        r := rand.New(s)
        ret := make([]Card, len(cards))
        for i, j := range r.Perm(len(cards)) {
            ret[i] = cards[j]
        }
        return ret
    }
}
```

Problem with this approach is client need to know how to create Source or Permer in case we pass Permer. So instead we are going to keep Shuffle as it is and use a global var for source as below:

```go
var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))

	for i, j := range shuffleRand.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}
```

However, it's not ideal solution to use package level variables, however in this case it's not exprted and it's not being used anywhere else, so it might be ok..

Now to test it we can use our own seed as below:

```go
func TestShuffle(t *testing.T) {
    // It always generate the same random slice as source is same
    // First call to shuffleRand.Perm(52) would be [40, 35, ...]
    shuffleRand = rand.New(rand.NewSource(0))
    orig := New()
    first := orig[40]
    second := orig[35]
    cards := New(Shuffle)

    if cards[0] != first {
        t.Errorf("Expected the first card to be %s, received: %s", first, cards[0])
    }

    if cards[1] != second {
        t.Errorf("Expected the second card to be %s, received: %s", second, cards[1])
    }

}
```
