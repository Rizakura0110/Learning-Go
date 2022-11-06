package main

import "fmt"

func main() {
	// スライス
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])
	}
	{
		var x []int
		fmt.Println(x)
	}

	{
		var x []int
		var y []int
		fmt.Println("x == nil:", x == nil)
		fmt.Println("y != nil:", y != nil)
	}
	// len
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("len(x):", len(x))
		var y = []int{}
		fmt.Println("len(y):", len(y))
	}
	fmt.Println("//////")
	{
		var x [][]int
		fmt.Println(x)

		var y = [][]int{{0, 1}, {2, 3}, {4, 5}}
		fmt.Println(y)
		fmt.Println(y[1])    // [2 3]
		fmt.Println(y[1][1]) // 3
	}
	// append
	{
		var x []int
		fmt.Println(x)
		x = append(x, 10)
		fmt.Println(x)
	}
	{
		var x = []int{1, 2, 3}
		x = append(x, 4)
		fmt.Println(x)
		x = append(x, 5, 6, 7)
		fmt.Println(x)
		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println(x)
	}
	// make
	{
		// 長さ5・キャパシティ5のintのスライス
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		fmt.Println("x[0], x[4]:", x[0], x[4])
	}
	{
		var data []int
		fmt.Println(data, len(data), cap(data))
		fmt.Println("data == nil:", data == nil)

		var x = []int{}
		fmt.Println(x, len(x), cap(x))
		fmt.Println("x == nil:", x == nil)

		data2 := []int{2, 4, 6, 8}
		fmt.Println(data2, len(data2), cap(data2))

	}
	{
		x := []int{1, 2, 3, 4}
		y := x[:3]
		z := x[1:]
		d := x[1:3]
		e := x[:]
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("d:", d)
		fmt.Println("e:", e)
	}
	{
		x := make([]int, 0, 5)
		x = append(x, 1, 2, 3, 4)
		y := x[:2:2]
		z := x[2:4:4]
		fmt.Println(cap(x), cap(y), cap(z))
		y = append(y, 30, 40, 50)
		x = append(x, 60)
		z = append(z, 70)
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	}
	{
		x := [...]int{5, 6, 7, 8}
		y := x[:2]
		z := x[2:]
		d := x[:]
		x[0] = 10
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)
		fmt.Println(d)
	}
	{
		x := []int{1, 2, 3, 4}
		y := make([]int, 4)
		num := copy(y, x)
		fmt.Println(y, num)
	}
	{
		x := []int{1, 2, 3, 4}
		d := [4]int{5, 6, 7, 8}
		fmt.Println(d)
		y := make([]int, 2)
		copy(y, d[:])
		fmt.Println(y)
		copy(d[:], x)
		fmt.Println(d)
	}

	{
		var nilMap map[string]int
		fmt.Println("nilmap == nil", nilMap == nil)
		// nilMap["abc"] = 3
	}

	{
		totalWins := map[string]int{}
		fmt.Println("totalWins == nil:", totalWins == nil)
		fmt.Println(totalWins["abc"])
		totalWins["abc"] = 3
		fmt.Println("totalWins[\"abc\"]:", totalWins["abc"])
	}

	{
		totalWins := map[string]int{
			"セネターズ":  14,
			"スパローズ":  15,
			"ファルコンズ": 22,
		}
		fmt.Println(totalWins)
	}
	{
		teams := map[string][]string{
			"ライターズ":    []string{"夏目", "森", "国木田"},
			"ナイツ":      []string{"武田", "徳川", "明智"},
			"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
		}
		fmt.Println(teams)
		fmt.Println(teams["ライターズ"])

		teams2 := map[string][]string{
			"シャチチーム":  []string{"謙信", "信長", "家康"},
			"ライオンチーム": []string{"レオ", "たか子", "カナ"},
			"猫チーム":    []string{"AKB", "MNB", "FNB"},
		}
		fmt.Println(teams2)
		fmt.Println(teams2["シャチチーム"])
		fmt.Println(teams2["チャチチーム"])
		fmt.Println(len(teams2["猫チーム"]))
	}

	{
		totalWins := map[string]int{}
		totalWins["ライターズ"] = 1
		totalWins["ナイツ"] = 2
		fmt.Println(totalWins["ライターズ"])
		fmt.Println(totalWins["ミュージシャンズ"])
		totalWins["ミュージシャンズ"]++
		fmt.Println(totalWins["ミュージシャンズ"])
	}
	// カンマokイディオム
	{
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok)
		v, ok = m["world"]
		fmt.Println(v, ok)
		v, ok = m["goodbye"]
		fmt.Println(v, ok)

		k := map[string]int{
			"hello": 5,
			"world": 10,
		}
		delete(k, "hello")
		fmt.Println(k)
	}
	{
		intSet := map[int]bool{}
		vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
		for _, v := range vals {
			intSet[v] = true
		}
		fmt.Println(len(vals), len(intSet))
		fmt.Println(intSet[5])
		fmt.Println(intSet[500])
		if intSet[100] {
			fmt.Println("100は含まれている")
		}
		if intSet[10] {
			fmt.Println("10は含まれている")
		}
	}
}
