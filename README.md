[![Quality And Tests](https://github.com/jmsilvadev/cycloid/actions/workflows/pull-requests.yml/badge.svg?branch=master)](https://github.com/jmsilvadev/cycloid/actions/workflows/pull-requests.yml)
[![Release](https://github.com/jmsilvadev/cycloid/actions/workflows/release.yml/badge.svg?branch=master)](https://github.com/jmsilvadev/cycloid/actions/workflows/release.yml)

## Introduction

The main focus in the package design was to provide a clear solution with a consumer focus and a code with quality and estability, this means that it should provide a clear and simple solution to use. To achieve this, the SOLID standard of single responsibility was followed, all the business logic was encapsulated, having public visibility only of the methods that the consumer needs to use. To grant the quality was use tools of code style and automated tests with 100% of coverage and tests scenarios to prevent known and minimize unkown flaws.

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

In progress

### Usage

In progress
