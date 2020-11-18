package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/dependencies"
	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func buildsystemsDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	c, err := createCompilerDigests(status, eval)
	if err != nil {
		return nil, err
	}
	digests = append(digests, *c)

	i, err := createImagesDigests(status, eval)
	if err != nil {
		return nil, err
	}
	digests = append(digests, *i)

	d, err := createContainerDependenciesDigests(status, eval)
	if err != nil {
		return nil, err
	}
	digests = append(digests, *d)

	return digests, nil
}

func createCompilerDigests(status *scanner.ScanStatus, eval *scans.Evaluation) (*Digest, error) {
	d := NewDigest(status, CompilersIndex, "compiler", "compilers")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.BuildsystemResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into buildss bytes")
		}

		if b.Compilers == nil {
			b.Compilers = make([]scans.Compiler, 0)
		}

		d.MarshalSourceData(b.Compilers, "compilers")

		switch len(b.Compilers) {
		case 0:
			err := d.AppendEval(eval, "chars", "none detected")
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}
		case 1:
			n := ""
			for _, c := range b.Compilers {
				n = c.Name
			}

			err := d.AppendEval(eval, "chars", n)
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}

			d.UseSingularTitle()
		default:
			err := d.AppendEval(eval, "count", len(b.Compilers))
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	return d, nil
}

func createImagesDigests(status *scanner.ScanStatus, eval *scans.Evaluation) (*Digest, error) {
	d := NewDigest(status, ContainerImagesIndex, "container image", "container images")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.BuildsystemResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into buildss bytes")
		}

		if b.Dockerfile.Images == nil {
			b.Dockerfile.Images = make([]scans.Image, 0)
		}

		d.MarshalSourceData(b.Dockerfile.Images, "container images")

		switch len(b.Dockerfile.Images) {
		case 0:
			err := d.AppendEval(eval, "chars", "none detected")
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}
		case 1:
			n := ""
			for _, c := range b.Dockerfile.Images {
				n = c.Name
			}

			err := d.AppendEval(eval, "chars", n)
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}

			d.UseSingularTitle()
		default:
			err := d.AppendEval(eval, "count", len(b.Dockerfile.Images))
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	return d, nil
}

func createContainerDependenciesDigests(status *scanner.ScanStatus, eval *scans.Evaluation) (*Digest, error) {
	d := NewDigest(status, ContainerDependenciesIndex, "container dependencies", "container dependencies")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.BuildsystemResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated container dep results into builds bytes")
		}

		if b.Dockerfile.Dependencies == nil {
			b.Dockerfile.Dependencies = make([]dependencies.Dependency, 0)
		}

		d.MarshalSourceData(b.Dockerfile.Dependencies, "container dependencies")

		switch len(b.Dockerfile.Dependencies) {
		case 0:
			err := d.AppendEval(eval, "chars", "none detected")
			if err != nil {
				return nil, fmt.Errorf("failed to create builds container dep digest: %v", err.Error())
			}
		case 1:
			n := ""
			for _, c := range b.Dockerfile.Dependencies {
				n = c.Name
			}

			err := d.AppendEval(eval, "chars", n)
			if err != nil {
				return nil, fmt.Errorf("failed to create builds container dep digest: %v", err.Error())
			}

			d.UseSingularTitle()
		default:
			err := d.AppendEval(eval, "count", len(b.Dockerfile.Dependencies))
			if err != nil {
				return nil, fmt.Errorf("failed to create builds container dep digest: %v", err.Error())
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	return d, nil
}
