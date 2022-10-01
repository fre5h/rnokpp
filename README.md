# RNOKPP (РНОКПП)

Helper functions to work with Ukrainian registration number of the taxpayer's account card (RNOKPP).  
Допоміжні функції для роботи з реєстраційним номером облікової картки платника податків (РНОКПП).

[![Build Status](https://img.shields.io/github/workflow/status/fre5h/rnokpp/CI/main?style=flat-square)](https://github.com/fre5h/rnokpp/actions?query=workflow%3ACI+branch%3Amain+)
[![CodeCov](https://img.shields.io/codecov/c/github/fre5h/rnokpp.svg?style=flat-square)](https://codecov.io/github/fre5h/rnokpp)
[![Go Report Card](https://goreportcard.com/badge/github.com/fre5h/rnokpp?style=flat-square)](https://goreportcard.com/report/github.com/fre5h/rnokpp)
[![License](https://img.shields.io/github/license/fre5h/rnokpp?style=flat-square)](https://pkg.go.dev/github.com/fre5h/rnokpp)
[![Gitter](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/fre5h/rnokpp)
[![GoDoc](https://pkg.go.dev/badge/github.com/fre5h/rnokpp)](https://pkg.go.dev/github.com/fre5h/rnokpp)


## Requirements 🧐

* GO >= 1.18

## Features 🎁

- [x] Get details about RNOKPP
- [x] Get gender
- [x] Check gender
- [x] Check validity
- [x] Generate RNOKPP by date and gender
- [x] Generate random RNOKPP

## Using 👨‍🎓

###### main.go

```go
package main

import (
    "fmt"
    "time"

    "github.com/fre5h/rnokpp"
)

func main() {
    // Get details about RNOKPP
    details, _ := rnokpp.GetDetails("3652504575")
    fmt.Println("details:", details) // valid, male, 01.01.2000

    // Get gender from RNOKPP
    gender1, _ := rnokpp.GetGender("3652504575")
    fmt.Println("gender1:", gender1) // male
    gender2, _ := rnokpp.GetGender("3068208400")
    fmt.Println("gender2:", gender2) // female

    // Check gender
    isMale, _ := rnokpp.IsMale("3652504575")
    fmt.Println("is male:", isMale) // true
    isFemale, _ := rnokpp.IsFemale("3652504575")
    fmt.Println("is female:", isFemale) // false

    // Check valid RNOKPP
    validRnokpp := rnokpp.IsValid("3652504575")
    invalidRnokpp := rnokpp.IsValid("1234567890")
    fmt.Println("rnokpp valid:", validRnokpp, invalidRnokpp) // true, false

    // Generate RNOKPP by date and gender
    birthday, _ := time.Parse("02.01.2006", "01.01.2000")
    generatedRnokppMale, _ := rnokpp.GenerateRnokpp(birthday, rnokpp.Male)
    fmt.Println("valid RNOKPP for male with birthday on 01.01.2000:", generatedRnokppMale) // valid RNOKPP for male with birthday on 01.01.2000, e.g. 3652322032
    generatedRnokppFemale, _ := rnokpp.GenerateRnokpp(birthday, rnokpp.Female)
    fmt.Println("valid RNOKPP for female with birthday on 01.01.2000:", generatedRnokppFemale) // valid RNOKPP for female with birthday on 01.01.2000, e.g. 3652347000

    // Generate a random RNOKPP
    generatedRandomRnokpp, _ := rnokpp.GenerateRandomRnokpp()
    fmt.Println("random rnokpp:", generatedRandomRnokpp) // e.g. random rnokpp: 3300507061
}
```

###### bash

```text
$ go get "github.com/fre5h/rnokpp"@v0.1.0
go: downloading github.com/fre5h/rnokpp v0.1.0
go: added github.com/fre5h/rnokpp v0.1.0

$ go run main.go
details: valid, male, 01.01.2000
gender1: male
gender2: female
is male: true
is female: false
rnokpp valid: true false
valid RNOKPP for male with birthday on 01.01.2000: 3652322032
valid RNOKPP for female with birthday on 01.01.2000: 3652347000
random rnokpp: 3300507061
```

## Contributing 🤝

See [CONTRIBUTING](https://github.com/fre5h/rnokpp/blob/master/.github/CONTRIBUTING.md) file.
