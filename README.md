# Jerry Maguire (1996) - Opinion on Code Quality

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

The total `342` lines of code are within the range of `0 - 66,000` lines. This is the `++` range of the Software Improvement Group (SIG).

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

Every unit size is below 30. Which means that the code is ranked into the `Simple` category. This is equivalent to the `++` category.

## Cyclomatic Complexity
I also counted the cyclomatic complexity manually. Each function/unit starts with a cyclomatic complexity of 1. Every if (else), case of a switch statement, and every return statement will increase this number by 1. I included return statements in my calculation because I make use of early returns in my code.

```
api\command\add.go AddAndShowMovie 6
api\command\add.go AddMovie 4
api\command\add.go createAddCommand 2

api\command\delete.go DeleteMovie 4
api\command\delete.go ShowMovieDeletion 4
api\command\delete.go createDeleteCommand 2

api\command\details.go ShowMovieDetails 4
api\command\details.go GetMovieDetails 4
api\command\details.go createDetailsCommand 2
api\command\details.go getMovieFromRow 4

api\command\handler.go parseArguments 1
api\command\handler.go HandleCommand 7

api\command\list.go ShowMovieList 5
api\command\list.go GetMovieList 4

api\command\summaries.go ShowMovieSummaries 6
api\command\summaries.go showMovieSummary 2
api\command\summaries.go addMovieSummary 8

api\command\util.go CreateNewCommand 2
api\command\util.go CreateImdbIdParameter 2

api\database\database.go OpenDatabase 4
api\database\database.go ExecDatabase 6
api\database\database.go QueryDatabase 8
api\database\database.go getResults 6

api\json_util\read.go ReadJSONRequest 6
api\json_util\read.go ParseInteger 6
api\json_util\read.go ParseFloat 6
api\json_util\read.go validateJSON 4
api\json_util\read.go decodeJSON 4

api\json_util\write.go WriteJSONResponse 4

api\omdb\api.go GetMovieDetails 7
api\omdb\api.go buildUrl 2
api\omdb\api.go getRequest 6

api\router\add.go AddMovie 4
api\router\add.go searchMovie 7

api\router\delete.go DeleteMovie 1

api\router\details.go MovieDetails 2
api\router\details.go writeMovieDetailsResponse 6

api\router\list.go ListMovies 1

api\router\routes.go StartAPI 1
api\router\routes.go registerRoutes 1
api\router\routes.go getAddress 3

api\router\util.go GetUrlId 2
api\router\util.go GetUrlParameter 2

error_util\error.go CheckError 2

main.go main 3
```

The cyclomatic complexity of all units is below 9, which falls within the `1-10` range, ranking it into the `simple, without much risk` risk level. This means that `0%` of the units fall into the `moderate`, `high`, or `very high risk` categories. This gives the code a `++` ranking.

## Duplication
According to the SIG, duplicated code should be 6 consecutive lines with an exact match when those lines are trimmed and concatenated.
Using this method I was unable to find any code duplication in the code base. To make sure I didn't overlook anything, I also tried to use PMD's copy-paste detector, as suggested by the assignment. But it was unable to find any code duplication (as defined in the paper) as well.

This would rank the code in the `0-3%` code duplication range, giving it the `++` rank.

## Final score
Based on all the rankins, the final score can be calculated. To achieve this, `++` will be converted to a 5, `+` to a 4, `o` to a 3, `-` to a 2, and `--` to a 1. After this is done, the average is taken to calculate the final score.

For my code, every indicator scored in the `++` category. Meaning that the final score of my code, according to the SIG model, is a 5.

## Comparison with Sonarqube
After I calculated all my scores manually, I compared the results to Sonarqube.
Sonarqube gave me the following results:
```
Reliability
- Bugs: 0
- Rating: A

Maintainability
- Code Smells: 0
- Debt: 0
- Debt Ratio: 0.0%
- Rating A

Duplications
- Density: 0.0%
- Duplicated Lines: 0
- Duplicated Blocks: 0
- Duplicated Files: 0

Size
- Lines of Code: 716
- Lines: 910
- Comment Lines: 4
- Comments: 0.6%

Complexity
- Cyclomatic Complexity: 128
- Cognitive Complexity: 66
```

These results did not differ in terms of quality assessment compared to the SIG Model. However they did use a different strategy to calculate the Lines of Code. And sometimes the Cyclomatic Complexity was higher than my own calculations. Which might have moved some units into the moderate risk category. Maybe my approach to calculate the Cyclomatic Complexity was wrong.

## Conclusion
Overall, the code quality of my application is pretty good. Although some comments might make it even easier to understand the application. And some units could be improved by decreasing their Cyclomatic Complexity. But it's sufficient for a college project.
