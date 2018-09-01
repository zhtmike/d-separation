# Go D-Seperation

[![GoDoc](https://godoc.org/github.com/zhtmike/d-separation?status.svg)](https://godoc.org/github.com/zhtmike/d-separation)
[![Build Status](https://travis-ci.org/zhtmike/d-separation.svg?branch=master)](https://travis-ci.org/zhtmike/d-separation)

An algorithm implementation for finding all D-separated nodes in a belief network

Reference: Daphne Koller and Nir Friedman, *Probabilistic Graphical Models Principles and Techniques*, p74-75.

## Examples

For a simple belief network, e.g. 1 --> 0 <-- 2, node-1 and node-2 is conditionally independent when node-0 is not observed.
