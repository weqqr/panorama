package mesh

import . "github.com/lord-server/panorama/pkg/linalg"

type CubeFaces uint8

const (
	CubeFaceNone CubeFaces = 0
	CubeFaceEast           = 1 << iota
	CubeFaceWest
	CubeFaceTop
	CubeFaceDown
	CubeFaceNorth
	CubeFaceSouth
)

//nolint:funlen // FIXME: embed and scale .obj file?
func Cuboid(x1, y1, z1, x2, y2, z2 float64, hiddenFaces CubeFaces) []Mesh {
	yp := NewMesh()
	yp.Vertices = []Vertex{
		{Position: Vec3(x1, y2, z1), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, 1.0, 0.0)},
		{Position: Vec3(x1, y2, z2), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(0.0, 1.0, 0.0)},
		{Position: Vec3(x2, y2, z2), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, 1.0, 0.0)},
		{Position: Vec3(x1, y2, z1), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, 1.0, 0.0)},
		{Position: Vec3(x2, y2, z1), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(0.0, 1.0, 0.0)},
		{Position: Vec3(x2, y2, z2), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, 1.0, 0.0)},
	}

	ym := NewMesh()
	ym.Vertices = []Vertex{
		{Position: Vec3(x1, y1, z1), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, -1.0, 0.0)},
		{Position: Vec3(x1, y1, z2), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(0.0, -1.0, 0.0)},
		{Position: Vec3(x2, y1, z2), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, -1.0, 0.0)},
		{Position: Vec3(x1, y1, z1), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, -1.0, 0.0)},
		{Position: Vec3(x2, y1, z1), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(0.0, -1.0, 0.0)},
		{Position: Vec3(x2, y1, z2), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, -1.0, 0.0)},
	}

	xp := NewMesh()
	xp.Vertices = []Vertex{
		{Position: Vec3(x2, y1, z1), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(1.0, 0.0, 0.0)},
		{Position: Vec3(x2, y1, z2), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(1.0, 0.0, 0.0)},
		{Position: Vec3(x2, y2, z2), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(1.0, 0.0, 0.0)},
		{Position: Vec3(x2, y1, z1), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(1.0, 0.0, 0.0)},
		{Position: Vec3(x2, y2, z1), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(1.0, 0.0, 0.0)},
		{Position: Vec3(x2, y2, z2), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(1.0, 0.0, 0.0)},
	}

	xm := NewMesh()
	xm.Vertices = []Vertex{
		{Position: Vec3(x1, y1, z1), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(-1.0, 0.0, 0.0)},
		{Position: Vec3(x1, y1, z2), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(-1.0, 0.0, 0.0)},
		{Position: Vec3(x1, y2, z2), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(-1.0, 0.0, 0.0)},
		{Position: Vec3(x1, y1, z1), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(-1.0, 0.0, 0.0)},
		{Position: Vec3(x1, y2, z1), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(-1.0, 0.0, 0.0)},
		{Position: Vec3(x1, y2, z2), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(-1.0, 0.0, 0.0)},
	}

	zp := NewMesh()
	zp.Vertices = []Vertex{
		{Position: Vec3(x1, y1, z2), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, 0.0, 1.0)},
		{Position: Vec3(x1, y2, z2), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(0.0, 0.0, 1.0)},
		{Position: Vec3(x2, y2, z2), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, 0.0, 1.0)},
		{Position: Vec3(x1, y1, z2), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, 0.0, 1.0)},
		{Position: Vec3(x2, y1, z2), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(0.0, 0.0, 1.0)},
		{Position: Vec3(x2, y2, z2), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, 0.0, 1.0)},
	}

	zm := NewMesh()
	zm.Vertices = []Vertex{
		{Position: Vec3(x1, y1, z1), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, 0.0, -1.0)},
		{Position: Vec3(x1, y2, z1), Texcoord: Vec2(0.0, 1.0), Normal: Vec3(0.0, 0.0, -1.0)},
		{Position: Vec3(x2, y2, z1), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, 0.0, -1.0)},
		{Position: Vec3(x1, y1, z1), Texcoord: Vec2(0.0, 0.0), Normal: Vec3(0.0, 0.0, -1.0)},
		{Position: Vec3(x2, y1, z1), Texcoord: Vec2(1.0, 0.0), Normal: Vec3(0.0, 0.0, -1.0)},
		{Position: Vec3(x2, y2, z1), Texcoord: Vec2(1.0, 1.0), Normal: Vec3(0.0, 0.0, -1.0)},
	}

	meshes := []Mesh{yp, ym, xp, xm, zp, zm}
	meshFaces := []CubeFaces{CubeFaceTop, CubeFaceDown, CubeFaceEast, CubeFaceWest, CubeFaceNorth, CubeFaceSouth}

	culledMesh := []Mesh{}

	for i, mesh := range meshes {
		if hiddenFaces&meshFaces[i] == 0 {
			culledMesh = append(culledMesh, mesh)
		}
	}

	return culledMesh
}

func Cube(hiddenFaces CubeFaces) *Model {
	model := NewModel()

	model.Meshes = append(model.Meshes, Cuboid(-0.5, -0.5, -0.5, 0.5, 0.5, 0.5, hiddenFaces)...)

	return &model
}
