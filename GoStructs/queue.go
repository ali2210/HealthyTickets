package GoStructs




type QueueList struct{
	ELement string
	Length int
}

	
type List struct{
	QueueList
	len int
}



func (q QueueList)Enque(ele string)QueueList{
	// Enque Code

	if q.isQueueNil(){
		q.ELement = ele
		q.Length = q.Length + 0
	}
	q.Length = q.Length + 1
	println("len queue:", q.Length)
	return q
}


func (q QueueList)DeQueue()(string){

	if q.isCapacityFull(q.Length+1){
		q.Length = (q.Length -1) 
	}else{
		if q.isQueueNil() {
			q.Length = 0
		}
		q.Length = (q.Length -1)
	}
	println("len DeQueue:", q.Length)
	return q.ELement

}

func (q QueueList)Print(){
	println("Data print:", q.ELement)
}

func (q QueueList)isCapacityFull(len int)(bool){

	if q.Length == 	len{
		return true
	}
	return false
}

func (q QueueList)isQueueNil()(bool){

	if q.Length == 0{
		return true
	}
	return false
}


func (l List)Add(q QueueList) (List, int){
	// mylist := List{}
		if l.len == 0{
		  	l.len = l.len + 1
		  	l.QueueList = q
		}else{
			l.len = l.len +1
			l.QueueList = q
		}
	return l, l.len
}


func (l List)Delete(q QueueList) *List{

	if l.Find(q) {
		l.QueueList.ELement = ""
		l.QueueList.Length-=1 
	}
	l.len= l.len -1
	return &l
}

func (l List)Find(q QueueList)bool{
	
	if l.QueueList == q{
		return true
	}
	return false 
}