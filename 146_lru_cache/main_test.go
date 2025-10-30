package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	t.Run("test_1", func(t *testing.T) {
		c := Constructor(2)
		c.Put(1, 1)
		c.Put(2, 2)
		require.Equal(t, 1, c.Get(1))
		c.Put(3, 3)
		require.Equal(t, -1, c.Get(2))
		c.Put(4, 4)
		require.Equal(t, -1, c.Get(1))
		require.Equal(t, 3, c.Get(3))
		require.Equal(t, 4, c.Get(4))
	})

	t.Run("test_2", func(t *testing.T) {
		c := Constructor(2)
		c.Put(2, 1)
		c.Put(2, 2)
		require.Equal(t, 2, c.Get(2))
		c.Put(1, 1)
		c.Put(4, 1)
		require.Equal(t, -1, c.Get(2))
	})

	t.Run("test_3", func(t *testing.T) {
		c := Constructor(2)
		require.Equal(t, -1, c.Get(2))
		c.Put(2, 6)
		require.Equal(t, -1, c.Get(1))
		c.Put(1, 5)
		c.Put(1, 2)
		require.Equal(t, 2, c.Get(1))
		require.Equal(t, 6, c.Get(2))
	})

	t.Run("test_4", func(t *testing.T) {
		c := Constructor(2)
		c.Put(2, 1)
		c.Put(3, 2)

		require.Equal(t, 2, c.Get(3))
		require.Equal(t, 1, c.Get(2))

		c.Put(4, 3)

		require.Equal(t, 1, c.Get(2))
		require.Equal(t, -1, c.Get(3))
		require.Equal(t, 3, c.Get(4))
	})
}
