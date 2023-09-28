# Dirty Harry - Opinion on Code Quality

In my opinion the code quality of this project is quite a bit better than during the last assignment, but it is still far from "perfect".

And here's why:
- Most errors are still generic, since I did not spend the time to fix this.
- Variable names, and the command handler have been improved.
- The `main.go` file has been split up into components to improve readability.
- All the separate files are still stored in one directory, since [Codegrade](https://app.codegra.de/) does not allow folders containing Go code to be uploaded.
- Not all endpoints use commands yet. Most commands must be rewritten to achieve this.
- The API router contains database logic.
- The API uses separate query methods from the `QueryDatabase` method because my `QueryDatabase` function does not return anything right now. I should use interfaces to allow this method to return anything I specify.
- All methods that are accessed by other files start with a capital letter, following Go's best practices.
- And I should split the handlers of the API up into multiple files. Instead of one `movies_router.go` file. Because this will make it easier to adjust such endpoints in the future and prevent one huge file if more endpoints have to be added. 
- I still haven't written any documentation, except for a few methods that shouldn't be there after another refactoring round. So it might still be difficult for other devs to understand my code.

> *But it still works ;)*
