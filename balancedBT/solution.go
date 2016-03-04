package balancedBT

type node struct {
  value int
  left  *node
  right *node
}

func isBinaryTreeSuperBalaced(root *node) bool {
  leafDepths := []int{}

  var isDepthValid = func(depth int) bool {
    for _, oneDepth := range leafDepths {
      if oneDepth == depth {
        return true
      }
    }
    // If it reaches here, the depth is not in the leafDepths slice.
    leafDepths = append(leafDepths, depth)
    switch len(leafDepths) {
    case 0:
      panic("We just added one element to the leafDepths, but the " +
        "leafDepths is empty! This should never happen!")
    case 1:
      // depth is the first depth we get.
      return true
    case 2:
      // There are two elements in the slice; see if they're only 1 apart.
      difference := leafDepths[0] - leafDepths[1]
      if difference == 1 || difference == -1 {
        return true
      } else {
        return false
      }
    default:
      // There are more than two different depths!
      return false
    }
  }

  var isBalanced func(int, *node) bool
  isBalanced = func(currentDepth int, root *node) bool {
    if root == nil {
      // Nothing to do.
      return true
    }
    if root.left == nil && root.right == nil { // This "root" is a leaf.
      if !isDepthValid(currentDepth) {
        return false
      }
    }
    result := isBalanced(currentDepth+1, root.left) &&
      isBalanced(currentDepth+1, root.right)
    return result
  }

  return isBalanced(0, root)
}
