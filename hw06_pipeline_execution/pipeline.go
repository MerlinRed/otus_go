package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func run(in In, done In) Out {
	out := make(Bi)

	go func(out Bi) {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case data, ok := <-in:
				if !ok {
					return
				}
				out <- data
			}
		}
	}(out)

	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = stage(run(in, done))
	}
	return in
}
