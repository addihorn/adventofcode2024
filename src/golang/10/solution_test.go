package main

import "testing"

func BenchmarkBFS(b *testing.B) {
	bfs_enables = true

	main()

}

func BenchmarkDFS(b *testing.B) {
	bfs_enables = false

	main()

}
