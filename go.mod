module github.com/scottfones/goqueens

require (
    game v1.0.0
	web v1.0.0
)

replace (
	game v1.0.0 => ./game
	web v1.0.0 => ./web
)