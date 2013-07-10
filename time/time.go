package time
import tm "time"
import "fmt"
import "errors"
type Time tm.Time

func (t *Time) Scan(ss fmt.ScanState, f rune) error {
	if f == 't' {
		var buf []rune
		token, err := ss.Token(false, func(r rune) bool {
			buf = append(buf, r)
			return !(len(buf) == 20)
		})
		if err != nil {
			panic("not empty err")
		}
		v, e := tm.Parse("2006-01-02 15:04:05", string(token))
		if e != nil {
			panic(e)
		}
		t1 := Time(v)
		t = &t1
		return nil
	}
	return errors.New("not supported fmt, only support %t, 2006-01-02 15:04:05")
}

