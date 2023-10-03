package main

// // TODO: Find a better way that makes use of interfaces in the queryDatabase function to query the database.
// // Since this would prevent a lot of code duplication.
// func queryMovies(movies *[]*types.Movie) {
// 	moviesDatabase := database.OpenMoviesDatabase()
// 	defer database.CloseMoviesDatabase(moviesDatabase)

// 	rows, err := moviesDatabase.Query("SELECT IMDb_id, Title, Rating, Year FROM movies;")
// 	error_util.CheckError(err)
// 	defer rows.Close()

// 	for rows.Next() {
// 		var movie types.Movie
// 		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
// 		*movies = append(*movies, &movie)
// 	}

// 	defer rows.Close()
// }

// // TODO: IDEM, ideally the commands would all return structs that are serialized to JSON.
// // But this requires a refactor, while I just want to get everything working for now.
// func queryMovie(id *string) *types.Movie {
// 	moviesDatabase := database.OpenMoviesDatabase()
// 	defer database.CloseMoviesDatabase(moviesDatabase)

// 	sql := fmt.Sprintf(`
// 		SELECT IMDb_id, Title, Rating, Year
// 		FROM movies
// 		WHERE IMDb_id='%s';
// 	`, *id)

// 	rows, err := moviesDatabase.Query(sql)
// 	error_util.CheckError(err)
// 	defer rows.Close()

// 	var movie types.Movie
// 	for rows.Next() {
// 		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
// 	}

// 	defer rows.Close()

// 	return &movie
// }
