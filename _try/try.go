package _try

type try struct {
	logic func()
	catch func(err any)
}

func New(logic func()) *try {
	return &try{
		logic: logic,
		catch: func(err any) {},
	}
}
func (this *try) Catch(catch func(err any)) *try {
	this.catch = catch
	return this
}
func (this *try) Do() {
	defer func() {
		if err := recover(); nil != err {
			this.catch(err)
		}
	}()
	this.logic()
}
