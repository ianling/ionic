package cyclonedx

import (
    "fmt"
    "github.com/CycloneDX/cyclonedx-go"
    "strings"
)

// FromJSONString takes a CycloneDX SBOM in JSON format, as a string, and returns a BOM object
func FromJSONString(sbomContents string) (*cyclonedx.BOM, error) {
    bom, err := fromString(sbomContents, cyclonedx.BOMFileFormatJSON)
    if err != nil {
        return nil, fmt.Errorf("failed to parse CycloneDX SBOM in JSON format: %s", err.Error())
    }

    return bom, nil
}

// FromXMLString takes a CycloneDX SBOM in XML format, as a string, and returns a BOM object
func FromXMLString(sbomContents string) (*cyclonedx.BOM, error) {
    bom, err := fromString(sbomContents, cyclonedx.BOMFileFormatXML)
    if err != nil {
        return nil, fmt.Errorf("failed to parse CycloneDX SBOM in XML format: %s", err.Error())
    }

    return bom, nil
}

func fromString(sbomContents string, fileType cyclonedx.BOMFileFormat) (*cyclonedx.BOM, error) {
    // need a stream-like object to decode
    sbomContentsStream := strings.NewReader(sbomContents)

    bom := new(cyclonedx.BOM)
    decoder := cyclonedx.NewBOMDecoder(sbomContentsStream, fileType)
    if err := decoder.Decode(bom); err != nil {
        return nil, err
    }

    return bom, nil
}
