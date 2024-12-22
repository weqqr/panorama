package mesh

import (
	. "github.com/lord-server/panorama/pkg/linalg"
)

type Vertex struct {
	Position Vector3
	Texcoord Vector2
	Normal   Vector3
}

type Mesh struct {
	Vertices []Vertex
}

func NewMesh() Mesh {
	return Mesh{
		Vertices: []Vertex{},
	}
}

type Model struct {
	Meshes []Mesh
}

func NewModel() Model {
	return Model{
		Meshes: []Mesh{},
	}
}
