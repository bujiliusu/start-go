package queue


//type Queue []int
type Queue []interface{} //支持任何类型
func (q *Queue) Push(v int) { //限定传入类型是int
	*q = append(*q, v)
}
//func (q *Queue) Push(v interface{}) {
//	*q = append(*q, v.(int)) 限定传入类型是int
//}
func (q *Queue) Pop() interface{}{
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) //限定穿出是int
}
func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}