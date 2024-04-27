package mathfunctions_test

import (
	"fmt"
	"mathfunctions"
	"sync"
	"testing"
)

func TestFindMin(t *testing.T) {
	t.Parallel()
	t.Run(`test find min number`, func(t *testing.T) {
		testTable := []struct {
			Values []int64
			Result int64
		}{
			{
				Values: []int64{-700, 0, 1, 10, 1, 2, 3, 4, 5, 6},
				Result: -700,
			},
			{
				Values: []int64{300, 0, 2, -100, 1, 3, 4, 5, 6, 7},
				Result: -100,
			},
			{
				Values: []int64{-100000, -178, 2, 3, 1, 3, 4, 5, 6, 7},
				Result: -100000,
			},
			{
				Values: []int64{90, 98, 78, 105, 100, 103, 104, 999},
				Result: 78,
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наименьшего числа
				result := mathfunctions.FindMinNumber(expect.Values...)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %d != %d`, expect.Result, result))
				}
			}()
		}

		wg.Wait()
	})

	t.Run(`test find max number`, func(t *testing.T) {
		testTable := []struct {
			Values []int64
			Result int64
		}{
			{
				Values: []int64{-700, 0, 1, 10, 1, 2, 3, 4, 5, 6},
				Result: 10,
			},
			{
				Values: []int64{300, 0, 2, -100, 1, 3, 4, 5, 6, 7},
				Result: 300,
			},
			{
				Values: []int64{0, 1, 2, 3, 2, 1, 2, 3, 1, 2},
				Result: 3,
			},
			{
				Values: []int64{90, 98, 78, 105},
				Result: 105,
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наибольшего числа
				result := mathfunctions.FindMaxNumber(expect.Values...)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %d != %d`, expect.Result, result))
				}
			}()
		}
		wg.Wait()

	})
}

func TestSumNumber(t *testing.T) {
	t.Parallel()
	t.Run(`test sum even number`, func(t *testing.T) {
		testTable := []struct {
			Values []int64
			Result int64
		}{
			{
				Values: []int64{-700, 0, 1, 10, 1, 2, 3, 4, 5, 6},
				Result: -678,
			},
			{
				Values: []int64{300, 0, 2, -100, 1, 3, 4, 5, 6, 7},
				Result: 212,
			},
			{
				Values: []int64{-100000, -178, 2, 3, 1, 3, 4, 5, 6, 7},
				Result: -100166,
			},
			{
				Values: []int64{90, 98, 78, 105, 100, 103, 104, 999},
				Result: 470,
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наименьшего числа
				result := mathfunctions.SumEvenNumber(expect.Values...)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %d != %d`, expect.Result, result))
				}
			}()
		}

		wg.Wait()
	})

	t.Run(`test sum no even number`, func(t *testing.T) {
		testTable := []struct {
			Values []int64
			Result int64
		}{
			{
				Values: []int64{-700, 0, 1, 10, 1, 2, 3, 4, 5, 6},
				Result: 10,
			},
			{
				Values: []int64{300, 0, 2, -100, 1, 3, 4, 5, 6, 7},
				Result: 16,
			},
			{
				Values: []int64{0, 1, 2, 3, 2, 1, 2, 3, 1, 2},
				Result: 9,
			},
			{
				Values: []int64{90, 98, 78, 105},
				Result: 105,
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наибольшего числа
				result := mathfunctions.SumNoEvenNumber(expect.Values...)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %d != %d`, expect.Result, result))
				}
			}()
		}
		wg.Wait()

	})
}
