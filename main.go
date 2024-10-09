package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	if err := run(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}

func run(w io.Writer, in io.Reader) error {
	ctx := context.Background()

	exporter, err := stdouttrace.New(stdouttrace.WithWriter(w))
	if err != nil {
		return err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tracerProvider)
	defer tracerProvider.Shutdown(ctx)

	line := 1

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		text := scanner.Text()
		if err := handle(ctx, text); err != nil {
			return fmt.Errorf("%w on line %d: %q", err, line, text)
		}
		line++
	}

	return scanner.Err()
}

func handle(ctx context.Context, text string) error {
	chunks := strings.SplitN(text, ",", 3)
	if len(chunks) != 3 {
		return fmt.Errorf("wanted 3 chunks, got %d", len(chunks))
	}
	start, err := time.Parse(time.UnixDate, strings.TrimSpace(chunks[0]))
	if err != nil {
		return fmt.Errorf("parsing start: %w", err)
	}
	end, err := time.Parse(time.UnixDate, strings.TrimSpace(chunks[1]))
	if err != nil {
		return fmt.Errorf("parsing end: %w", err)
	}

	ctx, span := otel.Tracer("logspan").Start(ctx, chunks[2], trace.WithTimestamp(start))
	defer span.End(trace.WithTimestamp(end))

	return nil
}
