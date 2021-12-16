package bank

type withdrawReq struct {
	amount int
	reply  chan bool
}

var deposits = make(chan int)          // send amount to deposit
var balances = make(chan int)          // receive balance
var withdraws = make(chan withdrawReq) // withdraw money

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	var reply = make(chan bool)
	withdraws <- withdrawReq{amount, reply}
	result := <-reply
	return result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case req := <-withdraws:
			if req.amount <= balance {
				balance -= req.amount
				req.reply <- true
			} else {
				req.reply <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
