package image

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	v1Types "github.com/google/go-containerregistry/pkg/v1/types"
)

type LayerMetadata struct {
	Index uint
	// sha256 digest of the layer contents (the docker "diff id")
	Digest    string
	MediaType v1Types.MediaType
	// Size in bytes of the layer content size
	Size int64
}

func readLayerMetadata(imgMetadata Metadata, layer v1.Layer, idx int) (LayerMetadata, error) {
	mediaType, err := layer.MediaType()
	if err != nil {
		return LayerMetadata{}, err
	}

	// digest = diff-id = a digest of the uncompressed layer content
	diffIDHash := imgMetadata.Config.RootFS.DiffIDs[idx]
	return LayerMetadata{
		Index:     uint(idx),
		Digest:    diffIDHash.String(),
		MediaType: mediaType,
	}, nil
}
