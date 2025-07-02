## 1.1.0 2025-07-02

- (DX) 💥 get rid the idea of "separate modules". It is very hard to interact with it. Try to simplify and remove it

## Dependency

- (feat) 🎉 add support js npm [pull request](https://github.com/dozer111/projectlinter-core/pull/19)
- (feat) extend example block. Add new field "description" [pull request](https://github.com/dozer111/projectlinter-core/pull/21)

- (chore) add json-schema autodetect in full-example files. [pull request](https://github.com/dozer111/projectlinter-core/pull/18)

- (bug) fix golang bump/substitute rules. Now they are checking only direct dependencies [pull request](https://github.com/dozer111/projectlinter-core/pull/20)


## Golang

- (chore) extend dependency struct. Now it has also method `IsIndirect() bool`

## PHP

### improved rules

- **composer**: improve rule "section is correct": [issue](https://github.com/dozer111/projectlinter-core/issues/15)
  - https://github.com/dozer111/projectlinter-core/pull/16
  - https://github.com/dozer111/projectlinter-core/pull/17
  - https://github.com/dozer111/projectlinter-core/pull/25
- **composer**: improve rule "section exists"
  - https://github.com/dozer111/projectlinter-core/pull/25
- **composer**: improve rule "section is absent"
  - https://github.com/dozer111/projectlinter-core/pull/25
- **composer**: ⚠ refactor rule "dependencies constaints are valid"
  - https://github.com/dozer111/projectlinter-core/pull/26
  - https://github.com/dozer111/projectlinter-core/pull/27


### parser

- become to parse section "config.bump-after-update" 
  - https://github.com/dozer111/projectlinter-core/pull/22
  - https://github.com/dozer111/projectlinter-core/pull/23
  - https://github.com/dozer111/projectlinter-core/pull/24