package main


func FindMinSubString(s1 *[]byte,s2 *[]byte) {
	hash_table := [256]int{}
	/*
		初始化 hash 表,s2中出现的字符标记为0,
		其他标记为负数
	*/
	for i:=0; i < 256; i++{
		hash_table[i] = -1;
	}
	for (p := s2; *p != '\0'; p++){
	hash_table[*p] = 0;
	}
	char* p1 = s1;
	char* p2 = s1;

//最短长度
int min_len = 2100000000;

//最小串的起止位置
char* min_p1 = s1;
char* min_p2 = s1;

//记录p1 p2之间的合法字符种类数 ,count == s2_len 包含s2
int count = 0;

int s2_len = strlen(s2);

/*
	注意p2到达s1的结束位置后，不能结束
	应当继续收缩，直到不含s2
*/
while(*p2 !='\0' || s2_len==count)
{

//p1...p2不包含s2
if (count<s2_len)
{
//属于s2中出现的字符
if (hash_table[*p2] == 0)
{
count++;
hash_table[*p2]++;
}
else if (hash_table[*p2] > 0)
{
hash_table[*p2]++;
}

p2++;
}

//不能用else 因为在上一个if语句中对count++;
if(count == s2_len)
{
if (p2-p1 < min_len)
{
min_p1 = p1;
min_p2 = p2;
min_len = p2-p1;
}

//收缩
hash_table[*p1]--;
if (hash_table[*p1] == 0)
{
count--;
}
p1++;
}
}

//输出
while(min_p1 < min_p2)
{
cout<<*min_p1;
min_p1++;
}
cout<<endl;

return min_len;
}
int main()
{
char str1[]="abdddcaabc";
char str2[]="abc";
cout<<FindMinSubString(str1,str2);

}
