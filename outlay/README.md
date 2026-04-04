# outlay

[![Go Reference](https://pkg.go.dev/badge/github.com/nanorele/gio-x/outlay.svg)](https://pkg.go.dev/github.com/nanorele/gio-x/outlay)

This package provides extra layouts for [gio](https://github.com/nanorele/gio).

## State

This package has no stable API, and should always be locked to a particular commit with
go modules.

## Layouts

### Grid

This layout allows placement of many items in a grid with to several different strategies for wrapping across lines. For examples, run:

### Radial

The radial layout allows you to lay out a set of widgets along an arc. The width and oritentation of the arc are configurable to allow for everything from a hand of cards to a full circle of widgets.

Known issues:
- The radial layout does not currently return correct dimensions for itself, which breaks most attempts to use it as part of a larger layout.

