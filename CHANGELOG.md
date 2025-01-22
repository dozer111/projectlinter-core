## 1.1.0 (on development)

- set the concrete projectlinter dependency version in all the submodules [commit](https://github.com/dozer111/projectlinter-core/commit/2264e64b3d47c86d36fda1308bbb4b3d7476ae3d)

## Dependency

- (chore) add json-schema autodetect in full-example files. [pull request](https://github.com/dozer111/projectlinter-core/pull/18)
- (feat) add support js npm [pull request](https://github.com/dozer111/projectlinter-core/pull/19)
- (bug) fix golang bump/substitute rules. Now they are checking only direct dependencies [pull request](https://github.com/dozer111/projectlinter-core/pull/20)
- (feat) extend example block. Add new field "description" [pull request](https://github.com/dozer111/projectlinter-core/pull/21)


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