package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func run(stage Stage, in In, done In) Out {
	out := make(Bi)
	stageOut := stage(in)

	go func(out Bi) {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case data, ok := <-stageOut:
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
	out := in
	for _, stage := range stages {
		out = run(stage, out, done)
	}

	return out
}
