package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		wasInCache = c.Set("", 0)
		require.False(t, wasInCache)

		wasInCache = c.Set("", 1)
		require.True(t, wasInCache)

		wasInCache = c.Set("ddd", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("fff", 20)
		require.False(t, wasInCache)

		wasInCache = c.Set("ggg", 50)
		require.False(t, wasInCache)

		wasInCache = c.Set("hhh", 60)
		require.False(t, wasInCache)

		wasInCache = c.Set("iii", 70)
		require.False(t, wasInCache)

		wasInCache = c.Set("uuu", 80)
		require.False(t, wasInCache)

		wasInCache = c.Set("ttt", 90)
		require.False(t, wasInCache)

		wasInCache = c.Set("rrr", 100)
		require.False(t, wasInCache)

		val, ok = c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("ddd")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		wasInCache = c.Set("", 0)
		require.False(t, wasInCache)

		wasInCache = c.Set("ddd", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("fff", 20)
		require.False(t, wasInCache)

		wasInCache = c.Set("ggg", 50)
		require.False(t, wasInCache)

		c.Clear()
		require.Equal(t, c.(*lruCache).capacity, 5)
		require.Equal(t, c.(*lruCache).queue.(*List).len, 0)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
