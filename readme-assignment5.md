## Volume
I wrote a very simple script in Node.JS that is capable of counting the Lines of Code according to some criteria:
1) Type structs are skipped, I don't see this type declaration as part of the code.
2) Empty lines are removed, since these do not contain any code.
3) Comment lines are removed, since these do not contain any code. I define comment lines as trimmed lines of code that start with `//` or `/*`.
4) Import lines are ignored, because they do not contain the core logic of the code. I define import lines as trimmed lines starting with one of the following: `import`, `package`, `""`, `_`.
5) Lines with closing brackets are ignored. I define these as trimmed lines that start with `}` or `)`.

The output of this Node.JS script is shown here:
```
api\command\add.go 30
api\command\delete.go 17
api\command\details.go 24
api\command\handler.go 23
api\command\list.go 17
api\command\summaries.go 36
api\command\util.go 4
api\database\database.go 39
api\json_util\read.go 40
api\json_util\write.go 6
api\omdb\api.go 20
api\router\add.go 27
api\router\delete.go 5
api\router\details.go 14
api\router\list.go 5
api\router\routes.go 21
api\router\util.go 4
api\types\movie.go 0
api\types\movie_details.go 0
error_util\error.go 3
main.go 7
Total lines of code in 21 .go files: 342
```

The total `342` lines of code are within the range of `0 - 66,000` lines. This is the ++ range of the Software Improvement Group (SIG).

## Unit size
To calculate the unit size, I followed the same approach I took during the calculation of the Volume of the code.
But instead of counting it for each file, I will count it for each method of each file. Excluding the function declaration this time. I did decide to do this manually, since it would save time.
The results are shown here:
```
api\command\add.go AddAndShowMovie 8
api\command\add.go AddMovie 5
api\command\add.go createAddCommand 8

api\command\delete.go DeleteMovie 5
api\command\delete.go ShowMovieDeletion 5
api\command\delete.go createDeleteCommand 4

api\command\details.go ShowMovieDetails 5
api\command\details.go GetMovieDetails 6
api\command\details.go createDetailsCommand 4
api\command\details.go getMovieFromRow 5

api\command\handler.go parseArguments 2
api\command\handler.go HandleCommand 19

api\command\list.go ShowMovieList 6
api\command\list.go GetMovieList 9

api\command\summaries.go ShowMovieSummaries 18
api\command\summaries.go showMovieSummary 5
api\command\summaries.go addMovieSummary 10

api\command\util.go CreateNewCommand 1
api\command\util.go CreateImdbIdParameter 1

api\database\database.go OpenDatabase 4
api\database\database.go ExecDatabase 8
api\database\database.go QueryDatabase 12
api\database\database.go getResults 9

api\json_util\read.go ReadJSONRequest 7
api\json_util\read.go ParseInteger 10
api\json_util\read.go ParseFloat 10
api\json_util\read.go validateJSON 4
api\json_util\read.go decodeJSON 4

api\json_util\write.go WriteJSONResponse 5

api\omdb\api.go GetMovieDetails 8
api\omdb\api.go buildUrl 1
api\omdb\api.go getRequest 7

api\router\add.go AddMovie 12
api\router\add.go searchMovie 12

api\router\delete.go DeleteMovie 4

api\router\details.go MovieDetails 5
api\router\details.go writeMovieDetailsResponse 7

api\router\list.go ListMovies 4

api\router\routes.go StartAPI 3
api\router\routes.go registerRoutes 7
api\router\routes.go getAddress 5

api\router\util.go GetUrlId 1
api\router\util.go GetUrlParameter 1

error_util\error.go CheckError 2

main.go main 6
```
