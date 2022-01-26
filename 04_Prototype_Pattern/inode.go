package main

type inode interface {
	clone() inode
	print(string)
}
