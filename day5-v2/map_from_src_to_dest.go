package main

import (
	"context"
	"log"

	"github.com/golang-collections/collections/queue"
)

type SrcToDestStage struct {
	srcToDestMap *SrcToDestMap
}

func (stdStage *SrcToDestStage) Process(ctx context.Context, in <-chan interface{}) <-chan interface{} {

	out := make(chan interface{})
	go func() {
		defer close(out)
		for v := range in {
			q1 := queue.New()
			q1.Enqueue(v.(*Range))
			q2 := queue.New()
			for _, stdrm := range stdStage.srcToDestMap.Map {
				q2.Enqueue(stdrm)
			}

			for q1.Len() > 0 && q2.Len() > 0 {
				// TODO: change with slog
				// log.Printf("Source: %s, Destination: %s\n",
				// 	stdStage.srcToDestMap.SourceString,
				// 	stdStage.srcToDestMap.DestinationString)
				select {
				case <-ctx.Done():
					// Context has been cancelled, exit the goroutine
					return
				default:
					srcVal := q1.Dequeue().(*Range)
					srcToDestRangeMap := q2.Dequeue().(*SingleRangeMap)

					overlap, nonoverlap, err := srcVal.GetOverlaps(srcToDestRangeMap.SourceRange)
					if err != nil {
						// TODO: change with slog
						// log.Printf("Source: %s, Destination: %s\n",
						// 	stdStage.srcToDestMap.SourceString,
						// 	stdStage.srcToDestMap.DestinationString)
						// log.Printf("srcVal: %s, srcToDestRangeMap.SourceRange: %s\n", srcVal, srcToDestRangeMap.SourceRange)
						log.Fatalf("Error: %s\n", err)
					}

					if overlap != nil {
						destRange, err := srcToDestRangeMap.Map(overlap)
						if err != nil {
							// TODO: change with slog
							// log.Printf("Source: %s, Destination: %s\n",
							// 	stdStage.srcToDestMap.SourceString,
							// 	stdStage.srcToDestMap.DestinationString)
							log.Fatalf("Error: %s\n", err)
						}
						out <- destRange
					}

					if len(nonoverlap) > 0 {
						for _, nr := range nonoverlap {
							q1.Enqueue(&Range{
								Start:  nr.Start,
								Length: nr.Length,
							})
						}
					}
				}
			}

			for q1.Len() > 0 {
				select {
				case <-ctx.Done():
					return
				default:
					srcVal := q1.Dequeue().(*Range)
					out <- &Range{
						Start:  srcVal.Start,
						Length: srcVal.Length,
					}
				}
			}
		}
	}()
	return out
}
