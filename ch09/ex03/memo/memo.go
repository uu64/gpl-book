package memo

import "fmt"

// Func is the type of the function to memoize.
type Func func(key string, cancelled ...<-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value       interface{}
	err         error
	isCancelled bool
}

type entry struct {
	res       result
	ready     chan struct{} // closed when res is ready
	cancelled <-chan struct{}
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key       string
	response  chan<- result // the client wants a single result
	cancelled <-chan struct{}
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancelled ...<-chan struct{}) (interface{}, error) {
	response := make(chan result)

	// cancelledの先頭しか使わない
	var c <-chan struct{} = nil
	if len(cancelled) > 0 {
		c = cancelled[0]
	}

	memo.requests <- request{key, response, c}
	res := <-response
	// NOTE: キャッシュを待っていたリクエストがキャンセルされたときにリトライしたいがうまくいかない
	// res := result{isCancelled: true}
	// for res.isCancelled {
	// 	memo.requests <- request{key, response, c}
	// 	res = <-response
	// }
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{
				ready:     make(chan struct{}),
				cancelled: req.cancelled,
			}
			go e.call(f, req.key) // call f(key)
			select {
			case <-e.ready:
				cache[req.key] = e
			case <-e.cancelled:
				cache[req.key] = nil
			}
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	e.res.isCancelled = false
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	select {
	// Wait for the ready condition.
	case <-e.ready:
		// Send the result to the client.
		response <- e.res
	case <-e.cancelled:
		response <- result{
			err:         fmt.Errorf("request was cancelled"),
			isCancelled: true,
		}
	}

}
