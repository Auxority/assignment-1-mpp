Opinion on Code Quality

- No error thrown when a movie is added with a duplicate IMDb id.
- Barely any validation, data types alone do not suffice.


<!-- # Apollo 13 - Opinion on Code Quality -->




<!-- The code quality has improved and degraded simultaneously this time around.

And here's why:
- I now use interfaces to prevent code duplication.
- I make use of custom errors to make debugging the code easier, following Go's best practices.
- I now only use the `CheckError(err)` *log.Fatalln(err)* method in my `main.go` file and in route handlers (since route handlers are not allowed to return errors).
- The `util.go` file in the `omdb` package should be moved somewhere else, since it is not specifically built for omdb. I just haven't been able to figure out a better package name yet.
- The naming of my functions has degraded, since I am struggling to find good names for functions that are similar.
- I now avoid SQL injection (https://go.dev/doc/database/sql-injection) by NOT using `fmt.Sprintf` to construct my SQL queries.
- I limit the number of goroutines to prevent running into rate limits. (https://stackoverflow.com/questions/25306073/always-have-x-number-of-goroutines-running-at-any-time)
- The endpoints now all make use of commands, instead of relying on their own database logic (e.g. The API router no longer contains database code).
- A lot of duplicate code has been removed.
- The API now uses separate files for the different routes, and I've made use of a routes.go file to define my routes.
- I'm no longer using Gin, since it was quite a large package. I've decided to use Gorilla Mux instead. (I first tried doing it with the net/http package and a switch statement, but in my opinion, this dependency makes my code cleaner)
- The QueryDatabase and ExecDatabase methods are quite large, and should be refactored into smaller methods.
- The `summaries.go` file uses concurrency, but the function is too large and should be refactored further to keep the code readable & maintainable.
- Ideally I would limit the number of if statements per function to one. But in many cases functions contain way more than that, increasing my code complexity.
- I should also limit the number of methods called from a function to two or three, depending on whether an if statement was used or not.

> *And again, it still works ;)* -->
