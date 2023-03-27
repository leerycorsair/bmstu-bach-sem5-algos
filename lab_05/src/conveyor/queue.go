package conveyor

type Queue []*Ration_t

func (queue *Queue) push(ration *Ration_t) {
	*queue = append(*queue, ration)
}

func (queue *Queue) pop() *Ration_t {
	var ration *Ration_t
	*queue, ration = (*queue)[1:], (*queue)[0]
	return ration
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}
