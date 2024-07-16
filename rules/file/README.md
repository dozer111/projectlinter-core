# github.com/dozer111/projectlinter-core/rules/file

```bash
go get -u github.com/dozer111/projectlinter-core/rules/file
```

The module that contains all the rules related to real physical files

| Function        | Description                                                                                  |
|-----------------|----------------------------------------------------------------------------------------------|
| `FileExists`    | Checks if the specified file is present in the project                                       |
| `FileIsAbsent`  | Checks if the project does not have the specified file                                       |
| `FilesAreSame`  | Checks that the file in the project is the same as the given reference file                  |
| `RenameFile`    | Checks if the specified file is in the project. If it is, suggests renaming it               |
| `SubstituteFile`| Checks if the specified file is in the project. If it is, suggests replacing it with another |

