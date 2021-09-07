[![Quality And Tests](https://github.com/jmsilvadev/cycloid/actions/workflows/pull-requests.yml/badge.svg)](https://github.com/jmsilvadev/cycloid/actions/workflows/pull-requests.yml)
[![Release](https://github.com/jmsilvadev/cycloid/actions/workflows/release.yml/badge.svg?branch=master)](https://github.com/jmsilvadev/cycloid/actions/workflows/release.yml)

## Introduction

The main focus in the package design was to provide a clear solution with a consumer focus and a code with quality and estability, this means that it should provide a clear and simple solution to use. To achieve this, the SOLID standard of single responsibility was followed, all the business logic was encapsulated, having public visibility only of the methods that the consumer needs to use. To grant the quality was use tools of code style and automated tests with 100% of coverage and tests scenarios to prevent known and minimize unkown flaws.

The idea was to create packages to be reused instead of creating compiled binaries for it. With packages users can easily create their binaries in a simple way using `go build`.

## Important Note

I would like to inform you that I created explanatory videos with the step-by-step creation of Kata02, my goal is to show my logic and skills in a didactic way so that you can evaluate my profile with more information, not only technical, but also my soft skills.

I'd love it if you watched the videos, even at 2x speed, because I'm looking for a company that values ​​professionals and that always have attention to their opinions and efforts, and you watching the videos, even if I don't have the profile for the company, you will demonstrate that you are on the right path in terms of valuing the professionals who are there.

You can see the list of videos in this [VIDEOLOG](VIDEOS.md)

## CI/CD And SemVer

The project uses the Devops concept of continuous integration and continuous delivery through pipelines. To guarantee the CI there is a check inside the pipelines that checks the code quality and runs all the tests. This process generates an artifact that can be viewed and analyzed by developers.

To ensure the continuous delivery, the pipeline uses the automatic semantic versioning creating a versioned realease after each merge in the branch master. This release are tags in the control version system.

The semver system uses the angular commit message model [Angular Commit Message Format](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).

## Helper

To facilitate the developers' work, a Makefile with containerized commands was created.

```bash
➜  interview-accountapi~$ make help
build                          Build docker image
deps                           Install dependencies
doc                            Show package documentation
down                           Stop docker container
fmt                            Applies standard formatting
lint                           Checks Code Style
logs                           Watch docker log files
rebuild                        Rebuild docker container
ssh                            Interactive access to container
test.coverage                  Check project test coverage
test                           Run all available tests
up                             Start docker container in daemon mode
vendor                         Install vendor
vet                            Finds issues in code

```

## Packages

1. Kata02: Package katas02 is a library that contains methods to solve the Kata02 problem, [Karate Chop](http://codekata.com/kata/kata02-karate-chop/).

2. Katas02 Encapsulation: Package katas02withencapsulation is a library that contains methods to solve the Kata02 problem, [Karate Chop](http://codekata.com/kata/kata02-karate-chop/). This library uses concepts of OOP and SOLID to implement best practices of programming.

3. Katas19: Package katas19 is a library that contains methods to solve the Kata19 problem, [Word Chains](http://codekata.com/kata/kata19-word-chains/).

### Usage

#### Kata02

- Instantiate the type of binary search and call the Chop method:

```bash
pos := iterable.Chop(5, []int{1, 3, 5})
```

- Types of binary search:

```bash
iterable: pos := iterable.Chop(5, []int{1, 3, 5})
recursive: pos = recursive.Chop(5, []int{1, 3, 5})
recursivereference: pos = recursivereference.Chop(5, []int{1, 3, 5})
slicerecursive: pos = slicerecursive.Chop(5, []int{1, 3, 5})
```

#### Kata02 With Encapsulation

- Instantiate the object of binary search and call the Chop method:

```bash
k := kataswithinternals.NewIterable()
pos := k.Chop(5, []int{1, 3, 5})
```

- Types of binary search:

```bash
iterable: k := kataswithinternals.NewIterable()
recursive: k := kataswithinternals.NewRecursive()
recursivereference: k := kataswithinternals.NewRecursiveReference()
slicerecursive: k := kataswithinternals.NewRecursiveSlice()
```

#### Kata19

- Set the path of the dictionary to be used:

```bash
pathToDictionary, _ := filepath.Abs("../../../storage/wordlist.txt")
```

- Instantiate the wordschain package and call the Discover method with the subject, target and path to dictionary:

```bash
chain, _ := wordschain.Discover("the", "end", pathToDictionary)
```
