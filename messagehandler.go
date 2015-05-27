package scrolls

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
)

type req map[string]interface{}

type messageHandler struct {
	net.Conn
	encoder   *json.Encoder
	callbacks map[string][]interface{}
	closer    chan struct{}
}

// connects to url via TCP and start listening for and relaying json messages
func newMessageHandler(url string) (*messageHandler, error) {
	c, err := net.Dial("tcp", url)
	if err != nil {
		return nil, err
	}

	conn := &messageHandler{
		c,
		json.NewEncoder(c),
		map[string][]interface{}{},
		make(chan struct{}),
	}

	// start listening for json messages
	go func() {
		r := bufio.NewReader(c)
		for {
			// messages are separated by a newline
			line, err := r.ReadBytes('\n')
			if err != nil {
				close(conn.closer)
				return
			}

			// sometimes we'll receive empty lines, ignore those
			if len(line) <= 1 {
				continue
			}

			// extract the 'msg' field
			var m struct{ Msg string }
			deny(json.Unmarshal(line, &m))

			// callbacks are cleared when called
			callbacks := conn.callbacks[m.Msg]
			conn.callbacks[m.Msg] = []interface{}{}

			if len(callbacks) == 0 {
				fmt.Println("unhandled message", m.Msg)
			}

			for _, f := range callbacks {
				m := reflect.New(reflect.TypeOf(f).In(0))
				deny(json.Unmarshal(line, m.Interface()))
				reflect.ValueOf(f).Call([]reflect.Value{m.Elem()})
			}
		}
	}()

	return conn, nil
}

// Sends a message to the Scrolls server.
// The parameter must be of a type defined in messagetypes.go
func (c messageHandler) send(m interface{}) {

	// encodes a struct as req by making sure that
	// all struct fields start with a lowercase character, as expected by
	// the Scrolls server.
	var encode func(reflect.Value) req
	encode = func(v reflect.Value) req {
		out := req{}

		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i)

			name := t.Field(i).Name
			name = strings.ToLower(name[:1]) + name[1:]

			switch val.Kind() {
			case reflect.String:
				out[name] = val.String()
			case reflect.Int:
				out[name] = strconv.Itoa(int(val.Int()))
			case reflect.Struct:
				out[name] = encode(val)
			case reflect.Slice:
				n := val.Len()
				arr := make([]interface{}, n)
				for i := 0; i < n; i++ {
					arr[i] = encode(val.Index(i))
				}
				out[name] = arr
			default:
				log.Fatalln("send: struct contains unhandled field type", val.Kind())
			}
		}

		return out
	}

	req := encode(reflect.ValueOf(m))
	req["msg"] = reflect.TypeOf(m).Name()[1:]

	c.encoder.Encode(req)
}

// Like send(), but waits for either an Ok message or a Fail message,
// in which case an error is returned.
func (c messageHandler) confirm(m interface{}) error {
	ch := make(chan error, 2)
	c.receive(func(m mOk) {
		ch <- nil
	})
	c.receive(func(m mFail) {
		ch <- errors.New(m.Info)
	})
	c.send(m)
	return <-ch
}

func (c messageHandler) request(data req, f interface{}) {
	if data == nil {
		data = req{}
	}
	data["msg"] = reflect.TypeOf(f).In(0).Name()[1:]

	c.receive(f)
	bytes, err := json.Marshal(data)
	deny(err)
	c.Write(bytes)
}

// Registers a callback that is executed once when a certain message
// is received. f must be of type func(M), where M is one of the
// types defined in messagetypes.go
func (c messageHandler) receive(f interface{}) {
	msg := reflect.TypeOf(f).In(0).Name()[1:]
	c.callbacks[msg] = append(c.callbacks[msg], f)
}
