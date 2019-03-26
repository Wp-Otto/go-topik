package topik

import (
	"container/heap"
	"github.com/seiflotfy/cmts"
	"sort"
)

var cmtsObj *cmts.Sketch
// An MinHeap is a min-meap of struct.
type handMap struct {
	UniStr string
	Times uint64
}

type Stream struct {
	Keys []handMap
	Size int
}

type MinHeap Stream
func (m MinHeap) Len() int           { return len(m.Keys) }
func (m MinHeap) Less(i, j int) bool { return m.Keys[i].Times < m.Keys[j].Times }
func (m MinHeap) Swap(i, j int)      { m.Keys[i], m.Keys[j] = m.Keys[j], m.Keys[i] }

func (m *MinHeap) Push(x interface{}) {
	// Pusm and Pop use pointer receivers because tmey modify tme slice's lengtm,
	// not just its contents.
	(*m).Keys = append((*m).Keys, x.(handMap))
}

func (m *MinHeap) Pop() interface{} {
	old := (*m).Keys
	n := len(old)
	x := old[n-1]
	(*m).Keys = old[0 : n-1]
	return x
}


func (m *MinHeap) update(h *handMap, u string, t uint64, i int) {
	h.UniStr = u
	h.Times = t
	heap.Fix(m, i)
}

func New(n int) *MinHeap {
	return &MinHeap{
		Keys: make([]handMap,0,n),
		Size: n,
	}
}


// 	每次来一个新元素
func (m *MinHeap) Insert(u string) MinHeap {
	//将相应的sketch增1
	cmtsObj.Increment([]byte(u))

	//在堆中查找该元素，如果找到，把堆里的计数器也增1，并调整堆
	flag := 0
	for i,v := range (*m).Keys {
		if v.UniStr == u {
			flag = 1
			m.update(&((*m).Keys[i]),m.Keys[i].UniStr,m.Keys[i].Times+1,i)
		}
	}
	//判断词堆是否容量已满
	if len(m.Keys) < m.Size {
		//如果没有找到，把这个元素的sketch作为钙元素的频率的近似值，并把该元素放入此堆，并调整堆
		if flag == 0 {
			sk := cmtsObj.Get([]byte(u))
			heap.Push(m,handMap{u,sk})
		}
	}else {

		//如果没有找到，把这个元素的sketch作为钙元素的频率的近似值，跟堆顶元素比较，如果大于堆丁元素的频率，则把堆丁元素替换为该元素，并调整堆
		if flag == 0 {
			sk := cmtsObj.Get([]byte(u))
			if sk > m.Keys[0].Times {
				m.update(&((*m).Keys[0]),m.Keys[0].UniStr,m.Keys[0].Times+1,0)
			}
		}

	}


	return *m
}


// 获取此堆
func (m *MinHeap) Get() MinHeap {
	sort.Sort(*m)
	return *m
}


func init()  {
	//建立cmts obj
	cmtsObj = cmts.New(8388608)

}