package functions

type Block struct {
	Try   func()
	Catch func(Exception)
}

type Exception interface{}

func Throw(err Exception) {
	ShowLog("Throw", err)
}

func (tcf Block) Do() {
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
