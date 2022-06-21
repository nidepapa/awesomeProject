/* Date:2022/2/17-2022
 * Author:zhaoyufan
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

var res = make([][]int,0)
var path []int
var tag = make([]bool,10)
var n int
var N,V int


func main() {


	fmt.Println()
}


func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i , j int ) bool{
		if intervals[i][0]==intervals[j][0] { return intervals[i][1]<=intervals[j][1]}
		return intervals[i][0]<=intervals[j][0]
	})
	preend:= math.MinInt64
	num:=0
	for _,v:=range intervals{
		if v[0]<preend {num++}else {preend=v[1]}
	}
	return num
}
func removeCoveredIntervals(intervals [][]int) int {
	quick_sort1(intervals,0,len(intervals)-1)
	var res = make([][]int,1,len(intervals))
	res[0] = intervals[0]
	var num = 0
	for i:=1;i<len(intervals);i++{
		if intervals[i][1]<=res[len(res)-1][1]{
			num++
		}else{
			res=append(res,intervals[i])
		}
	}
	return num
}

func quick_sort1(q [][]int,l,r int){
	if l<=r{return}
	i,j,x:=l-1,r+1,q[(l+r)>>1][0]
	for i<j{
		for{
			i++
			if q[i][0]>=x{break}
		}
		for {
			j--
			if q[j][0]<=x{break}
		}
		if i<j{q[i],q[j]=q[j],q[i]}
	}
	quick_sort1(q,l,j)
	quick_sort1(q,j+1,r)
}
func getANS(a,b string){
	x:=map[string]map[int]string{
		"a":{0: "a", 1: "ā", 2: "á", 3: "ǎ", 4: "à"},
		"ai":{0: "ai", 1: "āi", 2: "ái", 3: "ǎi", 4: "ài"},
		"an":{0: "an", 1: "ān", 2: "án", 3: "ǎn", 4: "àn"},
		"ang":{0: "ang", 1: "āng", 2: "áng", 3: "ǎng", 4: "àng"},
		"ao":{0: "ao", 1: "āo", 2: "áo", 3: "ǎo", 4: "ào"},

		"o":{0:"o", 1:"ō", 2:"ó", 3:"ǒ", 4:"ò"},
		"ong":{0:"ong", 1:"ōng", 2:"óng", 3:"ǒng", 4:"òng"},
		"ou":{0:"ou", 1:"ōu", 2:"óu", 3:"ǒu", 4:"òu"},

		"e":{0:"e", 1:"ē", 2:"é", 3:"ě", 4:"è"},
		"ei":{0:"ei", 1:"ēi", 2:"éi", 3:"ěi", 4:"èi"},
		"en":{0:"en", 1:"ēn", 2:"én", 3:"ěn", 4:"èn"},
		"eng":{0:"eng", 1:"ēng", 2:"éng", 3:"ěng", 4:"èng"},
		"er":{0:"er", 1:"ēr", 2:"ér", 3:"ěr", 4:"èr"},

		"i":{0:"i", 1:"ī", 2:"í", 3:"ǐ", 4:"ì"},
		"iy":{0:"i", 1:"ī", 2:"í", 3:"ǐ", 4:"ì"},
		"iz":{0:"i", 1:"ī", 2:"í", 3:"ǐ", 4:"ì"},
		"ix":{0:"i", 1:"ī", 2:"í", 3:"ǐ", 4:"ì"},
		"in":{0:"in", 1:"īn", 2:"ín", 3:"ǐn", 4:"ìn"},
		"ing":{0:"ing", 1:"īng", 2:"íng", 3:"ǐng", 4:"ìng"},
		"ian":{0: "ian", 1: "iān", 2: "ián", 3: "iǎn", 4: "iàn"},
		"iao":{0: "iao", 1: "iāo", 2: "iáo", 3: "iǎo", 4: "iào"},
		"ie":{0:"ie", 1:"iē", 2:"ié", 3:"iě", 4:"iè"},
		"iu":{0:"iu", 1:"iū", 2:"iú", 3:"iǔ", 4:"iù"},
		"ia":{0: "ia", 1: "iā", 2: "iá", 3: "iǎ", 4: "ià"},
		"iang":{0: "iang", 1: "iāng", 2: "iáng", 3: "iǎng", 4: "iàng"},
		"iong":{0:"iong", 1:"iōng", 2:"ióng", 3:"iǒng", 4:"iòng"},


		"u":{0:"u", 1:"ū", 2:"ú", 3:"ǔ", 4:"ù"},
		"ui":{0:"ui", 1:"ūi", 2:"úi", 3:"ǔi", 4:"ùi"},
		"un":{0:"un", 1:"ūn", 2:"ún", 3:"ǔn", 4:"ùn"},
		"uo":{0:"uo", 1:"ūo", 2:"úo", 3:"ǔo", 4:"ùo"},
		"ua":{0: "ua", 1: "uā", 2: "uá", 3: "uǎ", 4: "uà"},
		"uai":{0: "uai", 1: "uāi", 2: "uái", 3: "uǎi", 4: "uài"},
		"uan":{0: "uan", 1: "uān", 2: "uán", 3: "uǎn", 4: "uàn"},
		"uang":{0: "uang", 1: "uāng", 2: "uáng", 3: "uǎng", 4: "uàng"},
		"ue":{0:"ue", 1:"uē", 2:"ué", 3:"uě", 4:"uè"},


		"v":{0:"ü", 1:"ǖ", 2:"ǘ", 3:"ǚ", 4:"ǜ"},
		"ve":{0:"üe", 1:"üē", 2:"üé", 3:"üě", 4:"üè"},
		"van":{0: "van", 1: "vān", 2: "ván", 3: "vǎn", 4: "vàn"},
		"vn":{0:"un", 1:"ūn", 2:"ún", 3:"ǔn", 4:"ùn"},

	}

	ans:=make(map[int]string,5)
	ans[0]=a+x[b][0]
	ans[1]=a+x[b][1]
	ans[2]=a+x[b][2]
	ans[3]=a+x[b][3]
	ans[4]=a+x[b][4]
	fmt.Printf("\"%s\":{0:\"%s\",1:\"%s\",2:\"%s\",3:\"%s\",4:\"%s\"},\n",b,ans[0],ans[1],ans[2],ans[3],ans[4])
}

func coinChange(coins []int, amount int) int {
	var f=make([]int,amount+1)
	f[0]=0
	for i:=0;i<len(coins);i++{
		for j:=0;j<=amount;j++{
			if j==coins[i]{f[j]=1}
			f[j]=min(f[j],f[j-coins[i]]+1)
		}
	}
	return f[amount]
}
func findMaxForm(strs []string, m int, n int) int {
	var f = make([][][]int,len(strs))
	for i:=range f{
		f[i]=make([][]int,m+1)
		for j:=range f[i]{
			f[i][j]=make([]int,n+1)
		}
	}
	if getZeroNum(strs[0])<=m && len(strs[0])-getZeroNum(strs[0])<=n {f[0][getZeroNum(strs[0])][len(strs[0])-getZeroNum(strs[0])]=1}
	for i:=1;i<len(strs);i++{
		zeroNum:=getZeroNum(strs[i])
		oneNum:=len(strs[i])-zeroNum
		for j:=zeroNum;j<=m;j++{
			for k:=oneNum;k<=n;k++{
				f[i][j][k]=max(f[i-1][j][k],f[i-1][j-zeroNum][k-oneNum]+1)
			}
		}
	}
	return f[len(strs)-1][m][n]
}

func getZeroNum(x string) int {
	var  num int
	for i:= range x{
		if string(x[i])=="0"{num++}
	}
	return num
}
func findTargetSumWays(nums []int, target int) int {
	//sum = nums[0]+...nums[i]
	//f[i][j]= f[i-1][j] + f[i-1][j-nums[i]]
	var sum int
	for i:=range nums{
		sum+=nums[i]
	}
	bag:=(target+sum)/2
	if (bag)%2==1{return 0}
	if abs(target)>sum {return 0}
	var f = make([]int,bag+1)
	f[0]=1
	for i:=0;i<len(nums);i++{
		for j:=bag;j>=nums[i];j--{
			f[j]+=f[j-nums[i]]
		}
	}
	return f[bag]
}
func abs(x int) int {
	return int(math.Abs(float64(x)))
}


func bagMaxValue(v,w []int) int{
	//f[i][j]=max(f[i-1][j],f[i-1][j-v[i]]+w[i])

	var f = make([]int,V+1)
	for i:=0;i<N;i++{
		for j:=v[i];j<=V;j++{
			f[j]=max(f[j],f[j-v[i]]+w[i])
		}
	}
	return f[V]
}

func permuteUnique(nums []int) [][]int {
	n=len(nums)
	path=make([]int,n)
	quick_sort(nums,0,n-1)
	dfs(0,nums)
	return res
}

func dfs(u int,nums []int){
	if u==n{
		res=append(res,append([]int{},path...))
		return
	}else{
		for i:=range nums{
			if  !tag[i] {
				path[u]=nums[i]
				tag[i]=true
				dfs(u+1,nums)
				path[u]=0
				tag[i]=false
			}
		}
	}
}

func getFormattedTimeByMs(bptime int64) (t string) {
	var (
		hour_ms int64 = 60 * 60 * 1000
		min_ms  int64 = 60 * 1000
		sec_ms  int64 = 1000
	)
	h := bptime / hour_ms
	m := (bptime % hour_ms) / min_ms
	s := (bptime % min_ms) / sec_ms
	t = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	return
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var res = make([]int,0)
	i,j:=0,0
	for i<m && j<n{
		if nums1[i]<nums2[j]{
			res=append(res,nums1[i])
			i++
		}else{
			res=append(res,nums2[j])
			j++
		}
	}
	if i==m{res=append(res,nums2[j:n]...)}
	if j==n{res=append(res,nums1[i:m]...)}
	nums1=res
}

func reverse(q []int){
	var n=len(q)
	mid:=(n-1)>>1
	for i:=0;i<mid;i++{
		q[i],q[n-1-i]=q[n-1-i],q[i]
	}
}
func longestPalindrome(s string) string {
	n:=len(s)
	l,r:=0,0
	if n<=3 && s[0]==s[n-1]{return s}
	dp:=make([][]bool,n)
	for i:= range dp{
		dp[i]=make([]bool,n)
		dp[i][i]=true
	}

	for L:=2;L<=n;L++{
		for i:=2;i<n;i++{
			j:=L+i-1
			if j>=i && j<n-1{
				if s[i]==s[j]{
					if j-i<3{dp[i][j]=true}else{dp[i][j]=dp[i+1][j-1]}
				}else{
					dp[i][j]=false
				}
				if j-i+1>l-r+1 && dp[i][j]  {l,r=i,j}
			}

		}
	}
	return s[l:r+1]
}
func minPathSum(grid [][]int) int {
	var n = len(grid)
	var dp = make([][]int, n)
	var INX = int(^uint(0) >> 1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INX
		}
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+grid[i][j])
			}
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}
	return dp[n][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func quick_sort(q []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l-1, r+1, q[(l+ r)>>1]
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quick_sort(q, l, j)
	quick_sort(q, j+1, r)
}


func bsearch(q []int,target int) int {
	if len(q)==0{
		return -1
	}
	l,r:=0,len(q)-1
	for l<r{
		mid:=(l+r)>>1
		if q[mid]>= target{
			r=mid
		}else{
			l=mid+1
		}
	}
	return l
}



func lengthOfLongestSubstring(s string) int {
	var dic= make(map[byte]int,len(s))
	var res int
	for i,j:=0,0;i<len(s);i++{
		dic[s[i]]++
		for j<i &&dic[s[i]]>1{
			dic[s[j]]--
			j++
			res=max(res,i-j+1)
		}
	}
	return res
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b

}


func mySqrt(x int) int {
	var l,r,ans  =0,x,-1
	for r-l>0{
		mid:=(r+l)/2
		if mid*mid>x{
			r=mid
		}else{
			ans=mid
			l=mid+1
		}
	}
	return ans
}

type LRUCache struct {
	head *LinkedNode
	tail *LinkedNode
	dic map[int]*LinkedNode
	capacity int
}

type LinkedNode struct{
	key int
	val int
	next *LinkedNode
	pre *LinkedNode
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		dic:make(map[int]*LinkedNode,capacity),
		capacity:capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if v,ok:= this.dic[key];ok{
		if v==this.head{}else if v==this.tail{
			this.tail=v.pre
			v.next=this.head
			v.pre=nil
			this.head=v
		}else{
			v.pre.next=v.next
			v.next.pre=v.pre
			v.next=this.head
			v.pre=nil
			this.head=v
		}
		return v.val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if v,ok:=this.dic[key];ok{
		v.val=value
		v.key=key
		if v==this.head.next{}else if v==this.tail.pre{
			this.tail.pre=v.pre
			v.next=this.head.next
			v.pre=nil
			this.head.next=v
		}else{
			v.pre.next=v.next
			v.next.pre=v.pre
			v.next=this.head.next
			v.pre=nil
			this.head.next=v
		}
	}else{
		//not exist
		if len(this.dic)>=this.capacity{
			delete(this.dic,this.tail.pre.key)
			this.tail.pre.pre.next=nil
			this.tail.pre=this.tail.pre.pre
		}
		tmp:=new(LinkedNode)
		tmp.val=value
		tmp.key=key
		tmp.next=this.head.next
		this.head.next=tmp

		this.dic[key]=tmp
	}
}
