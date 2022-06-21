/* Date:2022/5/29-2022
 * Author:zhaoyufan
 */
package main

func LadderLength(beginWord string, endWord string, wordList []string) int {
	var dic = make(map[string]struct{})
	var queue = []string{beginWord}
	var step = 1
	for _,v:=range wordList{dic[v]=struct{}{}}

	for len(queue)!=0{
		for i:=len(queue);i>0;i--{
			word:=queue[0]
			if word == endWord{return step}
			for i:=0;i<len(word);i++{
				for j:=0;j<26;j++{
					replaceAtIndex(word,rune(j+'a'),i)
					if _,ok:=dic[word];ok{
						if word == endWord {return step+1}
						queue = append(queue,word)
						delete(dic,word)
						queue=queue[1:]
					}
				}
			}
		}
		step++
	}
	return 0
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
