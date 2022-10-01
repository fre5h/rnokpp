package internal

import "time"

const BaseYear = 1900

var BaseLocation, _ = time.LoadLocation("Europe/Kiev")
var BaseDate = time.Date(1900, 1, 1, 0, 0, 0, 0, BaseLocation)
