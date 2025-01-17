package copypasta

import (
	. "fmt"
	"math/bits"
)

/*
标准库 "math/bits" 包含了位运算常用的函数，如二进制中 1 的个数、二进制表示的长度等
注意：bits.Len(0) 返回的是 0 而不是 1
     bits.Len(x) 相当于 int(Log2(x)+eps)+1
     或者说 2^(Len(x)-1) <= x < 2^Len(x)

TIPS: & 和 | 在区间求和上具有单调性；^ 的区间求和见 strings.go 中的 trie.maxXor
** 代码和题目见下面的 logTrick 和 logTrickCnt

常用等式（若改变了计算的顺序，注意优先级！）
a|b = (a^b) + (a&b)    a^b = (a|b) - (a&b)
a+b = (a|b) + (a&b)
    = (a&b)*2 + (a^b)
    = (a|b)*2 - (a^b)
相关题目
https://codeforces.com/problemset/problem/1325/D
https://atcoder.jp/contests/abc050/tasks/arc066_b

结合律：(a&b)^(a&c) = a&(b^c)    其他符号类似
相关题目 https://leetcode-cn.com/contest/weekly-contest-237/problems/find-xor-sum-of-all-pairs-bitwise-and/

运算符优先级 https://golang.org/ref/spec#Operators
Precedence    Operator
    5         *  /  %  <<  >>  &  &^
    4         +  -  |  ^
    3         ==  !=  <  <=  >  >=
    2         &&
    1         ||

一些子集的枚举算法见 loopCollection
S∪{i}: S|1<<i
S\{i}:  S&^(1<<i)
构造 2^n-1，即 n 个 1 的另一种方法: ^(-1<<n)
检测是否只有一个 1：x&(x-1) == 0

https://oeis.org/A060142 每一段连续 0 的长度均为偶数的数：如 100110000100
Ordered set S defined by these rules: 0 is in S and if x is in S then 2x+1 and 4x are in S
0, 1, 3, 4, 7, 9, 12, 15, 16, 19, 25, 28, 31, 33, 36, 39, 48, 51, 57, 60, 63, 64, 67, 73, 76, 79, 97, 100
https://oeis.org/A086747 Baum-Sweet sequence
相关题目：蒙德里安的梦想 https://www.acwing.com/problem/content/293/

https://oeis.org/A047778 Concatenation of first n numbers in binary, converted to base 10
相关题目：https://leetcode-cn.com/contest/weekly-contest-218/problems/concatenation-of-consecutive-binary-numbers/
钱珀瑙恩数 Champernowne constant https://en.wikipedia.org/wiki/Champernowne_constant

异或和相关
https://oeis.org/A003987 异或矩阵
https://oeis.org/A003815 异或和 i  a(0)=0, a(4n+1)=1, a(4n+2)=4n+3, a(4n+3)=0, a(4n+4)=4n+4
    相关题目 https://codeforces.com/problemset/problem/1493/E
            https://codeforces.com/problemset/problem/460/D
https://oeis.org/A145768 异或和 i*i
https://oeis.org/A126084 异或和 质数
https://oeis.org/A018252 异或和 合数?
https://oeis.org/A072594 异或和 质因数分解 是积性函数 a(p^k)=p*(k&1)
	https://oeis.org/A072595 满足 A072594(n)=0 的数
https://oeis.org/A178910 异或和 因子
	https://oeis.org/A178911 满足 A178910(n)=n 的数 Perfex number

https://oeis.org/A038712 a(n) = n^(n-1) = 1, 3, 1, 7, 1, 3, 1, 15, 1, ...
https://oeis.org/A080277 A038712 的前缀和  =>  a(n) = n + 2*a(n/2)

二进制长度
https://oeis.org/A070939 a(0)=1, a(n)=bits.Len(n)
https://oeis.org/A083652 A070939 的前缀和

OnesCount 相当于二进制的 digsum
https://oeis.org/A000120 wt(n) = OnesCount(n)
https://oeis.org/A000788 前缀和 a(2^n)=n*2^(n-1)+1
https://oeis.org/A121853 前缀积 https://www.luogu.com.cn/problem/P4317
https://oeis.org/A092391 n+OnesCount(n)
	https://oeis.org/A010061 二进制自我数/哥伦比亚数（A092391 的补集）
https://oeis.org/A011371 n-OnesCount(n) Also highest power of 2 dividing n!
							a(n) = floor(n/2) + a(floor(n/2))
                         这同时是前 n 个数的质因子分解的 2 的幂次之和
https://oeis.org/A027868 Number of trailing zeros in n!; highest power of 5 dividing n!
                            a(n) = (n-A053824(n))/4, 其中 A053824(n) = Sum of digits of (n written in base 5)
推广至任意数：n! 的质因子分解中，p 的幂次为 (n-digsum_p(n))/(p-1)，其中 digsum_p(n) 表示 n 的 p 进制的数位和
https://oeis.org/A245788 n*OnesCount(n)
https://oeis.org/A049445 OnesCount(n)|n
	-  n/OnesCount(n)
https://oeis.org/A199238 n%OnesCount(n)
https://oeis.org/A010062 a(0)=1, a(n+1)=a(n)+OnesCount(a(n))
	https://oeis.org/A096303 从 n 出发不断执行 n+=OnesCount(n)，直到 n 在 A010062 中，所需要的迭代次数
	Number of iterations of n -> n + (number of 1's in binary representation of n) needed for the trajectory of n to join the trajectory of A010062
		https://oeis.org/A229743 Positions of records
		https://oeis.org/A229744 Values of records
	相关题目 https://www.luogu.com.cn/problem/P5891 https://class.luogu.com.cn/classroom/lgr66

https://oeis.org/A023416 Number of 0's in binary expansion of n
							a(n) = a(n/2) + 1 - n&1
https://oeis.org/A059015 A023416 的前缀和

十进制 digsum
https://oeis.org/A007953 digsum(n)
https://oeis.org/A062028 n+digsum(n)    质数 https://oeis.org/A047791    合数 https://oeis.org/A107743
	https://oeis.org/A003052 自我数/哥伦比亚数 Self number / Colombian number
	https://en.wikipedia.org/wiki/Self_number
	1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, 97, 108, ...
		https://oeis.org/A006378 自我质数 Self primes
https://oeis.org/A066568 n-digsum(n)
https://oeis.org/A057147 n*digsum(n)
https://oeis.org/A005349 digsum(n)|n   Niven (or Harshad) numbers
	https://oeis.org/A065877 digsum(n)∤n   Non-Niven (or non-Harshad) numbers
	https://oeis.org/A001101 Moran numbers: n such that (n / digsum(n)) is prime
https://oeis.org/A016052 a(1)=3, a(n+1)=a(n)+digsum(a(n))
https://oeis.org/A051885 Smallest number whose digsum = n
							int64(n%9+1) * int64(math.Pow10(n/9)) - 1
							相关题目 https://codeforces.com/contest/1373/problem/E
https://oeis.org/A077196 Smallest possible sum of the digits of a multiple of n https://oeis.org/A077194 https://oeis.org/A077195
							相关题目（0-1 最短路）https://atcoder.jp/contests/arc084/tasks/arc084_b
https://oeis.org/A118137 digsum(n)+digsum(n+1)
https://oeis.org/A003132 Sum of squares of digits of n
	https://oeis.org/A003621 Number of iterations until n reaches 1 or 4 under x goes to sum of squares of digits map
https://oeis.org/A055012 Sum of cubes of digits of n
https://oeis.org/A055013 Sum of 4th powers of digits of n
https://oeis.org/A055014 Sum of 5th powers of digits of n
https://oeis.org/A055015 Sum of 6th powers of digits of n
	相关题目 https://www.luogu.com.cn/problem/P1660

回文数
https://oeis.org/A002113 十进制回文数
	https://oeis.org/A043269 digsum(A002113(n))
	https://oeis.org/A070199 Number of palindromes of length <= n
https://oeis.org/A002779 回文平方数
	https://oeis.org/A002778 Numbers whose square is a palindrome
https://oeis.org/A002781 回文立方数
	https://oeis.org/A002780 Numbers whose cube is a palindrome
https://oeis.org/A002385 回文素数
	https://en.wikipedia.org/wiki/Palindromic_prime
https://oeis.org/A006567 反素数 emirp (primes whose reversal is a different prime)
	https://en.wikipedia.org/wiki/Emirp
https://oeis.org/A003459 绝对素数/可交换素数 Absolute primes (or permutable primes): every permutation of the digits is a prime
	https://en.wikipedia.org/wiki/Permutable_prime
https://oeis.org/A007500 Primes whose reversal in base 10 is also prime
https://oeis.org/A006995 二进制回文数

https://oeis.org/A090994 Number of meaningful differential operations of the n-th order on the space R^9
a(k+5) = a(k+4) + 4*a(k+3) - 3*a(k+2) - 3*a(k+1) + a(k)
相关题目 LC1215/双周赛10C https://leetcode-cn.com/contest/biweekly-contest-10/problems/stepping-numbers/

套路题 https://codeforces.com/problemset/problem/1415/D
按位归纳 https://codeforces.com/problemset/problem/925/C
*/

// 参考 strings/strings.go 中的 asciiSet
// 64 位系统下可以用 uint，把 5 改成 6，31 改成 63
type bitset []uint32 // b := make(bitset, n>>5+1)

func (b bitset) set(p int)           { b[p>>5] |= 1 << (p & 31) }
func (b bitset) reset(p int)         { b[p>>5] &^= 1 << (p & 31) }
func (b bitset) flip(p int)          { b[p>>5] ^= 1 << (p & 31) }
func (b bitset) contains(p int) bool { return 1<<(p&31)&b[p>>5] > 0 }

// 需要保证长度相同
func (b bitset) equals(c bitset) bool {
	for i, v := range b {
		if v != c[i] {
			return false
		}
	}
	return true
}
func (b bitset) hasSubset(c bitset) bool {
	for i, v := range b {
		if v|c[i] != v {
			return false
		}
	}
	return true
}

// 注：有关子集枚举的位运算技巧，见 search.go
func bitsCollection() {
	// 利用 -v = ^v+1
	lowbit := func(v int64) int64 { return v & -v }

	// 1,2,4,8,...
	isPow2 := func(v int64) bool { return v > 0 && v&(v-1) == 0 }

	bits31 := func(v int) []byte {
		bits := make([]byte, 31)
		for i := range bits {
			bits[i] = byte(v >> (30 - i) & 1)
		}
		return bits
	}
	_bits31 := func(v int) string { return Sprintf("%031b", v) }
	_bits32 := func(v uint) string { return Sprintf("%032b", v) }

	digitSum := func(v int) (s int) {
		for ; v > 0; v /= 10 {
			s += v % 10
		}
		return
	}

	// 对于数组的所有区间，返回 op(区间元素) 的各个结果
	// 利用操作的单调性求解
	// |: LC898/周赛100C https://leetcode-cn.com/contest/weekly-contest-100/problems/bitwise-ors-of-subarrays/
	// &: LC1521/周赛198D https://leetcode-cn.com/contest/weekly-contest-198/problems/find-a-value-of-a-mysterious-function-closest-to-target/
	// GCD: https://codeforces.com/edu/course/2/lesson/9/2/practice/contest/307093/problem/G
	//      https://codeforces.com/problemset/problem/475/D (见下面的 logTrickCnt)
	bitOpTrick := func(a []int, op func(x, y int) int) map[int]bool {
		ans := map[int]bool{} // 统计 op(一段区间) 的不同结果
		set := []int{}
		for _, x := range a {
			for i, v := range set {
				set[i] = op(v, x)
			}
			set = append(set, x)
			// 去重
			k := 0
			for _, w := range set[1:] {
				if set[k] != w {
					k++
					set[k] = w
				}
			}
			set = set[:k+1]
			for _, v := range set {
				// do v...
				ans[v] = true
			}
		}
		return ans
	}

	// 进阶：对于数组的所有区间，返回 op(区间元素) 的各个结果，及其出现次数
	// https://codeforces.com/problemset/problem/475/D
	bitOpTrickCnt := func(a []int, op func(x, y int) int) map[int]int64 {
		cnt := map[int]int64{}
		type pair struct{ v, l, r int }
		set := []pair{}
		for i, x := range a {
			for j, p := range set {
				set[j].v = op(p.v, x)
			}
			set = append(set, pair{x, i, i + 1})
			// 去重
			k := 0
			for _, q := range set[1:] {
				if set[k].v != q.v {
					k++
					set[k] = q
				} else {
					set[k].r = q.r
				}
			}
			set = set[:k+1]
			for _, p := range set {
				// do p...     [l,r)
				cnt[p.v] += int64(p.r - p.l)
			}
		}
		return cnt
	}

	// 找三个不同的在 [l,r] 范围内的数，其异或和为 0
	// 考虑尽可能地小化最大减最小的值，构造 (x, y, z) = (b*2-1, b*3-1, b*3), b=2^k
	// 相关题目 https://codeforces.com/problemset/problem/460/D
	zeroXorSum3 := func(l, r int64) []int64 {
		for b := int64(1); b*3 <= r; b <<= 1 {
			if x, y, z := b*2-1, b*3-1, b*3; l <= x && z <= r {
				return []int64{x, y, z}
			}
		}
		return nil
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 在 [low,high] 区间内找两个数字 A B，使其异或值最大且不超过 limit
	// 返回值保证 A <= B
	// 复杂度 O(log(high))
	maxXorWithLimit := func(low, high, limit int) (int, int) {
		n := bits.Len(uint(high ^ low))
		maxXor := 1<<n - 1
		mid := high&^maxXor | 1<<(n-1)
		if limit >= maxXor { // 无约束，相关题目 https://codeforces.com/problemset/problem/276/D
			return mid - 1, mid
		}
		if limit >= 1<<(n-1) { // A 和 B 能否在第 n-1 位不同的情况下，构造出一个满足要求的解？
			a, b := mid&(mid-1), mid
			for i := n - 2; i >= 0; i-- {
				bt := 1 << i
				if limit&bt > 0 { // a 取 1，b 取 0 总是优于 a 取 0，b 取 1
					a |= bt
				} else if high&(bt<<1-1) > ^low&(bt<<1-1) { // high 侧大，都取 1
					if high&bt == 0 { // b 没法取 1
						goto next
					}
					a |= bt
					b |= bt
				} else {            // low 侧大，都取 0
					if low&bt > 0 { // a 没法取 0
						goto next
					}
				}
				if (a^low)&bt > 0 { // a 不受 low 的约束
					a |= limit & (bt - 1)
					break
				}
				if (b^high)&bt > 0 { // b 不受 high 的约束
					a |= bt - 1
					b |= ^limit & (bt - 1)
					break
				}
			}
			return a, b
		}
		// A 和 B 在第 n-1 位上必须相同
	next:
		f := func(high int) (int, int) {
			n := bits.Len(uint(high ^ mid))
			maxXor := min(1<<n-1, limit)
			// 只有当 maxXor 为 0 时，返回值才必须相等
			if maxXor == 0 {
				return mid, mid
			}
			// maxXor 的最高位置于 B，其余置于 A
			mb := 1 << (bits.Len(uint(maxXor)) - 1)
			return mid | maxXor&^mb, mid | mb
		}
		if high-mid > mid-1-low { // 选区间长的一侧
			return f(high)
		}
		a, b := f(2*mid - 1 - low) // 对称到 high
		return 2*mid - 1 - b, 2*mid - 1 - a
	}

	_ = []interface{}{lowbit, isPow2, bits31, _bits31, _bits32, digitSum, bitOpTrick, bitOpTrickCnt, zeroXorSum3, maxXorWithLimit}
}

// https://halfrost.com/go_s2_de_bruijn/

// LC137 https://leetcode-cn.com/problems/single-number-ii/
// 除了某个元素只出现一次以外，其余每个元素均出现了三次。返回只出现了一次的元素
// 		定义两个集合 ones 和 twos，初始为空
// 		第一次出现就放在 ones 中
//		第二次出现就在 ones 中删除并放在 twos
//		第三次出现就从 twos 中删除
//		这样最终 ones 中就留下了最后的结果
func singleNumber(a []int) int {
	ones, twos := 0, 0
	for _, v := range a {
		ones = (ones ^ v) &^ twos
		twos = (twos ^ v) &^ ones
	}
	return ones
}
