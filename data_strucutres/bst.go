package main

import "fmt"

type node struct{
    data int
    left *node
    right *node
}

// Implementation of Binary Search Tree
func insert(root *node, val int)*node{
    temp := node{val, nil, nil}
    if root == nil{
        return &temp
    }
    if root.data < val {
        root.right = insert(root.right, val)
    }else {
        root.left = insert(root.left, val)
    }
    return root
}

// In order traversal of BST gives you a sorted array

func inorder(root *node){
    if root == nil{
        return
    }
    inorder(root.left)
    fmt.Println(root.data)
    inorder(root.right)
}

func main(){
    arr := [6]int{2,1,4,3,5,6}
    var root *node
    for _,v:= range arr{
        root = insert(root, v)
    }
    inorder(root)
}
