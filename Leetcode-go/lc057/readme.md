# 57.插入区间

## 1. 题目描述

给你一个 **无重叠的** *，* 按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

 

 **示例 1：** 

```

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]

```
 **示例 2：** 

```

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
```
 **示例 3：** 

```

输入：intervals = [], newInterval = [5,7]
输出：[[5,7]]

```
 **示例 4：** 

```

输入：intervals = [[1,5]], newInterval = [2,3]
输出：[[1,5]]

```
 **示例 5：** 

```

输入：intervals = [[1,5]], newInterval = [2,7]
输出：[[1,7]]

```
 

 **提示：** 
-  `0 <= intervals.length <= 10^4` 
-  `intervals[i].length == 2` 
-  `0 <= intervals[i][0] <= intervals[i][1] <= 10^5` 
-  `intervals` 根据 `intervals[i][0]` 按 **升序** 排列
-  `newInterval.length == 2` 
-  `0 <= newInterval[0] <= newInterval[1] <= 10^5` 
 
**标签**
`数组` 


## 2. 解题
设newInterval左端点为left，右端点为right。  
按顺序遍历，有以下几种情况：  
1.若left>intervals[i][1],说明新区间要加入的位置必定在i之后  
2.继续遍历，若right < intervals[i][0],说明新区间要加入的位置在i和i-1之间  
3.若right >= intervals[i][0],说明新区间和区间intervals[i]存在重复区间，需合并，合并后区间的左端点为两个区间的最小值，右端点为两个区间的最大值。  
4.继续遍历，合并完毕后，将剩余的区间加入到前面的区间，得到最终结果。
