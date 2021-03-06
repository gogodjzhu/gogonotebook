# 贪心算法

贪心算法的核心是将一个大的任务划分成小任务分别执行，在每个小任务的执行中使用最优策略，希望以局部的最优解推导出全局的最优解。

但是这种从局部最优到全局最优的推导必须满足`无后效`性条件，即每个任务之间不会互相影响彼此的选择。

比如说有任务1、任务2先后执行，任务1有选择\[A,B,C\]任务2有选择\[D,E,F\]，不论任务1执行任何选择，都不会影响任务2的选择，就叫做`无后效性`。

> #### 0-1背包问题(*ZeroOneKnapsack*)
> 
> 0-1背包问题是典型的可以使用贪心算法求近似解，但却无法求最优解的例子。因为每次选择之后，会影响背包剩余的容量，导致下一次选择的时候受到容量约束(`有后效性`)。
> 
> 但如果将物品的重量设定为可分割，则不存在`后效性`影响。

*PS. 事实上贪心算法适用的场景比较少。因为大多数问题都是有后效性的*
