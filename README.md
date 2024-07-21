# github.com/dozer111/projectlinter-core

The project is a part of [projectlinter cli app](https://github.com/dozer111/projectlinter-cli) solution

**IMPORTANT - READ this before read below https://github.com/dozer111/projectlinter-cli**

---

It contains all the main abstractions and rules from which you can create your own projectlinter

The rules are divided into separate submodules. This was made to push the users to use only the code which they need.

Download all the needed modules, configure it by your needs - profit.

---

The main submodule `github.com/dozer111/projectlinter-core` has no rules, only the main important abstractions and some helper code to make your interaction with `projectlinter` as easy as possible

---

The main things you need to know - 

- [rule - the most important thing in all the project](https://github.com/dozer111/projectlinter-core/blob/master/rules/rule.go)
- [rule_set.go - other important which are core for cli app](https://github.com/dozer111/projectlinter-core/blob/master/rules/rule_set.go)
- [util directory](https://github.com/dozer111/projectlinter-core/tree/master/util) - a reusable code for core and cli
- [printer director](https://github.com/dozer111/projectlinter-core/tree/master/printer) - a reusable code for failed messages

## Can you help the project?

Yes you can. This project was conceived as a multilanguage and multiConfiguration(something that is not programming language related but still a configuration) solution

- You can share at least your needs and ideas.(because now it helps you with a little number of possible problems)
- You can write your rules and ruleSets

Let's make your configuration codebase actual again)

