<p align="center"><font size=30>go-topik</font></p>

<p align="center">
<a href="https://godoc.org/github.com/Wp-Otto/go-topik"><img src="https://img.shields.io/badge/topk-godoc-brightgreen.svg" alt="godoc"></a>
</p>

Streaming Topk service using Count-min sketch probabilistic data structure 

Default Count-Min Sketch distribution 8M Memory

寻找数据流中出现最频繁的k个元素(find top k frequent items in a data stream)。这个问题也称为 Heavy Hitters.

这题也是从实践中提炼而来的，例如搜索引擎的热搜榜，找出访问网站次数最多的前10个IP地址，等等。

方案1: HashMap + Heap
用一个 HashMap<String, Long>，存放所有元素出现的次数，用一个小根堆，容量为k，存放目前出现过的最频繁的k个元素，

每次从数据流来一个元素，如果在HashMap里已存在，则把对应的计数器增1，如果不存在，则插入，计数器初始化为1
在堆里查找该元素，如果找到，把堆里的计数器也增1，并调整堆；如果没有找到，把这个元素的次数跟堆顶元素比较，如果大于堆丁元素的出现次数，则把堆丁元素替换为该元素，并调整堆
空间复杂度O(n)。HashMap需要存放下所有元素，需要O(n)的空间，堆需要存放k个元素，需要O(k)的空间，跟O(n)相比可以忽略不急，总的时间复杂度是O(n)
时间复杂度O(n)。每次来一个新元素，需要在HashMap里查找一下，需要O(1)的时间；然后要在堆里查找一下，O(k)的时间，有可能需要调堆，又需要O(logk)的时间，总的时间复杂度是O(n(k+logk))，k是常量，所以可以看做是O(n)。
如果元素数量巨大，单机内存存不下，怎么办？ 有两个办法，见方案2。

方案2: Count-Min Sketch + Heap
既然方案1中的HashMap太大，内存装不小，那么可以用Count-Min Sketch算法代替HashMap，

在数据流不断流入的过程中，维护一个标准的Count-Min Sketch 二维数组
维护一个小根堆，容量为k
每次来一个新元素，
将相应的sketch增1
在堆中查找该元素，如果找到，把堆里的计数器也增1，并调整堆；如果没有找到，把这个元素的sketch作为钙元素的频率的近似值，跟堆顶元素比较，如果大于堆丁元素的频率，则把堆丁元素替换为该元素，并调整堆
这个方法的时间复杂度和空间复杂度如下：

空间复杂度O(dm)。m是二维数组的列数，d是二维数组的行数，堆需要O(k)的空间，不过k通常很小，堆的空间可以忽略不计
时间复杂度O(nlogk)。每次来一个新元素，需要在二维数组里查找一下，需要O(1)的时间；然后要在堆里查找一下，O(logk)的时间，有可能需要调堆，又需要O(logk)的时间，总的时间复杂度是O(nlogk)。

该packages参考了许多github的大佬实现的  勿喷。
