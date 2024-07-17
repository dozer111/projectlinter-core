# bump dependency submodule

One of killer features

The main idea is: we configure a number of yaml configurations in which we say something like "the library \<libraryname> must be replaced to \<anotherLibrary>".

Then, we use this configs in our projects and check does we need to substitute those dependencies

---

The module contains all the necessary code

- own substitute json-schema for yaml configurations
- a parser for your yaml configurations
- go_build_generator - the code with which you can use go:generate to conveniently collect all configurations in one place and connect them to your program
- actual rules for all the currently supported languages



















