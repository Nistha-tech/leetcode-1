<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    Openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

## 114. Flatten Binary Tree to Linked List (Medium)

<p>Given a binary tree, flatten it to a linked list in-place.</p>

<p>For example, given the following tree:</p>

<pre>
    1
   / \
  2   5
 / \   \
3   4   6
</pre>

<p>The flattened tree should look like:</p>

<pre>
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
</pre>


### Related Topics
[[Tree](https://github.com/openset/leetcode/tree/master/tag/tree/README.md)]
[[Depth-first Search](https://github.com/openset/leetcode/tree/master/tag/depth-first-search/README.md)]

### Similar Questions
  1. [Flatten a Multilevel Doubly Linked List](https://github.com/openset/leetcode/tree/master/problems/flatten-a-multilevel-doubly-linked-list) (Medium)

### Hints
  1. If you notice carefully in the flattened tree, each node's right child points to the next node of a pre-order traversal.