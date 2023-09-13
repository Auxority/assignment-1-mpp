# Mission Impossible - Opinion on Code Quality

In my opinion the code quality of this project is bad.

And here's why:
The functions and variable names might be clear, but there's still room for improvement. There's a lot of code duplication in the project. The main function has way too many responsibilities, and should be refactored into multiple smaller methods. The code inside the `rows.Next()` for loops should be refactored into functions, to improve readability. Errors are printed when they occur, but not handled in any other way. This could be improved by creating custom error messages for the end user. The switch statement works for now, but is not built for expandability. If the project would need 200 commands, the switch statement would look like a huge mess. The `main.go` file contains all logic of the application, it would improve readability if the `sqlite` and `command` related code had their own Go files. I have not documented any line of code, or any method, which makes it harder for others to understand my code.

<details>
    <summary>Spoiler</summary>

    But it works ;)
</details>
