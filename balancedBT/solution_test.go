package balancedBT

import (
  "testing"
)

func TestNotBalancedTree(tt *testing.T) {
  root := &node{
    value: 0,
    left: &node{
      1, nil, nil},
    right: &node{
      2, &node{21, nil, nil}, &node{22, nil, &node{32, nil, nil}}}}

  if isBinaryTreeSuperBalaced(root) {
    tt.Errorf("The tree should NOT balanced.")
  }
}

func TestBalancedTree(tt *testing.T) {
  root := &node{
    value: 0,
    left: &node{
      1, &node{11, nil, nil}, nil},
    right: &node{
      2, &node{21, nil, nil}, &node{22, nil, &node{222, nil, nil}}}}

  if !isBinaryTreeSuperBalaced(root) {
    tt.Errorf("The tree should be balanced.")
  }
}

func TestEmptyTree(tt *testing.T) {
  root := &node{}

  if !isBinaryTreeSuperBalaced(root) {
    tt.Errorf("The tree should be balanced.")
  }
}

func TestOneElementTree(tt *testing.T) {
  root := &node{123, nil, nil}

  if !isBinaryTreeSuperBalaced(root) {
    tt.Errorf("The tree should be balanced.")
  }
}
