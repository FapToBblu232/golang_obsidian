from PIL import Image, ImageDraw
import sys
import math



def MdotM(M1, M2):
    M = [[0]*4 for _ in range(4)]
    for i in range(4):
        for j in range(4):
            for k in range(4):
                M[i][j] += M1[i][k] * M2[k][j]
    return M

def MdotV(M, V):
    R = [0]*4
    for i in range(4):
        for j in range(4):
            R[i] += M[i][j] * V[j]
    return R



def translate3d(dx, dy, dz):
    return [
        [1, 0, 0, dx],
        [0, 1, 0, dy],
        [0, 0, 1, dz],
        [0, 0, 0, 1]
    ]

def scale3d(sx, sy, sz):
    return [
        [sx, 0,  0,  0],
        [0,  sy, 0,  0],
        [0,  0,  sz, 0],
        [0,  0,  0,  1]
    ]

def rotateX(angle):
    c, s = math.cos(angle), math.sin(angle)
    return [
        [1, 0, 0, 0],
        [0, c, -s, 0],
        [0, s,  c, 0],
        [0, 0, 0, 1]
    ]

def rotateY(angle):
    c, s = math.cos(angle), math.sin(angle)
    return [
        [ c, 0, s, 0],
        [ 0, 1, 0, 0],
        [-s, 0, c, 0],
        [ 0, 0, 0, 1]
    ]

def rotateZ(angle):
    c, s = math.cos(angle), math.sin(angle)
    return [
        [c, -s, 0, 0],
        [s,  c, 0, 0],
        [0,  0, 1, 0],
        [0,  0, 0, 1]
    ]


def apply_matrix3d(M, p):
    x, y, z = p
    v = [x, y, z, 1]
    r = MdotV(M, v)
    return (r[0], r[1], r[2])

def transform_points3d(points, M):
    return [apply_matrix3d(M, p) for p in points]

def parse_file(filename):
    data = {'vertices': [], 'faces': []}

    with open(filename) as f:
        for line in f:
            if line.startswith("v "):
                x, y, z = line.split()[1:]
                data['vertices'].append((float(x), float(y), float(z)))

            elif line.startswith("f "):
                face = [int(x.split("/")[0]) for x in line.split()[1:]]
                data['faces'].append(face)

    return data

def draw_figure(data):
    img = Image.new('RGB', (400, 400), (0, 0, 0))
    dr = ImageDraw.Draw(img)

    # draw points
    for v in data['vertices']:
        x = int(v[0])
        y = int(v[1])
        if 0 <= x < 400 and 0 <= y < 400:
            img.putpixel((x, y), (255, 0, 0))

    # draw edges
    for face in data['faces']:
        n = len(face)

        for i in range(n):
            v1 = data['vertices'][face[i] - 1]
            v2 = data['vertices'][face[(i + 1) % n] - 1]
            #print(v1,v2)
            dr.line(
                [(v1[0], v1[1]), (v2[0], v2[1])],
                fill=(255, 255, 255), width=1
            )

    return img


def is_visible(vertices, face):

    A = vertices[face[0] - 1]
    B = vertices[face[1] - 1]
    C = vertices[face[-1] - 1]

    AB = (B[0] - A[0], B[1] - A[1], B[2] - A[2])
    AC = (C[0] - A[0], C[1] - A[1], C[2] - A[2])

    N = (
        AC[1]*AB[2] - AC[2]*AB[1],
        AC[2]*AB[0] - AC[0]*AB[2],
        AC[0]*AB[1] - AC[1]*AB[0]
    )

    Z = N[2] * -1
    print(Z)
    return Z > 0

def draw_figure_culled(data):
    img = Image.new('RGB', (400, 400), (0, 0, 0))
    dr = ImageDraw.Draw(img)

    for face in data['faces']:

        if not is_visible(data['vertices'], face):
            continue
        

        n = len(face)
        for i in range(n):
            v1 = data['vertices'][face[i] - 1]
            v2 = data['vertices'][face[(i + 1) % n] - 1]
            #print(v1,v2)
            dr.line(
                [(v1[0], v1[1]), (v2[0], v2[1])],
                fill=(255, 255, 255), width=1
            )

    return img

if __name__ == "__main__":
    filename = "pyramid (1).obj"
    if len(sys.argv) > 1:
        filename = sys.argv[1]

    data = parse_file(filename)

    # 3D transformations
    T = translate3d(200, 200, 0)
    S = scale3d(50, 50, 50)
    R = rotateY(math.pi / 4)

    M = MdotM(T, MdotM(R, S))

    data1 = data.copy()
    data1["vertices"] = transform_points3d(data["vertices"], M)

    img = draw_figure_culled(data1)
    #print("-------")
    #img = draw_figure(data1)
    img.show()
    img.save("figure.png")
