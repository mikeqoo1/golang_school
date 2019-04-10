package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

//FindMedianSortedArrays 尋找中位數
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	nums3 := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i = i + 1
		} else {
			nums3 = append(nums3, nums2[j])
			j = j + 1
		}
	}
	if i < len(nums1) {
		for ; i < len(nums1); i++ {
			nums3 = append(nums3, nums1[i])
		}
	}
	if j < len(nums2) {
		for ; j < len(nums2); j++ {
			nums3 = append(nums3, nums2[j])
		}
	}
	lenSums3 := len(nums3)
	if lenSums3%2 == 0 {
		return (float64(nums3[lenSums3/2-1]) + float64(nums3[lenSums3/2])) / 2.0
	} else {
		return float64(nums3[(lenSums3+1)/2-1])
	}
}

//Reverse 反轉整數 123 -> 321 -123 -> -321
func Reverse(x int) int {
	reverseans := 0
	if x < 0 {
		x = -x
		for x > 0 {
			fmt.Println("0.")
			reverseans *= 10
			fmt.Println("1.", "x =", x)
			fmt.Println("1.", "reverseans =", reverseans)
			reverseans += x % 10
			fmt.Println("2.", "x =", x)
			fmt.Println("2.", "reverseans =", reverseans)
			x /= 10
			fmt.Println("3.", "x =", x)
			fmt.Println("3.", "reverseans =", reverseans)
		}

		reverseans = -reverseans
		if reverseans < -2147483648 {
			return 0
		}

		return reverseans
	}

	for x > 0 {
		fmt.Println("0.")
		reverseans *= 10
		fmt.Println("1.", "x =", x)
		fmt.Println("1.", "reverseans =", reverseans)
		reverseans += x % 10
		fmt.Println("2.", "x =", x)
		fmt.Println("2.", "reverseans =", reverseans)
		x /= 10
		fmt.Println("3.", "x =", x)
		fmt.Println("3.", "reverseans =", reverseans)
	}

	if reverseans > 2147483647 {
		return 0
	}

	return reverseans
}

//MyAtoi MyAtoi
func MyAtoi(str string) int {
	//先判斷是不是正常的字串, 再判斷正負, 再合成新的字串, 接下來就可以轉了
	//假設 "-123"
	//找出有效起始位置, 開始組裝

	if len(str) == 0 {
		return 0
	}
	strByte := []byte(str)
	Head := true
	var newStr []byte
	for i := 0; i < len(strByte); i++ {
		if strByte[i] == ' ' {
			if Head {
				continue
			} else {
				break
			}
		}
		if strByte[i] >= '0' && strByte[i] <= '9' || (strByte[i] == '-' || strByte[i] == '+') && Head {
			Head = false
			newStr = append(newStr, strByte[i])
			continue
		}
		break
	}

	res, _ := strconv.Atoi(string(newStr))

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res

}

//IsPalindrome 判斷回文數
func IsPalindrome(x int) bool {
	// 先把原數字%10 在 /10 依照此步驟來反轉
	if x < 0 {
		return false
	}
	temp := x
	var x2 int
	for temp != 0 {
		pop := temp % 10
		//fmt.Println("pop", pop)
		temp = temp / 10
		//fmt.Println("temp", temp)
		x2 = x2*10 + pop
		//fmt.Println("x2", x2)
		//fmt.Println("><><<<><<<><>")
	}
	if x2 == x {
		return true
	}
	return false
}

//RegularExpressionMatching 10. Regular Expression Matching
func RegularExpressionMatching(s string, p string) bool {
	for true {
		r1, w1 := utf8.DecodeRuneInString(p)
		if w1 <= 0 {
			break
		}
		p = p[w1:]

		r2, w2 := utf8.DecodeRuneInString(p)
		if r2 == '*' {
			// Zero or more
			p = p[w2:]
			for true {
				if ok := RegularExpressionMatching(s, p); ok {
					return true
				}
				r3, w3 := utf8.DecodeRuneInString(s)
				if (r1 == '.' || r1 == r3) && w3 > 0 {
					s = s[w3:]
				} else {
					break
				}
			}
		} else {
			// One
			r3, w3 := utf8.DecodeRuneInString(s)
			if w3 <= 0 || !(r1 == '.' || r1 == r3) {

				return false
			}
			s = s[w3:]
		}
	}

	return len(s) <= 0

	/*快速版
		n, m := len(s), len(p)
	    dp := make([][]bool, n+1)
	    for i:=0;i<=n;i++ {
	        dp[i] = make([]bool, m+1)
	    }
	    dp[0][0] = true
	    for j:=2; j<=m; j++ {
	        if p[j-1] == '*' {
	            dp[0][j] = dp[0][j-2]
	        }
	    }
	    // fmt.Printf("%#v\n",dp[0])
	    for i:=1; i<=n; i++ {
	        for j:=1; j<=m; j++ {
	            if s[i-1] == p[j-1] || p[j-1] == '.' {
	                dp[i][j] = dp[i-1][j-1]
	                continue
	            }
	            if p[j-1] == '*' {
	                if p[j-2] != '.' && p[j-2] != s[i-1] {
	                    dp[i][j] = dp[i][j-2]
	                } else {
	                    dp[i][j] = dp[i][j-2] || dp[i-1][j]
	                }
	            }
	        }
	    }
	    // for _, r := range dp{
	    //     fmt.Printf("%v\n",r)
	    // }
		return dp[n][m]
	*/
}
