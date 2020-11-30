package tree

import (
	"encoding/json"
	"fmt"
)

//Datanode 範例結構
type Datanode struct {
	ID       int         `json:"id"`
	ParentID int         `json:"pid"`
	Name     string      `json:"name"`
	Child    []*Datanode `json:"child"`
}

//Jsondata 存取json資料
var Jsondata []byte

//makeTree 參數父節點和全部節點的Ptr
func makeTree(Allnode []*Datanode, node *Datanode) {
	childs, _ := haveChild(Allnode, node) //判斷節點是否有子節點並return
	if childs != nil {
		fmt.Println("父節點:", *node)
		fmt.Println("子節點:")

		for _, v := range childs {
			fmt.Println(">>", *v)
		}

		node.Child = append(node.Child, childs[0:]...) //添加子節點
		for _, v := range childs {                     //查詢子節點的子節點，並加到子節點
			_, has := haveChild(Allnode, v)
			if has {
				makeTree(Allnode, v) //遞迴檢查並增加節點
			}
		}
	}
}

//haveChild 檢查是否有子節點
func haveChild(Allnode []*Datanode, node *Datanode) (childs []*Datanode, yes bool) {
	for _, v := range Allnode {
		if v.ParentID == node.ID {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}

//transformjson struct資料轉成json格式
func transformjson(Data *Datanode) {
	Jsondata, _ = json.Marshal(Data)
	//fmt.Println(string(Jsondata))
}

//jsontotree json轉struct 並印出樹狀結構
func jsontotree(jsondata []byte) {
	var TreeNode *Datanode
	err := json.Unmarshal(jsondata, &TreeNode)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(TreeNode.ID, TreeNode.ParentID, TreeNode.Name)
		printTree(TreeNode)

	}
}

func printTree(TreeNode *Datanode) {
	for _, v := range TreeNode.Child {
		fmt.Printf("%d,%d,%s\n", v.ID, v.ParentID, v.Name)
		if len(v.Child) != 0 {
			printTree(v)
		}
	}
}

//ThisTreeDemo 示範樹狀結構
func ThisTreeDemo() {
	Data := make([]*Datanode, 0) //存储所有初始化struct
	a := Datanode{
		ID:       0,
		ParentID: -1,
		Name:     "動物目錄",
	}
	Data = append(Data, &a)

	b := Datanode{
		ID:       1,
		ParentID: 0,
		Name:     "一、犬科",
	}
	Data = append(Data, &b)

	c := Datanode{
		ID:       2,
		ParentID: 1,
		Name:     "1.哈士奇",
	}
	Data = append(Data, &c)

	d := Datanode{
		ID:       3,
		ParentID: 1,
		Name:     "2.黃金獵犬",
	}
	Data = append(Data, &d)

	e := Datanode{
		ID:       4,
		ParentID: 0,
		Name:     "二、貓科",
	}
	Data = append(Data, &e)

	f := Datanode{
		ID:       5,
		ParentID: 4,
		Name:     "1.賓士貓",
	}
	Data = append(Data, &f)

	g := Datanode{
		ID:       6,
		ParentID: 4,
		Name:     "2.橘貓",
	}
	Data = append(Data, &g)

	h := Datanode{
		ID:       7,
		ParentID: 6,
		Name:     "特點肥",
	}
	Data = append(Data, &h)

	i := Datanode{
		ID:       8,
		ParentID: 4,
		Name:     "3.豹貓",
	}
	Data = append(Data, &i)

	j := Datanode{
		ID:       9,
		ParentID: 0,
		Name:     "三、關係",
	}
	Data = append(Data, &j)

	Anode := Data[0]      //父節點
	makeTree(Data, Anode) //生成tree
	transformjson(Anode)  //struct轉成json
	jsontotree(Jsondata)  //json轉成struct
}
