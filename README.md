The Shining (1980) - Opinion on Code Quality

- Barely any validation, data types alone do not suffice.
- The application now uses the HTTPS OMDb API instead of the HTTP version. This improves security with the help of encryption.
- All functions are in proper packages now, maybe the code could be split up even further, but that would've been over engineering in my opinion.
- The QueryDatabase method has been refactored, so now it is more readable.
- The architecture of the application is lacking because no effort was put into researching the applicable design patterns. Even though some thought has still gone into the project file structure, there is still room for improvement.
- Concurrency is used when fetching the movie plots, the concurrent design allows for parallelism, which allows the plots of all movies to be added within a minute.
- The code was built using an iterative process. Starting with a CLI application, adding a REST API later, then using goroutines with an external API, and finally adding a frontend to it all. This process helped me to focus on the crucial parts first, and only changing these when required by the addition of new features.
- However, the complexity of my code is far from perfect. This could be improved by refactoring large functions into smaller ones, reducing the cyclomatic complexity of each function.
